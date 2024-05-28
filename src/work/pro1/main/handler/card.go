package h

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"sunny.ksw.kr/repo/card"
)

func Card(route fiber.Router) {

	cardroute := route.Group("/card")

	// cardroute.Get("/test", func(c *fiber.Ctx) error {
	// 	return c.SendString("member")
	// })

	cardroute.Get("/list", func(c *fiber.Ctx) error {

		limit, _ := strconv.Atoi(c.Query("limit", "10"))
		page, _ := strconv.Atoi(c.Query("page", "0"))

		search := card.SearchCard{}
		search.Limit = limit
		search.Page = page
		search.PageOffset = page - 1

		search.Finds()

		return c.JSON(search)

	})

	cardroute.Get("/test", func(c *fiber.Ctx) error {

		search := card.SearchCard{}

		search.Finds()

		companyCode_arr := []string{}
		companyCode_map := make(map[string]bool)

		for _, v := range search.Cards {
			if _, exists := companyCode_map[v.CompanyCode]; !exists {
				companyCode_arr = append(companyCode_arr, v.CompanyCode)
				companyCode_map[v.CompanyCode] = true
			}
		}

		return c.JSON(companyCode_arr)

	})

	cardroute.Get("/get/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")

		card, _ := card.FindCardById(id)

		return c.JSON(card)

	})

}
