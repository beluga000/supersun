package handler

import (
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"sunny.ksw.kr/repo/mokdon/bank"
)

func calculateDepositDetails(v bank.Deposit_Detail, monthlyAmount int, tax_rate float64, totalPeriod int) (total_Principal int, interest float64, tax float64, finalAmount float64, rejoinCount int) {
	yearlyInterestRate := v.Max_rate / 100
	deposit_period := v.Product_period
	total_Principal = monthlyAmount * totalPeriod
	rejoinCount = (totalPeriod + deposit_period - 1) / deposit_period // 반올림하여 재가입 횟수 계산

	for month := 1; month <= totalPeriod; month++ {
		interest += float64(monthlyAmount) * yearlyInterestRate * float64(totalPeriod-month+1) / 12
	}

	tax = interest * tax_rate
	finalAmount = float64(total_Principal) + interest - tax
	return
}

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
		model := bank.Recommand{}

		if err := c.BodyParser(&model); err != nil {
			return err
		}

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

		result := []bank.Recommand_Deposit{}
		sumAmount := 0
		tax_rate := 0.154

		period, _ := strconv.Atoi(model.Period)

		if model.MonthlyAmount == 0 {

			// 총 납입 원금
			total_Sum := 0
			// 총 세전이자
			total_Interest := 0
			// 총 세금
			total_Tax := 0
			// 총 만기금액
			total_FinalAmount := 0
			// 한달에 내야하는 월 적금 금액
			month_amount := 0

			for _, v := range search.Deposit_Details {
				total_Principal, interest, tax, finalAmount, rejoinCount := calculateDepositDetails(*v, v.Amount_max, tax_rate, period)
				sumAmount += int(finalAmount)

				deposit, _ := bank.FindDepositByCode(v.Code)

				total_Sum += total_Principal
				total_Interest += int(interest)
				total_Tax += int(tax)
				total_FinalAmount += int(finalAmount)

				result = append(result, bank.Recommand_Deposit{
					Deposit:        deposit,
					Deposit_Detail: *v,
					M_납입기간:         v.Product_period,
					M_월납입금:         v.Amount_max,
					M_원금:           total_Principal,
					M_이자:           int(interest),
					M_세금:           int(tax),
					M_만기금액:         int(finalAmount),
					M_재투자횟수:        rejoinCount,
				})

				if sumAmount > model.TargetAmount {
					break
				}
				month_amount += v.Amount_max
			}

			log.Print("월 납입 가능 금액이 0인 경우")
			log.Print("한달에 내야하는 월 적금 금액 : ", month_amount)

			return c.JSON(fiber.Map{
				"result":             result,
				"total_sum":          total_Sum,
				"total_interest":     total_Interest,
				"total_tax":          total_Tax,
				"total_final_amount": total_FinalAmount,
				"month_amount":       month_amount,
			})

		} else {

			log.Print("------월 납입 가능 금액이 0이 아닌 경우------")
			log.Print("입력한 목표 금액 : ", model.TargetAmount)
			log.Print("입력한 월 납입 금액 : ", model.MonthlyAmount)

			// 목표 기간  * 월 납입 가능 금액
			remainingAmount := model.MonthlyAmount

			// 총 원금 합계
			totalFinalAmount := 0

			// 총 원금 합계
			total_Sum := 0
			// 총 세전이자 합계
			total_Interest := 0
			// 총 세금 합계
			total_Tax := 0
			// 총 만기금액 합계
			total_FinalAmount := 0
			month_amount := 0

			for i, v := range search.Deposit_Details {
				if remainingAmount <= 0 {
					log.Print("남은 금액이 0 이하일 경우 종료")
					break
				}

				monthlyAmount := v.Amount_max
				if remainingAmount < monthlyAmount {
					monthlyAmount = remainingAmount
				}

				log.Print(i, "번째 상품 월 납입 금액 : ", monthlyAmount)

				total_Principal, interest, tax, finalAmount, rejoinCount := calculateDepositDetails(*v, monthlyAmount, tax_rate, period)
				totalFinalAmount += int(finalAmount)
				deposit, _ := bank.FindDepositByCode(v.Code)

				log.Printf("%d번째 상품 만기 수령액 : %.2f", i, finalAmount)

				total_Sum += total_Principal
				total_Interest += int(interest)
				total_Tax += int(tax)
				total_FinalAmount += int(finalAmount)

				result = append(result, bank.Recommand_Deposit{
					Deposit:        deposit,
					Deposit_Detail: *v,
					M_납입기간:         v.Product_period,
					M_월납입금:         monthlyAmount,
					M_원금:           total_Principal,
					M_이자:           int(interest),
					M_세금:           int(tax),
					M_만기금액:         int(finalAmount),
					M_재투자횟수:        rejoinCount,
				})

				month_amount += monthlyAmount
				remainingAmount -= monthlyAmount

				log.Print(i, "번째 상품 월 납입 금액을 제외한 남은 금액 : ", remainingAmount)
			}

			log.Print("목표 금액 중 월 납입 가능 금액으로 달성한 금액 : ", totalFinalAmount)

			// 월 납입 가능 금액으로 목표 금액을 달성한 경우
			if totalFinalAmount >= model.TargetAmount {
				log.Print("월 납입 가능 금액으로 목표 금액을 달성한 경우")
				log.Print("한달에 내야하는 월 적금 금액 : ", model.MonthlyAmount)

				return c.JSON(fiber.Map{
					"result":             result,
					"total_sum":          total_Sum,
					"total_interest":     total_Interest,
					"total_tax":          total_Tax,
					"total_final_amount": total_FinalAmount,
					"month_amount":       month_amount,
				})

			} else {
				// 추가로 필요한 월 납입금과 적합한 상품 계산

				log.Print("목표 금액 중 월 납입 가능 금액으로 달성하지 못한 경우")
				log.Print("목표 금액 : ", model.TargetAmount, " - 월 납입 가능 금액으로 가입한 상품의 만기 총합 : ", totalFinalAmount)

				neededAmount := model.TargetAmount - totalFinalAmount
				additionalDeposits := []bank.Recommand_Deposit{}

				additionalMonthlyTotal := 0
				additionalFinalAmount := 0

				extra_total_Sum := 0
				extra_total_Interest := 0
				extra_total_Tax := 0
				extra_total_FinalAmount := 0

				extra_month_amount := 0

				for _, v := range search.Deposit_Details {
					if neededAmount <= 0 {
						break
					}

					isAlreadyIncluded := false
					for _, r := range result {
						if r.Deposit_Detail.ID == v.ID {
							isAlreadyIncluded = true
							break
						}
					}
					if isAlreadyIncluded {
						continue
					}

					monthlyAmount := v.Amount_max
					if neededAmount < monthlyAmount {
						monthlyAmount = neededAmount
					}

					total_Principal, interest, tax, finalAmount, rejoinCount := calculateDepositDetails(*v, monthlyAmount, tax_rate, period)
					deposit, _ := bank.FindDepositByCode(v.Code)

					additionalDeposits = append(additionalDeposits, bank.Recommand_Deposit{
						Deposit:        deposit,
						Deposit_Detail: *v,
						M_납입기간:         v.Product_period,
						M_월납입금:         monthlyAmount,
						M_원금:           total_Principal,
						M_이자:           int(interest),
						M_세금:           int(tax),
						M_만기금액:         int(finalAmount),
						M_재투자횟수:        rejoinCount,
					})
					neededAmount -= int(finalAmount)
					additionalMonthlyTotal += monthlyAmount
					additionalFinalAmount += int(finalAmount)

					extra_total_Sum += total_Principal
					extra_total_Interest += int(interest)
					extra_total_Tax += int(tax)
					extra_total_FinalAmount += int(finalAmount)
					extra_month_amount += monthlyAmount
				}

				log.Print("월 납입 가능 금액이 0이 아닌 경우 + 제출한 월 납입 가능 금액으로 목표 금액을 달성하지 못한 경우")
				log.Print("납입 가능 금액 : ", model.MonthlyAmount)
				log.Print("추가로 필요한 월 납입 금액 : ", additionalMonthlyTotal)

				return c.JSON(fiber.Map{
					"result":                   result,
					"total_sum":                total_Sum,
					"total_interest":           total_Interest,
					"total_tax":                total_Tax,
					"total_final_amount":       total_FinalAmount,
					"month_amount":             month_amount,
					"extra_deposit":            additionalDeposits,
					"extra_total_sum":          extra_total_Sum,
					"extra_total_interest":     extra_total_Interest,
					"extra_total_tax":          extra_total_Tax,
					"extra_total_final_amount": extra_total_FinalAmount,
					"extra_month_amount":       extra_month_amount,
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
