import {
    createSignal,
} from "solid-js";

import type { models } from "../../wailsjs/go/models";


const [topology, setTopology] =
    createSignal<models.Topology | null>(null);


const [selectedDevice, setSelectedDevice] =
    createSignal<models.Node | null>(null);


export {
    topology,
    setTopology,

    selectedDevice,
    setSelectedDevice,
};