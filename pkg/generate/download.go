package generate

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

var unihanURLPattern = "https://unicode.org/Public/%s/ucd/Unihan.zip"

var unihanClient = &http.Client{
	Timeout: 5 * time.Minute,
}

func LatestVersion() string {
	return "UCD/latest"
}

func URL(version string) string {
	return fmt.Sprintf(unihanURLPattern, version)
}

func ReadVersion(version string) (*zip.Reader, error) {
	resp, err := unihanClient.Get(URL(version))
	if err != nil {
		return nil, fmt.Errorf("getting unihan database version %s (%s): %w", version, URL(version), err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting unihan database version %s (%s): %s", version, URL(version), resp.Status)
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading unihan database version %s (%s): %w", version, URL(version), err)
	}

	return zip.NewReader(bytes.NewReader(bs), int64(len(bs)))
}

func getFile(r *zip.Reader, filename string) (io.ReadCloser, error) {
	for _, f := range r.File {
		if f.Name != filename {
			continue
		}

		return f.Open()
	}

	return nil, fmt.Errorf("file %s not found in zip file", filename)
}
