package repositories

import "github.com/jinzhu/gorm"

// Repository the repository structure
type Repository struct {
	Db *gorm.DB
}

// CreateRepository create the repository
func (r *Repository) CreateRepository(db *gorm.DB) {
	r.Db = db
}
