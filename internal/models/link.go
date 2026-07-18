package models

type Link struct {
    From    string  `json:"from"`
    To      string  `json:"to"`
    Type    string  `json:"type"`
    Latency float64 `json:"latency"`
}