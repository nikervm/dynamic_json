package parser

import (
	"dynamic_json/parser/front"
	"dynamic_json/parser/model"
	"encoding/json"
	"errors"
	"github.com/buger/jsonparser"
)

var (
	InvalidJson = errors.New("invalid json")
)

func init() {
	model.ParsersType = make(map[model.Type]map[model.KeyType]model.ParseFunc)
	model.ParsersType[model.Front] = map[model.KeyType]model.ParseFunc{
		model.String:  front.ParseString,
		model.Boolean: front.ParseBoolean,
		model.Float:   front.ParseFloat,
		model.Integer: front.ParseInteger,
		model.Object:  front.ParseObject,
	}
}

type Parser struct {
	source []byte
}

func Init(data []byte) (*Parser, error) {
	if json.Valid(data) {
		return &Parser{source: data}, nil
	}
	return nil, InvalidJson
}

func (p *Parser) ToFront() ([]byte, error) {
	var errOut error
	var d []byte
	resultJson := make(map[string]interface{})

	// high level of json
	if errOut = jsonparser.ObjectEach(p.source, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		var err error

		if resultJson[string(key)], err = model.Parse(value, model.Front); err != nil {
			return err
		}

		return nil
	}); errOut != nil {
		return nil, errOut
	}

	if d, errOut = json.Marshal(resultJson); errOut != nil {
		return nil, errOut
	}
	return d, nil
}
