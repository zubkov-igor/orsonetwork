package models


type UDPService struct {

	IP string `json:"ip"`

	Port int `json:"port"`

	Protocol string `json:"protocol"`

	Service string `json:"service"`

	Info string `json:"info"`

}