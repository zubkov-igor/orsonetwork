import { createSignal } from "solid-js";
import { GetTopology } from "../../wailsjs/go/main/App";
import { setTopology } from "../store/topology";

export default function Header() {
    const [scanning, setScanning] = createSignal(false);

    const handleScan = async () => {
        if (scanning()) {
            return;
        }

        setScanning(true);

        try {
            const result = await GetTopology();

            setTopology(result);

            console.log("Topology received:", result);
        } catch (error) {
            console.error("Scan failed:", error);
        } finally {
            setScanning(false);
        }
    };

    return (
        <header class="header">
            <div class="header__body">
                <div class="header__logo-block">
                    <img
                        src="/logo.png"
                        class="header__logo"
                        alt="OrsoNetwork logo"
                    />

                    <span class="header__brand">
                        OrsoNetwork
                    </span>
                </div>

                <button
                    class="header__scan"
                    onClick={handleScan}
                    disabled={scanning()}
                >
                    {scanning() ? "Scanning..." : "Scan"}
                </button>
            </div>
        </header>
    );
}