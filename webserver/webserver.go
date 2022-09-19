package webserver

import (
	"fmt"
	"log"
	"net/http"
	"sesi6/webserver/controllers"

	"github.com/gin-gonic/gin"
)

const PORT = ":4000"

func Start() {
	r := gin.Default()
	r.GET("/users", controllers.GetUsersHandler)
	r.GET("/users/:id", controllers.GetUsersHandler)
	r.POST("/users", controllers.CreateUserHandler)

	log.Println("=========== Server started ===========")
	r.Run(PORT)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	fmt.Fprint(w, msg)
}
