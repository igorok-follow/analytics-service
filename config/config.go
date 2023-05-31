package config

type Config struct {
	Server    *Server    `yaml:"server"`
	Logger    *Logger    `yaml:"logger"`
	Gateway   *Gateway   `yaml:"gateway"`
	Analytics *Analytics `yaml:"analytics"`
	Tracing   *Tracing   `yaml:"tracing"`
	Database  *Database  `yaml:"database"`
}

type Server struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Gateway struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Logger struct {
	LogLevel string `yaml:"log_level"`
}

type Analytics struct {
	ApiKey string `yaml:"api_key"`
}

type Tracing struct {
	JaegerUri string `yaml:"jaeger_uri"`
}

type Database struct {
	Uri string `yaml:"uri"`
}
