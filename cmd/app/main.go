package main

import (
    "os"

    "developer-allocation-system/pkg/api/routers"
    "developer-allocation-system/pkg/db"
    "developer-allocation-system/pkg/services"
    "developer-allocation-system/pkg/utils"
    "developer-allocation-system/pkg/cache"
)

func main() {
    // Load configuration
    config := utils.LoadConfig()
    // if err != nil {
    //     log.Fatalf("Failed to load configuration: %v", err)
    // }

    // Initialize logger
    utils.SetupLogger(config.LogLevel)

    // Initialize database connection
    dbClient, err := db.NewDatabase(config.Database)
    if err != nil {
        utils.GetLogger().Fatalf("Failed to connect to database: %v", err)
    }

    // Initialize cache
    cacheClient, err := cache.NewCache(config.Cache)
    if err != nil {
        utils.GetLogger().Fatalf("Failed to connect to cache: %v", err)
    }

    // Initialize repositories
    developerRepo := db.NewDeveloperRepository(dbClient)
    taskRepo := db.NewTaskRepository(dbClient)
    cacheRepo := cache.NewCacheRepository(cacheClient)

    // Initialize services
    developerService := services.NewDeveloperService(
        developerRepo,
        cacheRepo,
        taskRepo,
    )

    taskService := services.NewTaskService(
        taskRepo,
        developerRepo,
        cacheRepo,
    )

    authService := services.NewAuthService(
        db.NewUserRepository(dbClient),
        config.JWTSecret,
    )

    // Setup router
    router := routers.SetupRouter(developerService, taskService, authService)

    // Start the server
    port := os.Getenv("PORT")
    if port == "" {
        port = config.ServerPort
    }
    utils.GetLogger().Infof("Starting server on port %s", port)
    if err := router.Run(":" + port); err != nil {
        utils.GetLogger().Fatalf("Failed to start server: %v", err)
    }
}
