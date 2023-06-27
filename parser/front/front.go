package front

import (
	"dynamic_json/parser/model"
	"errors"
	"github.com/buger/jsonparser"
)

// ????
//func init() {
//	model.ParsersType[model.Front] = map[model.KeyType]model.ParseFunc{
//		model.String:  ParseString,
//		model.Boolean: ParseBoolean,
//		model.Float:   ParseFloat,
//		model.Integer: ParseInteger,
//		model.Object:  ParseObject,
//	}
//}

var (
	SomeShitError = errors.New("some shit error [front]")
)

func ParseString(data []byte) (interface{}, error) {
	if v, err := jsonparser.GetBoolean(data, "userDefine"); err != nil && err != jsonparser.KeyPathNotFoundError {
		return "", err
	} else if v {
		return "()()()()()()()", nil
	}
	return jsonparser.GetString(data, "value")
}

func ParseBoolean(data []byte) (interface{}, error) {
	if v, err := jsonparser.GetBoolean(data, "userDefine"); err != nil && err != jsonparser.KeyPathNotFoundError {
		return "", err
	} else if v {
		return "()()()()()()()", nil
	}
	return jsonparser.GetBoolean(data, "value")
}

func ParseFloat(data []byte) (interface{}, error) {
	if v, err := jsonparser.GetBoolean(data, "userDefine"); err != nil && err != jsonparser.KeyPathNotFoundError {
		return "", err
	} else if v {
		return "()()()()()()()", nil
	}
	return jsonparser.GetFloat(data, "value")
}

func ParseInteger(data []byte) (interface{}, error) {
	if v, err := jsonparser.GetBoolean(data, "userDefine"); err != nil && err != jsonparser.KeyPathNotFoundError {
		return "", err
	} else if v {
		return "()()()()()()()", nil
	}
	return jsonparser.GetInt(data, "value")
}

func ParseObject(data []byte) (interface{}, error) {
	var obj []byte
	var errOut error
	// should check dateType (validate all this shit)
	if obj, _, _, errOut = jsonparser.Get(data, "props"); errOut != nil {
		return nil, SomeShitError
	}
	r := make(map[string]interface{})

	if errOut = jsonparser.ObjectEach(obj, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		var err error

		if r[string(key)], err = model.Parse(value, model.Front); err != nil {
			return err
		}

		return nil
	}); errOut != nil {
		return nil, errOut
	}
	return r, nil
}
