


### 대칭키 AES 암호화 함수

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


userroute.Post("/create", func(c *fiber.Ctx) error {
    model := encryption.User{}

    if err := c.BodyParser(&model); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    log.Print("입력받은 비밀번호 데이터 : ", model.Pwd)

    key := []byte("thisis32bitlongpassphraseimusing") // 32바이트 키 (AES-256)
    encryptedPwd, err := encryptAES(key, model.Pwd)
    if err != nil {
        return err
    }

    model.Pwd = encryptedPwd

    log.Print("암호화된 비밀번호 데이터 : ", model.Pwd)

    errMsg := model.Create()
    if errMsg.Failure {
        return c.JSON(errMsg)
    }

    return c.JSON(model)
})

### 대칭키 AES 복호화 함수


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

    key := []byte("thisis32bitlongpassphraseimusing")
    decryptedPwd, err := decryptAES(key, model.Pwd)
    if err != nil || decryptedPwd != login.Pwd {
        return c.Status(401).JSON(fiber.Map{"msg": "비밀번호가 일치하지 않습니다."})
    }

    return c.JSON(fiber.Map{"msg": "로그인 성공"})
})

### 비대칭키 RSA 암호화 함수

func encryptRSA(publicKey *rsa.PublicKey, plaintext string) (string, error) {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(plaintext), nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}


userroute.Post("/create", func(c *fiber.Ctx) error {
    model := encryption.User{}

    if err := c.BodyParser(&model); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    log.Print("입력받은 비밀번호 데이터 : ", model.Pwd)

    publicKey, _ := rsa.GenerateKey(rand.Reader, 2048) // RSA 키 생성
    encryptedPwd, err := encryptRSA(&publicKey.PublicKey, model.Pwd)
    if err != nil {
        return err
    }

    model.Pwd = encryptedPwd

    log.Print("암호화된 비밀번호 데이터 : ", model.Pwd)

    errMsg := model.Create()
    if errMsg.Failure {
        return c.JSON(errMsg)
    }

    return c.JSON(model)
})

### 비대칭키 RSA 복호화 함수

func decryptRSA(privateKey *rsa.PrivateKey, ciphertext string) (string, error) {
	cipherText, _ := base64.StdEncoding.DecodeString(ciphertext)
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, cipherText, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

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

    privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
    decryptedPwd, err := decryptRSA(privateKey, model.Pwd)
    if err != nil || decryptedPwd != login.Pwd {
        return c.Status(401).JSON(fiber.Map{"msg": "비밀번호가 일치하지 않습니다."})
    }

    return c.JSON(fiber.Map{"msg": "로그인 성공"})
})


### 해싱 Argon2 함수

func hashPasswordArgon2(password string) string {
	salt := []byte("somesalt") // 실제로는 랜덤한 솔트를 사용해야 함
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	return base64.RawStdEncoding.EncodeToString(hash)
}

userroute.Post("/create", func(c *fiber.Ctx) error {
    model := encryption.User{}

    if err := c.BodyParser(&model); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    log.Print("입력받은 비밀번호 데이터 : ", model.Pwd)

    hashedPwd := hashPasswordArgon2(model.Pwd)

    model.Pwd = hashedPwd

    log.Print("해싱된 비밀번호 데이터 : ", model.Pwd)

    errMsg := model.Create()
    if errMsg.Failure {
        return c.JSON(errMsg)
    }

    return c.JSON(model)
})


### 해싱 Argon2 검증 함수

func verifyPasswordArgon2(hashedPwd, plainPwd string) bool {
	salt := []byte("somesalt") // 해시할 때 사용한 솔트와 동일해야 함
	hash := argon2.IDKey([]byte(plainPwd), salt, 1, 64*1024, 4, 32)
	return base64.RawStdEncoding.EncodeToString(hash) == hashedPwd
}

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