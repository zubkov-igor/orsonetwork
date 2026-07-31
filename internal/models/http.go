package models

type HTTPInfo struct {
    Port        int    `json:"port"`
    Scheme      string `json:"scheme"`

    Server      string `json:"server"`
    Title       string `json:"title"`

    StatusCode  int    `json:"statusCode"`
    ContentType string `json:"contentType"`

    Scripts []string `json:"scripts"`

    Keywords []string `json:"keywords"`

    Fingerprint []string `json:"fingerprint"`
}