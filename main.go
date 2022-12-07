package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"unsafe"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

const (
	selectCurrentDbSql = "SELECT DATABASE()"
	allColumnInfoSql   = "SELECT * FROM information_schema.columns WHERE table_schema =? ORDER BY table_schema ASC,table_name ASC,ordinal_position ASC"
)

func main() {
	config := GetConfiguration()

	engine, err := xorm.NewEngine("mysql", config.Datasource)

	if err := engine.Ping(); err != nil {
		fmt.Println("can not create database engine,err:", err)

		return
	}

	if err != nil {
		fmt.Println("can not create database engine,err:", err)

		return
	}

	currentDb := ""

	if _, err := engine.SQL(selectCurrentDbSql).Get(&currentDb); err != nil {
		fmt.Println("can not get current database,err:", err)

		return
	}

	columns := make([]DataColumn, 0)

	if err := engine.SQL(allColumnInfoSql, currentDb).Find(&columns); err != nil {
		fmt.Println("can not get column information,err:", err)

		return
	}

	tableMap := make(map[string][]DataColumn)

	for _, column := range columns {
		tableName := column.TableName

		if _, ok := tableMap[tableName]; !ok {
			tableMap[tableName] = make([]DataColumn, 0)
		}

		tableMap[tableName] = append(tableMap[tableName], column)
	}

	funcMap := template.FuncMap{"upperCamelCase": UpperCamelCase, "lowerCamelCase": LowerCamelCase}

	tplName := fmt.Sprintf("%s.tpl", strings.ToLower(config.TargetType))
	tplFile := fmt.Sprintf("tpl/%s/%s.tpl", strings.ToLower(config.Lang), strings.ToLower(config.TargetType))

	if config.TplName != "" {
		tplName = config.TplName
	}

	if config.TplFile != "" {
		tplFile = config.TplFile
	}

	t, err := template.New(tplName).Funcs(funcMap).ParseFiles(tplFile)

	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}

	if err := os.RemoveAll(config.Output); err != nil {
		fmt.Println(err)

		return
	}

	for table, columns := range tableMap {
		if _, err := os.Stat(config.Output); os.IsNotExist(err) {
			if err := os.MkdirAll(config.Output, 0777); err != nil {
				fmt.Println("create out directory err:", err)

				return
			}
		}

		fileSb := new(strings.Builder)

		fileSb.WriteString(config.Output)
		fileSb.WriteString("/")

		if config.Lang == "go" {
			fileSb.WriteString(table)
		} else if config.Lang == "java" {
			if config.TargetType == "entity" {
				fileSb.WriteString(UpperCamelCase(table))
			} else if config.TargetType == "model" {
				fileSb.WriteString(UpperCamelCase(table))
			}
		}

		fileSb.WriteString(".")
		fileSb.WriteString(config.Lang)

		f, err := os.OpenFile(fileSb.String(), os.O_CREATE|os.O_WRONLY, 0666)

		defer f.Close()

		if err != nil {
			fmt.Println("can not create output file,err:", err)

			return
		}

		if err := t.Execute(f, &Config{TableName: table, Readonly: config.Readonly, PackageName: config.PackageName, Columns: columns}); err != nil {
			fmt.Println("There was an error:", err.Error())
		}
	}

	if err := exec.Command("gofmt", "-w", config.Output).Run(); err != nil {
		fmt.Println("format source  error:", err.Error())
	}

}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
