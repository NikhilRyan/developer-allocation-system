package models

import (
    "time"
)

// Task represents a task to be assigned to developers.
type Task struct {
    ID                    int               `json:"id" gorm:"primaryKey"`
    Title                 string            `json:"title" binding:"required"`
    Description           string            `json:"description"`
    Estimation            float64           `json:"estimation"` // Estimated time in hours
    RequiredSkills        map[string]int    `json:"required_skills" gorm:"-"`
    SkillsJSON            string            `json:"-" gorm:"column:required_skills"`
    RequiredKnowledge     float64           `json:"required_knowledge"`
    Priority              string            `json:"priority"`
    DeliveryDate          time.Time         `json:"delivery_date"`
    Status                string            `json:"status"`
    Dependencies          []int             `json:"dependencies" gorm:"-"`
    DependenciesJSON      string            `json:"-" gorm:"column:dependencies"`
    Stakeholders          []string          `json:"stakeholders" gorm:"-"`
    StakeholdersJSON      string            `json:"-" gorm:"column:stakeholders"`
    RiskAssessment        string            `json:"risk_assessment"`
    AssignedDeveloperID   *int              `json:"assigned_developer_id"`
    AssignedDeveloper     *Developer        `json:"assigned_developer" gorm:"foreignKey:AssignedDeveloperID"`
    CreatedAt             time.Time         `json:"created_at"`
    UpdatedAt             time.Time         `json:"updated_at"`
}

// BeforeSave is a GORM hook that runs before saving the task.
func (t *Task) BeforeSave() (err error) {
    // Serialize RequiredSkills map to JSON
    t.SkillsJSON, err = SerializeMap(t.RequiredSkills)
    if err != nil {
        return err
    }
    // Serialize Dependencies
    t.DependenciesJSON, err = SerializeIntSlice(t.Dependencies)
    if err != nil {
        return err
    }
    // Serialize Stakeholders
    t.StakeholdersJSON, err = SerializeSlice(t.Stakeholders)
    if err != nil {
        return err
    }
    return nil
}

// AfterFind is a GORM hook that runs after retrieving the task.
func (t *Task) AfterFind() (err error) {
    // Deserialize RequiredSkills JSON to map
    err = DeserializeMap(t.SkillsJSON, &t.RequiredSkills)
    if err != nil {
        return err
    }
    // Deserialize Dependencies
    err = DeserializeIntSlice(t.DependenciesJSON, &t.Dependencies)
    if err != nil {
        return err
    }
    // Deserialize Stakeholders
    err = DeserializeSlice(t.StakeholdersJSON, &t.Stakeholders)
    if err != nil {
        return err
    }
    return nil
}
