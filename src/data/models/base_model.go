package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id int `gorm:"primary_key"`

	CreatedAt time.Time `gorm:"type:Timestamp with time zone;not null"`
	ModifiedAt sql.NullTime `gorm:"type:Timestamp with time zone;null"`
	DeletedAt sql.NullTime `gorm:"type:Timestamp with time zone;null"`

	CreatedBy int `gorm:"not null"`
	ModifiedBy *sql.NullInt64 `gorm:"null"`
	DeletedBy  *sql.NullInt64  `gorm:"null"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	userId := -1
	if value != nil{
		userId = int(value.(float64))
	}
	m.CreatedAt = time.Now().UTC()
	m.CreatedBy = userId
	return
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	userId := &sql.NullInt64{Valid: false}
	if value != nil{
		userId = &sql.NullInt64{Valid: true, Int64: int64(value.(float64))}
	}
	m.ModifiedAt = sql.NullTime{Valid: true, Time: time.Now().UTC()}
	m.ModifiedBy = userId
	return
}

func (m *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	userId := &sql.NullInt64{Valid: false}
	if value != nil{
		userId = &sql.NullInt64{Valid: true, Int64: int64(value.(float64))}
	}
	m.DeletedAt = sql.NullTime{Valid: true, Time: time.Now().UTC()}
	m.DeletedBy = userId
	return
}

