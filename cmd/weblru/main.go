package main

import (
	"log"
	"net/http"

	"github.com/rileyr/golru/web"
	"github.com/rileyr/middleware"
	"github.com/rileyr/middleware/wares"

	"github.com/julienschmidt/httprouter"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	DefaultCacheSize = 100
	DefaultPort      = ":3030"
)

func main() {
	// Configuration:
	flag.Int("size", DefaultCacheSize, "upper limit to cache size")
	flag.String("port", DefaultPort, "port on which server will listen")
	flag.String("username", "", "basic auth username, optional")
	flag.String("password", "", "basic auth password, optional")
	flag.Parse()
	viper.BindPFlag("size", flag.Lookup("size"))
	viper.BindPFlag("port", flag.Lookup("port"))
	viper.BindPFlag("username", flag.Lookup("username"))
	viper.BindPFlag("password", flag.Lookup("password"))

	size := viper.GetInt("size")
	port := viper.GetString("port")
	username := viper.GetString("username")
	password := viper.GetString("password")

	// Middleware Stack:
	stack := middleware.NewStack()
	stack.Use(wares.RequestID)
	stack.Use(wares.Logging)
	if username != "" && password != "" {
		stack.Use(wares.BasicAuth(username, password))
	}

	// App Setup:

	handler := weblru.New(size)
	router := httprouter.New()

	router.GET("/cache", stack.Wrap(handler.Get))
	router.POST("/cache", stack.Wrap(handler.Add))
	router.DELETE("/cache", stack.Wrap(handler.Remove))

	// Start App:
	log.Printf("WebLRU starting on port %s with size of %d\n", port, size)
	log.Fatal(http.ListenAndServe(port, router))
}
