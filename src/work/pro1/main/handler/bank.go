package h

import (
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"sunny.ksw.kr/repo/bank"
)

func Bank(route fiber.Router) {

	bankroute := route.Group("/bank")

	// 적금 리스트
	bankroute.Get("/deposit/list", func(c *fiber.Ctx) error {

		limit, _ := strconv.Atoi(c.Query("limit", "10"))
		page, _ := strconv.Atoi(c.Query("page", "0"))

		bank_name := c.Query("bank_name", "")
		period := c.Query("period", "")
		categories := c.Query("categories", "")
		basic_rate_sort := c.Query("basic_rate_sort", "")
		max_rate_sort := c.Query("max_rate_sort", "")

		search := bank.SearchDeposit{}

		search.Limit = limit
		search.Page = page
		search.PageOffset = page - 1
		search.Bank_Name = bank_name
		search.Period = period
		if categories != "" {
			for _, v := range strings.Split(categories, ",") {
				switch v {
				case "특판":
					search.Categories = append(search.Categories, "specialOffer")
				case "방문없이가입":
					search.Categories = append(search.Categories, "online")
				case "정기적금":
					search.Categories = append(search.Categories, "savingFixed")
				case "자유적금":
					search.Categories = append(search.Categories, "savingFree")
				case "청년적금":
					search.Categories = append(search.Categories, "savingForYouth")
				case "청년도약계좌":
					search.Categories = append(search.Categories, "savingForJumpingYouth")
				case "군인적금":
					search.Categories = append(search.Categories, "savingForSoldier")
				case "주택청약":
					search.Categories = append(search.Categories, "housingSubscription")
				}
			}

			// search.Categories = strings.Split(categories, ",")
		}

		log.Print(search.Categories)

		search.Basic_Rate_Sort = basic_rate_sort
		search.Max_Rate_Sort = max_rate_sort
		search.Finds()

		return c.JSON(search)

	})

	// 적금 이율 상위 3개 리스트
	bankroute.Get("/deposit/top3", func(c *fiber.Ctx) error {

		search := bank.SearchDeposit{}

		search.Finds_Top3()

		return c.JSON(search)

	})

	// 적금 상세정보
	bankroute.Get("/deposit/get/:code", func(c *fiber.Ctx) error {

		type Result struct {
			Deposit        bank.Deposit
			Deposit_Detail bank.Deposit_Detail
		}

		code := c.Params("code")

		deposit, _ := bank.FindDepositByCode(code)

		deposit_detail, _ := bank.FindDeposit_DetailByCode(code)

		result := Result{
			Deposit:        deposit,
			Deposit_Detail: deposit_detail,
		}

		return c.JSON(result)

	})

	// 예금 리스트
	bankroute.Get("/instalment/list", func(c *fiber.Ctx) error {

		limit, _ := strconv.Atoi(c.Query("limit", "10"))
		page, _ := strconv.Atoi(c.Query("page", "0"))
		period := c.Query("period", "")
		categories := c.Query("categories", "")
		basic_rate_sort := c.Query("basic_rate_sort", "")
		max_rate_sort := c.Query("max_rate_sort", "")

		bank_name := c.Query("bank_name", "")

		search := bank.SearchInstalment_Savings{}

		search.Limit = limit
		search.Page = page
		search.PageOffset = page - 1
		search.Bank_Name = bank_name
		search.Period = period
		search.Basic_Rate_Sort = basic_rate_sort
		search.Max_Rate_Sort = max_rate_sort
		if categories != "" {
			for _, v := range strings.Split(categories, ",") {
				switch v {
				case "특판":
					search.Categories = append(search.Categories, "specialOffer")
				case "방문없이가입":
					search.Categories = append(search.Categories, "online")
				case "누구나가입":
					search.Categories = append(search.Categories, "anyone")
				case "정기적금":
					search.Categories = append(search.Categories, "savingFixed")
				case "자유적금":
					search.Categories = append(search.Categories, "savingFree")
				case "청년적금":
					search.Categories = append(search.Categories, "savingForYouth")
				case "청년도약계좌":
					search.Categories = append(search.Categories, "savingForJumpingYouth")
				case "군인적금":
					search.Categories = append(search.Categories, "savingForSoldier")
				case "주택청약":
					search.Categories = append(search.Categories, "housingSubscription")
				}
			}
		}

		search.Finds()

		return c.JSON(search)
	})

	bankroute.Get("/instalment/get/:code", func(c *fiber.Ctx) error {

		code := c.Params("code")

		type Result struct {
			Instalment_Savings        bank.Instalment_Savings
			Instalment_Savings_Detail bank.Instalment_Savings_Detail
		}

		instalment_saving, _ := bank.FindInstalment_SavingsByCode(code)

		instalmaent_saving_detail, _ := bank.FindInstalment_Savings_DetailByCode(code)

		result := Result{
			Instalment_Savings:        instalment_saving,
			Instalment_Savings_Detail: instalmaent_saving_detail,
		}

		return c.JSON(result)

	})

	bankroute.Get("/cma/test", func(c *fiber.Ctx) error {

		search := bank.SearchDeposit{}

		search.Finds()

		companyCode_arr := []string{}
		companyCode_map := make(map[string]bool)

		for _, v := range search.Deposits {
			if _, exists := companyCode_map[v.CompanyName]; !exists {
				companyCode_arr = append(companyCode_arr, v.CompanyName)
				companyCode_map[v.CompanyName] = true
			}
		}

		return c.JSON(companyCode_arr)

	})

	bankroute.Post("/recommand/deposit", func(c *fiber.Ctx) error {

		// 클라이언트에서 전달받은 데이터를 Recommand 구조체에 바인딩
		model := bank.Recommand{}

		// 바인딩 에러가 발생하면 에러 반환
		if err := c.BodyParser(&model); err != nil {
			return err
		}

		// 적금 상세정보 검색
		search := bank.SearchDeposit_Detail{
			Period:                  model.Period,
			Bank_name:               model.MainBank,
			Max_Rate_Sort:           "desc",
			Business:                model.Business,
			Children:                model.Children,
			Vulnerable_social_group: model.VulnerableSocialGroup,
			Young:                   model.Young,
			Soldier:                 model.Soldier,
			Old:                     model.Old,
		}
		search.Finds()

		// 추천 적금 리스트
		result := []bank.Recommand_Deposit{}

		// 총 만기금액 합계
		sumAmount := 0

		// 세금율 설정
		tax_rate := 0.154

		for _, v := range search.Deposit_Details {
			// 해당 적금 상품 월 최대 납입금액
			monthlyAmount := v.Amount_max
			// 해당 적금 상품 연간 최대 이율
			deposit_max_rate := v.Max_rate
			// 해당 적금 상품 연간 이율
			yearlyInterestRate := deposit_max_rate / 100
			// 해당 적금 상품 기간
			deposit_period := v.Product_period

			total_Principal := monthlyAmount * deposit_period

			var interest float64

			// 단리 이자 계산
			for month := 1; month <= deposit_period; month++ {
				interest += float64(monthlyAmount) * yearlyInterestRate * float64(deposit_period-month+1) / 12
			}

			tax := interest * tax_rate

			finalAmount := float64(total_Principal) + interest - tax

			sumAmount += int(finalAmount)

			if sumAmount > model.TargetAmount {
				break
			}

			result = append(result, bank.Recommand_Deposit{
				Deposit_Detail: *v,
				M_납입기간:         deposit_period,
				M_월납입금:         monthlyAmount,
				M_원금:           total_Principal,
				M_이자:           int(interest),
				M_세금:           int(tax),
				M_만기금액:         int(finalAmount),
			})

		}

		// 적금 개수 확인
		log.Print("적금 개수 : ", len(result))
		// 총 만기금액 확인
		log.Print("총 월 납입금액 : ", sumAmount)
		// 적금 이름 확인
		for _, v := range result {

			log.Print(v.Deposit_Detail.Product_name)
		}

		return c.JSON(result)

	})

	bankroute.Get("/deposit/list/test", func(c *fiber.Ctx) error {

		search := bank.SearchDeposit_Detail{}

		search.Max_Rate_Sort = "desc"

		search.Finds()

		return c.JSON(search)

	})

}
