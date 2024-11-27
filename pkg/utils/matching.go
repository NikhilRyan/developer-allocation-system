package utils

import (
	"developer-allocation-system/pkg/models"
	"math"
)

func CalculateMatchScore(dev models.Developer, task models.Task) float64 {
	score := 0.0
	weightSkillMatch := 0.5
	weightAvailability := 0.3
	weightOtherFactors := 0.2

	// Skill Match Score
	totalSkillScore := 0.0
	skillCount := 0
	for skill, requiredLevel := range task.RequiredSkills {
		if devLevel, ok := dev.SkillLevel[skill]; ok {
			skillScore := math.Min(float64(devLevel)/float64(requiredLevel), 1.0)
			totalSkillScore += skillScore
		}
		skillCount++
	}
	if skillCount > 0 {
		totalSkillScore = (totalSkillScore / float64(skillCount)) * 100
	}

	// Availability Score (Assuming 40 hours/week is full availability)
	availabilityScore := math.Min(dev.Availability/40.0, 1.0) * 100

	// Other Factors Score (e.g., SystemKnowledge)
	otherFactorsScore := dev.SystemKnowledge // Assuming it's a percentage from 0 to 100

	// Weighted Total Score
	score = (totalSkillScore * weightSkillMatch) +
		(availabilityScore * weightAvailability) +
		(otherFactorsScore * weightOtherFactors)

	// Normalize to a score between 0 and 100
	if score > 100 {
		score = 100
	}

	return score
}
