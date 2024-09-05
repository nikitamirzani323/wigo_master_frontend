package routers

import (
	"bitbucket.org/isbtotogroup/wigo_master_frontend/controllers"
	"bitbucket.org/isbtotogroup/wigo_master_frontend/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
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
	app.Get("/check/healthz", controllers.HealthCheck)
	app.Get("/check/dashboard", monitor.New())

	api := app.Group("/api", middleware.Gateway)

	api.Post("/login", controllers.CheckLogin)
	api.Post("/valid", controllers.Home)
	api.Post("/alladmin", controllers.Adminhome)
	api.Post("/detailadmin", controllers.AdminDetail)
	api.Post("/saveadmin", controllers.AdminSave)
	api.Post("/alladminrule", controllers.Adminrulehome)
	api.Post("/saveadminrule", controllers.AdminruleSave)
	api.Post("/curr", controllers.Currencyhome)
	api.Post("/currsave", controllers.CurrencySave)
	api.Post("/domain", controllers.Domainhome)
	api.Post("/domainsave", controllers.DomainSave)
	api.Post("/company", controllers.Companyhome)
	api.Post("/companyadmin", controllers.Companyadminhome)
	api.Post("/companyadminrule", controllers.Companyadminrulehome)
	api.Post("/companymoney", controllers.Companymoneyhome)
	api.Post("/companyconf", controllers.Companyconfhome)
	api.Post("/companysave", controllers.CompanySave)
	api.Post("/companyadminsave", controllers.CompanyadminSave)
	api.Post("/companyadminrulesave", controllers.CompanyadminruleSave)
	api.Post("/companymoneysave", controllers.CompanymoneySave)
	api.Post("/companymoneydelete", controllers.CompanymoneyDelete)
	api.Post("/companyconfsave", controllers.CompanyconfSave)

	return app
}
