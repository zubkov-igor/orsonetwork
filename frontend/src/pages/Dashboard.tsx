import { topology } from "../store/topology";
import TopologyGraph from "../components/TopologyGraph";

export default function Dashboard() {
    return (
        <div class="dashboard">
            <h1>Dashboard</h1>

            <div class="dashboard__stats">
                <div class="stat-card">
                    <span class="stat-card__label">
                        Nodes
                    </span>

                    <strong class="stat-card__value">
                        {topology()?.nodes.length ?? 0}
                    </strong>
                </div>

                <div class="stat-card">
                    <span class="stat-card__label">
                        Links
                    </span>

                    <strong class="stat-card__value">
                        {topology()?.links.length ?? 0}
                    </strong>
                </div>

                <div class="stat-card">
                    <span class="stat-card__label">
                        Networks
                    </span>

                    <strong class="stat-card__value">
                        {topology()?.networks.length ?? 0}
                    </strong>
                </div>
            </div>
            <div class="dashboard__topology">
    <h2>Network Topology</h2>

    <TopologyGraph compact />
</div>
            <h2>Devices</h2>

            <div class="dashboard__nodes">
                {topology()?.nodes.map((node) => (
                    <div class="node-card">
                        <h3>{node.label}</h3>

                        <span>
                            Type: {node.type}
                        </span>

                        <span>
                            IP: {node.ip}
                        </span>

                        <span>
                            MAC: {node.mac || "—"}
                        </span>

                        <span>
                            Vendor: {node.vendor || "—"}
                        </span>
                          <span>
                            Hostname: {node.hostname || "—"}
                        </span>
                          <span>
                            Sources:
{node.sources?.join(", ") || "—"}
                        </span>
                    </div>
                ))}
            </div>

        </div>
    );
}