sync-google-hosts
===============

windows下自动更新hosts文件，使本机能够正常访问google并使用其服务（gmail，chrome同步功能 etc.）

安装
----
```Bash
go get github.com/Unknwon/goconfig
go get github.com/jas0ns/sync-google-hosts
go install -ldflags="-H windowsgui" github.com/jas0ns/sync-google-hosts
``` 
请将源码目录下的`config.ini`文件与编译后的exe文件置于同一目录下

如果你需要将此同步程序在每次开机的时候自动运行，
可以将 %GOPATH%/bin/syngooglehosts.exe 以及同一目录下的`config.ini`拷贝到 开始菜单->所有程序->启动 文件夹下即可。
