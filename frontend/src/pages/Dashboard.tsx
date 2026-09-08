import { topology, scanning } from "../store/topology";

import TopologyGraph from "../components/TopologyGraph";

import { createMemo } from "solid-js";

export default function Dashboard() {
  const sortedNodes = createMemo(() => {
    return [...(topology()?.nodes || [])]
      .filter((node) => node.online)
      .sort((a, b) => {
        const aParts = a.ip.split(".").map(Number);
        const bParts = b.ip.split(".").map(Number);

        for (let i = 0; i < 4; i++) {
          if (aParts[i] !== bParts[i]) {
            return aParts[i] - bParts[i];
          }
        }

        return 0;
      });
  });

  return (
    <div class="dashboard">
      <h1>Dashboard</h1>

      <div class="dashboard__topology">
        <TopologyGraph compact />
      </div>

      <h3>Devices</h3>

      {scanning() && (
        <div class="scan-spinner">
          <span class="scan-spinner__icon"></span>
        </div>
      )}

      <div class="dashboard__nodes">
        {sortedNodes().map((node) => (
          <div class="node-card">
            <h3>{node.hostname || node.ip}</h3>

            {node.hostname && <span>{node.ip}</span>}

            {node.os && node.os.toLowerCase() !== "unknown" && (
              <span>{node.os}</span>
            )}

            {node.type && node.type.toLowerCase() !== "unknown" && (
              <span>{node.type}</span>
            )}

            {node.vendor && node.vendor.toLowerCase() !== "unknown" && (
              <span>{node.vendor}</span>
            )}

            <span
              class={`device-status status-dashboard ${
                node.online ? "device-status--online" : "device-status--offline"
              }`}
            ></span>
          </div>
        ))}
      </div>
    </div>
  );
}
