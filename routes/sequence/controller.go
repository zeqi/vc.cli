/**
 * Created by zeqi
 * @description
 * @module SequenceController
 * @version 1.0.0
 * @author Xijun Zhu <zhuzeqi2010@163.com>
 * @File controller
 * @Date 18-7-19
 * @Wechat zhuzeqi2010
 * @QQ 304566647
 * @Office-email zhuxj4@lenovo.com
 */

package sequence

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	sequenceService "vc.cli/services/sequence"
)

type Doc struct {
	Name     string `json:"name" xml:"name" binding:"required"`
	Comments string `json:"comments" xml:"comments"`
}

type Controller struct {
	SequenceService sequenceService.Service
}

func NewController() Controller {
	return Controller{SequenceService: sequenceService.NewService()}
}

func (o *Controller) Create(c *gin.Context) {
	doc := Doc{}
	err := c.ShouldBindWith(&doc, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器内部错误", "error": err})
		return
	}
	result, err := o.SequenceService.Create(doc.Name, doc.Comments)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器内部错误", "error": err})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (o *Controller) Find(c *gin.Context) {
	result, err := o.SequenceService.FInd(0, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器内部错误", "error": err})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (o *Controller) FindOne(c *gin.Context) {
	result, err := o.SequenceService.FIndOne(0, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器内部错误", "error": err})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (o *Controller) FindById(c *gin.Context) {
	result, err := o.SequenceService.FIndById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器内部错误", "error": err})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (o *Controller) IncByName(c *gin.Context) {
	result, err := o.SequenceService.IncByName(c.Param("name"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器内部错误", "error": err})
		return
	}
	c.JSON(http.StatusOK, result)
}
