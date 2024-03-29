package schema

import (
	"time"

	"github.com/yanzongzhen/magnetu/pkg/util"
)

// Cloud file management
type Cloud struct {
	ID        string    `json:"id" gorm:"size:20;primaryKey;"`            // Unique ID
	RawID     string    `json:"raw_id" gorm:"size:20;index;"`             // Raw ID
	Name      string    `json:"name" gorm:"size:255;index"`               // Name
	BTIH      string    `json:"btih" gorm:"size:64;"`                     // BTIH
	IDX       string    `json:"idx" gorm:"size:64;"`                      // IDX
	Source    string    `json:"source" gorm:"size:255;"`                  // Source
	URL       string    `json:"url" gorm:"size:255;"`                     // URL
	FileMeta  string    `json:"file_meta" gorm:"size:1024;"`              // Meta
	Size      int64     `json:"size" gorm:"size:20;"`                     // Size
	OSS       string    `json:"oss" gorm:"size:255;"`                     // OSS
	Way       string    `json:"way" gorm:"size:20;oneof=upload transfer"` // Way
	CreatedAt time.Time `json:"created_at" gorm:"index;"`                 // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;"`                 // Update time
}

// CloudQueryParam Defining the query parameters for the `Cloud` struct.
type CloudQueryParam struct {
	util.PaginationParam

	Name string `form:"name"` // Name
}

// CloudQueryOptions Defining the query options for the `Cloud` struct.
type CloudQueryOptions struct {
	util.QueryOptions
}

// CloudQueryResult Defining the query result for the `Cloud` struct.
type CloudQueryResult struct {
	Data       Clouds
	PageResult *util.PaginationResult
}

// Clouds Defining the slice of `Cloud` struct.
type Clouds []*Cloud

// CloudForm Defining the data structure for creating a `Cloud` struct.
type CloudForm struct {
	Name     string `form:"name" binding:"required"`      // Name
	BTIH     string `form:"btih" binding:"required"`      // BTIH
	IDX      string `form:"idx" binding:"required"`       // IDX
	Source   string `form:"source"`                       // Source
	URL      string `form:"url"`                          // URL
	Size     int64  `form:"size" binding:"required"`      // Size
	FileMeta string `form:"file_meta" binding:"required"` // Meta
	OSS      string `form:"oss"`                          // OSS
	RawID    string `form:"raw_id" binding:"required"`    // Raw ID
}

// Validate A validation function for the `CloudForm` struct.
func (a *CloudForm) Validate() error {
	return nil
}

// FillTo Convert `CloudForm` to `Cloud` object.
func (a *CloudForm) FillTo(cloud *Cloud) error {

	cloud.Name = a.Name
	cloud.BTIH = a.BTIH
	cloud.IDX = a.IDX
	cloud.Source = a.Source
	cloud.URL = a.URL
	cloud.Size = a.Size
	cloud.FileMeta = a.FileMeta
	cloud.RawID = a.RawID
	cloud.OSS = a.OSS

	return nil
}
