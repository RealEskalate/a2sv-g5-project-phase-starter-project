"use client";

import React, { useMemo } from "react";
import Image from "next/image";
import { usePathname, useRouter } from "next/navigation";

const Sidebar = () => {
  const pathname = usePathname();
  const router = useRouter();

  // Memoizing the isActive function to prevent unnecessary recalculations
  const isActive = useMemo(
    () => (path: string) => pathname === path,
    [pathname]
  );

  const menuItems = [
    {
      label: "Dashboard",
      url: "/",
      icon: "/assets/home-icon.svg",
      active: "/assets/home-icon-active.svg",
    },
    {
      label: "Transactions",
      url: "/transaction",
      icon: "/assets/transfer-icon.svg",
      active: "/assets/transfer-icon-active.svg",
    },
    {
      label: "Accounts",
      url: "/account",
      icon: "/assets/account-icon.svg",
      active: "/assets/account-icon-active.svg",
    },
    {
      label: "Investments",
      url: "/investment",
      icon: "/assets/invest-icon.svg",
      active: "/assets/invest-icon-active.svg",
    },
    {
      label: "Credit Cards",
      url: "/credit-cards",
      icon: "/assets/card-icon.svg",
      active: "/assets/card-icon-active.svg",
    },
    {
      label: "Loans",
      url: "/loan",
      icon: "/assets/loan-icon.svg",
      active: "/assets/loan-icon-active.svg",
    },
    {
      label: "Services",
      url: "/service",
      icon: "/assets/service-icon.svg",
      active: "/assets/service-icon-active.svg",
    },
    {
      label: "My Privileges",
      url: "/privilages",
      icon: "/assets/privi-icon.svg",
      active: "/assets/privi-icon-active.svg",
    },
    {
      label: "Setting",
      url: "/setting",
      icon: "/assets/setting-icon.svg",
      active: "/assets/setting-icon-active.svg",
    },
  ];

  return (
    <div className="py-6 px-5 w-[99.6%] h-screen flex flex-col gap-8 border-r border-r-[#E6EFF5] bg-white">
      <div className="flex gap-2 px-6">
        <Image src="/assets/logo.svg" alt="logo" width={36} height={36} />
        <h1 className="text-2xl font-extrabold text-[#343C6A]">BankDash.</h1>
      </div>

      {/* Menu */}
      <div className="w-full p-3 px-6 flex flex-col gap-2 text-base font-medium xs:hidden lg:flex">
        {menuItems.map((item, index) => (
          <button
            key={index}
            onClick={() => router.push(item.url)}
            className="flex items-center gap-x-6 relative py-3"
          >
            <div
              className={`${
                isActive(item.url) ? "visible" : "hidden"
              } flex w-6 h-[45px] rounded-[32px] bg-[#1814F3] absolute left-[-60px]`}
            ></div>
            <Image
              src={isActive(item.url) ? item.active : item.icon}
              alt={item.label}
              width={20}
              height={20}
            />
            <div
              className={`${
                isActive(item.url) ? "text-[#1814F3]" : "text-[#B1B1B1]"
              } hover:text-[#1814F3]`}
            >
              {item.label}
            </div>
          </button>
        ))}
      </div>
    </div>
  );
};

export default Sidebar;
