package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/maskfoder/ppapp/auth"
	"github.com/maskfoder/ppapp/database"
)

func LoadEnv(envfile string) {
	err := godotenv.Load(envfile)
	if err != nil {
		log.Fatal(err)
	}

}

func StartDB() {
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos",
		host, username, password, databaseName, port)

	database.Connect(dsn)
	database.Migrate()
}

func ServeAPI() {

	router := gin.Default()
	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/login", auth.Login)
	publicRoutes.POST("/register", auth.RegisterUser)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(auth.JWTAuthMiddleWare())
	protectedRoutes.POST("/project", auth.AddProject)
	protectedRoutes.POST("/task", auth.AddTask)
	protectedRoutes.GET("/all", auth.GetProjectsAndTasks)

	router.Run(os.Getenv("API_SERVER_PORT"))
}

func main() {
	LoadEnv(".env.local")
	StartDB()
	ServeAPI()

}
