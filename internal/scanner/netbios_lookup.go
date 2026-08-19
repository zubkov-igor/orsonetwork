package scanner

func LookupNetBIOS(
	ip string,
) (NetBIOSResult, error) {

	data, err := ProbeNetBIOS(ip)

	if err != nil {
		return NetBIOSResult{}, err
	}

	return ParseNetBIOSResponse(data), nil
}
