package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/biz"
	"github.com/yanzongzhen/magnetu/internal/mods/repository/schema"
	"github.com/yanzongzhen/magnetu/pkg/util"
)

// Repository management for Repository
type Repository struct {
	RepositoryBIZ *biz.Repository
}

// Query
// @Tags RepositoryAPI
// @Security ApiKeyAuth
// @Summary Query repository list
// @Param current query int true "pagination index" default(1)
// @Param pageSize query int true "pagination size" default(10)
// @Param user_id query string false "User ID"
// @Success 200 {object} util.ResponseResult{data=[]schema.Repository}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/repositories [get]
func (a *Repository) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.RepositoryQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.RepositoryBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// Get
// @Tags RepositoryAPI
// @Security ApiKeyAuth
// @Summary Get repository record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Repository}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/repositories/{id} [get]
func (a *Repository) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.RepositoryBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// Create
// @Tags RepositoryAPI
// @Security ApiKeyAuth
// @Summary Create repository record
// @Param body body schema.RepositoryForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Repository}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/repositories [post]
func (a *Repository) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.RepositoryForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.RepositoryBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// Update
// @Tags RepositoryAPI
// @Security ApiKeyAuth
// @Summary Update repository record by ID
// @Param id path string true "unique id"
// @Param body body schema.RepositoryForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/repositories/{id} [put]
func (a *Repository) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.RepositoryForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.RepositoryBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// Delete
// @Tags RepositoryAPI
// @Security ApiKeyAuth
// @Summary Delete repository record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/repositories/{id} [delete]
func (a *Repository) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.RepositoryBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
