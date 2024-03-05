package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/WIGO_MASTER_FRONTEND/controllers"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		// Set some security headers:
		// c.Set("Content-Security-Policy", "frame-ancestors 'none'")
		// c.Set("X-XSS-Protection", "1; mode=block")
		// c.Set("X-Content-Type-Options", "nosniff")
		// c.Set("X-Download-Options", "noopen")
		// c.Set("Strict-Transport-Security", "max-age=5184000")
		// c.Set("X-Frame-Options", "SAMEORIGIN")
		// c.Set("X-DNS-Prefetch-Control", "off")

		// Go to next middleware:
		return c.Next()
	})
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Static("/", "frontend/public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})
	app.Get("/ipaddress", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      "data",
			"BASEURL":     c.BaseURL(),
			"HOSTNAME":    c.Hostname(),
			"IP":          c.IP(),
			"IPS":         c.IPs(),
			"OriginalURL": c.OriginalURL(),
			"Path":        c.Path(),
			"Protocol":    c.Protocol(),
			"Subdomain":   c.Subdomains(),
		})
	})
	app.Get("/dashboard", monitor.New())

	app.Post("/api/login", controllers.CheckLogin)
	app.Post("/api/valid", controllers.Home)
	app.Post("/api/alladmin", controllers.Adminhome)
	app.Post("/api/detailadmin", controllers.AdminDetail)
	app.Post("/api/saveadmin", controllers.AdminSave)

	app.Post("/api/alladminrule", controllers.Adminrulehome)
	app.Post("/api/saveadminrule", controllers.AdminruleSave)
	app.Post("/api/curr", controllers.Currencyhome)
	app.Post("/api/currsave", controllers.CurrencySave)
	app.Post("/api/company", controllers.Companyhome)
	app.Post("/api/companyadmin", controllers.Companyadminhome)
	app.Post("/api/companyadminrule", controllers.Companyadminrulehome)
	app.Post("/api/companyconf", controllers.Companyconfhome)
	app.Post("/api/companysave", controllers.CompanySave)
	app.Post("/api/companyadminsave", controllers.CompanyadminSave)
	app.Post("/api/companyadminrulesave", controllers.CompanyadminruleSave)

	return app
}
