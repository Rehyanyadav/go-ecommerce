package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("i'm main function ")

	app := fiber.New()

	// Use a non-privileged port (>=1024). Port 900 is privileged on many systems
	// and will fail without elevated permissions. Also handle the returned error.
	if err := app.Listen(":9000"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
