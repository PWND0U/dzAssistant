package assistant

import (
	"context"
	"dzassistant/backend/models"
	"github.com/imroc/req/v3"
)

type Assistant struct {
	ctx context.Context
}

func NewAssistant() *Assistant {
	return &Assistant{}
}

// Startup is called when the app starts. The context is saved
func (a *Assistant) Startup(ctx context.Context) {
	a.ctx = ctx
}

// GetBiliFollower 获取B站粉丝数量
func (a *Assistant) GetBiliFollower() int {
	c := req.C()
	c.ImpersonateChrome()
	var info models.BiLiInfo
	resp, err := c.R().SetSuccessResult(&info).Get("https://api.bilibili.com/x/relation/stat?vmid=1532368659&jsonp=jsonp")
	if err != nil {
		return 0
	}
	if resp.IsSuccessState() {
		return info.Data.Follower
	}
	return 0
}
