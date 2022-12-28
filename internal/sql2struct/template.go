package sql2struct

import (
	"fmt"
	"os"
	"text/template"

	"github.com/geekr-dev/go-cli-app/internal/word"
)

const structTpl = `
type {{.TableName | ToCamelCase}} struct {
{{range .Columns}} {{$length := len .Comment}} {{if gt $length 0}} //
{{.Comment}} {{else}}// {{.Name}} {{end}}
	{{$typeLen := len .Type}} {{if gt $typeLen 0}}{{.Name | ToCamelCase}}
	{{.Type}} {{.Tag}} {{else}} {{.Name}} {{end}}
{{end}}}

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTamplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTamplate {
	return &StructTamplate{structTpl: structTpl}
}

func (t *StructTamplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     fmt.Sprintf("`json:%s`", column.ColumnName),
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

func (t *StructTamplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.SnakeToCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}
	return nil
}
