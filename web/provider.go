package web

import (
	"coco-tool/config/service"
	"coco-tool/config/utils"
	"github.com/gin-gonic/gin"
)

var providerCtl = &provider{}

type provider struct{}

func (p *provider) init() {
	router.GET("/", p.index)
	router.GET("/currentProvider", p.getCurrent)
	router.GET("/providers", p.providers)
	router.GET("/setProvider", p.setProvider)
}

func (p *provider) getCurrent(c *gin.Context) {
	currentProvider := service.ProviderSrv.GetCurrentProvider()
	c.JSON(200, utils.HttpResponse(0, "success", currentProvider))
}

func (p *provider) providers(c *gin.Context) {
	providers := service.ProviderSrv.GetProviders()
	c.JSON(200, utils.HttpResponse(0, "success", providers))
}

func (p *provider) setProvider(c *gin.Context) {
	service.ProviderSrv.SetProvider(c.Query("name"))
	c.JSON(200, utils.HttpResponse(0, "success", nil))
}

func (p *provider) index(c *gin.Context) {
	c.Writer.WriteString(service.ProviderSrv.Index())
}
