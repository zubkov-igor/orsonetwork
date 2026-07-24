package scanner

func LookupNetBIOS(ip string) (NetBIOSResult, error) {

    output, err := runNMBLookup(ip)

    if err != nil {
        return NetBIOSResult{}, err
    }

    return ParseNMBLookup(output), nil
}