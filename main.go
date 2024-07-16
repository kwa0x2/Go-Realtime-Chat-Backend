package main

import (	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/realtime-chat-backend/config"
	"github.com/kwa0x2/realtime-chat-backend/controller"
	"github.com/kwa0x2/realtime-chat-backend/repository"
	"github.com/kwa0x2/realtime-chat-backend/routes"
	"github.com/kwa0x2/realtime-chat-backend/service"
	"github.com/zishang520/socket.io/socket"
)

var userSockets = make(map[string]string)

func main() {
	config.LoadEnv()
	router := gin.New()
	config.PostgreConnection()
	io := socket.NewServer(nil, nil)
	store := config.RedisSession()
	router.Use(sessions.Sessions("connect.sid", store))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	userRepository := &repository.UserRepository{DB: config.DB}
	userService := &service.UserService{UserRepository: userRepository}
	userController := &controller.UserController{UserService: userService}

	authRepository := &repository.AuthRepository{DB: config.DB}
	authService := &service.AuthService{AuthRepository: authRepository}
	authController := &controller.AuthController{AuthService: authService}
	
	messageRepository := &repository.MessageRepository{DB: config.DB}
	messageService := &service.MessageService{MessageRepository: messageRepository}
	messageController := &controller.MessageController{MessageService: messageService}

	friendshipRepository := &repository.FriendshipRepository{DB: config.DB}
	friendshipService := &service.FriendshipService{FriendshipRepository: friendshipRepository}
	friendshipController := &controller.FriendshipController{FriendshipService: friendshipService, UserService: userService}

	routes.UserRoute(router, userController)
	routes.AuthRoute(router, authController)
	routes.MessageRoute(router, messageController)
	routes.FriendshipRoute(router, friendshipController)
	routes.SetupSocketIO(router, io, messageService,userService,friendshipService)

	if err := router.Run(":9000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}