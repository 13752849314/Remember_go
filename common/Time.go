package common

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

const DateTimeFormat = "2006-01-02 15:04:05"
const DateFormat = "2006-01-02"

type MyDateTime time.Time

type MyDate time.Time

func (m *MyDateTime) UnmarshalJSON(bytes []byte) (err error) {
	now, err := time.ParseInLocation(`"`+DateTimeFormat+`"`, string(bytes), time.Local)
	*m = MyDateTime(now)
	return
}

func (m MyDateTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(m)
	if tt.IsZero() {
		emptyStr, _ := json.Marshal("")
		return emptyStr, nil
	}
	b := make([]byte, 0, len(DateTimeFormat)+2)
	b = append(b, '"')
	b = tt.AppendFormat(b, DateTimeFormat)
	b = append(b, '"')
	return b, nil
}

func (m MyDateTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tt := time.Time(m)
	if tt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tt, nil
}

func (m *MyDateTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*m = MyDateTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (m MyDateTime) String() string {
	tt := time.Time(m)
	return tt.String()
}

// ==============================

func (m *MyDate) UnmarshalJSON(bytes []byte) (err error) {
	now, err := time.ParseInLocation(`"`+DateFormat+`"`, string(bytes), time.Local)
	*m = MyDate(now)
	return
}

func (m MyDate) MarshalJSON() ([]byte, error) {
	tt := time.Time(m)
	if tt.IsZero() {
		emptyStr, _ := json.Marshal("")
		return emptyStr, nil
	}
	b := make([]byte, 0, len(DateFormat)+2)
	b = append(b, '"')
	b = tt.AppendFormat(b, DateFormat)
	b = append(b, '"')
	return b, nil
}

func (m MyDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	tt := time.Time(m)
	if tt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tt, nil
}

func (m *MyDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*m = MyDate(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (m MyDate) String() string {
	tt := time.Time(m)
	return tt.String()
}
