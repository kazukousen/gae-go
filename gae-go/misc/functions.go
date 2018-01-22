package misc

import (
	"reflect"
	"strconv"
	"strings"
)

// ZeroOrNil checks if the argument is zero or null
func ZeroOrNil(obj interface{}) bool {
	value := reflect.ValueOf(obj)
	if !value.IsValid() {
		return true
	}
	if obj == nil {
		return true
	}
	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		return value.Len() == 0
	}
	zero := reflect.Zero(reflect.TypeOf(obj))
	if obj == zero.Interface() {
		return true
	}
	return false
}

// Split returns splitted strings
// @param s string, sep string, n int
func Split(s string, sep string, n int) []string {
	splitted := strings.SplitN(s, sep, n)
	parsed := make([]string, len(splitted))
	for i, val := range splitted {
		parsed[i] = strings.TrimSpace(val)
	}
	return parsed
}

// Uint16 returns casted uint16
// @oaram s string
func Uint16(s string) uint16 {
	var result uint16
	if s != "" {
		if u, err := strconv.ParseUint(s, 10, 16); err == nil {
			result = uint16(u)
		}
	}
	return result
}

// Atoi returns casted int
// @oaram s string
func Atoi(s string) int {
	var result int
	if s != "" {
		if i, err := strconv.Atoi(s); err == nil {
			result = i
		}
	}
	return result
}

// Bool returns casted bool
// @oaram s string
func Bool(s string) bool {
	result := false
	if s != "" {
		if b, err := strconv.ParseBool(s); err == nil {
			result = b
		}
	}
	return result
}
