package config

import (
	"fmt"
	"log"
	"os"
	"reflect"
	configs "server/configs"
	"strconv"

	"gopkg.in/yaml.v3"
)

var AppConfiguration configs.Configuration

func Load() {
	configs := loadAppConfig()
	overrideConfigWithEnv(&configs)
	if configs.API.Host == "" {
		log.Fatalf("Error loading configurations: EMPTY API HOST FROM CONFIGURATION FILE")
	}
	if configs.API.Port <= 0 {
		log.Fatalf("Error loading configurations: INVALID API PORT IN CONFIGURATION FILE")
	}
	ValidateConfig(configs)
	AppConfiguration = configs
}

func overrideConfigWithEnv(configs *configs.Configuration) {
	if host := os.Getenv("API_HOST"); host != "" {
		configs.API.Host = host
	}

	if port := os.Getenv("API_PORT"); port != "" {
		if value, err := strconv.Atoi(port); err == nil {
			configs.API.Port = value
		}
	}

	if dsn := os.Getenv("DB_DSN"); dsn != "" {
		configs.DB_DSN = dsn
	}

	if secret := os.Getenv("JWT_SECRET_KEY"); secret != "" {
		configs.JWT.SecretKey = secret
	}

	if expire := os.Getenv("JWT_TOKEN_EXPIRE"); expire != "" {
		configs.JWT.TokenExpire = expire
	}
}

func loadAppConfig() configs.Configuration {
	targetFileConfig := fmt.Sprintf("%s/configs/app.yaml", os.Getenv(ROOT_FOLDER_VAR))
	file, errO := os.Open(targetFileConfig)
	if errO != nil {
		log.Fatalf("Error opening App yaml config file [%s]: %s\n", targetFileConfig, errO.Error())
	}
	defer file.Close()

	var configs configs.Configuration
	decoder := yaml.NewDecoder(file)
	err2 := decoder.Decode(&configs)
	if err2 != nil {
		log.Fatalf("Error parsing App yaml config file [%s]: %s\n", targetFileConfig, err2.Error())
	}
	return configs
}

func ValidateConfig(config interface{}) {
	v := reflect.ValueOf(config)
	t := reflect.TypeOf(config)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if field.Kind() == reflect.Struct {
			ValidateConfig(field.Interface())
		} else {
			if IsEmptyValue(field) {
				log.Fatalf("Error loading configurations: EMPTY %s FROM CONFIGURATION FILE", fieldType.Name)
			}
		}
	}
}

func IsEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String, reflect.Slice, reflect.Map, reflect.Array:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return v.Interface() == reflect.Zero(v.Type()).Interface()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	}
	return false
}
