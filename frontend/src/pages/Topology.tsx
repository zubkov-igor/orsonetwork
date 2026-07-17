import TopologyGraph from "../components/TopologyGraph";

export default function Topology() {
    return (
        <section class="topology-page">
            <h1>Network Topology</h1>

            <div class="topology-page__graph">
                <TopologyGraph />
            </div>
        </section>
    );
}