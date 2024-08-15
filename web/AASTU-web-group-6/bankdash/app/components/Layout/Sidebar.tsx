"use client";

import React from "react";
import Image from "next/image";
import { usePathname } from "next/navigation";
import Link from "next/link";

const Sidebar = () => {
  const pathname = usePathname();

  const isActive = (path: string) => pathname === path;
  return (
    <div className="py-6 px-5 flex flex-col gap-8 ablsoute left-0 border-r border-r-[#E6EFF5] bg-white">
      <div className="flex gap-2">
        <Image src="/assets/logo.svg" alt="logo" width={36} height={36} />
        <h1 className="text-3xl font-bold text-[#343C6A]">BankDash</h1>
      </div>

      {/* Menu */}
      <div className="px-8 flex flex-col gap-y-8 font-medium">
        <div
          className={`flex items-center gap-x-6  ${
            isActive("/") ? "text-[#1814F3]" : "text-[#b1b1b1]"
          }`}
        >
          <Image src="/assets/home 2.svg" alt="logo" width={25} height={25}/>
          <h2 className="text-basex font-medium ">Dashboard</h2>
        </div>

        <div
          className={`flex items-center gap-x-6  ${
            isActive("/transaction") ? "text-[#1814F3]" : "text-[#b1b1b1]"
          }`}
        >
          <Image
            src="/assets/transfer 1.svg"
            alt="logo"
            width={25}
            height={25}
          />
          <h2 className="text-base   ">Transactions</h2>
        </div>
        <div
          className={`flex items-center gap-x-6  ${
            isActive("/Accounts") ? "text-[#1814F3]" : "text-[#b1b1b1]"
          }`}
        >
          <Image src="/assets/user 3 1.svg" alt="logo" width={25} height={25} />
          <h2 className="text-base   ">Accounts</h2>
        </div>
        <div
          className={`flex items-center gap-x-6  ${
            isActive("/Investments") ? "text-[#1814F3]" : "text-[#b1b1b1]"
          }`}
        >
          <Image
            src="/assets/economic-investment 1.svg"
            alt="logo"
            width={25}
            height={25}
          />
          <h2 className="text-base   ">Investments</h2>
        </div>
        <div
          className={`flex items-center gap-x-6  ${
            isActive("/credit-cards") ? "text-[#1814F3]" : "text-[#b1b1b1]"
          }`}
        >
          <Image
            src="/assets/credit-card 1.svg"
            alt="logo"
            width={25}
            height={25}
          />
          <h2 className="text-base   ">Credit Cards</h2>
        </div>
        <div
          className={`flex items-center gap-x-6  ${
            isActive("/loan") ? "text-[#1814F3]" : "text-[#b1b1b1]"
          }`}
        >
          <Image src="/assets/loan 1.svg" alt="logo" width={25} height={25} />
          <h2 className="text-base">Loans</h2>
        </div>
        <div
          className={`flex items-center gap-x-6  ${
            isActive("/service") ? "text-[#1814F3]" : "text-[#b1b1b1]"
          }`}
        >
          <Image
            src="/assets/service 1.svg"
            alt="logo"
            width={25}
            height={25}
          />
          <h2 className="text-base">Services</h2>
        </div>
        <div
          className={`flex items-center gap-x-6  ${
            isActive("/privilages") ? "text-[#1814F3]" : "text-[#b1b1b1]"
          }`}
        >
          <Image
            src="/assets/econometrics 1.svg"
            alt="logo"
            width={25}
            height={25}
          />
          <h2 className="text-base">My Privileges</h2>
        </div>
        <div
          className={`flex items-center gap-x-6  ${
            isActive("/setting") ? "text-[#1814F3]" : "text-[#b1b1b1]"
          }`}
        >
          <Image
            src="/assets/settings solid 1.svg"
            alt="logo"
            width={25}
            height={25}
          />
          <h2 className="text-base   ">Setting</h2>
        </div>
      </div>
    </div>
  );
};

export default Sidebar;
