package work

import (
	"strings"
	"unicode"

)

// 单词全部转换为大写
func ToUpper(s string) string  {
	return strings.ToUpper(s)
}

// 单词全部转换为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// 下划线单词转换为大写驼峰单词   eee_ccc_bbb  ----> EeeCccBbb
func UndersocoreToUpperCamelCase(s string) string {
	// 首先将下划线替换为空白字符
	s = strings.Replace(s,"_"," ",-1)
	// 将每个单词的首字母变大写
	s =strings.Title(s)
	// 将空白字符去掉
	return strings.Replace(s," ","",-1)
}

// 下划线单词转换为小写驼峰单词   EEE_CCC_BBB  ----> eEECCCBBB
func UndersocoreToLowerCamelCase(s string) string {
	// 首先调用UndersocoreToUpperCamelCase 函数将其变成大写驼峰单词
	s = UndersocoreToUpperCamelCase(s)
	// 然后将首字母小写
	return  string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// 驼峰单词转为下划线  AbcDde abc_def
func CamelCaseToUndersocore(s string) string  {
	// 定义临时存放数据的切片
	var output []rune
	for k,v := range s {
		// 将第一个字符转为小写后放到切片中,并退出本次循环
		if k == 0 {
			output = append(output,unicode.ToLower(rune(v)))
			continue
		}
		// 如果字符为大写，先在切片中插入下划线
		if unicode.IsUpper(rune(v)) {
			output = append(output,'_')
		}
		// 将字符转为小写后插入到切片中
		output = append(output,unicode.ToLower(rune(v)))
	}
	return string(output)
}