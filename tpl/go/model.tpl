package model

type {{upperCamelCase .TableName}} struct{
 	{{- range $column := .Columns}}
 	{{- if not .IsDelStatus}}
	{{upperCamelCase .ColumnName}} {{.GoLangType}} `json:"{{lowerCamelCase .ColumnName}}"` // {{.ColumnComment}} 
	{{- end -}}
 	{{end}}
}