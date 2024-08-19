package handler

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"sunny.ksw.kr/repo/mokdon/bank"
	"sunny.ksw.kr/repo/mokdon/card"
)

// Card API Group
func Card(route fiber.Router) {
	cardroute := route.Group("/card")

	// @Summary 카드 리스트 가져오기
	// @Description 카드 리스트를 가져옵니다.
	// @Tags Card
	// @Accept json
	// @Produce json
	// @Param limit query int false "페이지당 카드 수"  // 기본값 10
	// @Param page query int false "페이지 번호"  // 기본값 0
	// @Param code query string false "카드 코드"
	// @Param benefits query string false "카드 혜택 (콤마로 구분)"
	// @Param maxAnnualFee query int false "최대 연회비"
	// @Param basement query int false "기준"
	// @Param annualFeeSort query string false "연회비 정렬 방식 (asc, desc)"
	// @Param basementSort query string false "기준 정렬 방식 (asc, desc)"
	// @Success 200 {array} card.Card "카드 리스트"
	// @Router /api/v1/card/list [get]
	cardroute.Get("/list", func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "10"))
		page, _ := strconv.Atoi(c.Query("page", "0"))

		code := c.Query("code", "")
		benefits := c.Query("benefits", "")
		maxAnnualFee, _ := strconv.Atoi(c.Query("maxAnnualFee", "0"))
		basement, _ := strconv.Atoi(c.Query("basement", "0"))
		annualFeeSort := c.Query("annualFeeSort", "")
		basementSort := c.Query("basementSort", "")

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
		search.AnnualFeeSort = annualFeeSort
		search.BasementSort = basementSort
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
			case "BC":
				v.CompanyName = "비씨카드"
			}
		}

		return c.JSON(search)
	})

	// @Summary 카드 정보 가져오기
	// @Description 카드 ID로 카드 정보를 가져옵니다. 카드의 기본 정보와 추가 정보를 반환합니다.
	// @Tags Card
	// @Accept json
	// @Produce json
	// @Param card_id path string true "카드 ID"
	// @Success 200 {object} Result "카드 정보"
	// @Failure 404 {object} string "카드를 찾을 수 없음"
	// @Router /api/v1/card/get/{card_id} [get]
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

	// @Summary 카드 정보 목록 가져오기
	// @Description 모든 카드의 정보 리스트를 가져옵니다.
	// @Tags Card
	// @Accept json
	// @Produce json
	// @Success 200 {array} card.Card_Info "카드 정보 리스트"
	// @Router /api/v1/card/info/list [get]
	cardroute.Get("/info/list", func(c *fiber.Ctx) error {
		search := card.SearchCard_Info{}
		search.Finds()
		return c.JSON(search)
	})

	// @Summary 주차장 코드 목록 가져오기
	// @Description 주차장 코드를 가져옵니다. 중복을 제거하여 리스트로 반환합니다.
	// @Tags Card
	// @Accept json
	// @Produce json
	// @Success 200 {array} string "주차장 코드 리스트"
	// @Router /api/v1/card/test [get]
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

	// @Summary 최대 연회비 가져오기
	// @Description 카드 목록에서 최대 연회비를 계산하여 가져옵니다.
	// @Tags Card
	// @Accept json
	// @Produce json
	// @Success 200 {integer} integer "최대 연회비"
	// @Router /api/v1/card/test/max [get]
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
}
