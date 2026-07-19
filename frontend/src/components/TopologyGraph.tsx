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

onMount(() => {
    if (!container) {
        return;
    }

    const latencyGood = 10;
    const latencyWarning = 50;
    
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
    {
        selector: "node",

        style: {
            "background-color": colorNode,
            width: 70,
            height: 70,
            label: "data(label)",
            color: colorLight,
            "text-valign": "center",
            "text-halign": "center",
            "font-size": 11,
            "font-weight": "bold",
            "text-wrap": "wrap",
            "text-max-width": "80px",
            "border-width": 2,
            "border-color": colorNodeBorder,
        },
    },

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

    {
        selector: 'node[type = "gateway"]',

        style: {
            "background-color": colorGateway,
            width: 90,
            height: 90,
            "border-width": 3,
            "border-color": colorGatewayBorder,
            "font-size": 12,
        },
    },

        {
            selector: 'node[type = "host"]',

            style: {
                "background-color": colorNode,

                width: 70,
                height: 70,

                "border-color": colorNodeBorder,
            },
        },

        {
            selector: "edge",

            style: {
                width: 2,

                label: "data(latencyLabel)",

                color: colorLight,

                "text-background-color": colorLatencyLabelBg,

                "text-background-opacity": 1,
   
                "line-color": colorNodeBorder,

                "target-arrow-color": colorNodeBorder,

                "target-arrow-shape": "triangle",

                "curve-style": "bezier",
            },
        },
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

cy.on("tap", "node", (event) => {

    const node = event.target;

    const currentTopology = topology();

    if (!currentTopology) {
        return;
    }

    const device =
        currentTopology.nodes.find(
            (item) => item.id === node.id()
        );

    if (!device) {
        return;
    }

    setSelectedDevice(device);
});

requestAnimationFrame(() => {

    cy?.resize();

    cy?.layout({
        name: "breadthfirst",
        directed: true,
        padding: props.compact ? 20 : 40,
    }).run();

});

createEffect(() => {
    const currentTopology = topology();

    if (!currentTopology || !cy) {
        return;
    }

    cy.elements().remove();

    const nodes = currentTopology.nodes ?? [];
    const links = currentTopology.links ?? [];

    cy.add([
        ...nodes.map((node) => ({
            data: {
                id: node.id,
                label: `${node.type.toUpperCase()}\n${node.ip}`,
                type: node.type,
                ip: node.ip,
                mac: node.mac,
                vendor: node.vendor,
                hostname: node.hostname,
                sources: node.sources,
            },
        })),

      ...links.map((link, index) => ({
    data: {
        id: `link-${index}`,
        source: link.from,
        target: link.to,
        type: link.type,

        latencyLabel:
            link.latency > 0
                ? `${link.latency.toFixed(1)} ms`
                : "—",

        latencyStatus:
            link.latency <= 0
                ? "timeout"
                : link.latency < latencyGood
                    ? "good"
                    : link.latency <= latencyWarning
                        ? "warning"
                        : "critical",
    },
})),
    ]);

    cy.resize();

    cy.layout({
        name: "breadthfirst",
        directed: true,
        padding: props.compact ? 20 : 40,
    }).run();
});


createEffect(() => {
    const selected = selectedDevice();

    if (!cy) {
        return;
    }

    cy.nodes().removeClass(
        "node--selected"
    );

    if (!selected) {
        return;
    }

    const node =
        cy.getElementById(selected.id);

    if (node.length > 0) {
        node.addClass(
            "node--selected"
        );
    }
});



    onCleanup(() => {
        cy?.destroy();
    });

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