package scanner

import (
    
    "net"
    "time"
)


func ProbeNetBIOS(ip string) ([]byte, error) {

    conn, err := net.DialUDP(
        "udp",
        nil,
        &net.UDPAddr{
            IP: net.ParseIP(ip),
            Port:137,
        },
    )

    if err != nil {
        return nil, err
    }

    defer conn.Close()

    query := []byte{
    0x7b, 0x9d,
    0x00, 0x00,
    0x00, 0x01,
    0x00, 0x00,
    0x00, 0x00,
    0x00, 0x00,

    0x20,

    'C','K',
    'A','A','A','A','A','A',
    'A','A','A','A','A','A',
    'A','A','A','A','A','A',
    'A','A','A','A','A','A',
    'A','A','A','A',

    0x00,

    0x00, 0x21,
    0x00, 0x01,
}


    _, err = conn.Write(query)

    if err != nil {
        return nil, err
    }


    conn.SetReadDeadline(
        time.Now().Add(
            2*time.Second,
        ),
    )


    buf := make([]byte,512)

    n,_,err := conn.ReadFromUDP(buf)

    if err != nil {
        return nil, err
    }


    return buf[:n], nil
}

