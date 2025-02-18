package utils

import (
	"net/url"
	"regexp"
	"strings"
)

// HostStrip 将http://a.b.c:80/这种url去除不相关的字符，返回主机名
func HostStrip(u string) string {
	p, err := url.Parse(u)
	if err == nil && p.Host != "" {
		return p.Hostname()
	}
	host := strings.ReplaceAll(u, "https://", "")
	host = strings.ReplaceAll(host, "http://", "")
	host = strings.ReplaceAll(host, "/", "")
	host = strings.Split(host, ":")[0]

	return host
}

func CheckDomain(domain string) bool {
	domainPattern := `^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$`
	reg := regexp.MustCompile(strings.TrimSpace(domainPattern))
	return reg.MatchString(domain)
}
