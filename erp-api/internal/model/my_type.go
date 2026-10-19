package model

import (
	"database/sql/driver"
	"fmt"
)

// MyBitBool 适配 mysql bit(1)
type MyBitBool bool

func (b *MyBitBool) Scan(v any) error {
	if v == nil {
		*b = false
		return nil
	}
	var result bool
	switch val := v.(type) {
	case []byte:
		result = len(val) > 0 && val[0] == 1
	case int64:
		result = val == 1
	case bool:
		result = val
	default:
		return fmt.Errorf("bit(1) scan: unsupported type %T", v)
	}
	*b = MyBitBool(result)
	return nil
}

func (b *MyBitBool) Value() (driver.Value, error) {
	if *b {
		return int64(1), nil
	}
	return int64(0), nil
}

// NullMyBitBool 可空版本（nullable bit(1)）
type NullMyBitBool struct {
	MyBitBool
	Valid bool
}

func (n *NullMyBitBool) Scan(v any) error {
	if v == nil {
		n.Valid = false
		return nil
	}
	n.Valid = true
	return n.MyBitBool.Scan(v)
}

func (n *NullMyBitBool) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.MyBitBool.Value()
}
