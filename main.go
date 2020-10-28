package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/kiyonlin/dawn"
	"github.com/kiyonlin/dawn/config"
	"github.com/kiyonlin/dawn/fiberx"
	"github.com/kiyonlin/dawn/log"
)

func main() {
	config.Load(".")
	config.LoadEnv()

	log.InitFlags(nil)
	flag.Parse()
	defer log.Flush()

	sloop := dawn.Default()

	router := sloop.Router()
	router.Get("/", func(c *fiber.Ctx) error {
		return fiberx.Message(c, "Welcome to dawn 👋")
	})

	log.Infoln(0, sloop.Run(":3000"))
}
