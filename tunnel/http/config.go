package http

import "github.com/p4gefau1t/trojan-go/config"

type Config struct {
	Username string `json:"auth_user" yaml:"auth-user"`
	Password string `json:"auth_pass" yaml:"auth-pass"`
}

func init() {
	config.RegisterConfigCreator(Name, func() interface{} {
		return &Config{}
	})
}
