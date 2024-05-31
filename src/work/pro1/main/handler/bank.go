package h

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"sunny.ksw.kr/repo/bank"
)

func Bank(route fiber.Router) {

	bankroute := route.Group("/bank")

	bankroute.Get("/instalment/list", func(c *fiber.Ctx) error {

		limit, _ := strconv.Atoi(c.Query("limit", "10"))
		page, _ := strconv.Atoi(c.Query("page", "0"))

		companyname := c.Query("companyname", "")

		search := bank.SearchInstalment_Savings{}

		search.Limit = limit
		search.Page = page
		search.PageOffset = page - 1
		search.CompanyName = companyname
		search.Finds()

		return c.JSON(search)
	})

	bankroute.Get("/cma/test", func(c *fiber.Ctx) error {

		search := bank.SearchCma{}

		search.Finds()

		companyCode_arr := []string{}
		companyCode_map := make(map[string]bool)

		for _, v := range search.Cmas {
			if _, exists := companyCode_map[v.Code]; !exists {
				companyCode_arr = append(companyCode_arr, v.Code)
				companyCode_map[v.Code] = true
			}
		}

		return c.JSON(companyCode_arr)

	})

}
