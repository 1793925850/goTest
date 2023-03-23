package json2struct

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"tour/internal/word"
)

const (
	TYPEMAP_STRING_INTERFACE = "map[string]interface {}"
	TYPE_INTERFACE           = "[]interface {}"
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

	return strings.Join(append(p.Output, p.Children...), "\n")
}

func (p *Parser) toParentList(parentName string, parentValues []interface{}, isTop bool) {

}
