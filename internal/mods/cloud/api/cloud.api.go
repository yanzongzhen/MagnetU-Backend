package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yanzongzhen/magnetu/internal/mods/cloud/biz"
	"github.com/yanzongzhen/magnetu/internal/mods/cloud/schema"
	"github.com/yanzongzhen/magnetu/pkg/util"
)

// Cloud file management
type Cloud struct {
	CloudBIZ *biz.Cloud
}

// Query
// @Tags CloudAPI
// @Security ApiKeyAuth
// @Summary Query cloud list
// @Param current query int true "pagination index" default(1)
// @Param pageSize query int true "pagination size" default(10)
// @Success 200 {object} util.ResponseResult{data=[]schema.Cloud}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/clouds [get]
func (a *Cloud) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.CloudQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.CloudBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// Get
// @Tags CloudAPI
// @Security ApiKeyAuth
// @Summary Get cloud record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Cloud}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/clouds/{id} [get]
func (a *Cloud) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.CloudBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// Create
// @Tags CloudAPI
// @Security ApiKeyAuth
// @Summary Create cloud record
// @Param body body schema.CloudForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Cloud}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/clouds [post]
func (a *Cloud) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.CloudForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.CloudBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// Update
// @Tags CloudAPI
// @Security ApiKeyAuth
// @Summary Update cloud record by ID
// @Param id path string true "unique id"
// @Param body body schema.CloudForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/clouds/{id} [put]
func (a *Cloud) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.CloudForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.CloudBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// Delete
// @Tags CloudAPI
// @Security ApiKeyAuth
// @Summary Delete cloud record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/clouds/{id} [delete]
func (a *Cloud) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.CloudBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
