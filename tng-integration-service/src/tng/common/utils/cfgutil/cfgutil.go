package cfgutil

import (
	"os"
	"strconv"

	"github.com/astaxie/beego"
)

// Load gets config from ENV or file.
func Load(key string) string {
	val := os.Getenv(key)
	if val == "" {
		val = beego.AppConfig.String(key)
	}

	return val
}

// LoadInt loads configuration as Int.
func LoadInt(key string, value ...int) (int, error) {
	val := Load(key)
	if val == "" && len(value) > 0 {
		return value[0], nil
	}
	return strconv.Atoi(val)
}

// LoadFloat loads configuration as Float.
func LoadFloat(key string, def ...float64) float64 {
	valStr := Load(key)
	valFloat64, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		if len(def) > 0 {
			return def[0]
		}
		return 0
	}
	return valFloat64
}

// LoadBool loads configuration as Boolean.
func LoadBool(key string, def ...bool) bool {
	var (
		valStr       = Load(key)
		valBool, err = strconv.ParseBool(valStr)
	)
	if err != nil {
		if len(def) > 0 {
			return def[0]
		}
		return false
	}
	return valBool
}
