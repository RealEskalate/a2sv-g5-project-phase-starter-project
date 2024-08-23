"use client"
import { useState } from "react";
import Navbar from "./Navbar";
import Sidebar from "./Sidebar";

const LayoutWrapper = ({children}: {children: React.ReactNode}) => {
    const [isSidebarOpen, setIsSidebarOpen] = useState(true)
    const handleSidebar = () => {
        setIsSidebarOpen(!isSidebarOpen)

    }
  return (
    <div className="flex ">
        <div className={`h-full z-50  min-w-52 md:min-w-64 absolute sm:static sm:block bg-white  ${isSidebarOpen ? "block" : "hidden"} `}>
        <Sidebar setSidebarOpen={handleSidebar} sidebarOpen={isSidebarOpen} />

      </div>


      <div className="flex-grow-1  bg-white flex flex-col w-full h-screen  md:flex-grow-1   ">
        <div>
          <Navbar onMenuClick={handleSidebar} />
        </div>

        <div className="bg-background w-full flex-grow overflow-y-auto">
          {children}

        </div>

      </div>
    </div>
  )
}

export default LayoutWrapper