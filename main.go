package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	app := fiber.New()
	config := viper.New()
	config.SetConfigType("env")
	// config.SetConfigName(".env")
	// config.AddConfigPath(".")
	config.AutomaticEnv()
	err := config.ReadInConfig()

	if err != nil {
		fmt.Errorf("Fatal error config file: %w \n", err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, test!")
	})

	app.Get("/env", func(c *fiber.Ctx) error {
		name := config.GetString("APP_NAME")
		env := config.GetString("MODE")
		debug := config.GetBool("DEBUG")
		fmt.Println("test", config.GetString("app.name"))
		return c.JSON(fiber.Map{"name": name, "env": env, "debug": debug})
	})

	app.Get("/os-env", func(c *fiber.Ctx) error {
		name := os.Getenv("APP_NAME")
		env := os.Getenv("MODE")
		debug := os.Getenv("DEBUG")
		fmt.Println("app name", os.Getenv("APP_NAME"))
		return c.JSON(fiber.Map{"name": name, "env": env, "debug": debug})
	})

	app.Listen(":3001")
}
