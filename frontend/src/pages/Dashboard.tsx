import {
topology,
scanning,
} from "../store/topology";

import TopologyGraph from "../components/TopologyGraph";

import {
    createMemo,
} from "solid-js";

export default function Dashboard() {

const sortedNodes = createMemo(() => {
    return [
        ...(topology()?.nodes || [])
    ]
        .filter((node) => {
    return (
        node.mac ||
        node.hostname ||
        node.vendor ||
        (node.type && node.type !== "unknown")
    );
})
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
            <h2>Network Topology</h2>

            <TopologyGraph compact />
        </div>

        <h2>Devices</h2>

        {scanning() && (
            <div class="scan-spinner">
                <span class="scan-spinner__icon"></span>
            </div>
        )}

        <div class="dashboard__nodes">

            {sortedNodes().map((node) => (

                <div class="node-card">

                    <h3>
                        {node.hostname || node.ip}
                    </h3>

                    {node.hostname && (
                        <span>{node.ip}</span>
                    )}

                    <span>{node.type}</span>

                    <span>
                        {node.vendor || "Unknown vendor"}
                    </span>
                   <span
    class={`node-card__status ${
        node.online
            ? "node-card__status--online"
            : "node-card__status--offline"
    }`}
></span>
                </div>

            ))}

        </div>

    </div>
);

}
