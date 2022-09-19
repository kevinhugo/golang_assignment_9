package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sesi6/webserver/repositories"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUsersHandler(c *gin.Context) {

	userid := c.Param("id")
	userIdInt, err := strconv.Atoi(userid)

	fmt.Println(userIdInt)

	userAgent := c.Request.Header["Platform"][0]

	users, err := repositories.GetUsers(userIdInt)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Something went wrong. Ya Ndak Tau Ko Tanya Saya.",
		})
		return
	}

	if userAgent == "WEB" || userAgent == "" {
		// process as html
		// tpl, err := template.ParseFiles("webserver/views/static/index.html", "webserver/views/static/header.html")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Something went wrong. Ya Ndak Tau Ko Tanya Saya.",
			})
			return
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"users": "webserver/views/static/index.html",
		})
		return
	} else if userAgent != "API" {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Something went wrong. Ya Ndak Tau Ko Tanya Saya.",
		})
		return
	}

	// process as API
	responseData := map[string]interface{}{
		"payload": users,
	}

	c.JSON(http.StatusOK, responseData)
}

func CreateUserHandler(c *gin.Context) {

	var req repositories.User
	err := c.ShouldBind(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Something went wrong. Ya Ndak Tau Ko Tanya Saya.",
		})
		return
	}
	users, _ := repositories.GetUsers(0)
	req.ID = len(users) + 1
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()

	err = repositories.CreateUser(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Something went wrong. Ya Ndak Tau Ko Tanya Saya.",
		})
		return
	}
	responseData := map[string]interface{}{
		"payload": req,
	}

	c.JSON(http.StatusOK, responseData)
}

func writeJsonResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
