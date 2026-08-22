package models

type ScanResult struct {
    Topology Topology `json:"topology"`
    Duration int64    `json:"duration"`
}