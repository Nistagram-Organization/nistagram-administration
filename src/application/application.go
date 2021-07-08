package application

import (
	"github.com/Nistagram-Organization/nistagram-administration/src/clients/auth_grpc_client"
	"github.com/Nistagram-Organization/nistagram-administration/src/clients/post_grpc_client"
	administration2 "github.com/Nistagram-Organization/nistagram-administration/src/controllers/administration"
	"github.com/Nistagram-Organization/nistagram-administration/src/services/administration"
	"github.com/Nistagram-Organization/nistagram-shared/src/utils/jwt_utils"
	"github.com/Nistagram-Organization/nistagram-shared/src/utils/prometheus_handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	router        = gin.Default()
	requestsCount = prometheus_handler.GetHttpRequestsCounter()
	requestsSize  = prometheus_handler.GetHttpRequestsSize()
	uniqueUsers   = prometheus_handler.GetUniqueClients()
)

func configureCORS() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	router.Use(cors.New(corsConfig))
}

func registerPrometheusMiddleware() {
	prometheus.Register(requestsCount)
	prometheus.Register(requestsSize)
	prometheus.Register(uniqueUsers)

	router.Use(prometheus_handler.PrometheusMiddleware(requestsCount, requestsSize, uniqueUsers))
}

func StartApplication() {
	configureCORS()
	registerPrometheusMiddleware()

	authGrpcClient := auth_grpc_client.NewAuthGrpcClient()
	postGrpcClient := post_grpc_client.NewPostGrpcClient()
	administrationService := administration.NewAdministrationService(authGrpcClient, postGrpcClient)
	administrationController := administration2.NewAdministrationController(administrationService)

	router.POST("/administration/content", jwt_utils.GetJwtMiddleware(), jwt_utils.CheckRoles([]string{"admin"}), administrationController.DecideOnPost)

	router.GET("/metrics", prometheus_handler.PrometheusGinHandler())

	router.Run(":8088")
}
