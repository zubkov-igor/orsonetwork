import {
    onMount,
    onCleanup,
    createEffect,
} from "solid-js";

import cytoscape from "cytoscape";

import { topology } from "../store/topology";

export default function TopologyGraph() {
    let container: HTMLDivElement | undefined;

    let cy: cytoscape.Core | undefined;

    onMount(() => {
        if (!container) {
            return;
        }

        cy = cytoscape({
            container,

            style: [
  {
    selector: "node",

    style: {
        "background-color": "#d7dee7",
        label: "data(label)",
        color: "#0f1720",
        "text-valign": "center",
        "text-halign": "center",
    },
},

{
    selector: 'node[type = "gateway"]',

    style: {
        "background-color": "#f0a202",
        color: "#0b1117",
    },
},

{
    selector: 'node[type = "host"]',

    style: {
        "background-color": "#d7dee7",
        color: "#0f1720",
    },
},

{
    selector: "edge",

    style: {
        width: 2,
        "line-color": "#52606d",
    },
},
            ],

            layout: {
                name: "breadthfirst",
                directed: true,
                padding: 30,
            },
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
                        label: node.label,
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

            cy.layout({
                name: "breadthfirst",
                directed: true,
                padding: 30,
            }).run();
        });

        onCleanup(() => {
            cy?.destroy();
        });
    });

    return (
        <div
            ref={container}
            class="topology-graph"
        />
    );
}