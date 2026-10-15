package main

import (
	"path/filepath"
	"strings"

	"github.com/rainbrookx/erp/internal/infrastructure/config"
	"github.com/rainbrookx/erp/internal/infrastructure/database"
	"gorm.io/gen"
	"gorm.io/gen/field"
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
		name := strings.TrimPrefix(tableName, tablePrefix)
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
	//model := g.GenerateModel("jsh_account")
	//g.ApplyBasic(model)

	// 全部表一次性生成
	//g.ApplyBasic(g.GenerateAllTable()...)

	// 为了兼容！！！！！！jsh_tenant.enabled 的 DEFAULT b'1'
	// `enabled` bit(1) DEFAULT b'1' COMMENT '启用 0-禁用  1-启用',
	tables, err := db.Migrator().GetTables()
	if err != nil {
		panic("get tables error: " + err.Error())
	}

	for _, tbl := range tables {
		if !strings.HasPrefix(tbl, tablePrefix) {
			continue
		}
		model := g.GenerateModel(tbl,
			gen.FieldGORMTag("enabled", func(tag field.GormTag) field.GormTag {
				delete(tag, "default")
				return tag
			}),
		)
		g.ApplyBasic(model)
	}

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
