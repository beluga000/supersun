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

	bankroute.Get("/instalment/top3", func(c *fiber.Ctx) error {

		search := bank.SearchInstalment_Savings{}

		search.Finds_Top3()

		return c.JSON(search)

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

		if model.MonthlyAmount == 0 {

			test := 0

			// 총 납입 원금
			total_Sum := 0
			// 총 세전이자
			total_Interest := 0
			// 총 세금
			total_Tax := 0
			// 총 만기금액
			total_FinalAmount := 0

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

				deposit, _ := bank.FindDepositByCode(v.Code)

				// 총 납입 원금
				total_Sum += total_Principal
				// 총 세전이자
				total_Interest += int(interest)
				// 총 세금
				total_Tax += int(tax)
				// 총 만기금액
				total_FinalAmount += int(finalAmount)

				result = append(result, bank.Recommand_Deposit{
					Deposit:        deposit,
					Deposit_Detail: *v,
					M_납입기간:         deposit_period,
					M_월납입금:         monthlyAmount,
					M_원금:           total_Principal,
					M_이자:           int(interest),
					M_세금:           int(tax),
					M_만기금액:         int(finalAmount),
				})

				test += v.Amount_max

			}

			log.Print("월 납입 가능 금액이 0인 경우")
			log.Print("한달에 내야하는 월 적금 금액 : ", test)

			// return c.JSON(result)

			return c.JSON(fiber.Map{
				"result":  result,
				"총 납입 원금": total_Sum,
				"총 세전 이자": total_Interest,
				"총 세금":    total_Tax,
				"총 만기금액":  total_FinalAmount,
			})

		} else {
			// 월 납입금액이 0이 아닌 경우
			remainingAmount := model.MonthlyAmount
			totalFinalAmount := 0

			// 총 원금 합계
			total_Sum := 0
			// 총 세전이자 합계
			total_Interest := 0
			// 총 세금 합계
			total_Tax := 0
			// 총 만기금액 합계
			total_FinalAmount := 0

			for _, v := range search.Deposit_Details {
				if remainingAmount <= 0 {
					break
				}

				// 해당 적금 상품 월 최대 납입금액
				monthlyAmount := v.Amount_max
				if remainingAmount < monthlyAmount {
					monthlyAmount = remainingAmount
				}

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
				totalFinalAmount += int(finalAmount)
				deposit, _ := bank.FindDepositByCode(v.Code)

				// 총 납입 원금
				total_Sum += total_Principal
				// 총 세전이자
				total_Interest += int(interest)
				// 총 세금
				total_Tax += int(tax)
				// 총 만기금액
				total_FinalAmount += int(finalAmount)

				result = append(result, bank.Recommand_Deposit{
					Deposit:        deposit,
					Deposit_Detail: *v,
					M_납입기간:         deposit_period,
					M_월납입금:         monthlyAmount,
					M_원금:           total_Principal,
					M_이자:           int(interest),
					M_세금:           int(tax),
					M_만기금액:         int(finalAmount),
				})

				remainingAmount -= monthlyAmount
			}

			// 목표 금액 달성 여부 확인
			if totalFinalAmount >= model.TargetAmount {

				log.Print("월 납입 가능 금액이 0이 아닌 경우")
				log.Print("한달에 내야하는 월 적금 금액 : ", model.MonthlyAmount)

				return c.JSON(fiber.Map{
					"result":  result,
					"총 납입 원금": total_Sum,
					"총 세전 이자": total_Interest,
					"총 세금":    total_Tax,
					"총 만기금액":  total_FinalAmount,
				})

			} else {
				// 추가로 필요한 월 납입금과 적합한 상품 계산
				neededAmount := model.TargetAmount - totalFinalAmount
				additionalDeposits := []bank.Recommand_Deposit{}
				remainingAmount = model.MonthlyAmount
				additionalMonthlyTotal := 0
				additionalFinalAmount := 0

				for _, v := range search.Deposit_Details {
					if neededAmount <= 0 {
						break
					}

					// 기존에 선택된 상품은 제외
					isAlreadyIncluded := false
					for _, r := range result {
						if r.Deposit_Detail.ID == v.ID { // 고유한 식별자 필드를 비교
							isAlreadyIncluded = true
							break
						}
					}
					if isAlreadyIncluded {
						continue
					}

					monthlyAmount := v.Amount_max
					if remainingAmount < monthlyAmount {
						monthlyAmount = remainingAmount
					}

					deposit_max_rate := v.Max_rate
					yearlyInterestRate := deposit_max_rate / 100
					deposit_period := v.Product_period
					total_Principal := monthlyAmount * deposit_period
					var interest float64
					for month := 1; month <= deposit_period; month++ {
						interest += float64(monthlyAmount) * yearlyInterestRate * float64(deposit_period-month+1) / 12
					}
					tax := interest * tax_rate
					finalAmount := float64(total_Principal) + interest - tax
					deposit, _ := bank.FindDepositByCode(v.Code)
					additionalDeposits = append(additionalDeposits, bank.Recommand_Deposit{
						Deposit:        deposit,
						Deposit_Detail: *v,
						M_납입기간:         deposit_period,
						M_월납입금:         monthlyAmount,
						M_원금:           total_Principal,
						M_이자:           int(interest),
						M_세금:           int(tax),
						M_만기금액:         int(finalAmount),
					})
					neededAmount -= int(finalAmount)
					remainingAmount -= monthlyAmount
					additionalMonthlyTotal += monthlyAmount
					additionalFinalAmount += int(finalAmount)
				}

				log.Print("월 납입 가능 금액이 0이 아닌 경우 + 제출한 월 납입 가능 금액으로 목표 금액을 달성하지 못한 경우")
				log.Print("납입 가능 금액 : ", model.MonthlyAmount)
				log.Print("추가로 필요한 월 납입 금액 : ", additionalMonthlyTotal)

				return c.JSON(fiber.Map{
					"result":        result,
					"추가 상품 만기 총합":   additionalFinalAmount,
					"추가 상품 월 납입 금액": additionalMonthlyTotal,
					"추가 상품":         additionalDeposits,
				})
			}
		}

		return c.JSON("해당 적금 상품이 없습니다.")
	})

	bankroute.Get("/deposit/list/test", func(c *fiber.Ctx) error {

		search := bank.SearchDeposit_Detail{}

		search.Max_Rate_Sort = "desc"

		search.Finds()

		return c.JSON(search)

	})

}
