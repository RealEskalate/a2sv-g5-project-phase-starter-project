"use client";
import React from "react";
import { sidebarListItems } from "./sidebarListItems";
import Link from "next/link";
import Image from "next/image";
import { usePathname } from "next/navigation";

const Sidebar = ({ isOpen = false, toggleSidebar = () => { } }: { isOpen?: boolean; toggleSidebar?: () => void; }) => {
  const pathname = usePathname();

  return (
    <div
      className={`fixed top-0 left-0 border-r border-[#E6EFF5] w-[240px] bg-white z-50 transition-transform duration-300 ease-in-out ${
        isOpen ? 'translate-x-0' : '-translate-x-full'
      } sm:translate-x-0 sm:w-[240px] sm:block flex flex-col`}
      style={{ height: '100vh' }} // Ensure sidebar takes full viewport height
    >
      <div className="flex items-center p-4">
        <Image
          src={"/assets/navbar/credit-card.svg"}
          width={36}
          height={36}
          alt="bankDash logo"
          className="mr-3"
        />
        <p className="font-black lg:text-[25px] text-[#343C6A] md:text-[25px] sm:text-[20px]">
          BankDash.
        </p>
        <button
          className="sm:hidden ml-auto"
          onClick={toggleSidebar}
          aria-label="Toggle Sidebar"
        >
          <Image
            src={"/assets/navbar/hamburger.svg"}
            width={25}
            height={25}
            alt="hamburger"
          />
        </button>
      </div>
      <div className="flex flex-col flex-1 overflow-y-auto pt-2">
        {sidebarListItems.map((item, index) => (
          <Link key={index} href={item.path}>
            <div
              className={`flex items-center gap-8  py-1 ${
                pathname === item.path ? "text-[#2D60FF]" : "text-[#B1B1B1]"
              }`}
            >
              <div
                className={`h-[50px] w-[5px] ${
                  pathname === item.path ? "bg-[#2D60FF]" : "bg-transparent"
                } rounded-r-lg`}
              ></div>
              <div className="flex gap-8 items-center">
                <Image
                  src={pathname === item.path ? item.activeIcon : item.icon}
                  alt={item.name}
                  width={20}
                  height={20}
                />
                <h1
                  className={`font-medium text-[16px] ${
                    pathname === item.path ? "text-[#2D60FF]" : "text-[#B1B1B1]"
                  }`}
                >
                  {item.name}
                </h1>
              </div>
            </div>
          </Link>
        ))}
      </div>
    </div>
  );
};

export default Sidebar;
