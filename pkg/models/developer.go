package models

import (
    "time"
)

// Developer represents a software developer in the system.
type Developer struct {
    ID                   int               `json:"id" gorm:"primaryKey"`
    Name                 string            `json:"name" binding:"required"`
    Email                string            `json:"email" binding:"required,email" gorm:"unique"`
    Availability         float64           `json:"availability"` // Hours available during the sprint
    StartDate            time.Time         `json:"start_date"`
    SkillLevel           map[string]int    `json:"skill_level" gorm:"-"`
    SkillLevelJSON       string            `json:"-" gorm:"column:skill_level"`
    SystemKnowledge      float64           `json:"system_knowledge"` // Percentage
    Level                string            `json:"level"`
    CurrentWorkload      float64           `json:"current_workload"`
    OnCallRotationWeek   bool              `json:"on_call_rotation_week"`
    Responsibilities     []string          `json:"responsibilities" gorm:"-"`
    ResponsibilitiesJSON string            `json:"-" gorm:"column:responsibilities"`
    IsAvailable          bool              `json:"is_available"`
    TimeZone             string            `json:"time_zone"`
    LanguageProficiency  []string          `json:"language_proficiency" gorm:"-"`
    LanguageJSON         string            `json:"-" gorm:"column:language_proficiency"`
    CreatedAt            time.Time         `json:"created_at"`
    UpdatedAt            time.Time         `json:"updated_at"`
}

type DeveloperRecommendation struct {
    Developer Developer `json:"developer"`
    Score     float64   `json:"score"`
}

// BeforeSave is a GORM hook that runs before saving the developer.
func (d *Developer) BeforeSave() (err error) {
    // Serialize SkillLevel map to JSON
    d.SkillLevelJSON, err = SerializeMap(d.SkillLevel)
    if err != nil {
        return err
    }
    // Serialize Responsibilities
    d.ResponsibilitiesJSON, err = SerializeSlice(d.Responsibilities)
    if err != nil {
        return err
    }
    // Serialize LanguageProficiency
    d.LanguageJSON, err = SerializeSlice(d.LanguageProficiency)
    if err != nil {
        return err
    }
    return nil
}

// AfterFind is a GORM hook that runs after retrieving the developer.
func (d *Developer) AfterFind() (err error) {
    // Deserialize SkillLevel JSON to map
    err = DeserializeMap(d.SkillLevelJSON, &d.SkillLevel)
    if err != nil {
        return err
    }
    // Deserialize Responsibilities
    err = DeserializeSlice(d.ResponsibilitiesJSON, &d.Responsibilities)
    if err != nil {
        return err
    }
    // Deserialize LanguageProficiency
    err = DeserializeSlice(d.LanguageJSON, &d.LanguageProficiency)
    if err != nil {
        return err
    }
    return nil
}
