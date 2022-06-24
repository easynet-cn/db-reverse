package main

import (
	"fmt"
	"strings"
)

type DataColumn struct {
	TableSchema            string
	TableName              string
	ColumnName             string
	OrdinalPosition        int
	ColumnDefault          string
	IsNullable             string
	DataType               string
	CharacterMaximumLength string
	CharacterOctetLength   string
	NumericPrecision       string
	NumbericScale          string
	DatetimePrecision      string
	ColumnType             string
	ColumnKey              string
	Extra                  string
	ColumnComment          string
}

func (c *DataColumn) IsIdentity() bool {
	return strings.ToLower(c.Extra) == "auto_increment"
}

func (c *DataColumn) IsPrimary() bool {
	return strings.ToLower(c.ColumnKey) == "pri"
}

func (c *DataColumn) GoLangType() string {
	typeMapping := GetConfiguration().TypeMapping
	dataType := strings.ToLower(c.DataType)
	nullable := strings.ToLower(c.IsNullable) == "yes"

	if langType, ok := typeMapping[dataType]; ok && langType.Go != "" {
		if nullable {
			return fmt.Sprintf("*%s", langType.Go)
		}

		return langType.Go
	}

	if nullable {
		return fmt.Sprintf("*%s", dataType)
	}

	return dataType
}

func (c *DataColumn) JavaType() string {
	typeMapping := GetConfiguration().TypeMapping
	dataType := strings.ToLower(c.DataType)

	if langType, ok := typeMapping[dataType]; ok && langType.Java != "" {
		return langType.Java
	}

	return dataType
}

func (c *DataColumn) IsDateType() bool {
	return strings.ToLower(c.DataType) == "datetime"
}

func (c *DataColumn) IsDecimalType() bool {
	return strings.ToLower(c.DataType) == "decimal"
}

func (c *DataColumn) IsDelStatus() bool {
	return c.ColumnName == "del_status"
}

func (c *DataColumn) Tag() string {
	name := strings.ToLower(c.ColumnName)
	dataType := strings.ToLower(c.DataType)
	identity := strings.ToLower(c.Extra) == "auto_increment"
	primary := strings.ToLower(c.ColumnKey) == "pri"
	nullable := strings.ToLower(c.IsNullable) == "yes"
	ormTag := GetConfiguration().OrmTag

	sb := new(strings.Builder)

	sb.WriteString(fmt.Sprintf("`%s:\"", ormTag))
	sb.WriteString(dataType)
	sb.WriteString(" '")
	sb.WriteString(name)
	sb.WriteString("'")

	if identity {
		sb.WriteString(" autoincr")
	}

	if primary {
		sb.WriteString(" pk")
	}

	if nullable {
		sb.WriteString(" null")
	} else {
		sb.WriteString(" notnull")
	}

	sb.WriteString(" default(")

	if dataType == "varchar" || dataType == "text" || dataType == "longtext" {
		sb.WriteString("'")
	}

	sb.WriteString(c.ColumnDefault)

	if dataType == "varchar" || dataType == "text" || dataType == "longtext" {
		sb.WriteString("'")
	}

	sb.WriteString(")")

	sb.WriteString(" comment('")
	sb.WriteString(c.ColumnComment)
	sb.WriteString("')")

	sb.WriteString("\" json:\"")

	if name == "del_status" {
		sb.WriteString("-")
	} else {
		sb.WriteString(LowerCamelCase(c.ColumnName))
	}

	sb.WriteString("\"`")

	return sb.String()
}
