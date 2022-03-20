package v1

import (
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

// 对应路由的处理方法,编写好之后
type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a *Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

// @Summary 获取多篇文章
// @Produce json
// @Param title query string false "文章名称" minlength(3) maxlength(100)
// @Param state query int false "状态" Enum(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (a *Article) List(c *gin.Context) {}

// @Summary 创建文章
// @Produce json
// @Param title body string false "文章名称" minlength(3) maxlength(100)
// @Param state query int false "状态" Enum(0, 1) default(1)
// @Param created_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (a *Article) Create(c *gin.Context) {}

// @Summary 更新文章
// @Produce json
// @Param id path int true "文章ID"
// @Param title body string false "文章名称"
// @Param state body int false "状态" Enum(0, 1) default(1)
// @Param modified_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (a *Article) Update(c *gin.Context) {}

// @Summary 删除标签
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} model.ArticleSwagger  "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (a *Article) Delete(c *gin.Context) {}
