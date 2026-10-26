import { createSignal } from "solid-js";

import { DeepScanWifi } from "../../wailsjs/go/main/App";
import { wifi } from "../../wailsjs/go/models";

export default function Utility() {
  const [deepLoading, setDeepLoading] = createSignal(false);
  const [error, setError] = createSignal("");

  const [accessPoints, setAccessPoints] = createSignal<wifi.AccessPoint[]>([]);
  const [clients, setClients] = createSignal<wifi.Client[]>([]);

  const getAccessPointLabel = (bssid?: string) => {
    if (!bssid) {
      return "Not associated";
    }

    const ap = accessPoints().find(
      (item) => item.bssid.toLowerCase() === bssid.toLowerCase(),
    );

    if (!ap) {
      return bssid;
    }

    const ssid = ap.ssid || "<hidden>";

    return `${ssid} (${ap.bssid})`;
  };

  async function handleDeepScan() {
    setDeepLoading(true);
    setError("");

    try {
      const result = await DeepScanWifi();

      setAccessPoints(result.accessPoints ?? []);
      setClients(result.clients ?? []);

      console.log("Deep Scan:", result);
    } catch (e: any) {
      console.error("Deep Scan error:", e);
      setError(e?.message || String(e));
    } finally {
      setDeepLoading(false);
    }
  }

  return (
    <div class="utility-page">
      <div class="utility-page__header">
        <h2>Wi-Fi Scanner</h2>

    <button
    onClick={handleDeepScan}
    disabled={deepLoading()}
  >
    {deepLoading() ? "Deep Scanning..." : "Deep Scan"}
  </button>
</div>

{deepLoading() && (
  <div class="scan-spinner">
    <span class="scan-spinner__icon"></span>
  </div>
)}

{error() && <p class="error">{error()}</p>}

      {/* Deep Scan — Access Points */}
      {accessPoints().length > 0 && (
        <section class="wifi-deep-scan">
          <h3>Access Points ({accessPoints().length})</h3>

          <table>
            <thead>
              <tr>
                <th>SSID</th>
                <th>BSSID</th>
                <th>Signal</th>
                <th>Channel</th>
                <th>Encryption</th>
              </tr>
            </thead>

            <tbody>
              {accessPoints().map((ap) => (
                <tr>
                  <td>{ap.ssid || "<hidden>"}</td>
                  <td>{ap.bssid}</td>
                  <td>{ap.signal}</td>
                  <td>{ap.channel}</td>
                  <td>{ap.encryption}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </section>
      )}

      {/* Deep Scan — Clients */}
      {clients().length > 0 && (
        <section class="wifi-deep-scan">
          <h3>Clients ({clients().length})</h3>

          <table>
            <thead>
              <tr>
                <th>MAC</th>
                <th>Access Point</th>
                <th>Signal</th>
              </tr>
            </thead>

            <tbody>
              {clients().map((client) => (
                <tr>
                  <td>{client.mac}</td>
                  <td>{getAccessPointLabel(client.bssid)}</td>
                  <td>{client.signal}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </section>
      )}
    </div>
  );
}
