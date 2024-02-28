package controllers

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type response_pr struct {
	Status          int         `json:"status"`
	Message         string      `json:"message"`
	Record          interface{} `json:"record"`
	Listbranch      interface{} `json:"listbranch"`
	Listdepartement interface{} `json:"listdepartement"`
	Listcurr        interface{} `json:"listcurr"`
	Perpage         int         `json:"perpage"`
	Totalrecord     int         `json:"totalrecord"`
	Time            string      `json:"time"`
}

func PurchaseRequesthome(c *fiber.Ctx) error {
	type payload_prhome struct {
		Purchaserequest_search string `json:"purchaserequest_search"`
		Purchaserequest_status string `json:"purchaserequest_status"`
		Purchaserequest_page   int    `json:"purchaserequest_page"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_prhome)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(response_pr{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":        hostname,
			"purchaserequest_search": client.Purchaserequest_search,
			"purchaserequest_status": client.Purchaserequest_status,
			"purchaserequest_page":   client.Purchaserequest_page,
		}).
		Post(PATH + "api/purchaserequest")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*response_pr)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":          result.Status,
			"message":         result.Message,
			"record":          result.Record,
			"listbranch":      result.Listbranch,
			"listdepartement": result.Listdepartement,
			"listcurr":        result.Listcurr,
			"perpage":         result.Perpage,
			"totalrecord":     result.Totalrecord,
			"time":            time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func PurchaseRequestdetail(c *fiber.Ctx) error {
	type payload_prhome struct {
		Purchaserequest_id string `json:"purchaserequest_id"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_prhome)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"purchaserequest_id": client.Purchaserequest_id,
		}).
		Post(PATH + "api/purchaserequestdetail")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func PurchaseRequestdetail_view(c *fiber.Ctx) error {
	type payload_prhome struct {
		Purchaserequest_tipedoc  string `json:"purchaserequest_tipedoc"`
		Purchaserequest_idbranch string `json:"purchaserequest_idbranch"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_prhome)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"purchaserequest_tipedoc":  client.Purchaserequest_tipedoc,
			"purchaserequest_idbranch": client.Purchaserequest_idbranch,
		}).
		Post(PATH + "api/purchaserequestdetailview")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func PurchaseRequestSave(c *fiber.Ctx) error {
	type payload_prsave struct {
		Page                          string          `json:"page"`
		Sdata                         string          `json:"sdata" `
		Purchaserequest_search        string          `json:"purchaserequest_search" `
		Purchaserequest_page          int             `json:"purchaserequest_page" `
		Purchaserequest_id            string          `json:"purchaserequest_id" `
		Purchaserequest_idbranch      string          `json:"purchaserequest_idbranch" `
		Purchaserequest_iddepartement string          `json:"purchaserequest_iddepartement" `
		Purchaserequest_idemployee    string          `json:"purchaserequest_idemployee" `
		Purchaserequest_idcurr        string          `json:"purchaserequest_idcurr" `
		Purchaserequest_tipedoc       string          `json:"purchaserequest_tipedoc" `
		Purchaserequest_listdetail    json.RawMessage `json:"purchaserequest_listdetail" `
		Purchaserequest_totalitem     float32         `json:"purchaserequest_totalitem" `
		Purchaserequest_subtotal      float32         `json:"purchaserequest_subtotal" `
		Purchaserequest_remark        string          `json:"purchaserequest_remark" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_prsave)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":               hostname,
			"page":                          client.Page,
			"sdata":                         client.Sdata,
			"purchaserequest_search":        client.Purchaserequest_search,
			"purchaserequest_page":          client.Purchaserequest_page,
			"purchaserequest_id":            client.Purchaserequest_id,
			"purchaserequest_idbranch":      client.Purchaserequest_idbranch,
			"purchaserequest_iddepartement": client.Purchaserequest_iddepartement,
			"purchaserequest_idemployee":    client.Purchaserequest_idemployee,
			"purchaserequest_idcurr":        client.Purchaserequest_idcurr,
			"purchaserequest_tipedoc":       client.Purchaserequest_tipedoc,
			"purchaserequest_listdetail":    string(client.Purchaserequest_listdetail),
			"purchaserequest_totalitem":     client.Purchaserequest_totalitem,
			"purchaserequest_subtotal":      client.Purchaserequest_subtotal,
			"purchaserequest_remark":        client.Purchaserequest_remark,
		}).
		Post(PATH + "api/purchaserequestsave")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func PurchaseRequestStatusSave(c *fiber.Ctx) error {
	type payload_prsave struct {
		Purchaserequest_id     string `json:"purchaserequest_id" `
		Purchaserequest_status string `json:"purchaserequest_status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_prsave)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"purchaserequest_id":     client.Purchaserequest_id,
			"purchaserequest_status": client.Purchaserequest_status,
		}).
		Post(PATH + "api/purchaserequeststatussave")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
