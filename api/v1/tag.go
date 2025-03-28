package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/iszhaoxg/gin-web/models"
	"github.com/iszhaoxg/gin-web/pkg/e"
	"github.com/iszhaoxg/gin-web/pkg/setting"
	"github.com/iszhaoxg/gin-web/pkg/util"
	"github.com/unknwon/com"
	"net/http"
)

// GetTags
// @Tags    标签
// @Summary 查询标签
// @Param   name query string false "标签名"
// @Router  /api/v1/tags [get]
func GetTags(ctx *gin.Context) {
	name := ctx.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := ctx.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS
	data["list"] = models.GetTags(util.GetPage(ctx), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddTags
// @Tags    标签
// @Summary 新增标签
// @Param   name       query string true  "标签名"
// @Param   state      query int    false "状态"
// @Param   created_by query string false "创建人"
// @Router  /api/v1/tags [post]
func AddTags(ctx *gin.Context) {
	name := ctx.Query("name")
	state := com.StrTo(ctx.DefaultQuery("state", "0")).MustInt()
	createdBy := ctx.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0和1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExitTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// EditTags
// @Tags    标签
// @Summary 修改标签
// @Param   id          path  int    true  "标签ID"
// @Param   name        query string true  "标签名"
// @Param   state       query int    false "状态"
// @Param   modified_by query string true  "创建人"
// @Router  /api/v1/tags/{id} [put]
func EditTags(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()
	name := ctx.Query("name")
	modifiedBy := ctx.Query("modified_by")

	valid := validation.Validation{}
	var state int = -1
	if arg := ctx.Query("state"); arg != "" {
		state := com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只能是0和1")
	}
	valid.Required(id, "id").Message("id不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长100字符")
	valid.MaxSize(name, 100, "name").Message("名字最长100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExitTagById(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// DeleteTags
// @Tags    标签
// @Summary 删除标签
// @Param   id path int true "标签ID"
// @Router  /api/v1/tags/{id} [delete]
func DeleteTags(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于1")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExitTagById(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
