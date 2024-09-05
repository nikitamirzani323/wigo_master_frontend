package middleware

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"

	"bitbucket.org/isbtotogroup/wigo_master_frontend/configs"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type DomainResponse struct {
	DomainStatus string `json:"domainstatus"`
	Status       int    `json:"status"`
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func Gateway(c *fiber.Ctx) error {
	envsecret := os.Getenv("API_SECRET")
	// appmode := os.Getenv("APP_MODE")
	origin := c.Get("Origin")
	realip := c.IP()

	if len(c.IPs()) > 0 {
		realip = c.IPs()[0]
	}

	// log.Println(appmode)
	// if appmode == "development" {
	// 	if origin != "" && c.IP() == "127.0.0.1" {

	// 		log.Println("get origin from middleware ", origin)
	// 		log.Println("get request from Ip: ", realip)
	// 		return c.Next()
	// 	}
	// }

	// MD5 encrypt
	h := md5.New()
	io.WriteString(h, envsecret)
	convert := fmt.Sprintf("%x", h.Sum(nil))
	request := reverse(c.Get("x-endpoint-secret"))

	if request == convert {
		log.Println("get request from Ip: ", realip)
		return c.Next()
	}

	bodymaps := map[string]interface{}{
		"domain": origin,
		"tipe":   "master",
	}
	fmt.Println("maps body checkdomain: ", bodymaps)
	axios := resty.New()
	resp, err := axios.R().
		SetResult(DomainResponse{}).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Origin":       origin,
		}).
		SetBody(bodymaps).
		Post(configs.GetPathAPI() + "api/checkdomain")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*DomainResponse)
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	if result.Status == 200 {
		fmt.Println("get origin from middleware ", origin)
		return c.Next()
	}

	return c.JSON(fiber.Map{
		"status":  c.SendStatus(401),
		"message": "You can't access agent from this domain <br> ( Response Error: 401)",
		"record":  nil,
	})
}
