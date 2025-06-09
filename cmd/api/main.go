package main

import (
	"context"

	"github.com/mdcg/access-control-service/config"
	"github.com/mdcg/access-control-service/internal/http/gin"
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

			db := mongo.ConnectDB(envVars.MongoDBURI, envVars.MongoDBDatabase)
			repo := mongo_repo.NewRestrictionStore(ctx, db)

			restrictionService := restriction.NewService(repo)

			h := gin.Handlers(restrictionService)

			h.Run(":8080")
		},
	}
}

func main() {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(apiCommand())
	rootCmd.Execute()
}
