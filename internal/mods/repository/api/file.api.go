package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/biz"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/schema"
	"github.com/yanzongzhen/magnetu/pkg/util"
)

// File management for Repository
type File struct {
	FileBIZ *biz.File
}

// Query
// @Tags FileAPI
// @Security ApiKeyAuth
// @Summary Query File list
// @Param current query int true "pagination index" default(1)
// @Param pageSize query int true "pagination size" default(10)
// @Success 200 {object} util.ResponseResult{data=[]schema.File}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/files [get]
func (a *File) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.FileQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.FileBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// Get
// @Tags FileAPI
// @Security ApiKeyAuth
// @Summary Get File record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.File}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/files/{id} [get]
func (a *File) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.FileBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// Create
// @Tags FileAPI
// @Security ApiKeyAuth
// @Summary Create File record
// @Param body body schema.FileForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.File}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/files [post]
func (a *File) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.FileForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.FileBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// Update
// @Tags FileAPI
// @Security ApiKeyAuth
// @Summary Update File record by ID
// @Param id path string true "unique id"
// @Param body body schema.FileForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/files/{id} [put]
func (a *File) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.FileForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.FileBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// Delete
// @Tags FileAPI
// @Security ApiKeyAuth
// @Summary Delete File record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/files/{id} [delete]
func (a *File) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.FileBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
