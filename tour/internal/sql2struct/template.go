package sql2struct

import (
	"fmt"
	"html/template"
	"os"

	"tour/internal/word"
)

// gt 是 >，“大于”的意思
const structTpl = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
	{{ $typeLen := len .Type}} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string // 结构体某个属性的属性名
	Type    string // 该属性的类型
	Tag     string // 该属性的 json 表示
	Comment string // 该属性的注释
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

// NewStructTemplate 新建结构体模板
func NewStructTemplate() *StructTemplate {
	return &StructTemplate{
		structTpl: structTpl,
	}
}

// AssemblyColumns 将数据库中每个列对应的列名、数据类型、标签、注释更换成适用于 Go 的形式
func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))

	for _, column := range tbColumns {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName)

		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}

	return tplColumns
}

// Generate 根据结构体数组 tplColumns 和表名 tableName ，最终生成结构体
func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}

	err := tpl.Execute(os.Stdout, tplDB) // 这里是结构体生成的位置，目前是控制台(os.Stdout)
	if err != nil {
		return err
	}

	return nil
}
