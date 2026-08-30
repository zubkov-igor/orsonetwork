import {
    createSignal,
    onMount,
} from "solid-js";

import {
    GetScannerConfig,
    UpdateScannerConfig,
} from "../../wailsjs/go/main/App";


export default function Settings() {

    const [settings, setSettings] = createSignal({
        enable_icmp: true,
        enable_arp: true,
        enable_reverse_dns: true,
        enable_netbios: true,
        enable_mdns: true,
        enable_ssdp: true,
        enable_snmp: false,
        enable_tcp: true,
        enable_udp: true,
        workers: 20,
    });

    const [saved, setSaved] = createSignal(false);

    onMount(async () => {

        const config = await GetScannerConfig();

        setSettings(config);
    });

function updateSetting(
    key: string,
    value: boolean | number
) {
    setSettings({
        ...settings(),
        [key]: value,
    });
}
function saveSettings() {

    UpdateScannerConfig(
        settings()
    );

    setSaved(true);

    setTimeout(
        () => setSaved(false),
        3000
    );
}
    return (
        <div class="settings">

            <h1>Settings</h1>


            <section class="settings__section">

                <h2>
                    Discovery
                </h2>


                <div class="settings__row">
                    <span>
                        ICMP Ping
                    </span>

                    <input
    type="checkbox"
    checked={settings().enable_icmp}
    onChange={(e) =>
        updateSetting(
            "enable_icmp",
            e.currentTarget.checked
        )
    }
/>
                </div>


<div class="settings__row">

    <span>
        ARP
    </span>

    <input
        type="checkbox"
        checked={settings().enable_arp}
        onChange={(e) =>
            updateSetting(
                "enable_arp",
                e.currentTarget.checked
            )
        }
    />

</div>
                <div class="settings__row">
                    <span>
                        Reverse DNS
                    </span>

                    <input
    type="checkbox"
    checked={settings().enable_reverse_dns}
    onChange={(e) =>
        updateSetting(
            "enable_reverse_dns",
            e.currentTarget.checked
        )
    }
/>
                </div>


                <div class="settings__row">
                    <span>
                        NetBIOS
                    </span>

                    <input
    type="checkbox"
    checked={settings().enable_netbios}
    onChange={(e) =>
        updateSetting(
            "enable_netbios",
            e.currentTarget.checked
        )
    }
/>
                </div>

<div class="settings__row">

    <span>
        mDNS
    </span>

    <input
        type="checkbox"
        checked={settings().enable_mdns}
        onChange={(e) =>
            updateSetting(
                "enable_mdns",
                e.currentTarget.checked
            )
        }
    />

</div>


         <div class="settings__row">

    <span>
        SSDP
    </span>

    <input
        type="checkbox"
        checked={settings().enable_ssdp}
        onChange={(e) =>
            updateSetting(
                "enable_ssdp",
                e.currentTarget.checked
            )
        }
    />

</div>

<div class="settings__row">

    <span>
        SNMP
    </span>

    <input
        type="checkbox"
        checked={settings().enable_snmp}
        onChange={(e) =>
            updateSetting(
                "enable_snmp",
                e.currentTarget.checked
            )
        }
    />

</div>

<div class="settings__row">

    <span>
        TCP
    </span>

    <input
        type="checkbox"
        checked={settings().enable_tcp}
        onChange={(e) =>
            updateSetting(
                "enable_tcp",
                e.currentTarget.checked
            )
        }
    />

</div>

<div class="settings__row">

    <span>
        UDP
    </span>

    <input
        type="checkbox"
        checked={settings().enable_udp}
        onChange={(e) =>
            updateSetting(
                "enable_udp",
                e.currentTarget.checked
            )
        }
    />

</div>
            </section>


            <section class="settings__section">

                <h2>
                    Performance
                </h2>


                <div class="settings__row">

                    <span>
                        Discovery workers
                    </span>

                <input
    type="number"
    min="1"
    value={settings().workers}
    onInput={(e) =>
        updateSetting(
            "workers",
            Number(e.currentTarget.value)
        )
    }
/>

                </div>


            </section>


<button
            class="settings__save"
            onClick={saveSettings}
        >
            Save
        </button>

        {saved() && (
            <div class="settings__saved">
                Settings saved
            </div>
        )}

    </div>
    );
}