package schema

import (
	"time"

	"github.com/yanzongzhen/magnetu/pkg/util"
)

// Folder permissions for Repository
type Folder struct {
	FolderID       string `gorm:"primaryKey"`
	FolderName     string
	ParentFolderID *string    `gorm:"index;"`
	RepositoryID   string     `gorm:"index;foreignKey"`
	CreatedAt      time.Time  `json:"created_at"`                // Create time
	UpdateAt       time.Time  `json:"update_at"`                 // Update time
	Repository     Repository `gorm:"foreignKey:RepositoryID"`   // 与 Repository 的关联
	Files          []File     `gorm:"foreignKey:ParentFolderID"` // 与 File 的关联
}

// FolderQueryParam Defining the query parameters for the `Folder` struct.
type FolderQueryParam struct {
	util.PaginationParam

	FolderName string `form:"-"` // FolderName
}

// FolderQueryOptions Defining the query options for the `Folder` struct.
type FolderQueryOptions struct {
	util.QueryOptions
}

// FolderQueryResult Defining the query result for the `Folder` struct.
type FolderQueryResult struct {
	Data       Folders
	PageResult *util.PaginationResult
}

// Folders Defining the slice of `Folder` struct.
type Folders []*Folder

// FolderForm Defining the data structure for creating a `Folder` struct.
type FolderForm struct {
	FolderName string `form:"folder_name" binding:"required"` // FolderName
}

// Validate A validation function for the `FolderForm` struct.
func (a *FolderForm) Validate() error {
	return nil
}

// FillTo Convert `FolderForm` to `Folder` object.
func (a *FolderForm) FillTo(folder *Folder) error {
	folder.FolderName = a.FolderName
	return nil
}
