package service

import (
	"behappy-url-shortener/src/model"
	"github.com/sirupsen/logrus"
	url2 "net/url"
	"regexp"
)

func GetShortenUrl(r *model.GetShortenRequest) (error, string) {
	return nil, "nil"
}

func checkUrl(url string, checkDomain bool) bool {
	valid := true
	regexPattern := "/^(ftp|http|https):\\/\\/(\\w+:{0,1}\\w*@)?(\\S+)(:[0-9]+)?(\\/|\\/([\\w#!:.?+=&%@!\\-\\/]))?/"
	matched, _ := regexp.MatchString(regexPattern, "")
	if !matched {
		valid = false
	}
	if matched && checkDomain {
		loc, loc_err := url2.Parse(url)
		runOps, ori_err := url2.Parse(model.RunOpts.Url)
		if loc_err != nil && ori_err != nil {
			logrus.Error("解析url错误")
			valid = false
		}
		// 不解析当前域下的地址
		if loc.Hostname() == runOps.Hostname() {
			valid = false
		}
	}
	return valid
}
