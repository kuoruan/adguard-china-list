package service

import (
	"bufio"
	"net/http"
	"sync"

	"go.kuoruan.net/adguard-upstream/models"
)

var (
	urls = []string{
		"https://raw.githubusercontent.com/felixonmars/dnsmasq-china-list/master/accelerated-domains.china.conf",
		"https://raw.githubusercontent.com/felixonmars/dnsmasq-china-list/master/apple.china.conf",
		"https://raw.githubusercontent.com/felixonmars/dnsmasq-china-list/master/google.china.conf",
	}
)

type listType int

const (
	AcceleratedDomains listType = iota
	Apple
	Google
)

type urlType struct {
	url      string
	listType listType
}

type ChinaList struct {
	Client http.Client

	AcceleratedDomains []string
	Apple              []string
	Google             []string
}

func NewChinaListService() *ChinaList {
	return &ChinaList{
		Client: http.Client{},
	}
}

func (s *ChinaList) Update() error {
	var wg sync.WaitGroup

	for _, url := range []urlType{
		{url: urls[0], listType: AcceleratedDomains},
		{url: urls[1], listType: Apple},
		{url: urls[2], listType: Google},
	} {
		wg.Add(1)

		j := url

		go func() {
			defer wg.Done()

			req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/felixonmars/dnsmasq-china-list/master/accelerated-domains.china.conf", nil)

			if err != nil {
				return
			}

			req.Header.Set("User-Agent", "curl/7.64.1")

			res, err := s.Client.Do(req)
			if err != nil {
				return
			}

			defer res.Body.Close()

			m := make(map[string]interface{}, 10000)

			sc := bufio.NewScanner(res.Body)
			for sc.Scan() {
				line := sc.Text()

				if rule := models.ParseDnsmasqRule(line); rule != nil {
					m[rule.Domain] = nil
				}
			}

			if err := sc.Err(); err != nil {
				return
			}

			list := make([]string, 0, len(m))

			// map to slice
			for k := range m {
				list = append(list, k)
			}

			switch j.listType {
			case AcceleratedDomains:
				s.AcceleratedDomains = list
			case Apple:
				s.Apple = list
			case Google:
				s.Google = list
			}
		}()
	}

	wg.Wait()

	return nil
}

func (s *ChinaList) Transform() {

}
