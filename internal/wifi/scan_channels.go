package wifi

import (
	"context"
	"fmt"
	"time"
)

func ScanChannels(
	ctx context.Context,
	iface string,
	duration time.Duration,
) ([]AccessPoint, []Client, error) {
	accessPoints := make(map[string]AccessPoint)
	clients := make(map[string]Client)

	monitorIface, err := createMonitorInterface(iface)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"create monitor interface: %w",
			err,
		)
	}

	defer func() {
		if err := deleteMonitorInterface(monitorIface); err != nil {
			fmt.Printf(
				"warning: failed to delete monitor interface: %v\n",
				err,
			)
		}
	}()

	for _, channel := range scanChannels {
		fmt.Printf("Scanning Wi-Fi channel %d...\n", channel)

		aps, channelClients, err := AirodumpScan(
			ctx,
			monitorIface,
			channel,
			duration,
		)
		if err != nil {
			return nil, nil, fmt.Errorf(
				"scan channel %d: %w",
				channel,
				err,
			)
		}

		fmt.Printf(
			"Channel %d: %d APs, %d clients\n",
			channel,
			len(aps),
			len(channelClients),
		)

		for _, ap := range aps {
			accessPoints[ap.BSSID] = ap
		}

		for _, client := range channelClients {
			key := client.MAC + "|" + client.BSSID
			clients[key] = client
		}
	}

	result := make([]AccessPoint, 0, len(accessPoints))

	for _, ap := range accessPoints {
		result = append(result, ap)
	}

	clientResult := make([]Client, 0, len(clients))

	for _, client := range clients {
		clientResult = append(clientResult, client)
	}

	return result, clientResult, nil
}
