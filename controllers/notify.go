package controllers

import (
	"fmt"
	"talkGo/controllers"
)

type NotifyController struct {
	controllers.BaseController
}


// URLMapping ...
func (c *NotifyController) URLMapping() {
	c.Mapping("OnPublish", c.OnPublish)
	c.Mapping("OnPublishDone", c.OnPublishDone)
	c.Mapping("OnPlay", c.OnPlay)
	c.Mapping("OnPlayDone", c.OnPlayDone)
}

// 直播推流时通知...
// @Title OnPublish
// @Description up voice to server,chnage to text
// @Success 200 {object} models.Talk
// @Failure 403 :id is empty
// @router /onPublish [post]
func (c *NotifyController) OnPublish() {
	fmt.Println("OnPublish")
	fmt.Println("name = ", c.GetString("name"))
	fmt.Println("request body : ", string(c.Ctx.Input.RequestBody))
}


// 直播推流结束时通知...
// @Title OnPublishDone
// @Description up voice to server,chnage to text
// @Success 200 {object} models.Talk
// @Failure 403 :id is empty
// @router /onPublishDone [post]
func (c *NotifyController) OnPublishDone() {
	fmt.Println("OnPublishDone")
}


// 直播拉流时通知...
// @Title OnPlay
// @Description up voice to server,chnage to text
// @Success 200 {object} models.Talk
// @Failure 403 :id is empty
// @router /OnPlay [post]
func (c *NotifyController) OnPlay() {
	fmt.Println("OnPlay")
}

// 直播拉流结束时通知...
// @Title OnPlayDone
// @Description up voice to server,chnage to text
// @Success 200 {object} models.Talk
// @Failure 403 :id is empty
// @router /OnPlayDone [post]
func (c *NotifyController) OnPlayDone() {
	fmt.Println("OnPlayDone")
}
