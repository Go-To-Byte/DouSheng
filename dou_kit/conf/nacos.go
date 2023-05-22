package conf

type nacos struct {
	Host      string `mapstructure:"host" json:"host" yaml:"host"`
	Port      uint64 `mapstructure:"port" json:"port" yaml:"port"`
	Namespace string `mapstructure:"namespace" json:"namespace" yaml:"namespace"`
	User      string `mapstructure:"user" json:"user" yaml:"user"`
	Password  string `mapstructure:"password" json:"password" yaml:"password"`
	DataId    string `mapstructure:"dataid" json:"data_id" yaml:"data_id"`
	Group     string `mapstructure:"group" json:"group" yaml:"group"`
}

var Nacos = new(nacos)
