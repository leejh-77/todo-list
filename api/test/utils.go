package test

import (
	"strconv"
	"time"
)

func UniqueString(str string) string {
	return str + strconv.FormatInt(time.Now().UnixNano(), 10)
}

