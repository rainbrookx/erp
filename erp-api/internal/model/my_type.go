package model

import (
	"database/sql/driver"
	"errors"
)

// MyBitBool 适配 mysql bit(1)
type MyBitBool bool

func (b *MyBitBool) Scan(v any) error {
	if v == nil {
		*b = false
		return nil
	}
	bs, ok := v.([]byte)
	if !ok {
		return errors.New("bit(1) scan: expect []byte")
	}
	if len(bs) == 0 {
		*b = false
		return nil
	}
	*b = bs[0] == 1
	return nil
}

func (b *MyBitBool) Value() (driver.Value, error) {
	if *b {
		return []byte{1}, nil
	}
	return []byte{0}, nil
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
