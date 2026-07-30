package scanner

func LookupNetBIOS(ip string) (NetBIOSResult, error) {

	raw, err := ProbeNetBIOS(
		ip,
	)

	if err != nil {
		return NetBIOSResult{}, err
	}

	return ParseNetBIOSResponse(
		raw,
	), nil
}
