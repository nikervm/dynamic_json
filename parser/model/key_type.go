package model

import (
	"errors"
	"fmt"
	"github.com/buger/jsonparser"
)

type KeyType string

var (
	UnknownParserType = errors.New("unknown parser type")
)

const (
	String  KeyType = "string"
	Integer KeyType = "integer"
	Float   KeyType = "float"
	Object  KeyType = "object"
	Boolean KeyType = "boolean"
)

func GetKeyType(d []byte) KeyType {
	keyType, err := jsonparser.GetString(d, "type")
	if err != nil {
		return ""
	}

	return KeyType(keyType)
}

type ParseFunc func(data []byte) (interface{}, error)

type Type uint8

const (
	Front = iota
	Back
)

var ParsersType = map[Type]map[KeyType]ParseFunc{}

func Parse(data []byte, pt Type) (interface{}, error) {
	parsers := ParsersType[pt]
	if parsers == nil {
		return nil, UnknownParserType
	}

	cKeyType := GetKeyType(data)

	if f, exist := parsers[cKeyType]; !exist {
		fmt.Printf("not specified [%s]\n", cKeyType)
		//return nil, UnknownParserType
	} else {
		return f(data)
	}
	return nil, nil
}
