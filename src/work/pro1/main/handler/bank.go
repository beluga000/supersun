package h

import (
	"github.com/gofiber/fiber/v2"
	"sunny.ksw.kr/repo/bank"
)

func Bank(route fiber.Router) {

	bankroute := route.Group("/bank")

	bankroute.Get("/get", func(c *fiber.Ctx) error {

		search := bank.SearchDeposit_Detail{}

		search.Finds()

		return c.JSON(search)
	})

	bankroute.Get("/get2/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")

		deposit_detail, _ := bank.FindDeposit_DetailById(id)

		return c.JSON(deposit_detail)

	})

	bankroute.Get("/instalment/list", func(c *fiber.Ctx) error {

		search := bank.SearchInstalment_Savings{}

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
