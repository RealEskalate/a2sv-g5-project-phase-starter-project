"use client"
import { useState } from "react";
import Navbar from "./Navbar";
import Sidebar from "./Sidebar";
import { usePathname, useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import { useGetCurrentUserQuery } from "@/lib/redux/api/bankApi";

const LayoutWrapper = ({children}: {children: React.ReactNode}) => {
    const session = useSession();
    const [isSidebarOpen, setIsSidebarOpen] = useState(true)
    const handleSidebar = () => {
        setIsSidebarOpen(!isSidebarOpen)

    }
    const path = usePathname()
    const router = useRouter()

    if(path.includes("auth")){
      return <div className="h-screen overflow-y-auto">{children}</div>
    }

    if(session.status === "unauthenticated"){
      router.push('/')
      return (
        <div>
          
          <Navbar onMenuClick={handleSidebar} />
          
          
           {children}
        </div>
      )
    }
    

  return (
    <div className="flex ">
        <div className={`h-full z-10  min-w-52 md:min-w-64 absolute sm:static sm:block bg-white  ${isSidebarOpen ? "block" : "hidden"} `}>
        <Sidebar setSidebarOpen={handleSidebar} sidebarOpen={isSidebarOpen} />

      </div>


      <div className="flex-grow-1  bg-white flex flex-col w-full h-screen  md:flex-grow-1 overflow-hidden  ">
        <div>
          <Navbar onMenuClick={handleSidebar} />
        </div>

        <div className="bg-background p-4 w-full flex-grow overflow-y-auto">
          {children}
        </div>

      </div>
    </div>
  )
}

export default LayoutWrapper