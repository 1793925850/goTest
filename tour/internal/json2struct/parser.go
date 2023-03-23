package json2struct

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"tour/internal/word"
)

const (
	TYPE_MAP_STRING_INTERFACE = "map[string]interface {}"
	TYPE_INTERFACE            = "[]interface {}"
)

type Parser struct {
	Source     map[string]interface{}
	Output     Output
	Children   Output
	StructTag  string
	StructName string
}

type Output []string

// appendSegment 将 title 用函数处理，然后按照 format 的形式追加到 o 中
func (o *Output) appendSegment(format, title string, args ...interface{}) {
	s := []interface{}{}
	s = append(s, word.UnderscoreToUpperCamelCase(title))

	if len(args) != 0 {
		s = append(s, args...)
		format = "	" + format
	}

	*o = append(*o, fmt.Sprintf(format, s...))
}

// appendSuffix 添加后缀
func (o *Output) appendSuffix() {
	*o = append(*o, "}\n")
}

func NewParser(s string) (*Parser, error) {
	source := make(map[string]interface{})

	// Unmarshal 将 json编码型的数据进行解析并存到第二个参数里，第二个参数必须是指针
	// string 的底层是 []byte，因此可以互相转换
	if err := json.Unmarshal([]byte(s), &source); err != nil {
		return nil, err
	}

	return &Parser{
		Source:     source,
		StructTag:  "type %s struct {",
		StructName: "tour",
	}, nil
}

func (p *Parser) Json2Struct() string {
	p.Output.appendSegment(p.StructTag, p.StructName)

	for parentName, parentValues := range p.Source {
		valueType := reflect.TypeOf(parentValues).String()

		if valueType == TYPE_INTERFACE {
			p.toParentList(parentName, parentValues.([]interface{}), true)
		} else {
			var fields Fields

			fields.appendSegment(parentName, FieldSegment{
				Format: "%s",
				FieldValues: []FieldValue{
					{CamelCase: false, Value: valueType},
				},
			})
			p.Output.appendSegment("%s %s", fields[0].Name, fields[0].Type)
		}
	}
	p.Output.appendSuffix()

	// Join 将第一个参数里的每个元素之间都加入 sep ，并全部连接起来形成一个 string
	return strings.Join(append(p.Output, p.Children...), "\n")
}

func (p *Parser) toChildrenStruct(parentName string, values map[string]interface{}) {
	p.Children.appendSegment(p.StructTag, parentName)
	for fieldName, fieldValue := range values {
		p.Children.appendSegment("%s %s", fieldName, reflect.TypeOf(fieldValue).String())
	}
	p.Children.appendSuffix()
}

func (p *Parser) toParentList(parentName string, parentValues []interface{}, isTop bool) {
	var fields Fields

	for _, v := range parentValues {
		valueType := reflect.TypeOf(v).String()

		if valueType == TYPE_MAP_STRING_INTERFACE {
			fields = append(fields, p.handleParentTypeMapIface(v.(map[string]interface{}))...)
			p.Children.appendSegment(p.StructTag, parentName)

			for _, field := range fields.removeDuplicate() {
				p.Children.appendSegment("%s %s", field.Name, field.Type)
			}
			p.Children.appendSuffix()

			if isTop {
				valueType = word.UnderscoreToUpperCamelCase(parentName)
			}
		}

		if isTop {
			p.Output.appendSegment("%s %s%s", parentName, "[]", valueType)
		}
		break
	}
}

func (p *Parser) handleParentTypeMapIface(values map[string]interface{}) Fields {
	var fields Fields

	for fieldName, fieldValues := range values {
		var (
			fieldValueType = reflect.TypeOf(fieldValues).String()
			fieldSegment   = FieldSegment{
				Format: "%s",
				FieldValues: []FieldValue{
					{CamelCase: true, Value: fieldValueType},
				},
			}
		)

		switch fieldValueType {
		case TYPE_INTERFACE:
			fieldSegment = p.handleTypeIface(fieldName, fieldValues.([]interface{}))
		case TYPE_MAP_STRING_INTERFACE:
			fieldSegment = p.handleTypeMapIface(fieldName, fieldValues.(map[string]interface{}))
		}

		fields.appendSegment(fieldName, fieldSegment)
	}

	return fields
}

func (p *Parser) handleTypeIface(fieldName string, fieldValues []interface{}) FieldSegment {
	fieldSegment := FieldSegment{
		Format: "%s%s",
		FieldValues: []FieldValue{
			{CamelCase: false, Value: "[]"},
			{CamelCase: true, Value: fieldName},
		},
	}

	p.toParentList(fieldName, fieldValues, false)

	return fieldSegment
}

func (p *Parser) handleTypeMapIface(fieldName string, fieldValues map[string]interface{}) FieldSegment {
	fieldSegment := FieldSegment{
		Format: "%s",
		FieldValues: []FieldValue{
			{CamelCase: true, Value: fieldName},
		},
	}

	p.toChildrenStruct(fieldName, fieldValues)

	return fieldSegment
}
