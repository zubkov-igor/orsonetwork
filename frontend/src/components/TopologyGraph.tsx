import {
    onMount,
    onCleanup,
    createEffect,
} from "solid-js";

import cytoscape from "cytoscape";

import {
    topology,
    selectedDevice,
    setSelectedDevice,
} from "../store/topology";

interface TopologyGraphProps {
    compact?: boolean;
}

export default function TopologyGraph(
    props: TopologyGraphProps
) {
    let container: HTMLDivElement | undefined;

    let cy: cytoscape.Core | undefined;

    const latencyGood = 10;
    const latencyWarning = 50;

    function runLayout(compact: boolean) {
        if (!cy) return;

        const roots = cy.nodes('node[type = "gateway"]');

        cy.layout({
            name: "breadthfirst",
            directed: true,

            roots: roots.length > 0
                ? roots.map((n) => n.id())
                : undefined,

            padding: compact ? 10 : 40,

            spacingFactor: compact ? 0.8 : 1.25,

            circle: false,

            fit: true,

            animate: !compact,
            animationDuration: compact ? 0 : 300,
        }).run();

            if (compact) {

        roots.forEach((node) => {

            const position = node.position();

            node.position({
                x: position.x,
                y: position.y - 10,
            });

        });

    }
    }

    onMount(() => {
        if (!container) {
            return;
        }

        const styles = getComputedStyle(
            document.documentElement
        );

        const colorNode = styles
            .getPropertyValue("--color-node")
            .trim();

        const colorNodeBorder = styles
            .getPropertyValue("--color-node-border")
            .trim();

        const colorGateway = styles
            .getPropertyValue("--color-gateway")
            .trim();

        const colorGatewayBorder = styles
            .getPropertyValue("--color-gateway-border")
            .trim();

        const colorLight = styles
            .getPropertyValue("--color-light")
            .trim();

        const colorSwitch = styles
            .getPropertyValue("--color-switch")
            .trim();

        const colorSwitchBorder = styles
            .getPropertyValue("--color-switch-border")
            .trim();

        const colorServer = styles
            .getPropertyValue("--color-server")
            .trim();

        const colorServerBorder = styles
            .getPropertyValue("--color-server-border")
            .trim();

        const colorAP = styles
            .getPropertyValue("--color-ap")
            .trim();

        const colorAPBorder = styles
            .getPropertyValue("--color-ap-border")
            .trim();

        const colorLatencyGood = styles
            .getPropertyValue("--color-latency-good")
            .trim();

        const colorLatencyWarning = styles
            .getPropertyValue("--color-latency-warning")
            .trim();

        const colorLatencyCritical = styles
            .getPropertyValue("--color-latency-critical")
            .trim();

        const colorLatencyTimeout = styles
            .getPropertyValue("--color-latency-timeout")
            .trim();

        const colorLatencyLabelBg = styles
            .getPropertyValue("--color-latency-label-bg")
            .trim();

        cy = cytoscape({
            container,

            style: [
               // Базовый стиль узла
{
    selector: "node",
    style: {
        "background-color": colorNode,

        width: 18,
        height: 18,

        label: "data(label)",

        color: colorLight,

        "text-valign": "bottom",
        "text-halign": "center",

        "text-margin-y": 6,

        "font-size": 11,
        "font-weight": "bold",

        "text-wrap": "wrap",
        "text-max-width": "80px",

        "border-width": 2,
        "border-color": colorNodeBorder,
    },
},

// =========================
// Compact Dashboard
// =========================

    {
    selector: "node.compact",
    style: {
        width: 10,
        height: 10,
        "font-size": 8,
        "border-width": 1,

        "text-valign": "bottom",
        "text-halign": "center",
        "text-margin-y": 5,
    },
},

{
selector: 'node.compact[type = "gateway"]',
    style: {
        width: 10,
        height: 10,
        "font-size": 8,
        "border-width": 2,

        "text-valign": "top",
        "text-halign": "center",
        "text-margin-y": -5,
    },
},

// Выделенный узел
{
    selector: "node.node--selected",
    style: {
        "border-width": 5,
        "border-color": colorGatewayBorder,
        "overlay-color": colorGatewayBorder,
        "overlay-opacity": 0.25,
        "overlay-padding": 8,
    },
},

                // Выделенный узел
                {
                    selector: "node.node--selected",
                    style: {
                        "border-width": 5,
                        "border-color": colorGatewayBorder,
                        "overlay-color": colorGatewayBorder,
                        "overlay-opacity": 0.25,
                        "overlay-padding": 8,
                    },
                },

                // Hover на узле
                {
                    selector: "node.node--hover",
                    style: {
                        "border-width": 4,
                        "border-color": "#e74c3c",
                        "z-index": 999,
                    },
                },

                // Gateway — крупный, сверху
                {
    selector: 'node[type = "gateway"]',
    style: {
        "background-color": colorGateway,
        shape: "round-rectangle",

        width: 20,
        height: 20,

        "border-width": 3,
        "border-color": colorGatewayBorder,

        "font-size": 12,

        "text-valign": "top",
        "text-halign": "center",
        "text-margin-y": -8,
    },
},

                {
    selector: 'node.compact[type = "gateway"]',
    style: {
        width: 15,
        height: 15,
        "font-size": 8,
        "border-width": 2,
    },
},

                // Switch — прямоугольник
                {
                    selector: 'node[type = "switch"]',
                    style: {
                        "background-color": colorSwitch,
                        shape: "rectangle",
                        width: 70,
                        height: 50,
                        "border-color": colorSwitchBorder,
                    },
                },

                // Server — barrel
                {
                    selector: 'node[type = "server"]',
                    style: {
                        "background-color": colorServer,
                        shape: "barrel",
                        width: 60,
                        height: 80,
                        "border-color": colorServerBorder,
                    },
                },

                // AP — треугольник
                {
                    selector: 'node[type = "ap"]',
                    style: {
                        "background-color": colorAP,
                        shape: "triangle",
                        width: 60,
                        height: 60,
                        "border-color": colorAPBorder,
                    },
                },

                // Host — круг
                {
                    selector: 'node[type = "host"]',
                    style: {
                        "background-color": colorNode,
                        shape: "ellipse",
                        width: 10,
                        height: 10,
                        "border-color": colorNodeBorder,
                    },
                },

                {
    selector: 'node.compact[type = "host"]',
    style: {
        width: 18,
        height: 18,
        "border-width": 1,
    },
},

                // Базовый стиль ребра
                {
                    selector: "edge",
                    style: {
                        width: 2,
                        label: "data(latencyLabel)",
                        color: colorLight,
                        "text-background-color": colorLatencyLabelBg,
                        "text-background-opacity": 1,
                        "text-background-padding": "2",
                        "line-color": colorNodeBorder,
                        "curve-style": "bezier",
                        "text-rotation": "autorotate",
                        "font-size": "9px",
                    },
                },

                // Hover на ребре
                {
                    selector: "edge.edge--hover",
                    style: {
                        width: 5,
                        "line-color": "#e74c3c",
                        "target-arrow-color": "#e74c3c",
                        "z-index": 999,
                    },
                },

                // Физический линк — сплошная
                {
                    selector: 'edge[type = "physical"]',
                    style: {
                        "line-style": "solid",
                        width: 3,
                    },
                },

                // Туннель / VPN — пунктир
                {
                    selector: 'edge[type = "tunnel"]',
                    style: {
                        "line-style": "dashed",
                        "line-dash-pattern": [8, 4],
                        width: 2,
                    },
                },

                // Wi-Fi — точечная, без стрелки
                {
                    selector: 'edge[type = "wireless"]',
                    style: {
                        "line-style": "dotted",
                        width: 2,
                        "target-arrow-shape": "none",
                    },
                },

                // Trunk — толще
                {
                    selector: 'edge[type = "trunk"]',
                    style: {
                        width: 4,
                    },
                },

                // Статусы по задержке
                {
                    selector: 'edge[status = "good"]',
                    style: {
                        "line-color": colorLatencyGood,
                        "target-arrow-color": colorLatencyGood,
                    },
                },
                {
                    selector: 'edge[status = "warning"]',
                    style: {
                        "line-color": colorLatencyWarning,
                        "target-arrow-color": colorLatencyWarning,
                    },
                },
                {
                    selector: 'edge[status = "critical"]',
                    style: {
                        "line-color": colorLatencyCritical,
                        "target-arrow-color": colorLatencyCritical,
                    },
                },
                {
                    selector: 'edge[status = "timeout"]',
                    style: {
                        "line-color": colorLatencyTimeout,
                        "target-arrow-color": colorLatencyTimeout,
                    },
                },
            ],

            layout: {
                name: "breadthfirst",
                directed: true,
                padding: 30,
            },
        });

        // Клик по узлу — выбор устройства
        cy.on("tap", "node", (event) => {
            const node = event.target;

            const currentTopology = topology();

            if (!currentTopology) {
                return;
            }

            const device = currentTopology.nodes.find(
                (item) => item.id === node.id()
            );

            if (!device) {
                return;
            }

            setSelectedDevice(device);
        });

        // Hover на узел
        cy.on("mouseover", "node", (event) => {
            event.target.addClass("node--hover");
        });

        cy.on("mouseout", "node", (event) => {
            event.target.removeClass("node--hover");
        });

        // Hover на ребро
        cy.on("mouseover", "edge", (event) => {
            event.target.addClass("edge--hover");
        });

        cy.on("mouseout", "edge", (event) => {
            event.target.removeClass("edge--hover");
        });
    });

    createEffect(() => {
        const currentTopology = topology();

        if (!currentTopology || !cy) {
            return;
        }

    const nodes = currentTopology.nodes ?? [];
    const links = currentTopology.links ?? [];

    const visibleNodes = nodes.filter((node) => {
    return (
        node.type === "gateway" ||
        node.mac ||
        node.hostname ||
        node.vendor ||
        (node.type && node.type !== "unknown")
    );
});

const visibleNodeIds = new Set(
    visibleNodes.map((node) => node.id)
);

const visibleLinks = links.filter((link) => {
    return (
        visibleNodeIds.has(link.from) &&
        visibleNodeIds.has(link.to)
    );
});

cy.json({
    elements: {
       nodes: visibleNodes.map((node) => ({
    data: {
        id: node.id,
        label: node.ip,
        type: node.type,
        ip: node.ip,
        mac: node.mac,
        vendor: node.vendor,
        hostname: node.hostname,
        sources: node.sources,
        online: node.online,
        rtt: node.rtt,
    },

    classes: props.compact
        ? "compact"
        : "",
})),

        edges: visibleLinks.map((link, index) => ({
            data: {
                id: `link-${index}`,
                source: link.from,
                target: link.to,
                type: link.type,
                latencyLabel:
    props.compact
        ? ""
        : link.latency > 0
            ? `${link.latency.toFixed(1)} ms`
            : "—",
                status:
                    link.latency <= 0
                        ? "timeout"
                        : link.latency < latencyGood
                            ? "good"
                            : link.latency <= latencyWarning
                                ? "warning"
                                : "critical",
            },
        })),
    },
});

        cy.resize();
        runLayout(props.compact ?? false);
    });

    createEffect(() => {
        const selected = selectedDevice();

        if (!cy) {
            return;
        }

        cy.nodes().removeClass("node--selected");

        if (!selected) {
            return;
        }

        const node = cy.getElementById(selected.id);

        if (node.length > 0) {
            node.addClass("node--selected");
        }
    });

    onCleanup(() => {
        cy?.destroy();
    });

    return (
        <div
            ref={container}
            class={`topology-graph ${
                props.compact
                    ? "topology-graph--compact"
                    : ""
            }`}
        />
    );
}

