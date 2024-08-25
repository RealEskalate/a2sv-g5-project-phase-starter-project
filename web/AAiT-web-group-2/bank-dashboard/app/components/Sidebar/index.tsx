"use client";

import React, { useState } from "react";
import { usePathname } from "next/navigation";
import Link from "next/link";
import Image from "next/image";
import SidebarItem from "./SidebarItem";
import Home from "../../../public/iconI.png";
import Transaction from "../../../public/iconII.png";
import User from "../../../public/iconIII.png";
import Dollars from "../../../public/iconIV.png";
import CreditCard from "../../../public/iconV.png";
import LocalAtm from "../../../public/iconVI.png";
import Services from "../../../public/iconVII.png";
import Pereferences from "../../../public/iconVIII.png";
import Settings from "../../../public/iconVIIII.png";
import { PiSignOut } from "react-icons/pi";
import { useGetCurrentUserQuery } from "@/lib/redux/api/bankApi";
import { signOut } from "next-auth/react";

interface SidebarProps {
  sidebarOpen: boolean;
  setSidebarOpen: (arg: boolean) => void;
}

const menuGroups = [
  {
    name: "MENU",
    menuItems: [
      { icon: Home.src, label: "Dashboard", route: "/dashboard" },
      { icon: Transaction.src, label: "Transactions", route: "/transactions" },
      { icon: User.src, label: "Accounts", route: "/account" },
      { icon: Dollars.src, label: "Investments", route: "/investments" },
      { icon: CreditCard.src, label: "Credit Cards", route: "/credit-cards" },
      { icon: LocalAtm.src, label: "Loans", route: "/loans" },
      { icon: Services.src, label: "Services", route: "/services" },
      { icon: Settings.src, label: "Setting", route: "/settings" },
     
    ],
  },
];

const Sidebar: React.FC<SidebarProps> = ({ sidebarOpen, setSidebarOpen }) => {
  const pathname = usePathname();
  const [pageName, setPageName] = useState("dashboard")

  return (
      <aside
        className={`pt-6 border-r-2 border-[#E6EFF5]  left-0 min-w-52 md:w-64 sm:top-0 z-10 flex h-full flex-col overflow-y-hidden bg-white duration-300 ease-linear  sm:translate-x-0 ${
          sidebarOpen ? "translate-x-0" : "-translate-x-full"
        }`}
      >
        <div className="flex  items-center justify-between gap-2 px-6 py-5.5 lg:py-6.5">
          <Link href="/">
            <Image
              width={176}
              height={32}
              src={"/Logo.png"}
              alt="Logo"
              priority
            />
          </Link>
          <button
            onClick={() => setSidebarOpen(!sidebarOpen)}
            aria-controls="sidebar"
            className="block sm:hidden"
          >
            <svg
              className="fill-current"
              width="20"
              height="18"
              viewBox="0 0 20 18"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M19 8.175H2.98748L9.36248 1.6875C9.69998 1.35 9.69998 0.825 9.36248 0.4875C9.02498 0.15 8.49998 0.15 8.16248 0.4875L0.399976 8.3625C0.0624756 8.7 0.0624756 9.225 0.399976 9.5625L8.16248 17.4375C8.31248 17.5875 8.53748 17.7 8.76248 17.7C8.98748 17.7 9.17498 17.625 9.36248 17.475C9.69998 17.1375 9.69998 16.6125 9.36248 16.275L3.02498 9.8625H19C19.45 9.8625 19.825 9.4875 19.825 9.0375C19.825 8.55 19.45 8.175 19 8.175Z"
                fill=""
              />
            </svg>
          </button>
        </div>

        <div className="no-scrollbar flex flex-col overflow-y-auto duration-300 ease-linear">
          {/* <!-- Sidebar Menu --> */}
          <nav className="mt-5 px-4 py-4 lg:mt-9 lg:px-6">
            {menuGroups.map((group, groupIndex) => (
              <div key={groupIndex}>
                
                <ul className="mb-6 flex flex-col gap-1.5">
                  {group.menuItems.map((menuItem, menuIndex) => (
                    <SidebarItem
                      key={menuIndex}
                      item={menuItem}
                      pageName={pageName}
                      setPageName={setPageName}
                    />
                  ))}
                </ul>
              </div>
            ))}
            <div className="flex items-center font-medium text-lg gap-2.5 px-4 text-gray-500">
            <PiSignOut />
            <button onClick={() => signOut()}>Sign out</button>
            </div>
          </nav>
        </div>
      </aside>
      
    
  );
};

export default Sidebar;
