package main

import (
	"context"

	"github.com/booscaaa/go-signoz-otl/adapter/postgres"
	"github.com/booscaaa/go-signoz-otl/di"
	"github.com/booscaaa/go-signoz-otl/util"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	tracerProvider := util.InitTracer()
	ctx := context.Background()
	conn := postgres.GetConnection(ctx, tracerProvider)
	defer conn.Close()

	postgres.RunMigrations()

	productService := di.ConfigProductDI(conn)

	router := gin.Default()
	router.Use(otelgin.Middleware(util.ServiceName))
	router.GET("/product", productService.Fetch)
	router.Run(":3000")
}
