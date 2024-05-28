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

}
