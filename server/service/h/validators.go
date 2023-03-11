package h

import (
	"github.com/gin-gonic/gin/binding"
	localeEn "github.com/go-playground/locales/en"
	localeZhHans "github.com/go-playground/locales/zh_Hans"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

var Trans ut.Translator

// https://www.jianshu.com/p/51b9cd2006a8
func init() {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhHans := localeZhHans.New()
		en := localeEn.New()
		uni := ut.New(zhHans, en)
		Trans, _ = uni.GetTranslator("zhHans")
		//注册翻译器
		_ = zh.RegisterDefaultTranslations(validate, Trans)
		//注册自定义标签label来代替显示的StructFields字段
		validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			return field.Tag.Get("label")
		})
		//覆盖required校验规则
		_ = validate.RegisterTranslation("required", Trans, func(ut ut.Translator) error {
			return ut.Add("required", "请输入{0}", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, err := ut.T(fe.Tag(), fe.Field())
			if err != nil {
				return fe.(error).Error()
			}
			return t
		})

	}
}
