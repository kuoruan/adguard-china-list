package service

import (
	"bufio"
	"net/http"

	"go.kuoruan.net/adguard-upstream/models"
	"golang.org/x/sync/errgroup"
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

func NewChinaList() *ChinaList {
	return &ChinaList{
		Client: http.Client{},
	}
}

func (s *ChinaList) Update() error {
	var g errgroup.Group

	for _, url := range []urlType{
		{url: urls[0], listType: AcceleratedDomains},
		{url: urls[1], listType: Apple},
		{url: urls[2], listType: Google},
	} {
		j := url

		g.Go(func() error {
			req, err := http.NewRequest("GET", j.url, nil)

			if err != nil {
				return err
			}

			req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

			res, err := s.Client.Do(req)
			if err != nil {
				return err
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
				return err
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

			return nil
		})
	}

	return g.Wait()
}

func (s *ChinaList) Transform() {

}
