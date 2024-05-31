package h

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"sunny.ksw.kr/repo/bank"
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
		maxAnnualFee, _ := strconv.Atoi(c.Query("maxAnnualFee", "0"))
		basement, _ := strconv.Atoi(c.Query("basement", "0"))

		search := card.SearchCard{}
		search.Limit = limit
		search.Page = page
		search.PageOffset = page - 1
		search.Code = code
		search.MaxAnnualFee = maxAnnualFee
		if benefits != "" {
			search.Benefits = strings.Split(benefits, ",")
		}
		search.Basement = basement

		search.Finds()

		for _, v := range search.Cards {

			switch v.CompanyCode {
			case "SS":
				v.CompanyName = "삼성카드"
			case "KB":
				v.CompanyName = "국민카드"
			case "SH":
				v.CompanyName = "신한카드"
			case "HD":
				v.CompanyName = "현대카드"
			case "LO":
				v.CompanyName = "롯데카드"
			case "SK":
				v.CompanyName = "하나카드"
			case "WR":
				v.CompanyName = "우리카드"
			case "NH":
				v.CompanyName = "농협카드"
			case "IB":
				v.CompanyName = "기업카드"
			}
		}

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

	cardroute.Get("/test", func(c *fiber.Ctx) error {

		search := bank.SearchParking{}

		search.Finds()

		companyCode_arr := []string{}
		companyCode_map := make(map[string]bool)

		for _, v := range search.Parkings {
			if _, exists := companyCode_map[v.Code]; !exists {
				companyCode_arr = append(companyCode_arr, v.Code)
				companyCode_map[v.Code] = true
			}
		}

		return c.JSON(companyCode_arr)

	})

	cardroute.Get("/test/max", func(c *fiber.Ctx) error {

		search := card.SearchCard{}

		search.Finds()

		maxAnnualFee := 0

		for _, v := range search.Cards {

			if v.DomesticAnnualFee > maxAnnualFee {
				maxAnnualFee = v.DomesticAnnualFee
			}

		}

		return c.JSON(maxAnnualFee)

	})

	// cardroute.Get("/get/:id", func(c *fiber.Ctx) error {

	// 	id := c.Params("id")

	// 	card, _ := card.FindCardById(id)

	// 	return c.JSON(card)

	// })

}
