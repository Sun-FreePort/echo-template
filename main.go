package main

import (
	"encoding/json"
	"fmt"
	"github.com/Sun-FreePort/echo-game/cache"
	"github.com/Sun-FreePort/echo-game/db"
	"github.com/Sun-FreePort/echo-game/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"os"
	"time"
)

// echo-swagger middleware

//  @title           Game Swagger API
//  @version         0.1
//  @description     This is a sample server celler server.
//  @termsOfService  http://swagger.io/terms/

//  @contact.name   API Support
//  @contact.url    http://www.swagger.io/support
//  @contact.email  support@swagger.io

//  @license.name  Apache 2.0
//  @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host                       localhost:1323

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description "Type 'Bearer TOKEN' to correctly set the API Key"
func main() {
	// 读取配置
	envFile := ""
	if os.Getenv("CODENATION_ENV") == "prod" {
		envFile = "./env-prod.json"
	} else {
		envFile = "./env-dev.json"
	}
	var payload map[string]string
	content, err := os.ReadFile(envFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		panic(err)
	}

	// 实例化
	h := handler.NewHandler(cache.GetRedis(new(cache.Parameters)), db.New(db.Params{
		Username:  payload["DB_USERNAME"],
		Password:  payload["DB_PASSWORD"],
		Database:  payload["DB_DATABASE"],
		ParseTime: payload["DB_PARSE_TIME"],
	}))

	e := routeNew()
	router := e.Group("")
	h.Register(router)

	port := ":1323"
	if os.Getenv("CODENATION_ENV") == "prod" {
		port = ":8080"
	}

	e.Logger.Fatal(e.Start(port))
}

func routeNew() *echo.Echo {
	e := echo.New()

	e.GET("/doc/*", echoSwagger.WrapHandler)

	e.Pre(middleware.RemoveTrailingSlash())

	if os.Getenv("CODENATION_ENV") == "prod" {
		f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(fmt.Sprintf("error opening file: %v", err))
		}
		// Middleware
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
				`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
				`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
				`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
			Output: f,
		}))
	} else {
		e.Use(middleware.Logger())
	}

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 20 * time.Second,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://www.uiosun.com", "http://www.uiosun.com"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	return e
}