package dbModels

import "gorm.io/gorm"

type Joke struct {
	gorm.Model
	Text string `gorm:"not null"`
}
