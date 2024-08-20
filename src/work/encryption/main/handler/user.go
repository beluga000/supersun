package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io"
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/argon2"
	"sunny.ksw.kr/repo/encryption"
)

var keyStore = map[string]*rsa.PrivateKey{} // 사용자의 비밀키를 임시로 저장하는 맵 (실제 구현에서는 안전하게 저장해야 함)

// AES 복호화 함수
func decryptAES(key []byte, ciphertext string) (string, error) {
	cipherText, _ := hex.DecodeString(ciphertext)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

// AES 암호화 함수
func encryptAES(key []byte, plaintext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return hex.EncodeToString(ciphertext), nil
}

func encryptRSA(publicKey *rsa.PublicKey, plaintext string) (string, error) {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(plaintext), nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Argon2 해싱 함수
func hashPasswordArgon2(password string) string {
	salt := []byte("somesalt") // 실제로는 랜덤한 솔트를 사용해야 함
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	return base64.RawStdEncoding.EncodeToString(hash)
}

// Argon2 비밀번호 검증 함수
func verifyPasswordArgon2(hashedPwd, plainPwd string) bool {
	salt := []byte("somesalt") // 해시할 때 사용한 솔트와 동일해야 함
	hash := argon2.IDKey([]byte(plainPwd), salt, 1, 64*1024, 4, 32)
	return base64.RawStdEncoding.EncodeToString(hash) == hashedPwd
}

// RSA 복호화 함수
func decryptRSA(privateKey *rsa.PrivateKey, ciphertext string) (string, error) {
	cipherText, _ := base64.StdEncoding.DecodeString(ciphertext)
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, cipherText, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func User(route fiber.Router) {

	userroute := route.Group("/user")

	// 유저 데이터 생성
	userroute.Post("/create", func(c *fiber.Ctx) error {

		model := encryption.User{}

		if err := c.BodyParser(&model); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		log.Print("입력받은 비밀번호 데이터 : ", model.Pwd)

		hashedPwd := hashPasswordArgon2(model.Pwd)

		model.Pwd = hashedPwd

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

		if !verifyPasswordArgon2(model.Pwd, login.Pwd) {
			return c.Status(401).JSON(fiber.Map{"msg": "비밀번호가 일치하지 않습니다."})
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
