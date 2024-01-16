package schema

import (
	"time"

	"github.com/yanzongzhen/magnetu/pkg/util"
)

// Folder permissions for Repository
type Folder struct {
	FolderID       string    `json:"folder_id" gorm:"size:20;primaryKey"`                                                  // Unique ID
	FolderName     string    `json:"folder_name" gorm:"size:255;index"`                                                    // FolderName
	ParentFolderID string    `json:"parent_folder_id" gorm:"size:20;index;foreignKey:ParentFolderID;references:Folder.ID"` // From Folder.ID
	RepositoryID   string    `json:"repository_id" gorm:"size:20;foreignKey:RepositoryID;references:Repository.ID"`        // From Repository.ID
	CreatedAt      time.Time `json:"created_at" gorm:"index;"`                                                             // Create time
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
	FolderName     string `form:"folder_name" binding:"required"`      // FolderName
	ParentFolderID string `form:"parent_folder_id" binding:"required"` // ParentFolderID
	RepositoryID   string `form:"repository_id" binding:"required"`    // RepositoryID
}

// Validate A validation function for the `FolderForm` struct.
func (a *FolderForm) Validate() error {
	return nil
}

// FillTo Convert `FolderForm` to `Folder` object.
func (a *FolderForm) FillTo(folder *Folder) error {
	folder.FolderName = a.FolderName
	folder.ParentFolderID = a.ParentFolderID
	folder.RepositoryID = a.RepositoryID
	return nil
}
