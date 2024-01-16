package schema

import (
	"time"

	"github.com/yanzongzhen/magnetu/pkg/util"
)

// Repository management for Repository
type Repository struct {
	ID              string    `json:"id" gorm:"size:20;primaryKey;"`         // Unique ID
	UserID          string    `json:"user_id" gorm:"size:20;index"`          // From User.ID
	CurrentCapacity int       `json:"current_capacity" gorm:"size:20;index"` // Current capacity
	MaxCapacity     int       `json:"max_capacity" gorm:"size:20;index"`     // Max capacity
	Permissions     string    `json:"permissions" gorm:"size:255"`           // Permissions
	CreatedAt       time.Time `json:"created_at" gorm:"index;"`              // Create time
	UpdatedAt       time.Time `json:"updated_at" gorm:"index;"`              // Update time
}

// RepositoryQueryParam Defining the query parameters for the `Repository` struct.
type RepositoryQueryParam struct {
	util.PaginationParam

	UserID          string `form:"user_id"` // From User.ID
	CurrentCapacity int    `form:"-"`       // Current capacity
	MaxCapacity     int    `form:"-"`       // Max capacity
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
	CurrentCapacity int    `form:"current_capacity"`           // Current capacity
	MaxCapacity     int    `form:"max_capacity"`               // Max capacity
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
