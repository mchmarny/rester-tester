package image

import (
	"os"
	"testing"
)

const (
	nonYouTubeTestURL = "https://www.filestackapi.com/api/file/laj8RhF3Q9GJvBhEbdTI"
	youTubeTestURL    = "https://www.youtube.com/watch?v=DjByja9ejTQ"
)

func TestYouTubeURLParsing(t *testing.T) {

	urlResult, err := parseURL(youTubeTestURL)
	if err != nil {
		t.Fatalf("Error on parsing on thumbnail creation: %s -> %v",
			youTubeTestURL, err)
	}

	if youTubeTestURL == urlResult {
		t.Fatal("Didn't parse YouTube URL)",
			youTubeTestURL, urlResult)
	}

	t.Logf("YouTube URL: %s", urlResult)

	if !testing.Short() {

		testID := "test"

		tPath, err := makeThumbnail(testID, &ThumbnailRequest{
			Src:    urlResult,
			Width:  200,
			Height: 200,
		})
		if err != nil {
			t.Fatalf("Failed on thumbnail creation: %s -> %v", nonYouTubeTestURL, err)
		}

		if _, err := os.Stat(tPath); os.IsNotExist(err) {
			t.Fatalf("Thumbnail created but the resulting path does not exist: %s -> %v",
				tPath, err)
		}

		t.Logf("YouTube thumbnail path: %s", tPath)

	}

}

func TestNonYouTubeURLParsing(t *testing.T) {

	urlResult, err := parseURL(nonYouTubeTestURL)
	if err != nil {
		t.Fatalf("Error on parsing on thumbnail creation: %s -> %v",
			nonYouTubeTestURL, err)
	}

	if nonYouTubeTestURL != urlResult {
		t.Fatalf("Shouldn't rewrite non-YouTube URL (sent: %s got: %s)",
			nonYouTubeTestURL, urlResult)
	}

}

func TestThumbnailCreation(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping thumbnail creation in short mode.")
	}

	testID := "test"

	tPath, err := makeThumbnail(testID, &ThumbnailRequest{Src: nonYouTubeTestURL})
	if err != nil {
		t.Fatalf("Failed on thumbnail creation: %s -> %v", nonYouTubeTestURL, err)
	}

	if _, err := os.Stat(tPath); os.IsNotExist(err) {
		t.Fatalf("Thumbnail created but the resulting path does not exist: %s -> %v",
			tPath, err)
	}

	t.Logf("Test thumbnail path: %s", tPath)
}
