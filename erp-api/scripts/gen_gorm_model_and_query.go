package main

import (
	"path/filepath"
	"strings"

	"github.com/rainbrookx/erp/internal/infrastructure/config"
	"github.com/rainbrookx/erp/internal/infrastructure/database"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// DSN 连接你的mysql
	config.InitConfig()
	database.InitDatabase(config.C.Database)

	db := database.DB

	cwd, err := filepath.Abs(".")
	if err != nil {
		panic("获取当前工作目录失败: " + err.Error())
	}

	queryPath := filepath.Join(cwd, "internal", "query")
	modelPath := filepath.Join(cwd, "internal", "model")

	g := gen.NewGenerator(gen.Config{
		OutPath:          queryPath,
		ModelPkgPath:     modelPath,
		Mode:             gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithGeneric, // generate mode
		FieldNullable:    true,
		FieldCoverable:   false,
		FieldSignable:    false,
		FieldWithTypeTag: true,
	})
	// 表前缀，按实际情况替换
	const tablePrefix = "jsh_"

	// 指定 Model 名称的生成策略
	g.WithModelNameStrategy(func(tableName string) string {
		name := strings.TrimPrefix(tableName, "jsh_")
		if name == "" {
			return ""
		}
		// 下划线转大驼峰
		parts := strings.Split(name, "_")
		for i, p := range parts {
			if p == "" {
				continue
			}
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
		return strings.Join(parts, "")
	})

	g.Config.WithJSONTagNameStrategy(snakeToLowerCamel)

	g.UseDB(db)

	// ✅ bit 类型映射核心配置
	dataMap := map[string]func(columnType gorm.ColumnType) string{
		"bit": func(ct gorm.ColumnType) string {
			rawType, _ := ct.ColumnType()
			nullable, _ := ct.Nullable()
			// 只处理 bit(1)
			if strings.HasPrefix(rawType, "bit(1)") {
				if nullable {
					return "NullMyBitBool"
				}
				return "MyBitBool"
			}
			// 其他 bit(N) 你自己决定，比如返回 []byte
			return "[]byte"
		},
	}
	g.WithDataTypeMap(dataMap)

	// 👉 指定表名，生成model；多个表写多次
	//model := g.GenerateModel("jsh_platform_config")
	//g.ApplyBasic(model)

	// 全部表一次性生成
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}

// 将下划线分隔的列名转为小驼峰
func snakeToLowerCamel(s string) string {
	if s == "" {
		return s
	}
	parts := strings.Split(s, "_")
	for i := 1; i < len(parts); i++ {
		if parts[i] == "" {
			continue
		}
		parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
	}
	return strings.Join(parts, "")
}
