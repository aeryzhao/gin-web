package v1

import (
	"github.com/aeryzhao/gin-web/models"
	"github.com/aeryzhao/gin-web/pkg/e"
	"github.com/aeryzhao/gin-web/pkg/setting"
	"github.com/aeryzhao/gin-web/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

// GetArticle
// @Tags    文章
// @Summary 获取单篇文章
// @Param   id path int true "文章ID"
// @Router  /api/v1/articles/{id} [get]
func GetArticle(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于1")
	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExitArticleById(id) {
			data = models.GetArticleById(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

// GetArticles
// @Tags    文章
// @Summary 获取多篇文章
// @Param   tag_id query int false "标签ID"
// @Router  /api/v1/articles [get]
func GetArticles(ctx *gin.Context) {
	data := make(map[string]interface{})
	params := make(map[string]interface{})
	v := validation.Validation{}

	if arg := ctx.Query("tag_id"); arg != "" {
		tagId := com.StrTo(ctx.Query("tag_id")).MustInt()
		params["tag_id"] = tagId
		v.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}
	code := e.INVALID_PARAMS
	if !v.HasErrors() {
		code = e.SUCCESS
		data["list"] = models.GetArticles(util.GetPage(ctx), setting.PageSize, params)
		data["total"] = models.GetArticleTotal(params)
	} else {
		for _, err := range v.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddArticle
// @Tags    文章
// @Summary 添加文章
// @Param   tag_id     query int    true  "标签 ID"
// @Param   title      query string true  "标题"
// @Param   desc       query string true  "描述"
// @Param   content    query string true  "内容"
// @Param   created_by query string true  "创建人"
// @Param   state      query int    false "状态"
// @Router  /api/v1/articles [post]
func AddArticle(ctx *gin.Context) {
	tagId := com.StrTo(ctx.Query("tag_id")).MustInt()
	title := ctx.Query("title")
	desc := ctx.Query("desc")
	content := ctx.Query("content")
	createdBy := ctx.Query("created_by")
	ctx.Query("state")
	state := com.StrTo(ctx.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExitTagById(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state

			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

// EditArticle
// @Tags    文章
// @Summary 修改文章
// @Param   id          path  int    true  "文章ID"
// @Param   tag_id      query int    false "标签 ID"
// @Param   title       query string false "标题"
// @Param   desc        query string false "描述"
// @Param   content     query string false "内容"
// @Param   modified_by query string true  "修改人"
// @Router  /api/v1/articles/{id} [put]
func EditArticle(ctx *gin.Context) {
	valid := validation.Validation{}

	id := com.StrTo(ctx.Param("id")).MustInt()
	tagId := com.StrTo(ctx.Query("tag_id")).MustInt()
	title := ctx.Query("title")
	desc := ctx.Query("desc")
	content := ctx.Query("content")
	modifiedBy := ctx.Query("modified_by")

	var state int = -1
	if arg := ctx.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Required(id, "id").Message("文章 ID 不能为空")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExitArticleById(id) {
			if models.ExitTagById(tagId) {
				data := make(map[string]interface{})
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}

				data["modified_by"] = modifiedBy

				models.EditArticle(id, data)
				code = e.SUCCESS
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// DeleteArticle
// @Tags    文章
// @Summary 删除文章
// @Param   id path int true "文章ID"
// @Router  /api/v1/articles/{id} [delete]
func DeleteArticle(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExitArticleById(id) {
			models.DeleteArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
