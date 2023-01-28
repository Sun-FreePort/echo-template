# Echo Framework Template

使用原则：具体项目**必须**无需修改非 `handler / model` 文件夹，从而提供快速升级的解决方案。

## 目录结构

Ps: 加粗标识可修改

- \
  - **handler**\：业务逻辑代码与路由
    - aRoute.go：API 路由
    - handler.go：业务入口
  - **model**\：模型结构
  - **structure**\：纯数据结构（无关模型等）
  - dict\：i18n 字典，`custom` 文件夹内属于框架文件，不应修改或删除
  - db\：数据库文件，不应修改或删除
  - docs\：自动生成的文档
  - help\：辅助方法，`I` 打头的文件为框架文件，不应修改或删除
  - env.json：配置文件模板
  - **env-dev.json**：配置文件（开发环境）
  - **env-prod.json**：配置文件（生产环境）

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
