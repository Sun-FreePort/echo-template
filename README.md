## Todo

- [x] 注册
- [x] 登录
- [x] 数据字典

## 直接运行

刷新文档并运行项目：`swag init && go run main.go`

## 编译

- Windows：`GOOS=windows GOARCH=amd64 go build -o bin/app-amd64.exe main.go`
- macOS：`GOOS=darwin GOARCH=amd64 go build -o bin/app-amd64 main.go`
- Linux：`GOOS=linux GOARCH=amd64 go build -o bin/app-amd64 main.go`

更多打包可以参考：

[How to cross-compile Go programs for Windows, macOS, and Linux](https://freshman.tech/snippets/go/cross-compile-go-programs/)

## 常驻进程

> 你可以采用任意你喜欢的常驻进程管理软件。此处仅以 supervisor 为例：

采用 Pip 安装：

`pip install supervisor -i https://pypi.tuna.tsinghua.edu.cn/simple`

创建配置文件，并修正配置：

`echo_supervisord_conf > /etc/supervisord.conf`

```ini
[include]
files = /etc/supervisor/conf.d/*.conf
```

常见的文件：

```shell
# 启动：
supervisord -c /etc/supervisord.conf
# 获取进程：
supervisorctl status
# 刷新配置：
supervisorctl reload
# 关闭：
ps -ef | grep supervisord
kill -s SIGTERM 879
```

```ini
[program:go-game-api]
directory=/var/www/go-game
command=/var/www/go-game/app-amd64
autostart=true
autorestart=true
stderr_logfile=/var/www/go-game/app-amd64.err
stdout_logfile=/var/www/go-game/app-amd64.log
environment=CODENATION_ENV=prod
```
