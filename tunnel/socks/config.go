package socks

import "github.com/p4gefau1t/trojan-go/config"

type Config struct {
	LocalHost  string `json:"local_addr" yaml:"local-addr"`
	LocalPort  int    `json:"local_port" yaml:"local-port"`
	UDPTimeout int    `json:"udp_timeout" yaml:"udp-timeout"`
	Username   string `json:"auth_user" yaml:"auth-user"`
	Password   string `json:"auth_pass" yaml:"auth-pass"`
}

func init() {
	config.RegisterConfigCreator(Name, func() interface{} {
		return &Config{
			UDPTimeout: 60,
		}
	})
}
