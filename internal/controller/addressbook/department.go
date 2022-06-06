package addressbook

import (
	"github.com/gin-gonic/gin"
	"github.com/open-scrm/open-scrm/internal/global"
	"github.com/open-scrm/open-scrm/internal/vo"
	"github.com/open-scrm/open-scrm/lib/redislock"
	"github.com/open-scrm/open-scrm/pkg/addressbook/service"
	"time"
)

const (
	lockSyncCorpStructureKey = "addressbook.sync"
)

// SyncCorpStructure 同步组织架构到数据库
func SyncCorpStructure(ctx *gin.Context) {
	lock := redislock.NewRedisLock(global.GetRedis())
	ok, err := lock.GetLock(ctx.Request.Context(), lockSyncCorpStructureKey, 5*time.Minute)
	if err != nil {
		vo.SendError(ctx, err)
		return
	}
	if !ok {
		vo.SendFail(ctx, "同步频率过快，请5分钟后再尝试")
		return
	}
	// 开始同步.
	defer lock.UnLock(ctx, lockSyncCorpStructureKey)

	if err := service.NewSyncCorpStructureService().DoSync(ctx.Request.Context()); err != nil {
		vo.SendError(ctx, err)
		return
	}
	vo.SendOK(ctx, nil)
}
