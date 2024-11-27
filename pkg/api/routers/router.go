package routers

import (
    "github.com/gin-gonic/gin"
    "developer-allocation-system/pkg/api/handlers"
    "developer-allocation-system/pkg/api/middleware"
    "developer-allocation-system/pkg/services"
)

// SetupRouter configures the routes and middleware.
func SetupRouter(
    devService services.DeveloperService,
    taskService services.TaskService,
    authService services.AuthService,
) *gin.Engine {
    router := gin.Default()

    // Initialize handlers
    devHandler := handlers.NewDeveloperHandler(devService)
    taskHandler := handlers.NewTaskHandler(taskService)
    authHandler := handlers.NewAuthHandler(authService)

    // Apply middleware
    router.Use(middleware.LoggingMiddleware())

    // Public routes
    authRoutes := router.Group("/api/v1/auth")
    {
        authRoutes.POST("/register", authHandler.Register)
        authRoutes.POST("/login", authHandler.Login)
        authRoutes.POST("/refresh", authHandler.RefreshToken)
    }

    // Protected routes
    apiRoutes := router.Group("/api/v1")
    apiRoutes.Use(middleware.AuthMiddleware(authService))
    {
        // Developer routes
        devRoutes := apiRoutes.Group("/developers")
        {
            devRoutes.GET("/", devHandler.GetDevelopers)
            devRoutes.GET("/:id", devHandler.GetDeveloperByID)
            devRoutes.POST("/", devHandler.CreateDeveloper)
            devRoutes.PUT("/:id", devHandler.UpdateDeveloper)
            devRoutes.DELETE("/:id", devHandler.DeleteDeveloper)
            devRoutes.PATCH("/:id/availability", devHandler.UpdateAvailability)
            devRoutes.GET("/recommendations/:taskID", devHandler.GetDeveloperRecommendations)
        }

        // Task routes
        taskRoutes := apiRoutes.Group("/tasks")
        {
            taskRoutes.GET("/", taskHandler.GetTasks)
            taskRoutes.GET("/:id", taskHandler.GetTaskByID)
            taskRoutes.POST("/", taskHandler.CreateTask)
            taskRoutes.PUT("/:id", taskHandler.UpdateTask)
            taskRoutes.DELETE("/:id", taskHandler.DeleteTask)
            taskRoutes.POST("/assign", taskHandler.AssignTask)
            taskRoutes.GET("/predict-spillover", taskHandler.PredictSpillover)
        }
    }

    return router
}
