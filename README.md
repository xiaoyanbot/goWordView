goWordView 基于go语言的Word文件预览
===============================

本仓库工作基于 GoFileView 和 kkFileView 工作基础之上，并遵循相关开源协议。

感谢 GoFileView ( *https://github.com/leeli73/goFileView.git* ) 和 kkFileView( *https://gitee.com/kekingcn/file-online-preview.git* )

仅支持Linux系统，url暂时不支持中文。  已测试Ubuntu、CentOS

目前已经完成
============

Word转码为PDF，对PDF的在线预览

TODO
====

本地路径指定，省去下载步骤

部署编译
========

准备
----

安装Libreoffice

确保Libreoffice在path目录下

运行
----

下载源码

设置 GOPATH 为当前的目录

启动静态文件服务器

go run webs.go

启动预览服务

go run main.go

在浏览器里输入

http://localhost:8088/perview/onlinePreview?url=http://localhost:9090/123.docx


