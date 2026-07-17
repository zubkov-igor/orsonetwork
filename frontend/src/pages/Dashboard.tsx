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
            <TopologyGraph />
            <h2>Devices</h2>

            <div class="dashboard__nodes">
                {topology()?.nodes.map((node) => (
                    <div class="node-card">
                        <h3>{node.label}</h3>

                        <p>
                            Type: {node.type}
                        </p>

                        <p>
                            IP: {node.ip}
                        </p>

                        <p>
                            MAC: {node.mac || "—"}
                        </p>

                        <p>
                            Vendor: {node.vendor || "—"}
                        </p>
                    </div>
                ))}
            </div>
            
        </div>
    );
}