package utils

import (
    "developer-allocation-system/pkg/models"
    "fmt"
    "time"
)

func GenerateCacheKey(prefix string, id int) string {
    return fmt.Sprintf("%s_%d", prefix, id)
}

func EstimateCompletionTime(task models.Task, developer models.Developer) time.Time {
    // Simple estimation based on current time and task estimation
    return time.Now().Add(time.Duration(task.Estimation) * time.Hour)
}
