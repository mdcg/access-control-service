package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mdcg/access-control-service/config"
	_gin "github.com/mdcg/access-control-service/internal/http/gin"
	"github.com/mdcg/access-control-service/internal/logging"
	"github.com/mdcg/access-control-service/internal/mongo"
	"github.com/mdcg/access-control-service/restriction"
	mongo_repo "github.com/mdcg/access-control-service/restriction/mongo"
	"github.com/spf13/cobra"
)

func apiCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "Initialize API RESTful",
		Run: func(cmd *cobra.Command, args []string) {
			envVars := config.LoadEnvVars()
			ctx := context.Background()

			logger, err := logging.InitLog(ctx, envVars.OtelServiceName, map[string]string{
				"env":       envVars.Env,
				"component": "api",
			})
			if err != nil {
				log.Fatalf("error! %s", err)
			}

			logger.Info("🚀 Enviando log de teste para Loki")
			defer func() {
				if err := logging.Shutdown(ctx); err != nil {
					logger.Error("Shutdown do logger falhou", "err", err)
				}
			}()

			db := mongo.ConnectDB(envVars.MongoDBURI, envVars.MongoDBDatabase)
			repo := mongo_repo.NewRestrictionStore(ctx, db)

			restrictionService := restriction.NewService(repo)

			h := _gin.Handlers(restrictionService)
			h.GET("/health", func(c *gin.Context) {
				logger.Info("Health check recebido")
				c.JSON(200, gin.H{"ok": true})
			})

			h.Run(":8080")
		},
	}
}

func main() {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(apiCommand())
	rootCmd.Execute()
}
