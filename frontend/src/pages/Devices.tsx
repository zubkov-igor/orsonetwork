import {
    createSignal,
    For,
    Show,
} from "solid-js";

import {
    Router,
    Monitor,
} from "lucide-solid";

import {
    topology,
    selectedDevice,
    setSelectedDevice,
} from "../store/topology";

import Select from "../components/Select";


export default function Devices() {

    const [search, setSearch] =
        createSignal("");

    const [typeFilter, setTypeFilter] =
        createSignal("all");

    const devices = () => {

        const nodes =
            topology()?.nodes ?? [];


        return nodes
            .filter((node) => {

                return (
                    node.mac ||
                    node.hostname ||
                    node.vendor ||
                    (node.type &&
                        node.type !== "unknown")
                );

            })
            .filter((node) => {

                const text =
                    `${node.label}
                    ${node.ip}
                    ${node.mac ?? ""}
                    ${node.vendor ?? ""}`
                    .toLowerCase();


                return (
                    text.includes(
                        search().toLowerCase()
                    )
                    &&
                    (
                        typeFilter() === "all"
                        ||
                        node.type === typeFilter()
                    )
                );

            });

    };


    let rowElements: Record<string, HTMLTableRowElement> = {};

    return (
        <div class="devices">

            <h1>Devices</h1>


            <div class="devices__toolbar">

                <input
                    class="devices__search"
                    placeholder="Search devices..."
                    value={search()}
                    onInput={(e) =>
                        setSearch(
                            e.currentTarget.value
                        )
                    }
                />


                <Select
                    value={typeFilter()}
                    options={[
                        {
                            value: "all",
                            label: "All",
                        },
                        {
                            value: "gateway",
                            label: "Gateway",
                        },
                        {
                            value: "host",
                            label: "Host",
                        },
                    ]}
                    onChange={setTypeFilter}
                />

            </div>


            <div class="devices__layout">


                <div class="devices__table-wrapper">

                    <table class="devices__table">

                        <thead>

                            <tr>
                                <th>Status</th>
                                <th>Type</th>
                                <th>IP</th>
                                <th>MAC</th>
                                <th>Vendor</th>
                                <th>Hostname</th>
                            </tr>

                        </thead>


                        <tbody>

                            <For each={devices()}>

                                {(device) => (

                                    <tr
                                        ref={(element) => {
                                            rowElements[device.id] =
                                                element;
                                        }}

                                        class="devices__row"

                                        classList={{
                                            selected:
                                                selectedDevice()?.id ===
                                                device.id
                                        }}

                                        onClick={() => {

                                            if (
                                                selectedDevice()?.id ===
                                                device.id
                                            ) {

                                                setSelectedDevice(null);

                                                return;
                                            }


                                            setSelectedDevice(device);

                                        }}
                                    >

                                        <td>

                                            <span
                                                class={`device-status ${
                                                    device.online
                                                        ? "device-status--online"
                                                        : "device-status--offline"
                                                }`}
                                            >
                                                {device.online
                                                    ? "● Online"
                                                    : "● Offline"}
                                            </span>

                                        </td>


                                        <td>
                                            {device.type}
                                        </td>


                                        <td>
                                            {device.ip}
                                        </td>


                                        <td>
                                            {device.mac || "—"}
                                        </td>


                                        <td>
                                            {device.vendor || "—"}
                                        </td>


                                        <td>
                                            {device.hostname || "—"}
                                        </td>

                                    </tr>

                                )}

                            </For>

                        </tbody>

                    </table>

                </div>

<Show when={selectedDevice()}>

    {(device) => (

        <aside class="device-details">

            <div class="device-details__header">

                {device().type === "gateway"
                    ? <Router size={24} />
                    : <Monitor size={24} />
                }

                <div>

                    <h2>
                        {device().type.toUpperCase()}
                    </h2>

                    <span class="device-details__ip">
                        {device().ip}
                    </span>

                </div>

            </div>


            <div class="device-details__status">

                <span
                    class={`device-status ${
                        device().online
                            ? "device-status--online"
                            : "device-status--offline"
                    }`}
                >
                    {device().online
                        ? "● Online"
                        : "● Offline"}
                </span>

            </div>


            {/* NETWORK */}

            <div class="device-details__section">

                <h3>
                    Network
                </h3>

                <div class="device-details__row">
                    <span>Type</span>
                    <strong>
                        {device().type || "—"}
                    </strong>
                </div>

                <div class="device-details__row">
                    <span>IP</span>
                    <strong>
                        {device().ip || "—"}
                    </strong>
                </div>

                <div class="device-details__row">
                    <span>MAC</span>
                    <strong>
                        {device().mac || "—"}
                    </strong>
                </div>

                <div class="device-details__row">
                    <span>Hostname</span>
                    <strong>
                        {device().hostname || "—"}
                    </strong>
                </div>

                <div class="device-details__row">
                    <span>Vendor</span>
                    <strong>
                        {device().vendor || "—"}
                    </strong>
                </div>

                <div class="device-details__row">
                    <span>OS</span>
                    <strong>
                        {device().os || "—"}
                    </strong>
                </div>

                <div class="device-details__row">
                    <span>Latency</span>
                    <strong>
                        {device().online
                            ? `${(
                                device().rtt / 1_000_000
                            ).toFixed(1)} ms`
                            : "—"}
                    </strong>
                </div>

            </div>


{/* SERVICES */}

<div class="device-details__section">
    <h3>Services</h3>

    <Show
        when={
            device().ports &&
            device().ports.length > 0
        }
        fallback={
            <div class="device-details__empty">
                No open ports
            </div>
        }
    >
        <div class="device-details__services">
            <For each={device().ports}>
                {(port) => (
                    <div class="device-details__service">
                        <span class="device-details__service-port">
                            {port.protocol || "TCP"} / {port.number}
                        </span>

                        <strong class="device-details__service-name">
                            {port.service || "unknown"}
                        </strong>
                    </div>
                )}
            </For>
        </div>
    </Show>
</div>


            {/* HTTP */}

            <Show
                when={
                    device().http &&
                    device().http.length > 0
                }
            >

                <div class="device-details__section">

                    <h3>
                        HTTP
                    </h3>

                    <For each={device().http}>

                        {(http) => (

                            <div>

                                <div class="device-details__row">
                                    <span>Port</span>
                                    <strong>
                                        {http.port}
                                    </strong>
                                </div>

                                <div class="device-details__row">
                                    <span>Scheme</span>
                                    <strong>
                                        {http.scheme || "—"}
                                    </strong>
                                </div>

                                <div class="device-details__row">
                                    <span>Server</span>
                                    <strong>
                                        {http.server || "—"}
                                    </strong>
                                </div>

                                <div class="device-details__row">
                                    <span>Title</span>
                                    <strong>
                                        {http.title || "—"}
                                    </strong>
                                </div>

                                <div class="device-details__row">
                                    <span>Status</span>
                                    <strong>
                                        {http.statusCode || "—"}
                                    </strong>
                                </div>

                            </div>

                        )}

                    </For>

                </div>

            </Show>


            {/* UDP SERVICES */}

            <Show
                when={
                    device().udpServices &&
                    device().udpServices.length > 0
                }
            >

                <div class="device-details__section">

                    <h3>
                        UDP Services
                    </h3>

                    <For each={device().udpServices}>

                        {(service) => (

                            <div class="device-details__row">

                                <span>
                                    UDP / {service.port}
                                </span>

                                <strong>
                                    {service.service || "unknown"}
                                </strong>

                            </div>

                        )}

                    </For>

                </div>

            </Show>


            {/* SNMP */}

            <Show
                when={
                    device().snmp &&
                    device().snmp.length > 0
                }
            >

                <div class="device-details__section">

                    <h3>
                        SNMP
                    </h3>

                    <For each={device().snmp}>

                        {(info) => (

                            <div class="device-details__row">

                                <span>
                                    SNMP
                                </span>

                                <strong>
                                    {JSON.stringify(info)}
                                </strong>

                            </div>

                        )}

                    </For>

                </div>

            </Show>


            {/* mDNS */}

            <Show
                when={
                    device().mdns &&
                    device().mdns.length > 0
                }
            >

                <div class="device-details__section">

                    <h3>
                        mDNS
                    </h3>

                    <For each={device().mdns}>

                        {(service) => (

                            <div class="device-details__row">

                                <span>
                                    Service
                                </span>

                                <strong>
                                    {JSON.stringify(service)}
                                </strong>

                            </div>

                        )}

                    </For>

                </div>

            </Show>

        </aside>

    )}

</Show>

            </div>

        </div>
    );
}

