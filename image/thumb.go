package image

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
)

func parseURL(u string) (validURL string, err error) {

	logger.Printf("Parsing URL from %s...", u)

	testURL, err := url.Parse(u)
	if err != nil {
		return "", fmt.Errorf("Not properly formatted URL: %s -> %v", u, err)
	}

	logger.Printf("URL host: %s", testURL.Host)

	// if youtube extract the actual video URL
	if testURL.Host == "www.youtube.com" || testURL.Host == "youtu.be" {
		logger.Print("Apears to be YouTube video, extracting actual URL...")

		videoID := testURL.Query().Get("v")
		if len(videoID) == 0 {
			return "", fmt.Errorf("YouTube URL but it has no no video id: %v", testURL)
		}

		urls, err := getYouTubeVideo(videoID)
		if err != nil {
			return "", fmt.Errorf("Error while getting youtube video: %s -> %v", videoID, err)
		}

		if urls == nil || len(urls) < 1 {
			return "", fmt.Errorf("Youtube video has urls: %s -> %v", videoID, urls)
		}

		return urls[0].URL, nil

	}

	return testURL.String(), nil

}

func makeThumbnail(key string, req *ThumbnailRequest) (path string, err error) {

	logger.Printf("Creating thumbnail from %s:%v...", key, req)

	if key == "" || len(key) < 3 {
		return "", fmt.Errorf("Thumbnail key (min 3 characters) is required: %v", key)
	}

	if req == nil {
		return "", fmt.Errorf("Null request: %s", key)
	}

	if req.Width < minThumbnailWidth || req.Width > maxThumbnailWidth {
		logger.Printf("Width outside of permitted range: %d", req.Width)
		req.Width = defaultThumbnailWidth
	}

	if req.Height < minThumbnailHeight || req.Height > maxThumbnailHeight {
		logger.Printf("Height outside of permitted range: %d", req.Height)
		req.Height = defaultThumbnailHeight
	}

	validURL, err := parseURL(req.Src)
	if err != nil {
		return "", fmt.Errorf("Not properly formatted URL: %s -> %v", req.Src, err)
	}

	// temp dir
	tmpDir, err := ioutil.TempDir("", key)
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("img_%s.png", key)
	thumbnailPath := filepath.Join(tmpDir, fileName)
	thumbnailFile, err := os.Create(thumbnailPath)
	if err != nil {
		return "", err
	}

	resolution := fmt.Sprintf("%dx%d", req.Width, req.Height)

	// ffmpeg -ss 1 -i test.mp4 -t 1 -s 300x300 -vframes 1 new.png
	logger.Print("Executing native ffmpeg command...")
	cmd := exec.Command("ffmpeg", "-y", "-ss", "1", "-i", validURL,
		"-t", "1", "-s", resolution, "-vframes", "1", thumbnailFile.Name())

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return thumbnailPath, nil
}
