import { createSignal } from "solid-js";

import { DeepScanWifi } from "../../wailsjs/go/main/App";
import { wifi } from "../../wailsjs/go/models";

export default function Utility() {
  const [deepLoading, setDeepLoading] = createSignal(false);
  const [error, setError] = createSignal("");

  const [accessPoints, setAccessPoints] = createSignal<wifi.AccessPoint[]>([]);
  const [clients, setClients] = createSignal<wifi.Client[]>([]);

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

  const groupedClients = () => {
    const groups = new Map<
      string,
      {
        accessPoint?: wifi.AccessPoint;
        clients: wifi.Client[];
      }
    >();

    const unassociated: wifi.Client[] = [];

    for (const client of clients()) {
      const bssid = client.bssid?.trim();

      if (!bssid || bssid.toLowerCase() === "(not associated)") {
        unassociated.push(client);
        continue;
      }

      const key = bssid.toLowerCase();

      if (!groups.has(key)) {
        const accessPoint = accessPoints().find(
          (ap) => ap.bssid.toLowerCase() === key,
        );

        groups.set(key, {
          accessPoint,
          clients: [],
        });
      }

      groups.get(key)!.clients.push(client);
    }

    return {
      associated: Array.from(groups.values()),
      unassociated,
    };
  };

  return (
    <div class="utility-page">
      <div class="utility-page__header">
        <h2>Wi-Fi Scanner</h2>

        <button onClick={handleDeepScan} disabled={deepLoading()}>
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
                <th>VENDOR</th>
              </tr>
            </thead>

            <tbody>
              {accessPoints().map((ap) => (
                <tr>
                  <td>{ap.ssid || "<Hidden network>"}</td>
                  <td>{ap.bssid}</td>
                  <td>{ap.signal}</td>
                  <td>{ap.channel}</td>
                  <td>{ap.encryption}</td>
                  <td>{ap.vendor || "—"}</td>
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

          {groupedClients().associated.length > 0 && (
            <div class="wifi-client-groups">
              {groupedClients().associated.map((group) => (
                <div class="wifi-client-group">
                  <div class="wifi-client-group__header">
                    <div class="wifi-client-group__info">
                      <strong>{group.accessPoint?.ssid || "Unknown AP"}</strong>

                      <span class="wifi-client-group__bssid">
                        BSSID:{" "}
                        {group.accessPoint?.bssid || group.clients[0]?.bssid}
                      </span>
                    </div>

                    <span class="wifi-client-group__count">
                      {group.clients.length}{" "}
                      {group.clients.length === 1 ? "client" : "clients"}
                    </span>
                  </div>

                  <table>
                    <thead>
                      <tr>
                        <th>Client MAC</th>
                        <th>VENDOR</th>
                        <th>Brand</th>
                        <th>Model</th>
                        <th>Signal</th>
                      </tr>
                    </thead>

                    <tbody>
                      {group.clients.map((client) => (
                        <tr>
                          <td>{client.mac}</td>
                          <td>{client.vendor || "—"}</td>
                          <td>{client.brand || "—"}</td>
                          <td>{client.model || "—"}</td>
                          <td>
                            {client.signal && client.signal !== "-1"
                              ? client.signal
                              : "—"}
                          </td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
                </div>
              ))}
            </div>
          )}

          {groupedClients().unassociated.length > 0 && (
            <div class="wifi-client-group wifi-client-group--unassociated">
              <div class="wifi-client-group__header">
                <div class="wifi-client-group__info">
                  <strong>Not associated</strong>

                  <span class="wifi-client-group__bssid">
                    No associated access point
                  </span>
                </div>

                <span class="wifi-client-group__count">
                  {groupedClients().unassociated.length}{" "}
                  {groupedClients().unassociated.length === 1
                    ? "client"
                    : "clients"}
                </span>
              </div>

              <table>
                <thead>
                  <tr>
                    <th>Client MAC</th>
                    <th>VENDOR</th>
                    <th>Brand</th>
                    <th>Model</th>
                    <th>Signal</th>
                  </tr>
                </thead>

                <tbody>
                  {groupedClients().unassociated.map((client) => (
                    <tr>
                      <td>{client.mac}</td>
                      <td>{client.vendor || "—"}</td>
                      <td>{client.brand || "—"}</td>
                      <td>{client.model || "—"}</td>
                      <td>
                        {client.signal && client.signal !== "-1"
                          ? client.signal
                          : "—"}
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </section>
      )}
    </div>
  );
}
