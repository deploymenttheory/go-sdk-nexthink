package ui_events

import (
	"fmt"
	"net/url"
	"strings"
)

func messagePath(href string) (string, error) {
	if href == "" {
		return Endpoint + "/messages", nil
	}
	u, err := url.Parse(href)
	if err != nil {
		return "", fmt.Errorf("invalid next link: %w", err)
	}
	if (u.Scheme != "" && u.Scheme != "https") || (u.Host != "" && u.Scheme == "") || u.User != nil || u.Fragment != "" || strings.Contains(u.Path, "..") || u.RawPath != "" {
		return "", fmt.Errorf("next link must be a messages URL without credentials or fragment")
	}
	if u.Path != "messages" && u.Path != Endpoint+"/messages" {
		return "", fmt.Errorf("next link must target the UI messages endpoint")
	}
	// Absolute links retain their origin so the shared transport can reject cross-origin URLs.
	if u.IsAbs() {
		if u.Path != Endpoint+"/messages" {
			return "", fmt.Errorf("absolute next link must target the UI messages endpoint")
		}
		return u.String(), nil
	}
	path := Endpoint + "/messages"
	if u.RawQuery != "" {
		path += "?" + u.RawQuery
	}
	return path, nil
}
