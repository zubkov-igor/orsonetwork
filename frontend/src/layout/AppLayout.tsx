import { route } from "../store/router";

import Header from "../components/Header";
import Aside from "../components/Aside";
import Footer from "../components/Footer";

import Dashboard from "../pages/Dashboard";
import Topology from "../pages/Topology";
import Devices from "../pages/Devices";
import Settings from "../pages/Settings";

import { sidebarExpanded } from "../store/ui";

function renderPage() {
  switch (route()) {
    case "dashboard":
      return <Dashboard />;
    case "topology":
      return <Topology />;
    case "devices":
      return <Devices />;
    case "settings":
      return <Settings />;
    default:
      return <Dashboard />;
  }
}

export default function AppLayout() {
  return (
    <div
  class="app__layout"
  classList={{
    expanded: sidebarExpanded(),
  }}
>
      <Header />
      <Aside />

      <main class="main">
        {renderPage()}
      </main>

      <Footer />
    </div>
  );
}