package models

import (
	"net"
	"net/url"
	"strings"
)

type DnsmasqRule struct {
	Domain string
	DNS    string
}

func ParseDnsmasqRule(rule string) *DnsmasqRule {
	/** server=/0-6.com/114.114.114.114 */

	// 解析规则
	// 1. 以 server= 开头
	// 2. 以 ip 结尾
	// 3. 中间是域名

	splits := strings.Split(rule, "/")

	if len(splits) != 3 {
		return nil
	}

	if splits[0] != "server=" {
		return nil
	}

	if _, err := url.Parse(splits[1]); err != nil {
		return nil
	}

	if ip := net.ParseIP(splits[2]); ip == nil {
		return nil
	}

	return &DnsmasqRule{
		Domain: splits[1],
		DNS:    splits[2],
	}
}
