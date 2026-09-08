import { createSignal } from "solid-js";

export type Route =
  "dashboard" | "topology" | "devices" | "utility" | "settings";

export const [route, setRoute] = createSignal<Route>("dashboard");
