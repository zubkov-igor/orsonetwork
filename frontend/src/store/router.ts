import { createSignal } from "solid-js";


export type Route =
    | "dashboard"
    | "topology"
    | "devices"
    | "settings";


export const [
    route,
    setRoute
] = createSignal<Route>("dashboard");