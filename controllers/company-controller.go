package controllers

import (
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type response_company struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Record      interface{} `json:"record"`
	Listcurr    interface{} `json:"listcurr"`
	Perpage     int         `json:"perpage"`
	Totalrecord int         `json:"totalrecord"`
	Time        string      `json:"time"`
}
type response_companyadmin struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	Record   interface{} `json:"record"`
	Listrule interface{} `json:"listrule"`
	Time     string      `json:"time"`
}
type response_companyadminrule struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
	Time    string      `json:"time"`
}

func Companyhome(c *fiber.Ctx) error {
	type payload_companyhome struct {
		Company_search string `json:"company_search"`
		Company_page   int    `json:"company_page"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_companyhome)
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
		SetResult(response_company{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"company_search":  client.Company_search,
			"company_page":    client.Company_page,
		}).
		Post(PATH + "api/company")
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
	result := resp.Result().(*response_company)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":      result.Status,
			"message":     result.Message,
			"record":      result.Record,
			"listcurr":    result.Listcurr,
			"perpage":     result.Perpage,
			"totalrecord": result.Totalrecord,
			"time":        time.Since(render_page).String(),
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
func Companyadminhome(c *fiber.Ctx) error {
	type payload_companyadminhome struct {
		Companyadmin_idcompany string `json:"companyadmin_idcompany"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_companyadminhome)
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
		SetResult(response_companyadmin{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":        hostname,
			"companyadmin_idcompany": client.Companyadmin_idcompany,
		}).
		Post(PATH + "api/companyadmin")
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
	result := resp.Result().(*response_companyadmin)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":   result.Status,
			"message":  result.Message,
			"record":   result.Record,
			"listrule": result.Listrule,
			"time":     time.Since(render_page).String(),
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
func Companyadminrulehome(c *fiber.Ctx) error {
	type payload_companyadminhome struct {
		Companyadmin_idcompany string `json:"companyadmin_idcompany"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_companyadminhome)
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
		SetResult(response_companyadminrule{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":        hostname,
			"companyadmin_idcompany": client.Companyadmin_idcompany,
		}).
		Post(PATH + "api/companyadminrule")
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
	result := resp.Result().(*response_companyadminrule)
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
func CompanySave(c *fiber.Ctx) error {
	type payload_companysave struct {
		Page           string  `json:"page"`
		Sdata          string  `json:"sdata" `
		Company_search string  `json:"company_search" `
		Company_page   int     `json:"company_page" `
		Company_id     string  `json:"company_id" `
		Company_idcurr string  `json:"company_idcurr" `
		Company_name   string  `json:"company_name" `
		Company_owner  string  `json:"company_owner" `
		Company_phone1 string  `json:"company_phone1" `
		Company_phone2 string  `json:"company_phone2" `
		Company_email  string  `json:"company_email" `
		Company_minfee float64 `json:"company_minfee" `
		Company_url1   string  `json:"company_url1" `
		Company_url2   string  `json:"company_url2" `
		Company_status string  `json:"company_status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_companysave)
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
			"client_hostname": hostname,
			"page":            client.Page,
			"sdata":           client.Sdata,
			"company_search":  client.Company_search,
			"company_page":    client.Company_page,
			"company_id":      client.Company_id,
			"company_idcurr":  client.Company_idcurr,
			"company_name":    client.Company_name,
			"company_owner":   client.Company_owner,
			"company_phone1":  client.Company_phone1,
			"company_phone2":  client.Company_phone2,
			"company_email":   client.Company_email,
			"company_minfee":  client.Company_minfee,
			"company_url1":    client.Company_url1,
			"company_url2":    client.Company_url2,
			"company_status":  client.Company_status,
		}).
		Post(PATH + "api/companysave")
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
func CompanyadminSave(c *fiber.Ctx) error {
	type payload_companyadminsave struct {
		Page                   string `json:"page"`
		Sdata                  string `json:"sdata" `
		Companyadmin_id        int    `json:"companyadmin_id" `
		Companyadmin_idcompany string `json:"companyadmin_idcompany" `
		Companyadmin_idrule    int    `json:"companyadmin_idrule" `
		Companyadmin_username  string `json:"companyadmin_username" `
		Companyadmin_password  string `json:"companyadmin_password" `
		Companyadmin_name      string `json:"companyadmin_name" `
		Companyadmin_status    string `json:"companyadmin_status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_companyadminsave)
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
			"companyadmin_id":        client.Companyadmin_id,
			"companyadmin_idrule":    client.Companyadmin_idrule,
			"companyadmin_idcompany": client.Companyadmin_idcompany,
			"companyadmin_username":  client.Companyadmin_username,
			"companyadmin_password":  client.Companyadmin_password,
			"companyadmin_name":      client.Companyadmin_name,
			"companyadmin_status":    client.Companyadmin_status,
		}).
		Post(PATH + "api/companyadminsave")
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
func CompanyadminruleSave(c *fiber.Ctx) error {
	type payload_companyadminsave struct {
		Page                         string `json:"page"`
		Sdata                        string `json:"sdata" `
		Companyadminrule_id          int    `json:"companyadminrule_id" `
		Companyadminrule_idcompany   string `json:"companyadminrule_idcompany" `
		Companyadminrule_nmruleadmin string `json:"companyadminrule_nmruleadmin" `
		Companyadminrule_ruleadmin   string `json:"companyadminrule_ruleadmin" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_companyadminsave)
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
			"client_hostname":              hostname,
			"page":                         client.Page,
			"sdata":                        client.Sdata,
			"companyadminrule_id":          client.Companyadminrule_id,
			"companyadminrule_idcompany":   client.Companyadminrule_idcompany,
			"companyadminrule_nmruleadmin": client.Companyadminrule_nmruleadmin,
			"companyadminrule_ruleadmin":   client.Companyadminrule_ruleadmin,
		}).
		Post(PATH + "api/companyadminrulesave")
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
