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
    <div className="flex">
        <div className={` min-w-52 md:min-w-64 absolute sm:static sm:block bg-white  ${isSidebarOpen ? "block" : "hidden"} `}>
        <Sidebar setSidebarOpen={handleSidebar} sidebarOpen={isSidebarOpen} />

      </div>


      <div className="flex-grow-1  bg-white flex flex-col w-full md:flex-grow-1   ">
        <Navbar onMenuClick={handleSidebar} />

        <div className="bg-background min-h-screen w-full">
          {children}

        </div>

      </div>
    </div>
  )
}

export default LayoutWrapper