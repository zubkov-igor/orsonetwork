import {
    createSignal,
} from "solid-js";

import {
    UpdateScannerConfig,
} from "../../wailsjs/go/main/App";


export default function Settings() {

const [settings, setSettings] = createSignal({
    EnableICMP: true,
    EnableARP: true,
    EnableReverseDNS: true,
    EnableNetBIOS: true,
    EnableMDNS: true,
    EnableSSDP: true,
    EnableSNMP: true,

    EnableTCP: true,
    EnableUDP: true,

    Workers: 20,
});

const [saved, setSaved] = createSignal(false);

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
    checked={settings().EnableICMP}
    onChange={(e) =>
        updateSetting(
            "EnableICMP",
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
        checked={settings().EnableARP}
        onChange={(e) =>
            updateSetting(
                "EnableARP",
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
    checked={settings().EnableReverseDNS}
    onChange={(e) =>
        updateSetting(
            "EnableReverseDNS",
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
    checked={settings().EnableNetBIOS}
    onChange={(e) =>
        updateSetting(
            "EnableNetBIOS",
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
        checked={settings().EnableMDNS}
        onChange={(e) =>
            updateSetting(
                "EnableMDNS",
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
        checked={settings().EnableSSDP}
        onChange={(e) =>
            updateSetting(
                "EnableSSDP",
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
        checked={settings().EnableSNMP}
        onChange={(e) =>
            updateSetting(
                "EnableSNMP",
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
        checked={settings().EnableTCP}
        onChange={(e) =>
            updateSetting(
                "EnableTCP",
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
        checked={settings().EnableUDP}
        onChange={(e) =>
            updateSetting(
                "EnableUDP",
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
    value={settings().Workers}
    onInput={(e) =>
        updateSetting(
            "Workers",
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