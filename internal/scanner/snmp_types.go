package scanner

type SNMPVersion byte


const (
	SNMPv1 SNMPVersion = 0
	SNMPv2c SNMPVersion = 1
)


type SNMPRequest struct {

	Version SNMPVersion

	Community string

	RequestID int

	OID []int
}