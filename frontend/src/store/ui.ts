import { createSignal } from "solid-js";

export const [
    sidebarExpanded,
    setSidebarExpanded
] = createSignal(false);


export function toggleSidebar() {
    setSidebarExpanded(!sidebarExpanded());
}