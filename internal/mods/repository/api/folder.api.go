package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/biz"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/schema"
	"github.com/yanzongzhen/magnetu/pkg/util"
)

// Folder management for Repository
type Folder struct {
	FolderBIZ *biz.Folder
}

// Query
// @Tags FolderAPI
// @Security ApiKeyAuth
// @Summary Query Folder list
// @Param current query int true "pagination index" default(1)
// @Param pageSize query int true "pagination size" default(10)
// @Success 200 {object} util.ResponseResult{data=[]schema.Folder}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/folders [get]
func (a *Folder) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.FolderQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.FolderBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// Get
// @Tags FolderAPI
// @Security ApiKeyAuth
// @Summary Get Folder record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Folder}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/folders/{id} [get]
func (a *Folder) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.FolderBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// Create
// @Tags FolderAPI
// @Security ApiKeyAuth
// @Summary Create Folder record
// @Param body body schema.FolderForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Folder}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/folders [post]
func (a *Folder) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.FolderForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.FolderBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// Update
// @Tags FolderAPI
// @Security ApiKeyAuth
// @Summary Update Folder record by ID
// @Param id path string true "unique id"
// @Param body body schema.FolderForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/folders/{id} [put]
func (a *Folder) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.FolderForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.FolderBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// Delete
// @Tags FolderAPI
// @Security ApiKeyAuth
// @Summary Delete Folder record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/folders/{id} [delete]
func (a *Folder) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.FolderBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
