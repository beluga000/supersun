package h

import (
	"github.com/gofiber/fiber/v2"
	"sunny.ksw.kr/repo/member"
)

func Member(route fiber.Router) {

	memberroute := route.Group("/member")

	memberroute.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("member")
	})

	memberroute.Post("/insert", func(c *fiber.Ctx) error {

		member := new(member.Member)

		c.BodyParser(member)

		errMsg := member.Create()
		if errMsg.Failure {
			return c.JSON(errMsg)
		}

		return c.JSON(member)

	})

}
