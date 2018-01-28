package image

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	urlFormat = "http://www.youtube.com/get_video_info?&video_id="
)

type youTubeVideoURL struct {
	Type string
	URL  string
}

func getYouTubeVideo(videoID string) ([]*youTubeVideoURL, error) {
	queryString, err := fetchMeta(videoID)
	if err != nil {
		return nil, err
	}

	u, err := parseMeta(queryString)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func fetchMeta(videoID string) (string, error) {
	resp, err := http.Get(urlFormat + videoID)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	queryString, _ := ioutil.ReadAll(resp.Body)
	return string(queryString), nil
}

func parseMeta(queryString string) ([]*youTubeVideoURL, error) {
	u, _ := url.Parse("?" + queryString)
	query := u.Query()

	// no such video
	if query.Get("errorcode") != "" || query.Get("status") == "fail" {
		return nil, errors.New(query.Get("reason"))
	}

	// decode the format data some more
	formatParams := strings.Split(query.Get("url_encoded_fmt_stream_map"), ",")

	urls := make([]*youTubeVideoURL, 0)

	// load multiple format choices
	for _, f := range formatParams {
		furl, _ := url.Parse("?" + f)
		fquery := furl.Query()
		ytURL := &youTubeVideoURL{
			Type: fquery.Get("type"),
			URL:  fquery.Get("url") + "&signature=" + fquery.Get("sig"),
		}
		//logger.Printf("[%s] %s", ytURL.Type, ytURL.URL)
		urls = append(urls, ytURL)
	}

	return urls, nil
}
