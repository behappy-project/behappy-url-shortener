package service

import (
	"behappy-url-shortener/src/common"
	"behappy-url-shortener/src/model"
	"behappy-url-shortener/src/util"
	"github.com/golang-module/carbon/v2"
	"github.com/mitchellh/mapstructure"
	"regexp"
	"time"
)

func ShortenUrl(r *model.ShortenRequest) (error, model.ShortenResponse) {
	var shortenRes model.ShortenResponse
	ok := checkLongUrl(r.LongUrl)
	if !ok {
		return common.ServerErrorWithMsg("url is not correct!"), shortenRes
	}
	endDateIsZero := r.EndDate.IsZero()
	startDateIsZero := r.StartDate.IsZero()
	endDateBeforeNow := !endDateIsZero && r.EndDate.Lt(carbon.Now())
	startDateBeforeNow := !startDateIsZero && r.StartDate.Lt(carbon.Now())
	endDateBeforeStartDate := !startDateIsZero && !endDateIsZero && r.StartDate.Gt(r.EndDate.Carbon)
	if endDateBeforeNow || startDateBeforeNow || endDateBeforeStartDate {
		return common.ParameterError("time param is not correct!"), shortenRes
	}
	var expired time.Duration
	if !endDateIsZero {
		expired = r.EndDate.Carbon2Time().Sub(util.ZeroTime(time.Now()))
	}
	var currentErr error
	model.Set(r.LongUrl, r.StartDate, r.EndDate, expired, r.CNew, func(err error, reply map[string]string) {
		if err != nil {
			currentErr = err
			return
		}
		mapstructure.Decode(reply, &shortenRes)
	})
	return currentErr, shortenRes
}

func HandleShortenUrl(shortUrl string) (error, string) {
	longUrlRes := ""
	ok := checkShortUrl(shortUrl)
	if !ok {
		return common.ServerErrorWithMsg("url is not correct!"), longUrlRes
	}
	var currentErr error
	model.Get(shortUrl, func(err error, reply map[string]string) {
		if err != nil {
			currentErr = err
			return
		}
		startDate := reply["start_date"]
		endDate := reply["end_date"]
		longUrl := reply["url"]
		if startDate != "" && endDate != "" {
			dbStartDate := carbon.Parse(startDate)
			dbEndDate := carbon.Parse(endDate)
			now := carbon.Now()
			if !(now.Gt(dbStartDate) && now.Lt(dbEndDate)) {
				currentErr = common.ServerErrorWithMsg("url can't be used!")
			}
		}
		longUrlRes = longUrl
	})
	return currentErr, longUrlRes
}

func checkLongUrl(url string) bool {
	regexPattern := `^(http|ftp|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?`
	matched, _ := regexp.MatchString(regexPattern, url)
	// Url correct
	return matched
}
func checkShortUrl(url string) bool {
	regexPattern := `^[\w=]+$`
	matched, _ := regexp.MatchString(regexPattern, url)
	// Url correct
	return matched
}
