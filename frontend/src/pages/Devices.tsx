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


        return nodes.filter((node) => {

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
                            <th>Sources</th>
                        </tr>
                    </thead>


                    <tbody>

                        <For each={devices()}>

                            {(device) => (
<tr
    class="devices__row"
    classList={{
        selected:
            selectedDevice()?.id === device.id
    }}
   onClick={() => {

    if (
        selectedDevice()?.id === device.id
    ) {
        setSelectedDevice(null);

        return;
    }

    setSelectedDevice(device);}}>

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

                                    <td>
                                         
    {device.sources?.length
        ? device.sources.join(", ")
        : "—"}
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

<div class="device-details__section">

    <h3>
        Network
    </h3>

    <div class="device-details__row">
        <span>Type</span>
        <strong>
            {device().type}
        </strong>
    </div>

    <div class="device-details__row">
        <span>IP</span>
        <strong>
            {device().ip}
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


<div class="device-details__section">

    <h3>
        Hardware
    </h3>

    <div class="device-details__row">
        <span>MAC</span>
        <strong>
            {device().mac || "—"}
        </strong>
    </div>

    <div class="device-details__row">
        <span>Vendor</span>
        <strong>
            {device().vendor || "—"}
        </strong>
    </div>

</div>



            </aside>



        )}

    </Show>

        </div>
    </div>
    );
}