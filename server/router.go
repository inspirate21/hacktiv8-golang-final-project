package server

import (
	"hacktiv8-golang-final-project/controller/controllercomment"
	"hacktiv8-golang-final-project/controller/controllerphoto"
	"hacktiv8-golang-final-project/controller/controllersocialmedia"
	"hacktiv8-golang-final-project/controller/controlleruser"
	_ "hacktiv8-golang-final-project/docs"
	"hacktiv8-golang-final-project/middleware"
	"hacktiv8-golang-final-project/repository/repositorycomment"
	"hacktiv8-golang-final-project/repository/repositoryphoto"
	"hacktiv8-golang-final-project/repository/repositorysocialmedia"
	"hacktiv8-golang-final-project/repository/repositoryuser"
	"hacktiv8-golang-final-project/service/servicecomment"
	"hacktiv8-golang-final-project/service/servicephoto"
	"hacktiv8-golang-final-project/service/servicesocialmedia"
	"hacktiv8-golang-final-project/service/serviceuser"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {

	repoUser := repositoryuser.New(db)
	srvUser := serviceuser.New(repoUser)
	ctrlUser := controlleruser.New(srvUser)

	routeUser := r.Group("/users")

	// route user
	routeUser.POST("/register", ctrlUser.Create)
	routeUser.POST("/login", ctrlUser.Login)
	routeUser.PUT("", middleware.Authorization, ctrlUser.Update)
	routeUser.DELETE("", middleware.Authorization, ctrlUser.DeleteByID)

	// route photos
	repoPhoto := repositoryphoto.New(db)
	srvPhoto := servicephoto.New(repoPhoto)
	ctrlPhoto := controllerphoto.New(srvPhoto)

	r.GET("/", ctrlPhoto.GetPhotos)

	r.GET("photos", middleware.Authorization, ctrlPhoto.GetPhotos)
	r.POST("photos", middleware.Authorization, ctrlPhoto.Create)
	r.PUT("photos/:photoID", middleware.Authorization, ctrlPhoto.Update)
	r.DELETE("photos/:photoID", middleware.Authorization, ctrlPhoto.Delete)

	// route comment
	repoComment := repositorycomment.New(db)
	srvComment := servicecomment.New(repoComment)
	ctrlComment := controllercomment.New(srvComment)
	r.GET("comments", middleware.Authorization, ctrlComment.Get)
	r.POST("comments", middleware.Authorization, ctrlComment.Create)
	r.PUT("comments/:commentID", middleware.Authorization, ctrlComment.Update)
	r.DELETE("comments/:commentID", middleware.Authorization, ctrlComment.Delete)

	// route social media
	repoSocialmedia := repositorysocialmedia.New(db)
	srvSocialmedia := servicesocialmedia.New(repoSocialmedia, repoPhoto)
	ctrlSocialmedia := controllersocialmedia.New(srvSocialmedia)
	routerSocialmedia := r.Group("/socialmedias")
	routerSocialmedia.POST("", middleware.Authorization, ctrlSocialmedia.Create)
	routerSocialmedia.GET("", middleware.Authorization, ctrlSocialmedia.GetList)
	routerSocialmedia.PUT("/:socialmediaid", middleware.Authorization, ctrlSocialmedia.UpdateByID)
	routerSocialmedia.DELETE("/:socialmediaid", middleware.Authorization, ctrlSocialmedia.DeleteByID)

	// routing docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
