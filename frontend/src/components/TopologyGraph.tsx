import {
    onMount,
    onCleanup,
    createEffect,
} from "solid-js";

import cytoscape from "cytoscape";

import {
    topology,
    selectedDevice,
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

                "line-color": colorNodeBorder,

                "target-arrow-color": colorNodeBorder,

                "target-arrow-shape": "triangle",

                "curve-style": "bezier",
            },
        },
    ],

    layout: {
        name: "breadthfirst",
        directed: true,
        padding: 30,
    },
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
            },
        })),

        ...links.map((link, index) => ({
            data: {
                id: `link-${index}`,
                source: link.from,
                target: link.to,
                type: link.type,
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