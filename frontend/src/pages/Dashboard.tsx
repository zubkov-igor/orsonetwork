import { topology } from "../store/topology";
import TopologyGraph from "../components/TopologyGraph";

export default function Dashboard() {
    return (
        <div class="dashboard">
            <h1>Dashboard</h1>
            <div class="dashboard__topology">
    <h2>Network Topology</h2>

    <TopologyGraph compact />
</div>
            <h2>Devices</h2>

    <div class="dashboard__nodes">
    {topology()?.nodes.map((node) => (
        <div class="node-card">
         <h3>
    {node.hostname || node.ip}
</h3>

{node.hostname && (
    <span>{node.ip}</span>
)}

<span>{node.type}</span>

<span>{node.vendor || "Unknown vendor"}</span>
        </div>
    ))}
</div>

        </div>
    );
}