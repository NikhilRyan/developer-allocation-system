package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "developer-allocation-system/pkg/models"
    "developer-allocation-system/pkg/services"
    "developer-allocation-system/pkg/utils"
)

// DeveloperHandler handles developer-related HTTP requests.
type DeveloperHandler struct {
    DeveloperService services.DeveloperService
}

// NewDeveloperHandler creates a new DeveloperHandler.
func NewDeveloperHandler(devService services.DeveloperService) *DeveloperHandler {
    return &DeveloperHandler{
        DeveloperService: devService,
    }
}

// GetDevelopers handles GET requests to retrieve all developers.
func (h *DeveloperHandler) GetDevelopers(c *gin.Context) {
    developers, err := h.DeveloperService.GetAllDevelopers()
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to retrieve developers")
        return
    }
    utils.RespondWithSuccess(c, http.StatusOK, developers)
}

// GetDeveloperByID handles GET requests to retrieve a developer by ID.
func (h *DeveloperHandler) GetDeveloperByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid developer ID")
        return
    }

    developer, err := h.DeveloperService.GetDeveloperByID(id)
    if err != nil {
        utils.RespondWithError(c, http.StatusNotFound, "Developer not found")
        return
    }
    utils.RespondWithSuccess(c, http.StatusOK, developer)
}

// CreateDeveloper handles POST requests to create a new developer.
func (h *DeveloperHandler) CreateDeveloper(c *gin.Context) {
    var dev models.Developer
    if err := c.ShouldBindJSON(&dev); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid input data")
        return
    }

    err := h.DeveloperService.CreateDeveloper(&dev)
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create developer")
        return
    }
    utils.RespondWithSuccess(c, http.StatusCreated, dev)
}

// UpdateDeveloper handles PUT requests to update an existing developer.
func (h *DeveloperHandler) UpdateDeveloper(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid developer ID")
        return
    }

    var dev models.Developer
    if err := c.ShouldBindJSON(&dev); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid input data")
        return
    }
    dev.ID = id

    err = h.DeveloperService.UpdateDeveloper(&dev)
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update developer")
        return
    }
    utils.RespondWithSuccess(c, http.StatusOK, dev)
}

// DeleteDeveloper handles DELETE requests to remove a developer.
func (h *DeveloperHandler) DeleteDeveloper(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid developer ID")
        return
    }

    err = h.DeveloperService.DeleteDeveloper(id)
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to delete developer")
        return
    }
    utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Developer deleted successfully"})
}

// UpdateAvailability handles PATCH requests to update a developer's availability.
func (h *DeveloperHandler) UpdateAvailability(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid developer ID")
        return
    }

    var availabilityUpdate struct {
        Availability float64 `json:"availability" binding:"required"`
    }
    if err := c.ShouldBindJSON(&availabilityUpdate); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid input data")
        return
    }

    err = h.DeveloperService.UpdateAvailability(id, availabilityUpdate.Availability)
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update availability")
        return
    }
    utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Availability updated successfully"})
}

func (h *DeveloperHandler) GetDeveloperRecommendations(c *gin.Context) {
    taskIDParam := c.Param("taskID")
    taskID, err := strconv.Atoi(taskIDParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

    recommendations, err := h.DeveloperService.GetDeveloperRecommendations(taskID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, recommendations)
}
