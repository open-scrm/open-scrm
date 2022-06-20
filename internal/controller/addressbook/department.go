package addressbook

import (
	"github.com/gin-gonic/gin"
	"github.com/open-scrm/open-scrm/internal/global"
	"github.com/open-scrm/open-scrm/internal/vo"
	"github.com/open-scrm/open-scrm/lib/log"
	"github.com/open-scrm/open-scrm/lib/redislock"
	"github.com/open-scrm/open-scrm/pkg/addressbook/service"
	"github.com/open-scrm/open-scrm/pkg/response"
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
		response.SendError(ctx, err)
		return
	}
	if !ok {
		response.SendFail(ctx, "同步频率过快，请5分钟后再尝试")
		return
	}
	// 开始同步.
	defer lock.UnLock(ctx, lockSyncCorpStructureKey)

	if err := service.NewSyncCorpStructureService().DoSync(ctx.Request.Context()); err != nil {
		response.SendError(ctx, err)
		return
	}
	response.SendOK(ctx, nil)
}

// DepartmentList 部门管理：部门列表
func DepartmentList(ctx *gin.Context) {
	req := new(vo.AddressBookListDeptRequest)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithContext(ctx.Request.Context()).WithError(err).Errorf("DepartmentList 参数错误")
		response.SendFail(ctx, "参数错误")
		return
	}
	if err := req.Validate(); err != nil {
		response.SendError(ctx, err)
		return
	}
	res, err := service.NewDeptService().ListDepartmentTree(ctx, nil, req.Name)
	if err != nil {
		log.WithContext(ctx.Request.Context()).WithError(err).Errorf("DepartmentList 查询列表失败")
		response.SendError(ctx, err)
		return
	}
	response.SendOK(ctx, res)
}
