package h

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"sunny.ksw.kr/repo/card"
)

func Card(route fiber.Router) {

	cardroute := route.Group("/card")

	// cardroute.Get("/test", func(c *fiber.Ctx) error {
	// 	return c.SendString("member")
	// })

	// 카드 리스트 조회
	cardroute.Get("/list", func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "10"))
		page, _ := strconv.Atoi(c.Query("page", "0"))

		code := c.Query("code", "")
		benefits := c.Query("benefits", "")

		search := card.SearchCard{}
		search.Limit = limit
		search.Page = page
		search.PageOffset = page - 1
		search.Code = code
		if benefits != "" {
			search.Benefits = strings.Split(benefits, ",")
		}

		search.Finds()

		return c.JSON(search)
	})

	cardroute.Get("/get/:card_id", func(c *fiber.Ctx) error {

		card_id := c.Params("card_id")

		card_info, _ := card.FindCard_InfoByCardID(card_id)

		card_, _ := card.FindCardByCardID(card_id)

		type Result struct {
			Card     card.Card
			CardInfo card.Card_Info
		}

		result := Result{
			Card:     card_,
			CardInfo: card_info,
		}

		return c.JSON(result)

	})

	cardroute.Get("/info/list", func(c *fiber.Ctx) error {

		search := card.SearchCard_Info{}

		search.Finds()

		return c.JSON(search)
	})

	// cardroute.Get("/test", func(c *fiber.Ctx) error {

	// 	search := card.SearchCard{}

	// 	search.Finds()

	// 	companyCode_arr := []string{}
	// 	companyCode_map := make(map[string]bool)

	// 	for _, v := range search.Cards {
	// 		if _, exists := companyCode_map[v.CompanyCode]; !exists {
	// 			companyCode_arr = append(companyCode_arr, v.CompanyCode)
	// 			companyCode_map[v.CompanyCode] = true
	// 		}
	// 	}

	// 	return c.JSON(companyCode_arr)

	// })

	// cardroute.Get("/get/:id", func(c *fiber.Ctx) error {

	// 	id := c.Params("id")

	// 	card, _ := card.FindCardById(id)

	// 	return c.JSON(card)

	// })

}
