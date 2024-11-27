package services

import (
    "developer-allocation-system/pkg/models"
    "developer-allocation-system/pkg/repositories"
    "developer-allocation-system/pkg/utils"
    "errors"
)

type TaskService interface {
    GetAllTasks() ([]models.Task, error)
    GetTaskByID(id int) (*models.Task, error)
    CreateTask(task *models.Task) error
    UpdateTask(task *models.Task) error
    DeleteTask(id int) error
    AssignTask(taskID int, developerID int) error
    PredictSpillover() ([]models.Task, error)
}

type taskService struct {
    taskRepo      repositories.TaskRepository
    devRepo       repositories.DeveloperRepository
    cache         repositories.CacheRepository
}

func NewTaskService(taskRepo repositories.TaskRepository, devRepo repositories.DeveloperRepository, cache repositories.CacheRepository) TaskService {
    return &taskService{
        taskRepo: taskRepo,
        devRepo:  devRepo,
        cache:    cache,
    }
}

func (s *taskService) GetAllTasks() ([]models.Task, error) {
    // Try to get from cache
    var tasks []models.Task
    err := s.cache.Get("all_tasks", &tasks)
    if err == nil {
        return tasks, nil
    }

    // Fetch from database
    tasks, err = s.taskRepo.GetAll()
    if err != nil {
        return nil, err
    }

    // Set cache
    s.cache.Set("all_tasks", tasks)

    return tasks, nil
}

func (s *taskService) GetTaskByID(id int) (*models.Task, error) {
    // Try to get from cache
    var task models.Task
    cacheKey := utils.GenerateCacheKey("task", id)
    err := s.cache.Get(cacheKey, &task)
    if err == nil {
        return &task, nil
    }

    // Fetch from database
    taskPtr, err := s.taskRepo.GetByID(id)
    if err != nil {
        return nil, err
    }

    // Set cache
    s.cache.Set(cacheKey, taskPtr)

    return taskPtr, nil
}

func (s *taskService) CreateTask(task *models.Task) error {
    err := s.taskRepo.Create(task)
    if err != nil {
        return err
    }

    // Invalidate cache
    s.cache.Delete("all_tasks")

    return nil
}

func (s *taskService) UpdateTask(task *models.Task) error {
    err := s.taskRepo.Update(task)
    if err != nil {
        return err
    }

    // Invalidate cache
    s.cache.Delete("all_tasks")
    cacheKey := utils.GenerateCacheKey("task", task.ID)
    s.cache.Delete(cacheKey)

    return nil
}

func (s *taskService) DeleteTask(id int) error {
    err := s.taskRepo.Delete(id)
    if err != nil {
        return err
    }

    // Invalidate cache
    s.cache.Delete("all_tasks")
    cacheKey := utils.GenerateCacheKey("task", id)
    s.cache.Delete(cacheKey)

    return nil
}

func (s *taskService) AssignTask(taskID int, developerID int) error {
    task, err := s.taskRepo.GetByID(taskID)
    if err != nil {
        return err
    }

    developer, err := s.devRepo.GetByID(developerID)
    if err != nil {
        return err
    }

    // Check if developer is available
    if !developer.IsAvailable {
        return errors.New("developer is not available")
    }

    // Assign task
    task.AssignedDeveloperID = &developerID
    err = s.taskRepo.Update(task)
    if err != nil {
        return err
    }

    // Update developer's workload and availability
    developer.CurrentWorkload += task.Estimation
    developer.Availability -= task.Estimation
    if developer.Availability <= 0 {
        developer.IsAvailable = false
    }
    err = s.devRepo.Update(developer)
    if err != nil {
        return err
    }

    // Invalidate cache
    s.cache.Delete(utils.GenerateCacheKey("task", taskID))
    s.cache.Delete(utils.GenerateCacheKey("developer", developerID))

    return nil
}

func (s *taskService) PredictSpillover() ([]models.Task, error) {
    // Fetch all tasks and developers
    tasks, err := s.taskRepo.GetAll()
    if err != nil {
        return nil, err
    }
    developers, err := s.devRepo.GetAll()
    if err != nil {
        return nil, err
    }

    var spilloverTasks []models.Task

    for _, task := range tasks {
        if task.Status == "Completed" {
            continue
        }
        // Estimate if task will spill over
        var dev *models.Developer
        if task.AssignedDeveloperID != nil {
            devID := *task.AssignedDeveloperID
            for _, d := range developers {
                if d.ID == devID {
                    dev = &d
                    break
                }
            }
        }
        if dev == nil {
            // No developer assigned, likely to spill over
            spilloverTasks = append(spilloverTasks, task)
            continue
        }
        estimatedCompletion := utils.EstimateCompletionTime(task, *dev)
        if estimatedCompletion.After(task.DeliveryDate) {
            spilloverTasks = append(spilloverTasks, task)
        }
    }

    return spilloverTasks, nil
}
