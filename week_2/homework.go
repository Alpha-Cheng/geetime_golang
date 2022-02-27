package mian

import (
	"database/sql"
	"fmt"
	"strings"
)

var notFoundCode = 40001
var systemErr = 50001

func Biz2() error {
	err := Dao2("")
	if IsNoRow(err) {

	} else if err != nil {

	}
	return nil
}

func Dao2(query string) error {
	err := mockError()
	if err == sql.ErrNoRows {
		return fmt.Errorf("%d,not found", notFoundCode)
	} else if err != nil {
		return fmt.Errorf("%d not found", systemErr)
	}
}

func IsNoRow(err error) bool {
	return strings.HasPrefix(err.Error(), fmt.Sprintf("%d", notFoundCode))
}
