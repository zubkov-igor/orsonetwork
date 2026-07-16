import { For } from "solid-js";

import { 
    setRoute,
    route,
    type Route 
} from "../store/router";

import {
    sidebarExpanded,
    toggleSidebar,
} from "../store/ui";

const menu: {
    id: Route;
    title: string;
    icon: string;
}[] = [
    {
        id: "dashboard",
        title: "Dashboard",
        icon: "/icons/dashboard.svg",
    },
    {
        id: "topology",
        title: "Topology",
        icon: "/icons/topology.svg",
    },
    {
        id: "devices",
        title: "Devices",
        icon: "/icons/devices.svg",
    },
    {
        id: "settings",
        title: "Settings",
        icon: "/icons/settings.svg",
    },
];


export default function Aside() {


    return (
        <aside
            class="sidebar"
            classList={{
                expanded: sidebarExpanded(),
            }}
        >

            <div class="sidebar__header">

                <button
                    class="sidebar__toggle"
                    classList={{
                        rotated: sidebarExpanded(),
                    }}
                    onClick={toggleSidebar}
                >
                    <img
                        src="/icons/spinner.svg"
                        alt="toggle sidebar"
                    />
                </button>

            </div>


            <nav class="sidebar__menu">

                <For each={menu}>
                    {(item) => (

                        <button
                            class="sidebar__item"
                            classList={{
                                active: route() === item.id,
                            }}
                            onClick={() => setRoute(item.id)}
                            title={
                                sidebarExpanded()
                                    ? undefined
                                    : item.title
                            }
                        >

                            <img
                                class="sidebar__icon"
                                src={item.icon}
                                alt=""
                            />


                            <span class="sidebar__text">
                                {item.title}
                            </span>

                        </button>

                    )}
                </For>

            </nav>

        </aside>
    );
}