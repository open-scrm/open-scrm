package global

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v8"
	"github.com/open-scrm/open-scrm/configs"
	"github.com/open-scrm/open-scrm/lib/wxwork"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	redisClient *redis.Client

	wxWorkClient *wxwork.Client

	mongoDriver *mongo.Client

	snowflakeNode *snowflake.Node
)

func InitGlobal(ctx context.Context, conf *configs.Config) error {
	var (
		cli *redis.Client
	)
	{
		cli = redis.NewClient(&redis.Options{
			Addr:     conf.Redis.Addr,
			Password: conf.Redis.Password, // no password set
			DB:       conf.Redis.Db,       // use default DB
		})
		if err := cli.Ping(ctx).Err(); err != nil {
			logrus.WithContext(ctx).WithError(err).Errorf("ping redis failed")
			return err
		}
		redisClient = cli
	}
	// 企微客户端
	{
		akStore := wxwork.NewRedisStorage(cli, "cache.")
		wxWorkClient = wxwork.NewClient(akStore, wxwork.DefaultHttpClient)
	}
	// mongo driver
	{
		opt := options.Client().
			ApplyURI(conf.Mongo.Host).
			SetAuth(options.Credential{
				AuthMechanism: "SCRAM-SHA-1",
				AuthSource:    conf.Mongo.AdminDatabase,
				Username:      conf.Mongo.Username,
				Password:      conf.Mongo.Password,
				PasswordSet:   true,
			}).
			SetMinPoolSize(conf.Mongo.PoolSize).
			SetMaxPoolSize(conf.Mongo.MaxPoolSize).
			SetConnectTimeout(time.Duration(conf.Mongo.Timeout) * time.Second).
			SetAppName("open-scrm")
		driver, err := mongo.Connect(ctx, opt)
		if err != nil {
			return err
		}
		if err := driver.Ping(ctx, readpref.Primary()); err != nil {
			return err
		}
		mongoDriver = driver
	}
	{
		snowflakeNode, _ = snowflake.NewNode(1)
	}

	return nil
}

func GetRedis() *redis.Client {
	return redisClient
}

func GetWxWorkClient() *wxwork.Client {
	return wxWorkClient
}

func GetMongoDriver() *mongo.Client {
	return mongoDriver
}

func GetSnowflakeNode() *snowflake.Node {
	return snowflakeNode
}
