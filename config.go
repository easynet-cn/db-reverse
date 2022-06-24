package main

type Config struct {
	TableName   string
	Readonly    bool
	PackageName string
	Columns     []DataColumn
}

func (c *Config) PrimaryColumnDataType() string {
	for _, column := range c.Columns {
		if column.IsPrimary() {
			return column.JavaType()
		}
	}

	return ""
}

func (c *Config) HasDelStatus() bool {
	for _, column := range c.Columns {
		if column.IsDelStatus() {
			return true
		}
	}

	return false
}

func (c *Config) HasDecimalType() bool {
	for _, column := range c.Columns {
		if column.IsDecimalType() {
			return true
		}
	}

	return false
}

func (c *Config) HasDateType() bool {
	for _, column := range c.Columns {
		if column.IsDateType() {
			return true
		}
	}

	return false
}

func (c *Config) HasEnterpriseId() bool {
	return c.HasColumn("enterprise_id")
}

func (c *Config) HasCode() bool {
	return c.HasColumn("code")
}

func (c *Config) HasStatus() bool {
	return c.HasColumn("status")
}

func (c *Config) HasColumn(name string) bool {
	for _, column := range c.Columns {
		if column.ColumnName == name {
			return true
		}
	}

	return false
}
