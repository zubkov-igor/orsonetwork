package models

type Topology struct {
    Nodes    []Node    `json:"nodes"`
    Links    []Link    `json:"links"`
    Networks []Network `json:"networks"`
}