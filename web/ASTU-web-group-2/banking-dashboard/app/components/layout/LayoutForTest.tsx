import React, { useState } from "react";
import Sidebar from "./sidebar/Sidebar";
import MenuBarMobile from "./navbar/Navbar";

export default function LayoutForSidebarAndNavbar() {

  const [showSidebar, setShowSidebar] = useState(false);

  return (
    <div className="min-h-screen">
      <div className="flex">
        <MenuBarMobile setter={setShowSidebar} />
        <Sidebar show={showSidebar} setter={setShowSidebar} />
      </div>
    </div>
  );
}
