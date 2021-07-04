package application

import (
	"github.com/Nistagram-Organization/nistagram-administration/src/clients/auth_grpc_client"
	"github.com/Nistagram-Organization/nistagram-administration/src/clients/post_grpc_client"
	administration2 "github.com/Nistagram-Organization/nistagram-administration/src/controllers/administration"
	"github.com/Nistagram-Organization/nistagram-administration/src/services/administration"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	router.Use(cors.New(corsConfig))

	authGrpcClient := auth_grpc_client.NewAuthGrpcClient()
	postGrpcClient := post_grpc_client.NewPostGrpcClient()
	administrationService := administration.NewAdministrationService(authGrpcClient, postGrpcClient)
	administrationController := administration2.NewAdministrationController(administrationService)

	router.POST("/administration/content", administrationController.DecideOnPost)

	router.Run(":8088")
}
