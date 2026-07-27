package models

type Port struct {
    Number   int    `json:"number"`
    Protocol string `json:"protocol"`
    Service  string `json:"service"`
    Open     bool   `json:"open"`
}