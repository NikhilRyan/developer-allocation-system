package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "developer-allocation-system/pkg/models"
    "developer-allocation-system/pkg/services"
    "developer-allocation-system/pkg/utils"
)

// TaskHandler handles task-related HTTP requests.
type TaskHandler struct {
    TaskService services.TaskService
}

// NewTaskHandler creates a new TaskHandler.
func NewTaskHandler(taskService services.TaskService) *TaskHandler {
    return &TaskHandler{
        TaskService: taskService,
    }
}

// GetTasks handles GET requests to retrieve all tasks.
func (h *TaskHandler) GetTasks(c *gin.Context) {
    tasks, err := h.TaskService.GetAllTasks()
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to retrieve tasks")
        return
    }
    utils.RespondWithSuccess(c, http.StatusOK, tasks)
}

// GetTaskByID handles GET requests to retrieve a task by ID.
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid task ID")
        return
    }

    task, err := h.TaskService.GetTaskByID(id)
    if err != nil {
        utils.RespondWithError(c, http.StatusNotFound, "Task not found")
        return
    }
    utils.RespondWithSuccess(c, http.StatusOK, task)
}

// CreateTask handles POST requests to create a new task.
func (h *TaskHandler) CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid input data")
        return
    }

    err := h.TaskService.CreateTask(&task)
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create task")
        return
    }
    utils.RespondWithSuccess(c, http.StatusCreated, task)
}

// UpdateTask handles PUT requests to update an existing task.
func (h *TaskHandler) UpdateTask(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid task ID")
        return
    }

    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid input data")
        return
    }
    task.ID = id

    err = h.TaskService.UpdateTask(&task)
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update task")
        return
    }
    utils.RespondWithSuccess(c, http.StatusOK, task)
}

// DeleteTask handles DELETE requests to remove a task.
func (h *TaskHandler) DeleteTask(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid task ID")
        return
    }

    err = h.TaskService.DeleteTask(id)
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to delete task")
        return
    }
    utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// AssignTask handles POST requests to assign a task to a developer.
func (h *TaskHandler) AssignTask(c *gin.Context) {
    var assignment struct {
        TaskID       int `json:"task_id" binding:"required"`
        DeveloperID  int `json:"developer_id" binding:"required"`
    }
    if err := c.ShouldBindJSON(&assignment); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid input data")
        return
    }

    err := h.TaskService.AssignTask(assignment.TaskID, assignment.DeveloperID)
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to assign task")
        return
    }
    utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Task assigned successfully"})
}

// PredictSpillover handles GET requests to predict task spillover.
func (h *TaskHandler) PredictSpillover(c *gin.Context) {
    predictions, err := h.TaskService.PredictSpillover()
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to predict spillover")
        return
    }
    utils.RespondWithSuccess(c, http.StatusOK, predictions)
}
