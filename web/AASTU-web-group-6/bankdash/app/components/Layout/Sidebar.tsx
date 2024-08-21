"use client";

import React, { useMemo } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faBars,
  faClose,
  faRightFromBracket,
  faSearch,
  faMoon,
  faSun,
} from "@fortawesome/free-solid-svg-icons";
import { faBell } from "@fortawesome/free-regular-svg-icons";
import Image from "next/image";
import { usePathname, useRouter } from "next/navigation";
import { useAppSelector, useAppDispatch } from "@/app/Redux/store/store";
import { toggleDarkMode } from "@/app/Redux/slices/darkModeSlice";

const Sidebar = ({
  isOpen,
  closeSidebar,
}: {
  isOpen: boolean;
  closeSidebar: () => void;
}) => {
  const pathname = usePathname();
  const router = useRouter();
  const dispatch = useAppDispatch();

  const isDarkMode = useAppSelector((state) => state.darkMode.darkMode);
  const onDarkMode = () => {
    dispatch(toggleDarkMode());
  };

  // Memoizing the isActive function to prevent unnecessary recalculations
  const isActive = useMemo(
    () =>
      (path: string, additionalPaths: string[] = []) => {
        if (pathname === path) return true;
        return additionalPaths.some(
          (additionalPath) => pathname === additionalPath
        );
      },
    [pathname]
  );
  const flag = 1;

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
      label: "Settings",
      url: "/settings/editprofile",
      icon: "/assets/setting-icon.svg",
      active: "/assets/setting-icon-active.svg",
      additionalActivePaths: [
        "/settings/editprofile",
        "/settings/preference",
        "/settings/security",
      ],
    },
    {
      label: "LogOut",
      url: "/login",
      icon: "/assets/logout-icon.svg",
      active: "/assets/logout-icon-active.svg",
    },
  ];

  // Apply dark mode class directly
  const darkModeClass = isDarkMode ? "dark" : "";
  const logo = isDarkMode ? "/assets/logo-white.svg" : "/assets/logo-blue.svg";

  return (
    <div className="py-6 px-5 w-[99.6%] h-screen flex flex-col gap-8 border-r border-r-[#E6EFF5] border-white dark:border-r-gray-700 dark:bg-[#232328] relative">
      <div className="flex gap-2 px-[4%] relative">
        <Image
          src={logo || "/assets/logo-blue.svg"}
          alt="logo"
          width={36}
          height={36}
        />
        <h1 className="text-2xl font-extrabold text-[#343C6A] dark:text-white">
          BankDash.
        </h1>
        <button
          onClick={closeSidebar}
          className="bg-[#F5F7FA] rounded-[12px] border-2 border-solid border-slate-200 p-3 py-2 flex items-center left-60 hover:bg-[#d0e6f6] lg:hidden"
        >
          <FontAwesomeIcon icon={faClose} className="text-2xl text-gray-700" />
        </button>
      </div>

      {/* Menu */}
      <div className="p-3 px-6 xxs:py-1 flex overflow-y-auto overflow-x-hidden flex-col gap-2 text-base font-medium text-nowrap">
        {menuItems.map((item, index) => (
          <div key={index}>
            {/* Dark Mode Toggle Button for Small Devices */}
            {flag === index + 1 && (
              <button
                onClick={onDarkMode}
                className="flex w-full mb-3 mt-[-2px] items-center bg-[#F5F7FA] dark:bg-gray-700 rounded-full p-2 hover:bg-gray-200 dark:hover:bg-gray-600 lg:hidden"
              >
                <FontAwesomeIcon
                  icon={isDarkMode ? faSun : faMoon}
                  className="text-[#718EBF] dark:text-yellow-400 text-xl px-1"
                />
                <span className="ml-2 text-[#343C6A] dark:text-white">
                  {isDarkMode ? "Light Mode" : "Dark Mode"}
                </span>
              </button>
            )}
            <button
              onClick={() => {
                router.push(item.url);
                closeSidebar();
              }}
              className="flex items-center gap-x-6 relative py-3"
            >
              <div
                className={`${
                  isActive(item.url, item.additionalActivePaths || [])
                    ? "visible"
                    : "hidden"
                } flex z-20 w-6 h-[45px] rounded-[32px] bg-[#1814F3] absolute md:left-[-40px]`}
              ></div>
              <Image
                src={
                  isActive(item.url, item.additionalActivePaths || [])
                    ? item.active
                    : item.icon
                }
                alt={item.label}
                width={20}
                height={20}
              />
              <div
                className={`${
                  isActive(item.url)
                    ? "text-[#1814F3] dark:text-white"
                    : "text-[#B1B1B1]"
                } hover:text-[#1814F3]`}
              >
                {item.label}
              </div>
            </button>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Sidebar;
