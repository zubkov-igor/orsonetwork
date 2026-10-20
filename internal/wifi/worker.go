package wifi

import (
	"context"
	"fmt"
	"time"
)

type WiFiWorker struct {
	Interface string
	Duration  time.Duration
	Interval  time.Duration
}

func NewWiFiWorker(
	iface string,
	duration time.Duration,
	interval time.Duration,
) *WiFiWorker {
	return &WiFiWorker{
		Interface: iface,
		Duration:  duration,
		Interval:  interval,
	}
}

func (w *WiFiWorker) Scan(
	ctx context.Context,
) ([]AccessPoint, []Client, error) {
	if w.Interface == "" {
		return nil, nil, fmt.Errorf("wireless interface is empty")
	}

	if w.Duration <= 0 {
		return nil, nil, fmt.Errorf("scan duration must be greater than zero")
	}

	return ScanChannels(
		ctx,
		w.Interface,
		w.Duration,
	)
}

func (w *WiFiWorker) Run(ctx context.Context) error {
	if w.Interface == "" {
		return fmt.Errorf("wireless interface is empty")
	}

	if w.Duration <= 0 {
		return fmt.Errorf("scan duration must be greater than zero")
	}

	if w.Interval <= 0 {
		return fmt.Errorf("scan interval must be greater than zero")
	}

	for {
		_, _, err := w.Scan(ctx)
		if err != nil {
			return err
		}

		timer := time.NewTimer(w.Interval)

		select {
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()

		case <-timer.C:
		}
	}
}
