package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
