package util

import (
	"image/color"
	"strings"
	"uuid"

	"github.com/mojocn/base64Captcha"
)

// Captcha 验证码相关配置常量
const (
	captchaLength = 4
	captchaWidth  = 105
	captchaHeight = 35
	captchaSource = "123456789abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
)

var (
	// rawDriver 原始的驱动
	rawDriver *base64Captcha.DriverString

	// captchaDriver 验证码驱动
	captchaDriver base64Captcha.Driver

	// captchaStore 内存存储 todo 改成 redis
	captchaStore base64Captcha.Store

	// captcha 验证码生成工具
	captcha *base64Captcha.Captcha
)

func init() {
	bgColor := &color.RGBA{R: 200, G: 200, B: 200, A: 200}

	driverString := base64Captcha.DriverString{
		Length:          captchaLength,
		Width:           captchaWidth,
		Height:          captchaHeight,
		Source:          captchaSource,
		BgColor:         bgColor,
		ShowLineOptions: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine,
	}

	rawDriver = driverString.ConvertFonts()
	captchaDriver = &CustomUUIDDriver{
		DriverString: rawDriver,
	}
	captchaStore = base64Captcha.DefaultMemStore
	captcha = base64Captcha.NewCaptcha(captchaDriver, captchaStore)
}

func GenerateCaptcha() (id, b64s, answer string, err error) {
	id, b64s, answer, err = captcha.Generate()
	return
}

func VerifyCaptcha(id, answer string) bool {
	return captchaStore.Verify(id, answer, true)
}

// === Start CustomUUIDDriver ===

type CustomUUIDDriver struct {
	*base64Captcha.DriverString
}

func (d *CustomUUIDDriver) GenerateIdQuestionAnswer() (id, content, answer string) {
	_, content, answer = d.DriverString.GenerateIdQuestionAnswer()
	id = strings.ReplaceAll(uuid.New().String(), "-", "")
	return
}

// === End CustomUUIDDriver ===
