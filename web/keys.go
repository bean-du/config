package web

import (
	"coco-tool/config/model/entity"
	"coco-tool/config/service"
	"coco-tool/config/utils"
	"github.com/gin-gonic/gin"
)

var keys = &keysCtl{}

type keysCtl struct{}

func (k *keysCtl) init() {
	router.GET("/keys", k.keys)
	router.POST("/keyDetails", k.keyDetails)
	router.POST("/set", k.set)
	router.POST("/get", k.get)
	router.POST("/apply", k.apply)
	router.POST("/del", k.del)
}

func (k *keysCtl) keys(c *gin.Context) {
	strings, err := service.KeysSrv.Keys(c.Request.Context())
	if err != nil {
		c.JSON(200, utils.HttpResponse(1, "internal error "+err.Error(), nil))
		return
	}
	c.JSON(200, utils.HttpResponse(0, "success", strings))
}

func (k *keysCtl) keyDetails(c *gin.Context) {
	record := new(entity.Record)
	if err := c.ShouldBindJSON(record); err != nil {
		c.JSON(200, utils.HttpResponse(1, "params error "+err.Error(), nil))
		return
	}
	if record.Key == "" {
		c.JSON(200, utils.HttpResponse(1, "params error", nil))
		return
	}
	details, err := service.KeysSrv.KeyDetails(c.Request.Context(), record.Key)
	if err != nil {
		c.JSON(200, utils.HttpResponse(1, "internal error "+err.Error(), nil))
		return
	}
	c.JSON(200, utils.HttpResponse(0, "success", details))
}

func (k *keysCtl) set(c *gin.Context) {
	record := new(entity.Record)
	if err := c.ShouldBindJSON(record); err != nil {
		c.JSON(200, utils.HttpResponse(1, "params error "+err.Error(), nil))
		return
	}
	if record.Key == "" || record.Value == "" {
		c.JSON(200, utils.HttpResponse(1, "params error", nil))
		return
	}
	if err := service.KeysSrv.Set(c.Request.Context(), record.Key, record.Value); err != nil {
		c.JSON(200, utils.HttpResponse(1, "internal error "+err.Error(), nil))
		return
	}
	c.JSON(200, utils.HttpResponse(0, "ok", nil))
}

func (k *keysCtl) apply(c *gin.Context) {
	record := new(entity.Record)
	if err := c.ShouldBindJSON(record); err != nil {
		c.JSON(200, utils.HttpResponse(1, "params error "+err.Error(), nil))
		return
	}
	if record.Key == "" || record.Version == "" || record.Value == "" {
		c.JSON(200, utils.HttpResponse(1, "params error ", nil))
		return
	}
	if err := service.KeysSrv.Apply(c.Request.Context(), record.Key, record.Version, record.Value); err != nil {
		c.JSON(200, utils.HttpResponse(1, "internal error "+err.Error(), nil))
		return
	}
	c.JSON(200, utils.HttpResponse(0, "ok", nil))
}

func (k *keysCtl) del(c *gin.Context) {
	record := new(entity.Record)
	if err := c.ShouldBindJSON(record); err != nil {
		c.JSON(200, utils.HttpResponse(1, "params error "+err.Error(), nil))
		return
	}
	if record.Key == "" {
		c.JSON(200, utils.HttpResponse(1, "params error ", nil))
		return
	}
	if err := service.KeysSrv.Del(c.Request.Context(), record.Key, record.Version); err != nil{
		c.JSON(200, utils.HttpResponse(1, "internal error "+err.Error(), nil))
		return
	}
	c.JSON(200, utils.HttpResponse(0, "ok", nil))
}

func (k *keysCtl) get(c *gin.Context) {
	record := new(entity.Record)
	if err := c.ShouldBindJSON(record); err != nil {
		c.JSON(200, utils.HttpResponse(1, "params error "+err.Error(), nil))
		return
	}
	if record.Key == "" {
		c.JSON(200, utils.HttpResponse(1, "params error", nil))
		return
	}
	res, err := service.KeysSrv.Get(c.Request.Context(), record.Key, record.Version)
	if err != nil {
		c.JSON(200, utils.HttpResponse(1, "internal error "+err.Error(), nil))
		return
	}
	c.JSON(200, utils.HttpResponse(0, "success", res))
}


