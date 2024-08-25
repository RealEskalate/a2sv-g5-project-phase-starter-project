"use client";

import React from "react";
import { useRouter, usePathname } from "next/navigation";

interface MenuItem {
  label: string;
  url: string;
}

const Navigation: React.FC = () => {
  const router = useRouter();
  const pathname = usePathname();

  const menuItems: MenuItem[] = [
    { label: "Edit Profile", url: "/settings/editprofile" },
    { label: "Preferences", url: "/settings/preference" },
    { label: "Security", url: "/settings/security" },
  ];

  const isActive = (path: string) => pathname === path;

  return (
    <div className=" w-full max-h-16 flex font-Inter justify-between items-center border-b dark:border-gray-500 xs:mt-6 sm:mt-2 pt-4 px-2 min-h-6">
      <div className="flex space-x-8 px-7">
        {menuItems.map((item, index) => (
          <button
            key={index}
            onClick={() => router.push(item.url)}
            className={`border-b-[6px] rounded cursor-pointer w-24 dark:text-white ${
              isActive(item.url) ? "border-[#1814F3]" : "border-transparent"
            }`}
          >
            {item.label}
          </button>
        ))}
      </div>
    </div>
  );
};

export default Navigation;
