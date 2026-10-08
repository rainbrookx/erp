package tool

import (
	"image/color"

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
	// captchaDriver 验证码驱动单例，全局唯一
	captchaDriver base64Captcha.Driver
	// captchaStore 内存存储单例
	captchaStore base64Captcha.Store
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

	captchaDriver = driverString.ConvertFonts()
	captchaStore = base64Captcha.DefaultMemStore
}

func GenerateCaptcha() (id, b64s, answer string, err error) {
	captcha := base64Captcha.NewCaptcha(captchaDriver, captchaStore)
	id, b64s, answer, err = captcha.Generate()
	return
}

func VerifyCaptcha(id string, value string) bool {
	return captchaStore.Verify(id, value, true)
}
