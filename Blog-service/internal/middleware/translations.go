package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	// 多语言包,与universal-translator配套使用
	"github.com/go-playground/locales/en" 
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	// 通用翻译器
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	// validator的翻译器
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/en"
)


// 自定义中间件 Translation， 
func Translations() gin.HandlerFunc{
	
	return func(c *gin.Context){
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		locale := c.GetHeader("locale")		// 获取预定的header参数locale
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validato)
		if ok{
			switch locale{
			case "zh":
				// 将验证器和对应语言类型的Translator注册进来
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v, trans)
				break
			default:
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			}
			// 将Translator存储到全局上下文中
			c.Set("trans", trans)
		}
	}

	c.Next()
}