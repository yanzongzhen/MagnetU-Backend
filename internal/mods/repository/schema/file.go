package schema

import (
	"time"

	"github.com/yanzongzhen/magnetu/pkg/util"
)

// File permissions for Repository
type File struct {
	FileID        string  `gorm:"primaryKey"`
	FileName      string  `gorm:"index;"`
	ParentFileID  *string `gorm:"index;"`
	RepositoryID  string  `gorm:"index;foreignKey"`
	FileExtension string
	FileType      string
	FileSize      int64
	StoragePath   string
	CloudPath     string
	Remark        string     `gorm:"size:1023;"`
	IsFolder      bool       `gorm:"default:false"`
	Files         Files      `gorm:"foreignKey:ParentFileID"` // 与 Folder 的关联
	Repository    Repository `gorm:"foreignKey:RepositoryID"` // 与 Repository 的关联
	CreateAt      time.Time  `json:"create_at"`               // Create time
	UpdateAt      time.Time  `json:"update_at"`               // Update time
}

// FileQueryParam Defining the query parameters for the `File` struct.
type FileQueryParam struct {
	util.PaginationParam

	RepositoryID string `form:"repository_id"` // From Repository.ID
	ParentFileID string `form:"parent_file_id"`
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
	FileName      string `form:"file_name" binding:"required"`      // FileName
	FileExtension string `form:"file_extension" binding:"required"` // FileExtension
	FileType      string `form:"file_type" binding:"required"`      // FileType
	FileSize      int64  `form:"file_size" binding:"required"`      // FileSize
	DownloadCount int    `form:"download_count"`                    // DownloadCount
	StoragePath   string `form:"storage_path" binding:"required"`   // StoragePath
	CloudPath     string `form:"cloud_path"`                        // CloudPath
	Remark        string `form:"remark"`                            // Remark
	IsFolder      bool   `form:"is_folder"`                         // IsFolder
	ParentFileID  string `form:"parent_file_id"`                    // ParentFileID
	RepositoryID  string `form:"repository_id" binding:"required"`  // RepositoryID
	CreateAt      time.Time
	UpdateAt      time.Time
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
	file.StoragePath = a.StoragePath
	file.CloudPath = a.CloudPath
	file.CreateAt = a.CreateAt
	file.UpdateAt = a.UpdateAt
	file.Remark = a.Remark
	file.IsFolder = a.IsFolder
	file.ParentFileID = &a.ParentFileID
	file.RepositoryID = a.RepositoryID
	return nil
}
