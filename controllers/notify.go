package controllers

import (
	"fmt"
	"talkGo/controllers"
	"os/exec"
	"io/ioutil"
	"bytes"
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
	roomName := c.GetString("name")
	tCurl := c.GetString("tcurl")

	cmd := exec.Command("ffmpeg", " -i " + tCurl + "/" + roomName ,  "-f image2",  "-ss 5", "-vframes 1",  "-s 220*220",  "/root/go/src/talkGo/static/hlsCover/" + roomName + "_cover.png")
	stdout, err := cmd.StdoutPipe()
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return
	}
	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return
	}
	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error(), stderr.String())
		return
	}
	fmt.Printf("stdout:\n\n %s", bytes)
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
