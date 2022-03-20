package service

// 入参校验增加绑定/验证结构体
type CountArticleRequest struct{
	Title string `form:"title" binding:"max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct{
	Title string `form:"title" binding:"max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct{
	Title string `form: "title" binding:"required, min =3, max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct{
	ID uint32 `form: "id" binding: "required,gte=1"`
	Title string `form: "title" binding:"required, min =3, max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct{
	ID uint32 `form: "id" binding: "required,gte=1"`
}
