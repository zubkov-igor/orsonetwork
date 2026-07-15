package models

type Node struct {
    ID     string `json:"id"`
    Label  string `json:"label"`
    Type   string `json:"type"`
    IP     string `json:"ip"`
    MAC    string `json:"mac"`
    Vendor string `json:"vendor"`
}