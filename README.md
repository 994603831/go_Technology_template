# Go语言学习技术模板

突然有个想法，把当前知道的 Golang 技术框架的模板整理。这只是我个人写的demo，其中也有许多不足和错误，希望各位朋友可以通过发起 Fork + Pull Request 的方式进行补充和更新。

## 交流社区

### 中文社区

- [Golang 中国](http://www.golangtc.com/)：国内较早的 Go 社区，汇聚各类信息与服务

- [Study Golang](http://studygolang.com/)：国内 Go 社区先驱，同样汇聚各类信息与服务

- [Revel 交流论坛](http://gorevel.cn/)：[Revel](https://github.com/revel/revel) 框架的中文社区

- [GoCN Forum](https://gocn.vip/)：Go 语言爱好者中文交流论坛

  

### 英文社区

- [Go Forum](https://forum.golangbridge.org/)：Go 语言爱好者英文交流论坛
- [golang-nuts 邮件列表](https://groups.google.com/forum/#!forum/golang-nuts)：Go 语言官方指定邮件列表讨论区
- [Go+ 社区](https://plus.google.com/u/0/communities/114112804251407510571)：Go 语言官方指定 G+ 社区

### 知识图谱

- [Go Knowledge Graph](https://github.com/gocn/knowledge)：Go 知识图谱

## 网址导航

- 官方：

  - [Go 中国站点](https://golang.google.cn/): Go 语言中国官方站点(无需翻墙)
  - [Playground](http://play.golang.org/)：Go 语言代码在线运行

- 国内镜像：

  - [Go 指南国内镜像](http://tour.golangtc.com/)
  - [Go 语言国内下载镜像](http://www.golangtc.com/download)
  - [Go 官方网站国内镜像](http://docs.studygolang.com/)

- Web 框架：

  - [Macaron](https://go-macaron.com/)：模块化 Web 框架
  - [Beego](http://beego.me/)：重量级 Web 框架
  - [Revel](https://github.com/revel/revel)：较早成熟的重量级 Web 框架
  - [Martini](https://github.com/go-martini/martini): 一个强大为了编写模块化 Web 应用而生的 Go 语言框架
  - [Echo](https://echo.labstack.com/): 功能模块齐全, 上手容易, 文档示例齐全
  - [Gin](https://github.com/gin-gonic/gin)：轻量级 HTTP Web 框架

- ORM 以及数据库驱动：

  - [xorm](https://github.com/go-xorm/xorm)：支持 MySQL、PostgreSQL、SQLite3 以及 MsSQL
  - [mgo](http://labix.org/mgo)：MongoDB 官方推荐驱动
  - [gorm](https://github.com/jinzhu/gorm): 全功能 ORM (无限接近) 支持 MySQL、PostgreSQL、SQLite3 以及 MsSQL

- 辅助站点：

  - [Go Walker](https://gowalker.org/)：Go 语言在线 API 文档
  - [gobuild.io](http://gobuild.io/)：Go 语言在线二进制编译与下载
  - [Rego](http://regoio.herokuapp.com/)：Go 语言正则在线测试
  - [gopm.io](https://gopm.io/)：科学下载第三方包
  - [Json To Go struct](https://mholt.github.io/json-to-go/):Convert JSON to Go struct在线工具

- 开发工具：

  - [Emacs24](http://ftp.gnu.org/gnu/emacs/)：[配置脚本](https://github.com/wackonline/hack/blob/master/install-mint-dev/install-emacs.d.sh) / [(中文社区)](http://emacser.com/)
  - [LiteIDE](https://github.com/visualfc/liteide)
  - [Sublime Text 2/3](http://sublimetext.com/)：[配置教程](http://my.oschina.net/Obahua/blog/110767)
  - [GoLand](https://www.jetbrains.com/go/?fromMenu)
  - [Atom](https://atom.io/)：[配置插件](https://atom.io/packages/go-plus)（感觉还不错，类似 Sublime，配置比较简单）
  - [VIM](http://www.vim.org/)：[配置插件](https://github.com/humiaozuzu/dot-vimrc)（嫌 vim 配置麻烦的童鞋可以直接用这个）

- 学习站点：

  - [Go by Example](https://gobyexample.com/)
  - [Go database/sql tutorial](http://go-database-sql.org/)

- 支持 Go 的云平台：

  - [Koding](https://koding.com/)

  - [Nitrous.IO](https://www.nitrous.io/)

  - [Get up and running with Go on Google Cloud Platform](https://cloud.google.com/go/)

  - [AWS SDK for Go - Developer Preview](http://aws.amazon.com/cn/sdk-for-go/):=>[github](https://github.com/aws/aws-sdk-go)

  - azure sdk for go

    :=>

    github

    - [How to Use CoreOS on Azure](https://azure.microsoft.com/zh-cn/documentation/articles/virtual-machines-linux-coreos-how-to/)
    - [Create Azure Web app with GoLang extension](https://azure.microsoft.com/zh-cn/documentation/templates/101-webapp-with-golang/)

  - Qiniu

    - [Qiniu SDK for Go](http://developer.qiniu.com/docs/v6/sdk/go-sdk.html):=>[github](https://github.com/qiniu/api.v6)

- 其它站点：

  - [Golang 杂志](https://flipboard.com/section/the-golang-magazine-bJ1GqB)：[订阅说明](http://bbs.go-china.org/post/476)
  - [Reddit](http://www.reddit.com/r/golang/)
  - [Newspaper.IO](http://www.newspaper.io/golang)：Golang 新闻订阅
  - [Go Newsletter](http://www.golangweekly.com/)：Golang 新闻订阅

## 资料汇总

### 中文资料

- 书籍：
  - [《深入解析Go》](https://github.com/tiancaiamao/go-internals)
  - [《Go实战开发》](https://github.com/astaxie/Go-in-Action)
  - [《Go入门指南》](https://github.com/Unknwon/the-way-to-go_ZH_CN)
  - [《Go语言标准库》](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example)
  - [《Go Web 编程》](https://github.com/astaxie/build-web-application-with-golang)
  - [《Go语言博客实践》](https://github.com/achun/Go-Blog-In-Action)
  - [《Go语言学习笔记》](https://github.com/qyuhen/book)
- 翻译：
  - [Effective Go](https://golang.org/doc/effective_go.html) 英文版
  - [The Way to Go](https://github.com/Unknwon/the-way-to-go_ZH_CN) 中文版
  - [《Learning Go》](https://github.com/miekg/gobook)英文版:=>[《Learning Go》](https://github.com/mikespook/Learning-Go-zh-cn) 中文版
- 教程：
  - [《Go编程基础》](https://github.com/Unknwon/go-fundamental-programming)
  - [《Go Web基础》](https://github.com/Unknwon/go-web-foundation)
  - [《Go名库讲解》](https://github.com/Unknwon/go-rock-libraries-showcases)
  - [Go 命令教程](https://github.com/hyper-carrot/go_command_tutorial)