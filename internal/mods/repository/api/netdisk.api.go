package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/biz"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/schema"
	"github.com/yanzongzhen/magnetu/pkg/util"
)

type NetDisk struct {
	NetDiskBIZ *biz.NetDisk
}

// Query
// @Tags NetDiskAPI
// @Security ApiKeyAuth
// @Summary Query NetDisk list
// @Param current query int true "pagination index" default(1)
// @Param pageSize query int true "pagination size" default(10)
// @Param parent_file_id query string false "parent_file_id"
// @Success 200 {object} util.ResponseResult{data=[]schema.File}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/netdisk [get]
func (a *NetDisk) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.FileQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}
	// 设置本人的 RepositoryID
	params.RepositoryID = util.FromUserID(ctx)
	result, err := a.NetDiskBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// Get
// @Tags NetDiskAPI
// @Security ApiKeyAuth
// @Summary Get NetDisk record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.File}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/netdisk/{id} [get]
func (a *NetDisk) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.NetDiskBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// Create
// @Tags NetDiskAPI
// @Security ApiKeyAuth
// @Summary Create NetDisk record
// @Param body body schema.FileForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.File}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/netdisk [post]
func (a *NetDisk) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.FileForm
	if err := util.ParseJSON(c, &item); err != nil {
		util.ResError(c, err)
		return
	}
	item.RepositoryID = util.FromUserID(ctx)

	result, err := a.NetDiskBIZ.Create(ctx, &item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// Update
// @Tags NetDiskAPI
// @Security ApiKeyAuth
// @Summary Update NetDisk record by ID
// @Param id path string true "unique id"
// @Param body body schema.FileForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/netdisk/{id} [put]
func (a *NetDisk) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.FileForm
	if err := util.ParseJSON(c, &item); err != nil {
		util.ResError(c, err)
		return
	}
	// 设置本人的 RepositoryID
	item.RepositoryID = util.FromUserID(ctx)
	err := a.NetDiskBIZ.Update(ctx, c.Param("id"), &item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, nil)
}

// Delete
// @Tags NetDiskAPI
// @Security ApiKeyAuth
// @Summary Delete NetDisk record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/netdisk/{id} [delete]
func (a *NetDisk) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	if err := a.NetDiskBIZ.Delete(ctx, c.Param("id")); err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, nil)
}
