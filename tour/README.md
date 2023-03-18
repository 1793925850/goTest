# 项目：tour（工具集）

项目类型：命令行应用

# 目录结构

- tour：项目工作目录
  - main.go：项目入口函数，启动项目
  - go.mod
    - go.sum
  - cmd
    - word.go：用于放置单词格式转换的子命令word
    - root.go：用于放置根命令
  - internal
    - word
      - word.go：具体编写单词各种转换方法
  - pkg
  - README.md

# 功能清单

- 单词格式转换
- 便利的时间工具
- SQL语句到结构体的转换
- JSON到结构体的转换