package {{.PackageName}}.entity;

{{if .HasDecimalType}}
import java.math.BigDecimal;
{{- end -}}
{{if .HasDateType}}
import java.time.LocalDateTime;
{{- end}}

public class {{upperCamelCase .TableName}} {
	{{- range $column := .Columns}}

	/**
	 * {{.ColumnComment}}
	 */
	private {{.JavaType}} {{lowerCamelCase .ColumnName}};
	{{end}}

    {{- range $column := .Columns}}
	public {{.JavaType}} get{{upperCamelCase .ColumnName}}() {
		return {{lowerCamelCase .ColumnName}};
	}

	public void set{{upperCamelCase .ColumnName}}({{.JavaType}} {{lowerCamelCase .ColumnName}}) {
		this.{{lowerCamelCase .ColumnName}} = {{lowerCamelCase .ColumnName}};
	}
	{{end}}
}