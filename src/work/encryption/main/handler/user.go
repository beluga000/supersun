package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"sunny.ksw.kr/repo/encryption"
)

func User(route fiber.Router) {

	userroute := route.Group("/user")

	// 유저 데이터 생성
	userroute.Post("/create", func(c *fiber.Ctx) error {

		model := encryption.User{}

		if err := c.BodyParser(&model); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		log.Print("입력받은 비밀번호 데이터 : ", model.Pwd)

		bytes, err := bcrypt.GenerateFromPassword([]byte(model.Pwd), 14)
		if err != nil {
			return err
		}

		model.Pwd = string(bytes)

		log.Print("암호화된 비밀번호 데이터 : ", model.Pwd)

		errMsg := model.Create()
		if errMsg.Failure {
			return c.JSON(errMsg)
		}

		return c.JSON(model)

	})

	// 로그인 확인
	userroute.Post("/login", func(c *fiber.Ctx) error {

		type Login struct {
			Id  string `json:"id"`
			Pwd string `json:"pwd"`
		}

		login := Login{}

		if err := c.BodyParser(&login); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		model, _ := encryption.FindUserData(login.Id)

		err := bcrypt.CompareHashAndPassword([]byte(model.Pwd), []byte(login.Pwd))
		if err != nil {
			return c.JSON(fiber.Map{"msg": "비밀번호가 일치하지 않습니다."})
		}

		return c.JSON(fiber.Map{"msg": "로그인 성공"})

	})

	// 유저 데이터 조회
	userroute.Get("/read/:id", func(c *fiber.Ctx) error {

		model := encryption.User{}

		errMsg := model.GetById(c.Params("id"))
		if errMsg.Failure {
			return c.JSON(errMsg)
		}

		return c.JSON(model)

	})
}
