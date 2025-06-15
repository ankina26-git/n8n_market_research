package analytics

import "gorm.io/gorm"

type Apis struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}
