import { createSignal } from "solid-js";

const [expanded, setExpanded] = createSignal(false);

export const sidebarExpanded = expanded;

export function toggleSidebar() {
    setExpanded(!expanded());
}