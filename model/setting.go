package model

type DataSource struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	Database string `yaml:"database"`
}

type Settings struct {
	Datasource DataSource    `yaml:"datasource"`
}
