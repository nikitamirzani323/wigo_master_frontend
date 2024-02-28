package controllers

import (
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type response_employee struct {
	Status          int         `json:"status"`
	Message         string      `json:"message"`
	Record          interface{} `json:"record"`
	Listdepartement interface{} `json:"listdepartement"`
	Perpage         int         `json:"perpage"`
	Totalrecord     int         `json:"totalrecord"`
	Time            string      `json:"time"`
}

func Employeehome(c *fiber.Ctx) error {
	type payload_employeehome struct {
		Employee_search string `json:"employee_search"`
		Employee_page   int    `json:"employee_page"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_employeehome)
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
		SetResult(response_employee{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"employee_search": client.Employee_search,
			"employee_page":   client.Employee_page,
		}).
		Post(PATH + "api/employee")
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
	result := resp.Result().(*response_employee)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":          result.Status,
			"message":         result.Message,
			"record":          result.Record,
			"listdepartement": result.Listdepartement,
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
func Employeeshare(c *fiber.Ctx) error {
	type payload_employeehome struct {
		Employee_iddepartement string `json:"employee_iddepartement"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_employeehome)
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
			"client_hostname":        hostname,
			"employee_iddepartement": client.Employee_iddepartement,
		}).
		Post(PATH + "api/employeeshare")
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
func EmployeeSave(c *fiber.Ctx) error {
	type payload_employeesave struct {
		Page                   string `json:"page"`
		Sdata                  string `json:"sdata" `
		Employee_search        string `json:"employee_search" `
		Employee_page          int    `json:"employee_page" `
		Employee_id            string `json:"employee_id" `
		Employee_iddepartement string `json:"employee_iddepartement" `
		Employee_name          string `json:"employee_name" `
		Employee_alamat        string `json:"employee_alamat" `
		Employee_email         string `json:"employee_email" `
		Employee_phone1        string `json:"employee_phone1" `
		Employee_phone2        string `json:"employee_phone2" `
		Employee_status        string `json:"employee_status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_employeesave)
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
			"client_hostname":        hostname,
			"page":                   client.Page,
			"sdata":                  client.Sdata,
			"employee_search":        client.Employee_search,
			"employee_page":          client.Employee_page,
			"employee_id":            client.Employee_id,
			"employee_iddepartement": client.Employee_iddepartement,
			"employee_name":          client.Employee_name,
			"employee_alamat":        client.Employee_alamat,
			"employee_email":         client.Employee_email,
			"employee_phone1":        client.Employee_phone1,
			"employee_phone2":        client.Employee_phone2,
			"employee_status":        client.Employee_status,
		}).
		Post(PATH + "api/employeesave")
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
