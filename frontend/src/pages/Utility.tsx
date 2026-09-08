import { createSignal } from "solid-js";
import { ScanWifi } from "../../wailsjs/go/main/App";

interface WifiNetwork {
  ssid: string;
  bssid?: string;
  signal?: string;
  security?: string;
  channel?: string;
  freq?: string;
  rate?: string;
  mode?: string;
}

function getWifiRisk(n: WifiNetwork): "critical" | "high" | "medium" | "low" {
  const sec = (n.security || "").toUpperCase();

  if (!sec || sec === "--" || sec.includes("OPEN")) return "critical";
  if (sec.includes("WEP")) return "critical";
  if (sec.includes("WPA") && !sec.includes("WPA2") && !sec.includes("WPA3"))
    return "high";
  if (sec.includes("TKIP")) return "high";
  if (sec.includes("WPA2") && !sec.includes("WPA3")) return "medium";
  return "low";
}

const riskOrder = {
  critical: 0,
  high: 1,
  medium: 2,
  low: 3,
} as const;

function getRiskLabel(risk: keyof typeof riskOrder): string {
  switch (risk) {
    case "critical":
      return "No encryption or WEP";
    case "high":
      return "Outdated WPA / TKIP";
    case "medium":
      return "WPA2 — consider upgrading to WPA3";
    case "low":
      return "Modern protection (WPA3)";
  }
}

export default function Utility() {
  const [networks, setNetworks] = createSignal<WifiNetwork[]>([]);
  const [loading, setLoading] = createSignal(false);
  const [error, setError] = createSignal("");

  const sortedNetworks = () =>
    [...networks()].sort(
      (a, b) => riskOrder[getWifiRisk(a)] - riskOrder[getWifiRisk(b)],
    );

  const riskStats = () => {
    const stats = { critical: 0, high: 0, medium: 0, low: 0 };

    for (const n of networks()) {
      stats[getWifiRisk(n)]++;
    }

    return stats;
  };

  async function handleScan() {
    setLoading(true);
    setError("");

    try {
      const result = await ScanWifi();
      setNetworks(result ?? []);
    } catch (e: any) {
      console.error(e);
      setError(e?.message || "Wi-Fi scanning error");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div class="utility-page">
      <div class="utility-page__header">
        <h2>Wi‑Fi Scanner</h2>
        <button onClick={handleScan} disabled={loading()}>
          {loading() ? "Scan..." : "Scanning Wi‑Fi"}
        </button>
      </div>

      {error() && <p class="error">{error()}</p>}

      {/* Счётчики рисков */}
      <div class="wifi-summary">
        <div class="risk-metric risk-metric--critical">
          <span class="risk-metric__value">{riskStats().critical}</span>
          <span class="risk-metric__label">Critical</span>
        </div>

        <div class="risk-metric risk-metric--high">
          <span class="risk-metric__value">{riskStats().high}</span>
          <span class="risk-metric__label">High</span>
        </div>

        <div class="risk-metric risk-metric--medium">
          <span class="risk-metric__value">{riskStats().medium}</span>
          <span class="risk-metric__label">Medium</span>
        </div>

        <div class="risk-metric risk-metric--low">
          <span class="risk-metric__value">{riskStats().low}</span>
          <span class="risk-metric__label">Low</span>
        </div>
      </div>

      <table>
        <thead>
          <tr>
            <th>SSID</th>
            <th>Signal</th>
            <th>Security</th>
            <th>Risk</th>
            <th>Channel</th>
            <th>BSSID</th>
          </tr>
        </thead>
        <tbody>
          {sortedNetworks().map((n) => {
            const risk = getWifiRisk(n);

            return (
              <tr>
                <td>{n.ssid}</td>
                <td>{n.signal}</td>
                <td>{n.security}</td>
                <td>
                  <span
                    class={`risk-badge risk-badge--${risk}`}
                    title={getRiskLabel(risk)}
                  >
                    {risk}
                  </span>
                </td>
                <td>{n.channel}</td>
                <td>{n.bssid}</td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
}
