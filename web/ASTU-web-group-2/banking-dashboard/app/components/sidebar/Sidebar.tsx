"use client";
import React from "react";
import { sidebarListItems } from "./sidebarListItems";
import Link from "next/link";
import Image from "next/image";
import { usePathname } from "next/navigation";

const Sidebar = () => {
  const pathname = usePathname();
  console.log(pathname);

  return (
    <div className="hidden sm:block absolute top-0 left-0 bottom-0 border-r border-[#E6EFF5] w-[240px]">
      <div className="flex flex-col w-[189px] h-[420px]  pt-[114px]">
        {sidebarListItems.map((item, index) => (
          <div key={index} className="flex gap-6">
            <Link href={item.path}>
              <div
                className={`flex items-center gap-5 pt-3 ${
                  pathname === item.path ? "text-[#2D60FF]" : "text-[#B1B1B1]"
                }`}
              >
                <div
                  className={`h-[50px] w-[5px] ${
                    pathname === item.path ? "bg-[#2D60FF]" : "bg-transparent"
                  } rounded-r-lg`}
                ></div>
                <div className="flex gap-5">
                  <img
                    src={pathname === item.path ? item.activeIcon : item.icon}
                    alt={item.name}
                    width={20}
                    height={20}
                  />
                  <h1
                    className={`font-medium text-[16px] ${
                      pathname === item.path
                        ? "text-[#2D60FF]"
                        : "text-[#B1B1B1]"
                    }`}
                  >
                    {item.name}
                  </h1>
                </div>
              </div>
            </Link>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Sidebar;
