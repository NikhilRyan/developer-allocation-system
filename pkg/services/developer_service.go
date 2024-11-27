package services

import (
    "errors"
    "sort"

    "developer-allocation-system/pkg/models"
    "developer-allocation-system/pkg/repositories"
    "developer-allocation-system/pkg/utils"
)

type DeveloperService interface {
    GetAllDevelopers() ([]models.Developer, error)
    GetDeveloperByID(id int) (*models.Developer, error)
    CreateDeveloper(dev *models.Developer) error
    UpdateDeveloper(dev *models.Developer) error
    DeleteDeveloper(id int) error
    UpdateAvailability(id int, availability float64) error
    GetDeveloperRecommendations(taskID int) ([]models.DeveloperRecommendation, error)
}

type developerService struct {
    repo     repositories.DeveloperRepository
    cache    repositories.CacheRepository
    taskRepo repositories.TaskRepository
}

func NewDeveloperService(repo repositories.DeveloperRepository, cache repositories.CacheRepository, taskRepo repositories.TaskRepository) DeveloperService {
    return &developerService{
        repo:     repo,
        cache:    cache,
        taskRepo: taskRepo,
    }
}

func (s *developerService) GetAllDevelopers() ([]models.Developer, error) {
    // Try to get from cache
    var developers []models.Developer
    err := s.cache.Get("all_developers", &developers)
    if err == nil {
        return developers, nil
    }

    // Fetch from database
    developers, err = s.repo.GetAll()
    if err != nil {
        return nil, err
    }

    // Set cache
    s.cache.Set("all_developers", developers)

    return developers, nil
}

func (s *developerService) GetDeveloperByID(id int) (*models.Developer, error) {
    // Try to get from cache
    var developer models.Developer
    cacheKey := utils.GenerateCacheKey("developer", id)
    err := s.cache.Get(cacheKey, &developer)
    if err == nil {
        return &developer, nil
    }

    // Fetch from database
    developerPtr, err := s.repo.GetByID(id)
    if err != nil {
        return nil, err
    }

    // Set cache
    s.cache.Set(cacheKey, developerPtr)

    return developerPtr, nil
}

func (s *developerService) CreateDeveloper(dev *models.Developer) error {
    err := s.repo.Create(dev)
    if err != nil {
        return err
    }

    // Invalidate cache
    s.cache.Delete("all_developers")

    return nil
}

func (s *developerService) UpdateDeveloper(dev *models.Developer) error {
    err := s.repo.Update(dev)
    if err != nil {
        return err
    }

    // Invalidate cache
    s.cache.Delete("all_developers")
    cacheKey := utils.GenerateCacheKey("developer", dev.ID)
    s.cache.Delete(cacheKey)

    return nil
}

func (s *developerService) DeleteDeveloper(id int) error {
    err := s.repo.Delete(id)
    if err != nil {
        return err
    }

    // Invalidate cache
    s.cache.Delete("all_developers")
    cacheKey := utils.GenerateCacheKey("developer", id)
    s.cache.Delete(cacheKey)

    return nil
}

func (s *developerService) UpdateAvailability(id int, availability float64) error {
    err := s.repo.UpdateAvailability(id, availability)
    if err != nil {
        return err
    }

    // Invalidate cache
    cacheKey := utils.GenerateCacheKey("developer", id)
    s.cache.Delete(cacheKey)

    return nil
}

func (s *developerService) GetDeveloperRecommendations(taskID int) ([]models.DeveloperRecommendation, error) {
    // Fetch the task
    task, err := s.taskRepo.GetByID(taskID)
    if err != nil {
        return nil, errors.New("task not found")
    }

    // Fetch all developers
    developers, err := s.repo.GetAll()
    if err != nil {
        return nil, errors.New("failed to fetch developers")
    }

    // Calculate match scores
    var recommendations []models.DeveloperRecommendation
    for _, dev := range developers {
        if !dev.IsAvailable {
            continue
        }

        score := utils.CalculateMatchScore(dev, *task)
        recommendations = append(recommendations, models.DeveloperRecommendation{
            Developer: dev,
            Score:     score,
        })
    }

    // Sort recommendations by score descending
    sort.SliceStable(recommendations, func(i, j int) bool {
        return recommendations[i].Score > recommendations[j].Score
    })

    return recommendations, nil
}

