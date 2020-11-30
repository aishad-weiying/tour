# 使用 Cobra 实现的简单的单词转换

支持一下几种格式的单词转换

1. 全部单词转换为小写
2. 全部单词转换为大写
3. 下划线单词转换为大写驼峰单词
4. 下划线单词转换为小写驼峰单词
5. 驼峰单词转换为下划线单词

## 使用方式

项目编译之后,直接在命令行执行
```bash
# 帮助
./tour word -h
# 使用
./tour -s "需要转换的字符串" -m 转换模式(1-5)
```
