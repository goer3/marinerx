package utils

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// 解析结构体类型，支持指针、切片等多层嵌套，返回最终的结构体类型
func resolveStructType(obj any) reflect.Type {
	// 拿到运行时类型，空对象直接返回
	t := reflect.TypeOf(obj)
	if t == nil {
		return nil
	}

	// 连续解引用指针，拿到真实类型
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	// 如果是切片/数组，继续取元素类型
	if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		elem := t.Elem()
		// 元素也可能是指针（如 []*Req），继续解引用
		for elem.Kind() == reflect.Pointer {
			elem = elem.Elem()
		}
		// 元素最终是结构体才是合法请求类型
		if elem.Kind() == reflect.Struct {
			return elem
		}
		return nil
	}

	// 非切片场景，直接判断是否是结构体
	if t.Kind() == reflect.Struct {
		return t
	}

	// 其他类型（map/string/int等）都不支持
	return nil
}

// 获取验证错误信息
func getTagMessageByValidationErrors(vErrs validator.ValidationErrors, objType reflect.Type) string {
	for _, e := range vErrs {
		if f, exist := objType.FieldByName(e.Field()); exist {
			// 优先返回自定义 msg 标签
			if msg := f.Tag.Get("msg"); msg != "" {
				return msg
			}
			// 没有 msg 时回退到默认错误
			return e.Error()
		}
	}
	return ""
}

// 获取验证错误信息
func getMessageByError(err error, objType reflect.Type) string {
	// 直接是 validator.ValidationErrors
	if errs, ok := err.(validator.ValidationErrors); ok {
		if msg := getTagMessageByValidationErrors(errs, objType); msg != "" {
			return msg
		}
	}

	// 错误被包装过，使用 errors.As 解开
	var vErrs validator.ValidationErrors
	if errors.As(err, &vErrs) {
		if msg := getTagMessageByValidationErrors(vErrs, objType); msg != "" {
			return msg
		}
	}

	return ""
}

// 解析切片对象，支持指针、切片等多层嵌套，返回最终的切片值和是否成功解析的标志
func resolveSliceValue(obj any) (reflect.Value, bool) {
	// 拿到反射值
	v := reflect.ValueOf(obj)
	if !v.IsValid() {
		return reflect.Value{}, false
	}

	// 连续解引用指针（例如 *[]Req -> []Req）
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return reflect.Value{}, false
		}
		v = v.Elem()
	}

	// 只接受切片或数组
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return reflect.Value{}, false
	}

	return v, true
}

// 对切片对象逐条校验，返回首条失败数据的真实序号和错误信息
func getFirstSliceItemMessage(obj any) (int, string, bool) {
	// 先解析出切片值
	v, ok := resolveSliceValue(obj)
	if !ok {
		return 0, "", false
	}

	// 复用 Gin 注册过的 validator（含自定义规则）
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok || validate == nil {
		return 0, "", false
	}

	// 逐条校验，命中首个错误即返回
	for i := 0; i < v.Len(); i++ {
		item := v.Index(i)
		// 元素可能是指针，继续解引用
		for item.Kind() == reflect.Pointer {
			if item.IsNil() {
				break
			}
			item = item.Elem()
		}

		// 只处理结构体元素
		if item.Kind() != reflect.Struct {
			continue
		}

		// 校验失败后组装更友好的提示
		if itemErr := validate.Struct(item.Interface()); itemErr != nil {
			msg := getMessageByError(itemErr, item.Type())
			if msg == "" {
				msg = itemErr.Error()
			}
			return i + 1, msg, true
		}
	}

	return 0, "", false
}

// 获取验证错误信息
func GetValidateErrorMessage(err error, obj any) string {
	// 空错误直接返回
	if err == nil {
		return ""
	}

	// 解析请求结构体类型，用于读取字段 msg 标签
	objType := resolveStructType(obj)
	if objType == nil {
		return err.Error()
	}

	// 切片校验错误（例如 []Request）
	if errs, ok := err.(binding.SliceValidationError); ok {
		// 优先返回“第N条数据”的提示
		if index, msg, found := getFirstSliceItemMessage(obj); found {
			return fmt.Sprintf("第%d条数据，%s", index, msg)
		}

		// 兜底返回第一条原始错误
		if len(errs) > 0 {
			// 这里可能是 validator 默认英文错误
			return errs[0].Error()
		}
		return err.Error()
	}

	// 非切片场景：直接提取 msg
	if msg := getMessageByError(err, objType); msg != "" {
		return msg
	}

	// 最终兜底
	return err.Error()
}
