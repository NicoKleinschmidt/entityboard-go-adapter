package adapter

import (
	"fmt"
	"reflect"
	"strings"
)

var settingsParseMap = map[string]interface{}{
	"string":    "",
	"number":    int64(0),
	"boolean":   false,
	"object":    map[string]interface{}{},
	"file":      "",
	"string[]":  []string{},
	"number[]":  []int64{},
	"boolean[]": []bool{},
	"object[]":  []map[string]interface{}{},
	"file[]":    []string{},
}

// ParseSettingsTemplate generates the settings struct from the template.
// This can be used for example to validate the json sent by the server.
func ParseSettingsTemplate(settings interface{}) (interface{}, error) {
	return parseSettingsTemplate(reflect.ValueOf(settings))
}

func parseSettingsTemplate(settingsValue reflect.Value) (interface{}, error) {
	settingsType := settingsValue.Type()

	if settingsType.Kind() != reflect.Map {
		return nil, fmt.Errorf("settings has to be a map")
	}

	if settingsType.Key().Kind() != reflect.String {
		return nil, fmt.Errorf("settings key must be string")
	}

	if settingsType.Elem().Kind() != reflect.Interface {
		return nil, fmt.Errorf("settings value must be interface{}")
	}

	fields := []reflect.StructField{}
	iter := settingsValue.MapRange()

	for iter.Next() {
		k := iter.Key().String()
		v := iter.Value().Elem()
		fieldType, err := getType(v)

		if err != nil {
			return nil, err
		}

		fields = append(fields, reflect.StructField{
			Name: strings.ReplaceAll(k, " ", "_"),
			Type: fieldType,
		})
	}

	return reflect.New(reflect.StructOf(fields)).Interface(), nil
}

func getType(value reflect.Value) (reflect.Type, error) {
	if value.Kind() == reflect.String {
		fieldVal, err := parseString(value.String())

		if err != nil {
			return nil, err
		}

		return reflect.TypeOf(fieldVal), nil
	} else if value.Kind() == reflect.Map {
		valueType, err := parseSettingsTemplate(value)

		if err != nil {
			return nil, err
		}

		return reflect.TypeOf(valueType), nil
	} else {
		return nil, fmt.Errorf("map value must be string or map, is %v", value.Kind())
	}
}

func parseString(k string) (interface{}, error) {
	valueType, ok := settingsParseMap[k]

	if ok {
		return valueType, nil
	}

	if strings.HasPrefix(k, "enum:") {
		return parseEnum(k)
	} else if strings.HasPrefix(k, "flag:") {
		return parseFlag(k)
	} else {
		return nil, fmt.Errorf("%s is not a valid value", k)
	}
}

func parseEnum(enum string) (interface{}, error) {
	return int64(0), nil
}

func parseFlag(flag string) (interface{}, error) {
	return int64(0), nil
}
