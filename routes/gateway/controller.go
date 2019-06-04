/**
 * Created by zeqi
 * @description
 * @module GatwayController
 * @version 1.0.0
 * @author Xijun Zhu <zhuzeqi2010@163.com>
 * @File controller
 * @Date 18-7-19
 * @Wechat zhuzeqi2010
 * @QQ 304566647
 * @Office-email zhuxj4@lenovo.com
 */

package gateway

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"vc.cli/utils"
)

type Controller struct {
}

func NewController() Controller {
	return Controller{}
}

// Status check
func (o *Controller) App(c *gin.Context) {
	// return html
	instance := utils.GetInstance()
	c.HTML(http.StatusOK, instance.Config.Server.MainPage, gin.H{})
}

func (o *Controller) Profile(c *gin.Context) {
	// return 200 with empty JSON body
	session := sessions.Default(c)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count += 1
	}
	session.Set("count", count)
	session.Save()
	c.JSON(http.StatusOK, gin.H{"count": count})
}
