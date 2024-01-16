package schema

import (
	"time"

	"github.com/yanzongzhen/magnetu/pkg/util"
)

// File permissions for Repository
type File struct {
	FileID         string    `json:"file_id" gorm:"size:20;primaryKey"`                                                    // Unique ID
	FileName       string    `json:"file_name" gorm:"size:255;index"`                                                      // FileName
	FileExtension  string    `json:"file_extension" gorm:"size:255;index"`                                                 // FileExtension
	FileType       string    `json:"file_type" gorm:"size:255;index"`                                                      // FileType
	FileSize       int       `json:"file_size" gorm:"size:20;index"`                                                       // FileSize
	DownloadCount  int       `json:"download_count" gorm:"size:20;index"`                                                  // DownloadCount
	StoragePath    string    `json:"storage_path" gorm:"size:255;index"`                                                   // StoragePath
	CloudPath      string    `json:"cloud_path" gorm:"size:255;index"`                                                     // CloudPath
	ParentFolderID string    `json:"parent_folder_id" gorm:"size:20;index;foreignKey:ParentFolderID;references:Folder.ID"` // From Folder.ID
	RepositoryID   string    `json:"repository_id" gorm:"size:20;foreignKey:RepositoryID;references:Repository.ID"`        // From Repository.ID
	UploadTime     time.Time `json:"upload_time" gorm:"index;"`
}

// FileQueryParam Defining the query parameters for the `File` struct.
type FileQueryParam struct {
	util.PaginationParam
}

// FileQueryOptions Defining the query options for the `File` struct.
type FileQueryOptions struct {
	util.QueryOptions
}

// FileQueryResult Defining the query result for the `File` struct.
type FileQueryResult struct {
	Data       Files
	PageResult *util.PaginationResult
}

// Files Defining the slice of `File` struct.
type Files []*File

// FileForm Defining the data structure for creating a `File` struct.
type FileForm struct {
	FileName       string    `form:"file_name" binding:"required"`        // FileName
	FileExtension  string    `form:"file_extension" binding:"required"`   // FileExtension
	FileType       string    `form:"file_type" binding:"required"`        // FileType
	FileSize       int       `form:"file_size" binding:"required"`        // FileSize
	DownloadCount  int       `form:"download_count"`                      // DownloadCount
	StoragePath    string    `form:"storage_path" binding:"required"`     // StoragePath
	CloudPath      string    `form:"cloud_path"`                          // CloudPath
	ParentFolderID string    `form:"parent_folder_id" binding:"required"` // ParentFolderID
	RepositoryID   string    `form:"repository_id" binding:"required"`    // RepositoryID
	UploadTime     time.Time `form:"upload_time"`                         // Upload Time
}

// Validate A validation function for the `FileForm` struct.
func (a *FileForm) Validate() error {
	return nil
}

// FillTo Convert `FileForm` to `File` object.
func (a *FileForm) FillTo(file *File) error {
	file.FileName = a.FileName
	file.FileExtension = a.FileExtension
	file.FileType = a.FileType
	file.FileSize = a.FileSize
	file.DownloadCount = a.DownloadCount
	file.StoragePath = a.StoragePath
	file.CloudPath = a.CloudPath
	file.ParentFolderID = a.ParentFolderID
	file.RepositoryID = a.RepositoryID
	file.UploadTime = a.UploadTime
	return nil
}
