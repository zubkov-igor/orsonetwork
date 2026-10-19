import { createSignal, For, Show } from "solid-js";

import { Router, Monitor } from "lucide-solid";

import { topology, selectedDevice, setSelectedDevice } from "../store/topology";

import Select from "../components/Select";

function formatSMBShareType(type: string): string {
  const value = Number(type);

  if (!Number.isFinite(value)) {
    return type || "Unknown";
  }

  const unsigned = value >>> 0;
  const baseType = unsigned & 0xff;

  let name: string;

  switch (baseType) {
    case 0:
      name = "Disk share";
      break;
    case 1:
      name = "Print queue";
      break;
    case 2:
      name = "Device";
      break;
    case 3:
      name = "IPC";
      break;
    default:
      name = `Unknown (${baseType})`;
  }

  const suffixes: string[] = [];

  if ((unsigned & 0x80000000) !== 0) {
    suffixes.push("special");
  }

  if ((unsigned & 0x02000000) !== 0) {
    suffixes.push("temporary");
  }

  return suffixes.length > 0 ? `${name} (${suffixes.join(", ")})` : name;
}

function formatDeviceType(type: string): string {
  switch (type) {
    case "access_point":
      return "Access Point";
    case "computer":
      return "Computer";
    case "gateway":
      return "Gateway";
    case "server":
      return "Server";
    case "phone":
      return "Phone";
    case "tablet":
      return "Tablet";
    case "printer":
      return "Printer";
    case "camera":
      return "Camera";
    case "nas":
      return "NAS";
    case "tv":
      return "TV";
    case "console":
      return "Console";
    case "iot":
      return "IoT";
    case "router":
      return "Router";
    case "unknown":
      return "Unknown";
    default:
      return type;
  }
}

export default function Devices() {
  const [search, setSearch] = createSignal("");
  const [typeFilter, setTypeFilter] = createSignal("all");

  const deviceTypes = () => {
    const types = new Set(
      (topology()?.nodes ?? []).map((node) => node.type).filter(Boolean),
    );

    return Array.from(types).sort();
  };

  const devices = () => {
    const nodes = topology()?.nodes ?? [];

    return nodes.filter((node) => {
      const text = `${node.label}
                  ${node.ip}
                  ${node.mac ?? ""}
                  ${node.vendor ?? ""}`.toLowerCase();

      return (
        text.includes(search().toLowerCase()) &&
        (typeFilter() === "all" || node.type === typeFilter())
      );
    });
  };

  let rowElements: Record<string, HTMLTableRowElement> = {};

  return (
    <div class="devices">
      <h1>Devices</h1>

      <div class="devices__toolbar">
        <input
          class="devices__search"
          placeholder="Search devices..."
          value={search()}
          onInput={(e) => setSearch(e.currentTarget.value)}
        />

        <Select
          value={typeFilter()}
          options={[
            {
              value: "all",
              label: "All",
            },
            ...deviceTypes().map((type) => ({
              value: type,
              label: formatDeviceType(type),
            })),
          ]}
          onChange={setTypeFilter}
        />
      </div>

      <div class="devices__layout">
        <div class="devices__table-wrapper">
          <table class="devices__table">
            <thead>
              <tr>
                <th>Status</th>
                <th>Type</th>
                <th>IP</th>
                <th>MAC</th>
                <th>Vendor</th>
                <th>Hostname</th>
              </tr>
            </thead>

            <tbody>
              <For each={devices()}>
                {(device) => (
                  <tr
                    ref={(element) => {
                      rowElements[device.id] = element;
                    }}

                    class="devices__row"

                    classList={{
                      selected: selectedDevice()?.id === device.id,
                    }}

                    onClick={() => {
                      if (selectedDevice()?.id === device.id) {
                        setSelectedDevice(null);

                        return;
                      }

                      setSelectedDevice(device);
                    }}
                  >
                    <td>
                      <span
                        class={`device-status ${
                          device.online
                            ? "device-status--online"
                            : "device-status--offline"
                        }`}
                      ></span>
                    </td>

                    <td>{device.type}</td>

                    <td>{device.ip}</td>

                    <td>{device.mac || "—"}</td>

                    <td>{device.vendor || "—"}</td>

                    <td>{device.hostname || "—"}</td>
                  </tr>
                )}
              </For>
            </tbody>
          </table>
        </div>

        <Show when={selectedDevice()}>
          {(device) => (
            <aside class="device-details">
              <div class="device-details__header">
                {device().type === "gateway" ? (
                  <Router size={24} />
                ) : (
                  <Monitor size={24} />
                )}

                <div>
                  <h2>{device().type.toUpperCase()}</h2>

                  <span class="device-details__ip">{device().ip}</span>
                </div>

                <span
                  class={`device-status ${
                    device().online
                      ? "device-status--online"
                      : "device-status--offline"
                  }`}
                ></span>
              </div>

              {/* NETWORK */}

              <div class="device-details__section">
                <div class="device-details__row">
                  <span>Type</span>
                  <strong>{device().type || "—"}</strong>
                </div>

                <div class="device-details__row">
                  <span>IP</span>
                  <strong>{device().ip || "—"}</strong>
                </div>

                <div class="device-details__row">
                  <span>MAC</span>
                  <strong>{device().mac || "—"}</strong>
                </div>

                <div class="device-details__row">
                  <span>Hostname</span>
                  <strong>{device().hostname || "—"}</strong>
                </div>

                <div class="device-details__row">
                  <span>Vendor</span>
                  <strong>{device().vendor || "—"}</strong>
                </div>

                <div class="device-details__row">
                  <span>OS</span>
                  <strong>{device().os || "—"}</strong>
                </div>

                <div class="device-details__row">
                  <span>Latency</span>
                  <strong>
                    {device().online
                      ? `${(device().rtt / 1_000_000).toFixed(1)} ms`
                      : "—"}
                  </strong>
                </div>
              </div>

              {/* SERVICES */}

              <details class="device-details__accordion">
                <summary>Services</summary>

                <div class="device-details__section">
                  <Show
                    when={device().ports && device().ports.length > 0}
                    fallback={
                      <div class="device-details__empty">No open ports</div>
                    }
                  >
                    <div class="device-details__services">
                      <For each={device().ports}>
                        {(port) => (
                          <div class="device-details__service">
                            <span class="device-details__service-port">
                              {port.protocol || "TCP"} / {port.number}
                            </span>

                            <strong class="device-details__service-name">
                              {port.service || "Unknown"}
                            </strong>
                          </div>
                        )}
                      </For>
                    </div>
                  </Show>
                </div>
              </details>

              {/* SMB SHARES */}

              <Show
                when={device().smb_shares && device().smb_shares.length > 0}
              >
                <details class="device-details__accordion">
                  <summary>SMB Shares</summary>

                  <div class="device-details__section">
                    <div class="device-details__shares">
                      <For each={device().smb_shares}>
                        {(share) => (
                          <div class="device-details__share">
                            <div class="device-details__share-header">
                              <strong class="device-details__share-name">
                                {share.name || "—"}
                              </strong>

                              <span class="device-details__share-type">
                                {share.type
                                  ? formatSMBShareType(share.type)
                                  : "Unknown"}
                              </span>
                            </div>

                            <Show when={share.comment}>
                              <div class="device-details__share-comment">
                                {share.comment}
                              </div>
                            </Show>
                          </div>
                        )}
                      </For>
                    </div>
                  </div>
                </details>
              </Show>

              {/* HTTP */}

              <Show when={device().http && device().http.length > 0}>
                <details class="device-details__accordion">
                  <summary>HTTP</summary>

                  <div class="device-details__section">
                    <For each={device().http}>
                      {(http) => (
                        <div>
                          <div class="device-details__row">
                            <span>Endpoint</span>
                            <strong>
                              {(http.scheme || "http").toUpperCase()} /{" "}
                              {http.port || "—"}
                            </strong>
                          </div>

                          <div class="device-details__row">
                            <span>Server</span>
                            <strong>{http.server || "—"}</strong>
                          </div>

                          <div class="device-details__row">
                            <span>Title</span>
                            <strong>{http.title || "—"}</strong>
                          </div>

                          <div class="device-details__row">
                            <span>Status</span>
                            <strong>{http.statusCode || "—"}</strong>
                          </div>

                          <div class="device-details__row">
                            <span>Content-Type</span>
                            <strong>{http.contentType || "—"}</strong>
                          </div>
                        </div>
                      )}
                    </For>
                  </div>
                </details>
              </Show>

              {/* UDP SERVICES */}

              <Show
                when={device().udpServices && device().udpServices.length > 0}
              >
                <div class="device-details__section">
                  <h3>UDP Services</h3>

                  <For each={device().udpServices}>
                    {(service) => (
                      <div class="device-details__row">
                        <span>UDP / {service.port}</span>

                        <strong>{service.service || "Unknown"}</strong>
                      </div>
                    )}
                  </For>
                </div>
              </Show>

              {/* SNMP */}

              <Show when={device().snmp && device().snmp.length > 0}>
                <div class="device-details__section">
                  <h3>SNMP</h3>

                  <For each={device().snmp}>
                    {(info) => (
                      <div>
                        <div class="device-details__row">
                          <span>Version</span>
                          <strong>{info.Version || "—"}</strong>
                        </div>

                        <div class="device-details__row">
                          <span>Community</span>
                          <strong>{info.Community || "—"}</strong>
                        </div>

                        <div class="device-details__row">
                          <span>System</span>
                          <strong>{info.SysName || "—"}</strong>
                        </div>

                        <div class="device-details__row">
                          <span>Description</span>
                          <strong>{info.SysDescr || "—"}</strong>
                        </div>

                        <div class="device-details__row">
                          <span>Location</span>
                          <strong>{info.SysLocation || "—"}</strong>
                        </div>
                      </div>
                    )}
                  </For>
                </div>
              </Show>

              {/* mDNS */}

              <Show when={device().mdns && device().mdns.length > 0}>
                <div class="device-details__section">
                  <h3>mDNS</h3>

                  <div class="device-details__mdns">
                    <For each={device().mdns}>
                      {(service) => (
                        <div class="device-details__mdns-item">
                          <div class="device-details__mdns-header">
                            <strong class="device-details__mdns-name">
                              {service.name || "—"}
                            </strong>

                            <span class="device-details__mdns-service">
                              {service.service || "Unknown"}
                            </span>
                          </div>

                          <div class="device-details__mdns-meta">
                            {service.host || "—"}
                            {service.ip ? ` · ${service.ip}` : ""}
                            {service.port ? ` · ${service.port}` : ""}
                          </div>
                        </div>
                      )}
                    </For>
                  </div>
                </div>
              </Show>
            </aside>
          )}
        </Show>
      </div>
    </div>
  );
}
