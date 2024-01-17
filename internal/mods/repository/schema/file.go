package schema

import (
	"time"

	"github.com/yanzongzhen/magnetu/pkg/util"
)

// File permissions for Repository
type File struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	FileName     string    `json:"file_name" gorm:"index;"`
	ParentID     string    `json:"parent_id" gorm:"index;"`
	RepositoryID string    `json:"repository_id" gorm:"index;foreignKey"`
	FileMeta     string    `json:"file_meta" gorm:"size:1024;"`
	FileType     string    `json:"file_type" gorm:"index;"`
	FileSize     int64     `json:"file_size" gorm:"default:0"`
	URL          string    `json:"url"`
	TransURL     string    `json:"trans_url"`
	IsFolder     bool      `json:"is_folder" gorm:"default:false"`
	Files        Files     `json:"files" gorm:"foreignKey:ParentID"` // 与 Folder 的关联
	CreatedAt    time.Time `json:"created_at"`                       // Create time
	UpdatedAt    time.Time `json:"updated_at"`                       // Update time
}

func (f *File) TableName() string {
	return "file"
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

func (a Files) Len() int {
	return len(a)
}

func (a Files) Less(i, j int) bool {
	return a[i].CreatedAt.Unix() > a[j].CreatedAt.Unix()
}

func (a Files) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Files) ToMap() map[string]*File {
	m := make(map[string]*File)
	for _, item := range a {
		m[item.ID] = item
	}
	return m
}

func (a Files) ToTree() Files {
	var list Files
	m := a.ToMap()
	for _, item := range a {
		if item.ParentID == "" {
			list = append(list, item)
			continue
		}
		if parent, ok := m[item.ParentID]; ok {
			if parent.Files == nil {
				children := Files{item}
				parent.Files = children
				continue
			}
			parent.Files = append(parent.Files, item)
		}
	}
	return list
}

// FileForm Defining the data structure for creating a `File` struct.
type FileForm struct {
	FileName     string `json:"file_name"  binding:"required"` // FileName
	FileType     string `json:"file_type"`                     // FileType
	FileSize     int64  `json:"file_size"`                     // FileSize
	URL          string `json:"url"`                           // URL
	TransURL     string `json:"trans_url"`                     // TransURL
	IsFolder     bool   `json:"is_folder"`                     // IsFolder
	ParentID     string `json:"parent_id"`                     // ParentFileID
	RepositoryID string `json:"repository_id"`                 // RepositoryID
}

// Validate A validation function for the `FileForm` struct.
func (a *FileForm) Validate() error {
	return nil
}

// FillTo Convert `FileForm` to `File` object.
func (a *FileForm) FillTo(file *File) error {
	file.FileName = a.FileName
	file.FileType = a.FileType
	file.FileSize = a.FileSize
	file.URL = a.URL
	file.TransURL = a.TransURL
	file.IsFolder = a.IsFolder
	file.ParentID = a.ParentID
	file.RepositoryID = a.RepositoryID
	return nil
}
