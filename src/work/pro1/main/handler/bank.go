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

		companyname := c.Query("companyname", "")

		search := bank.SearchInstalment_Savings{}

		search.Limit = limit
		search.Page = page
		search.PageOffset = page - 1
		search.CompanyName = companyname
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

}
