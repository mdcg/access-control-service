package restriction

import (
	"fmt"
	"time"
)

type TimeRange struct {
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
}

type Restriction struct {
	Key   string               `json:"key" binding:"required,min=3"`
	Value string               `json:"value" binding:"required"`
	Rules map[string]TimeRange `json:"rules" binding:"required,min=1"`
}

func (r *Restriction) ValidateDates() error {
	for service, rule := range r.Rules {
		if !rule.StartDate.Before(rule.EndDate) {
			return fmt.Errorf("start_date must be before end_date for service '%s'", service)
		}
	}
	return nil
}

type RestrictionRepository interface {
	Create(r *Restriction) error
}

type UseCase interface {
	CreateRestriction(restriction *Restriction) error
}
