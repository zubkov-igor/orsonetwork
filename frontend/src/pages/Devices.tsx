import {
    createSignal,
    For,
    Show,
} from "solid-js";

    import {
    topology,
    selectedDevice,
    setSelectedDevice,
} from "../store/topology";


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


                <select
                    class="devices__filter"
                    value={typeFilter()}
                    onChange={(e) =>
                        setTypeFilter(
                            e.currentTarget.value
                        )
                    }
                >

                    <option value="all">
                        All
                    </option>

                    <option value="gateway">
                        Gateway
                    </option>

                    <option value="host">
                        Host
                    </option>

                </select>

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
                                        <span class="device-status device-status--online">
                                            ● Online
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

                                </tr>

                            )}

                        </For>

                    </tbody>

                </table>

            </div>

 <Show when={selectedDevice()}>

        {(device) => (

            <aside class="device-details">

                <h2>
                    {device().label}
                </h2>


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

            </aside>

        )}

    </Show>

        </div>
    </div>
    );
}