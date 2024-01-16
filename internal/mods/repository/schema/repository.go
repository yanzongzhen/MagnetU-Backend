package schema

import (
	"github.com/yanzongzhen/magnetu/pkg/util"
	"time"
)

// Repository management for Repository
type Repository struct {
	RepositoryID    string `gorm:"primaryKey"`
	UserID          string // 外键，关联到用户表（如果有）
	CurrentCapacity int64
	MaxCapacity     int64
	Permissions     string
	Folders         []Folder  `gorm:"foreignKey:RepositoryID"` // 与 Folder 的关联
	Files           []File    `gorm:"foreignKey:RepositoryID"` // 与 File 的关联
	CreatedAt       time.Time `json:"created_at"`              // Create time
	UpdatedAt       time.Time `json:"updated_at"`
}

// RepositoryQueryParam Defining the query parameters for the `Repository` struct.
type RepositoryQueryParam struct {
	util.PaginationParam

	UserID          string `form:"user_id"` // From User.ID
	CurrentCapacity int64  `form:"-"`       // Current capacity
	MaxCapacity     int64  `form:"-"`       // Max capacity
	Permissions     string `form:"-"`       // Permissions
}

// RepositoryQueryOptions Defining the query options for the `Repository` struct.
type RepositoryQueryOptions struct {
	util.QueryOptions
}

// RepositoryQueryResult Defining the query result for the `Repository` struct.
type RepositoryQueryResult struct {
	Data       Repositories
	PageResult *util.PaginationResult
}

// Repositories Defining the slice of `Repository` struct.
type Repositories []*Repository

// RepositoryForm Defining the data structure for creating a `Repository` struct.
type RepositoryForm struct {
	UserID          string `form:"user_id" binding:"required"` // From User.ID
	CurrentCapacity int64  `form:"current_capacity"`           // Current capacity
	MaxCapacity     int64  `form:"max_capacity"`               // Max capacity
	Permissions     string `form:"permissions"`                // Permissions
}

// Validate A validation function for the `RepositoryForm` struct.
func (a *RepositoryForm) Validate() error {
	return nil
}

// FillTo Convert `RepositoryForm` to `Repository` object.
func (a *RepositoryForm) FillTo(repository *Repository) error {
	repository.UserID = a.UserID
	repository.CurrentCapacity = a.CurrentCapacity
	repository.MaxCapacity = a.MaxCapacity
	repository.Permissions = a.Permissions
	return nil
}
