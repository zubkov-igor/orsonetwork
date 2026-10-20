package wifi

import (
	"context"
	"time"
)

func DeepScan(
	ctx context.Context,
	iface string,
	duration time.Duration,
) ([]AccessPoint, []Client, error) {
	return ScanChannels(
		ctx,
		iface,
		duration,
	)
}
