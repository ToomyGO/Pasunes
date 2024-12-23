package main

import (
	auth "cmd/internal/auth"
	shared "cmd/internal/shared"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	shared.InitDB()

	app := fiber.New()
	app.Static("/"   )

	//app.Post("/signup", auth.SignupHanler)

	fmt.Println("Server is Running...")
	log.Fatal(app.Listen(":3000"))
}
