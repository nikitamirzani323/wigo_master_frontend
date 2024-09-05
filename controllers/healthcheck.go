package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {

	realip := "0.0.0.0"
	if len(c.IPs()) > 0 {
		realip = c.IPs()[0]
	}

	if realip == "0.0.0.0" {
		realip = c.IP()
	}
	log.Println("realIp: ", realip)
	res := map[string]interface{}{
		"status":       fiber.StatusOK,
		"data":         "Server is up and running",
		"container_ip": c.IP(),
		"ip_list":      c.IPs(),
		"real_ip":      realip,
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
