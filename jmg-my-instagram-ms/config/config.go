package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Type     string `mapstructure:"type"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
		SSLMode  string `yaml:"sslmode"`
	} `mapstructure:"database"`

	Server struct {
		Port  string `mapstructure:"port"`
		Debug bool   `mapstructure:"debug"`
	} `mapstructure:"server"`
}

var Configuration *Config

func LoadConfig(path string) (*Config, error) {

	// Configurar Viper
	viper.AddConfigPath(path)     // Directorio donde está el archivo de configuración
	viper.SetConfigName("config") // Nombre del archivo de configuración (sin la extensión)
	viper.SetConfigType("yaml")   // Tipo de archivo de configuración (yaml en este caso)

	// Leer la configuración
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error al leer el archivo de configuración: %w", err)
	}

	// Unmarshal a la estructura de configuración
	if err := viper.Unmarshal(&Configuration); err != nil {
		return nil, fmt.Errorf("error al deserializar configuración: %w", err)
	}

	return Configuration, nil
}
