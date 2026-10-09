package util

import (
	"context"
	"errors"
	"image/color"
	"log/slog"
	"strings"
	"time"
	"uuid"

	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
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

	// captchaStore Redis 存储
	captchaStore base64Captcha.Store

	// captcha 验证码生成工具
	captcha *base64Captcha.Captcha
)

func InitRandomImageUtil(rdb *redis.Client) {
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
	captchaStore = NewRedisStore(rdb, "captcha_codes:", 2*time.Minute)
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

// === Start RedisStore ===

// RedisStore 实现 base64Captcha.Store 接口
type RedisStore struct {
	client     *redis.Client
	expiration time.Duration
	keyPrefix  string
}

// NewRedisStore 创建一个新的 RedisStore
func NewRedisStore(client *redis.Client, prefix string, expiration time.Duration) *RedisStore {
	return &RedisStore{
		client:     client,
		expiration: expiration,
		keyPrefix:  prefix,
	}
}

func (s *RedisStore) Set(id string, value string) error {
	ctx := context.Background()
	key := s.keyPrefix + id
	return s.client.Set(ctx, key, value, s.expiration).Err()
}

// Get 从 Redis 获取验证码答案，并可根据 clear 参数决定是否删除
func (s *RedisStore) Get(id string, clear bool) string {
	ctx := context.Background()
	key := s.keyPrefix + id
	val, err := s.client.Get(ctx, key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			slog.Error("验证码缓存到Redis", "err", err)
		}
		return "" // 如果键不存在或出错，返回空字符串
	}

	if clear {
		if err := s.client.Del(ctx, key).Err(); err != nil {
			slog.Error("验证码缓存到Redis", "err", err)
		}
	}
	return val
}

// Verify 验证用户提交的答案是否正确
func (s *RedisStore) Verify(id, answer string, clear bool) bool {
	// 直接调用 Get 方法，传入 clear 参数
	storedAnswer := s.Get(id, clear)
	return storedAnswer == answer
}

// === End RedisStore ===
