package model

import "gorm.io/gorm"

type Model struct{}

func New() *Model {
	return &Model{}
}

func (m *Model) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
	)
}

func (m *Model) Seeds(db *gorm.DB) error {
	return nil
}
