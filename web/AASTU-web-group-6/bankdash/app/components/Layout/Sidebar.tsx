"use client";

import React from "react";
import Image from "next/image";
import { usePathname } from "next/navigation";
import Link from "next/link";

const Sidebar = () => {
  const pathname = usePathname();

  const isActive = (path: string) => pathname === path;

  return (
    <div className="py-6 px-5 flex flex-col gap-8 border-r border-r-[#E6EFF5] bg-white">
      <div className="flex gap-2">
        <Image src="/assets/logo.svg" alt="logo" width={36} height={36} />
        <h1 className="text-3xl font-bold text-[#343C6A]">BankDash</h1>
      </div>

      {/* Menu */}
      <div className="px-8 flex flex-col gap-y-8 font-medium">
        <Link href="/" className="flex items-center gap-x-6">
          <Image src="/assets/home 2.svg" alt="logo" width={25} height={25} />
          <h2
            className={`text-base font-medium ${
              isActive("/") ? "text-[#1814F3]" : "text-[#b1b1b1]"
            } hover:text-blue-500`}
          >
            Dashboard
          </h2>
        </Link>

        <Link href="/transaction" className="flex items-center gap-x-6">
          <Image
            src="/assets/transfer 1.svg"
            alt="logo"
            width={25}
            height={25}
          />
          <h2
            className={`text-base ${
              isActive("/transaction") ? "text-[#1814F3]" : "text-[#b1b1b1]"
            } hover:text-blue-500`}
          >
            Transactions
          </h2>
        </Link>

        <Link href="/accounts" className="flex items-center gap-x-6">
          <Image
            src="/assets/user 3 1.svg"
            alt="logo"
            width={25}
            height={25}
          />
          <h2
            className={`text-base ${
              isActive("/accounts") ? "text-[#1814F3]" : "text-[#b1b1b1]"
            } hover:text-blue-500`}
          >
            Accounts
          </h2>
        </Link>

        <Link href="/investments" className="flex items-center gap-x-6">
          <Image
            src="/assets/economic-investment 1.svg"
            alt="logo"
            width={25}
            height={25}
          />
          <h2
            className={`text-base ${
              isActive("/investments") ? "text-[#1814F3]" : "text-[#b1b1b1]"
            } hover:text-blue-500`}
          >
            Investments
          </h2>
        </Link>

        <Link href="/credit-cards" className="flex items-center gap-x-6">
          <Image
            src="/assets/credit-card 1.svg"
            alt="logo"
            width={25}
            height={25}
          />
          <h2
            className={`text-base ${
              isActive("/credit-cards") ? "text-[#1814F3]" : "text-[#b1b1b1]"
            } hover:text-blue-500`}
          >
            Credit Cards
          </h2>
        </Link>

        <Link href="/loan" className="flex items-center gap-x-6">
          <Image src="/assets/loan 1.svg" alt="logo" width={25} height={25} />
          <h2
            className={`text-base ${
              isActive("/loan") ? "text-[#1814F3]" : "text-[#b1b1b1]"
            } hover:text-blue-500`}
          >
            Loans
          </h2>
        </Link>

        <Link href="/service" className="flex items-center gap-x-6">
          <Image
            src="/assets/service 1.svg"
            alt="logo"
            width={25}
            height={25}
          />
          <h2
            className={`text-base ${
              isActive("/service") ? "text-[#1814F3]" : "text-[#b1b1b1]"
            } hover:text-blue-500`}
          >
            Services
          </h2>
        </Link>

        <Link href="/privilages" className="flex items-center gap-x-6">
          <Image
            src="/assets/econometrics 1.svg"
            alt="logo"
            width={25}
            height={25}
          />
          <h2
            className={`text-base ${
              isActive("/privilages") ? "text-[#1814F3]" : "text-[#b1b1b1]"
            } hover:text-blue-500`}
          >
            My Privileges
          </h2>
        </Link>

        <Link href="/setting" className="flex items-center gap-x-6">
          <Image
            src="/assets/settings solid 1.svg"
            alt="logo"
            width={25}
            height={25}
          />
          <h2
            className={`text-base ${
              isActive("/setting") ? "text-[#1814F3]" : "text-[#b1b1b1]"
            } hover:text-blue-500`}
          >
            Setting
          </h2>
        </Link>
      </div>
    </div>
  );
};

export default Sidebar;
