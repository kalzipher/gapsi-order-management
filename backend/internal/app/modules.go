package app

import "github.com/gin-gonic/gin"

func (a *App) registerPublicModules(api *gin.RouterGroup) {
	a.registerAuthModule(api)
}

func (a *App) registerProtectedModules(protected *gin.RouterGroup) {
	a.registerOrdersModule(protected)
	a.registerStatsModule(protected)
}
