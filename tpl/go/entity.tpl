package entity

type {{upperCamelCase .TableName}} struct{
 	{{range $column := .Columns -}}
 	{{upperCamelCase .ColumnName }} {{.GoLangType}} {{.Tag}} 
 	{{end -}}
}