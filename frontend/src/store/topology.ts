import { createSignal } from "solid-js";
import type { models } from "../../wailsjs/go/models";

const [topology, setTopology] =
    createSignal<models.Topology | null>(null);

export {
    topology,
    setTopology,
};