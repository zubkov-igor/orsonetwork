import { createSignal, onCleanup } from "solid-js";

import { GetTopology } from "../../wailsjs/go/main/App";

import {
  topology,
  setTopology,
  scanning,
  setScanning,
} from "../store/topology";

const [scanDuration, setScanDuration] = createSignal(0);

const [lastScan, setLastScan] = createSignal<number | null>(null);

const [scanElapsed, setScanElapsed] = createSignal(0);

let scanTimer: ReturnType<typeof setInterval> | undefined;

let scanStartedAt = 0;

export default function Header() {
  const handleScan = async () => {
    if (scanning()) {
      return;
    }

    setScanning(true);

    scanStartedAt = Date.now();

    setScanElapsed(0);

    scanTimer = setInterval(() => {
      setScanElapsed(Date.now() - scanStartedAt);
    }, 100);

    try {
      const result = await GetTopology();

      setTopology(result.topology);
      setScanDuration(result.duration);
      setLastScan(result.lastScan);

      console.log("Topology received:", result);
    } catch (error) {
      console.error("Scan failed:", error);
    } finally {
      if (scanTimer) {
        clearInterval(scanTimer);

        scanTimer = undefined;
      }

      setScanning(false);
    }
  };
  onCleanup(() => {
    if (scanTimer) {
      clearInterval(scanTimer);
    }
  });

  return (
    <header class="header">
      <div class="header__body">
        <div class="header__logo-block">
          <img src="/logo.png" class="header__logo" alt="OrsoNetwork logo" />

          <span class="header__brand">OrsoNetwork</span>
        </div>

        <div class="header__stats">
          <div class="header__stat">
            <span class="header__stat-label">Nodes:</span>

            <strong class="header__stat-value">
              {topology()?.nodes.length ?? 0}
            </strong>
          </div>

          <div class="header__stat">
            <span class="header__stat-label">Links:</span>

            <strong class="header__stat-value">
              {topology()?.links.length ?? 0}
            </strong>
          </div>

          <div class="header__stat">
            <span class="header__stat-label">Networks:</span>

            <strong class="header__stat-value">
              {topology()?.networks.length ?? 0}
            </strong>
          </div>

          <div class="header__stat">
            <span class="header__stat-label">Scan duration:</span>

            <strong class="header__stat-value">
              {scanning()
                ? `${(scanElapsed() / 1000).toFixed(1)}s`
                : scanDuration() > 0
                  ? `${(scanDuration() / 1000).toFixed(1)}s`
                  : "—"}
            </strong>
          </div>

          <div class="header__stat">
            <span class="header__stat-label">Last scan</span>

            <strong class="header__stat-value">
              {lastScan() !== null
                ? new Date(lastScan()! * 1000).toLocaleTimeString([], {
                    hour: "2-digit",
                    minute: "2-digit",
                    second: "2-digit",
                    hour12: false,
                  })
                : "—"}
            </strong>
          </div>
        </div>

        <button class="header__scan" onClick={handleScan} disabled={scanning()}>
          {scanning() ? "Scanning..." : "Scan"}
        </button>
      </div>
    </header>
  );
}
