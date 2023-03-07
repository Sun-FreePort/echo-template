package main

import (
	"encoding/json"
	"fmt"
	"github.com/Sun-FreePort/echo-template/cache"
	"github.com/Sun-FreePort/echo-template/db"
	"github.com/Sun-FreePort/echo-template/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"os"
	"strconv"
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
	dbHost, err := strconv.Atoi(payload["CACHE_DATABASE"])
	if err != nil {
		panic(err)
	}

	h := handler.NewHandler(cache.GetRedis(cache.Params{
		Db:   dbHost,
		Port: payload["CACHE_PORT"],
		Host: payload["CACHE_HOST"],
	}), db.New(db.Params{
		Host:      payload["DB_HOST"],
		Port:      payload["DB_PORT"],
		Username:  payload["DB_USERNAME"],
		Password:  payload["DB_PASSWORD"],
		Database:  payload["DB_DATABASE"],
		ParseTime: payload["DB_PARSE_TIME"],
	}))

	e := routeNew(payload)
	router := e.Group("")
	h.Register(router)

	port := ":1323"
	if os.Getenv("CODENATION_ENV") == "prod" {
		port = ":8080"
	}

	e.Logger.Fatal(e.Start(port))
}

func routeNew(payload map[string]string) *echo.Echo {
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
			Format: payload["LOG_FORMAT"],
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
