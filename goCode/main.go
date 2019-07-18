package main

import (
	"net/http"
	"os"
	"time"

	"github.com/bxcodec/faker"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// initialize the gin framework
	port := os.Getenv("PORT")
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 15 * time.Second,
	}))
	if port == "" {
		port = "8000"
	}
	r.GET("/employees/:offset", GetEmployees)

	http.ListenAndServe(":"+port, r)

}

func GetEmployees(c *gin.Context) {
	var employees []Employee

	for i := 0; i < 20; i++ {
		var emp Employee
		// initialize fake data
		faker.FakeData(&emp)
		employees = append(employees, emp)
	}
	var data Data
	data.Data = employees
	c.JSON(200, data)
}

type Data struct {
	Data []Employee `json:"data"`
}

type Employee struct {
	// RowID     int    `gorm:"primary_key"`
	FirstName string `gorm:"type:varchar(255)"`
	LastName  string `gorm:"type:varchar(255)"`
	// Position  string `gorm:"type:varchar(255)"`
	// Email     string `gorm:"type:varchar(255)"`
	// Office    string `gorm:"type:varchar(255)"`
	// Extn      string `gorm:"type:varchar(255)"`
	// Age       int
	// Salary    string `gorm:"type:varchar(255)"`
	// StartDate string `gorm:"type:varchar(255)"`
}
