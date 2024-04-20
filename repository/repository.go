package repository

import (
	"gorm.io/gorm"
)

// PageFormInterface page info
type PageFormInterface interface {
	GetPage() int
	GetSize() int
	GetQuery() string
	GetSort() string
	IsIgnoreSort() bool
	GetReverse() bool
}

// Repo repo struct
type Repo struct {
}

// FindOneObjectByID find one
func (r *Repo) FindOneObjectByID(db *gorm.DB, id uint, i interface{}) error {
	return r.FindOneObjectByField(db, "id", id, i)
}

// FindOneObjectByField find one
func (r *Repo) FindOneObjectByField(db *gorm.DB, field string, value interface{}, i interface{}) error {
	if err := db.Where(field+" = ?", value).First(i).Error; err != nil {
		return err
	}
	return nil
}

// Create create
func (r *Repo) Create(db *gorm.DB, i interface{}) error {
	if err := db.
		Set("gorm:association_autoupdate", false).
		Set("gorm:association_autocreate", false).
		Create(i).Error; err != nil {
		return err
	}
	return nil
}

// Update update
func (r *Repo) Update(db *gorm.DB, i interface{}) error {
	if err := db.
		Set("gorm:association_autoupdate", false).
		Set("gorm:association_autocreate", false).
		Save(i).Error; err != nil {
		return err
	}
	return nil
}

// Delete update
func (r *Repo) Delete(db *gorm.DB, i interface{}) error {
	if err := db.
		Set("gorm:association_autoupdate", false).
		Set("gorm:association_autocreate", false).
		Delete(i).Error; err != nil {
		return err
	}
	return nil
}

// FindOneByID find one by id
func (r *Repo) FindOneByID(db *gorm.DB, id uint, i interface{}) error {
	db = db.Set("gorm:auto_preload", true)
	if err := r.FindOneObjectByID(db, id, i); err != nil {
		return err
	}
	return nil
}
