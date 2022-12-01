package data

import (
	"fmt"
	"io"
	"net/http"
)

var session = "53616c7465645f5f4aef647575fbb58e3ccfbe04aaa03d15c3dadaa3d3383dc370a753256c3e5b03e9145d0f71cdea2e38a7b6d95121c52ca4081487718bd84d"

func GetData(day int) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", day), nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Cookie", "session="+session)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
