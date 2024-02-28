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

	app.Post("/api/employee", controllers.Employeehome)
	app.Post("/api/employeeshare", controllers.Employeeshare)
	app.Post("/api/employeesave", controllers.EmployeeSave)
	app.Post("/api/departement", controllers.Departementhome)
	app.Post("/api/departementsave", controllers.DepartementSave)
	app.Post("/api/catevendor", controllers.Catevendorhome)
	app.Post("/api/catevendorsave", controllers.CatevendorSave)
	app.Post("/api/vendor", controllers.Vendorhome)
	app.Post("/api/vendorshare", controllers.Vendorshare)
	app.Post("/api/vendorsave", controllers.VendorSave)
	app.Post("/api/curr", controllers.Currencyhome)
	app.Post("/api/currsave", controllers.CurrencySave)
	app.Post("/api/uom", controllers.Uomhome)
	app.Post("/api/uomshare", controllers.Uomshare)
	app.Post("/api/uomsave", controllers.UomSave)
	app.Post("/api/branch", controllers.Branchhome)
	app.Post("/api/branchsave", controllers.BranchSave)
	app.Post("/api/warehouse", controllers.Warehousehome)
	app.Post("/api/warehousesave", controllers.WarehouseSave)
	app.Post("/api/warehousestorage", controllers.WarehouseStorage)
	app.Post("/api/warehousestoragesave", controllers.WarehouseStorageSave)
	app.Post("/api/warehousestoragebin", controllers.WarehouseStorageBin)
	app.Post("/api/warehousestoragebinsave", controllers.WarehouseStorageBinSave)
	app.Post("/api/merek", controllers.Merekhome)
	app.Post("/api/merekshare", controllers.Merekshare)
	app.Post("/api/mereksave", controllers.MerekSave)
	app.Post("/api/categoryitem", controllers.Cateitemhome)
	app.Post("/api/categoryitemsave", controllers.CateitemSave)
	app.Post("/api/item", controllers.Itemhome)
	app.Post("/api/itemshare", controllers.Itemshare)
	app.Post("/api/itemuom", controllers.Itemuom)
	app.Post("/api/itemsave", controllers.ItemSave)
	app.Post("/api/itemuomsave", controllers.ItemuomSave)
	app.Post("/api/itemuomdelete", controllers.ItemuomDelete)

	app.Post("/api/purchaserequest", controllers.PurchaseRequesthome)
	app.Post("/api/purchaserequestdetail", controllers.PurchaseRequestdetail)
	app.Post("/api/purchaserequestdetailview", controllers.PurchaseRequestdetail_view)
	app.Post("/api/purchaserequestsave", controllers.PurchaseRequestSave)
	app.Post("/api/purchaserequeststatussave", controllers.PurchaseRequestStatusSave)
	app.Post("/api/rfq", controllers.Rfqhome)
	app.Post("/api/rfqdetail", controllers.Rfqdetail)
	app.Post("/api/rfqsave", controllers.RfqSave)
	app.Post("/api/rfqstatussave", controllers.RfqStatusSave)
	app.Post("/api/posave", controllers.PoSave)

	return app
}
