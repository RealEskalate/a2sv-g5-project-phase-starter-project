import React, { useState } from "react";
import Image from "next/image";
import {
  MdHome,
  MdAttachMoney,
  MdAccountBalance,
  MdAssessment,
  MdCreditCard,
  MdPayment,
  MdStar,
  MdSettings,
} from "react-icons/md";
import { FaTimes } from "react-icons/fa";
import { RiHandCoinLine } from "react-icons/ri";
import { usePathname, useRouter } from "next/navigation";
import SidebarElements from "./SidebarElements";

interface SidebarProps {
  toggle: boolean;
  handleClose: () => void;
}
const Sidebar = ({ toggle, handleClose }: SidebarProps) => {
  
const map = new Map();
map.set("/", "Dashboard");
map.set("/dashboard", "Dashboard");
map.set("/transactions", "Transactions");
map.set("/accounts", "Accounts");
map.set("/investments", "Investments");
map.set("/creditCards", "Credit Cards");
map.set("/loans", "Loans");
map.set("/bankingServices", "Services");
map.set("/bankingSettings", "Settings");

const pathName = usePathname();
const router = useRouter();
const [active, setActive] = useState(map.get(pathName));

  const elements = [
    {
      id: 1,
      text: "Dashboard",
      destination: "./dashboard",
      icon: MdHome,
    },
    {
      id: 2,
      text: "Transactions",
      destination: "./transactions",
      icon: MdAttachMoney,
    },
    {
      id: 3,
      text: "Accounts",
      destination: "./accounts",
      icon: MdAccountBalance,
    },
    {
      id: 4,
      text: "Investments",
      destination: "./investments",
      icon: MdAssessment,
    },
    {
      id: 5,
      text: "Credit Cards",
      destination: "./creditCards",
      icon: MdCreditCard,
    },
    {
      id: 6,
      text: "Loans",
      destination: "./loans",
      icon: RiHandCoinLine,
    },
    {
      id: 7,
      text: "Services",
      destination: "./bankingServices",
      icon: MdPayment,
    },
    {
      id: 8,
      text: "Settings",
      destination: "./bankingSettings",
      icon: MdSettings,
    },
  ];

  const handleNav = async (destination: string) => {
    router.push(destination);
  };
  const handleActive = (element: string) => {
    setActive(element);
    handleClose();
  };

  return (
    <>
      <div className="hidden md:flex md:flex-col md:gap-5 py-7 border-r h-svh sticky top-0 dark:bg-[#020817]">
        <div className="px-5 py-2">
          <Image src="/Logo.png" width={183} height={36} alt="Logo" />
        </div>
        <SidebarElements
          handleNav={handleNav}
          handleActive={handleActive}
          elements={elements}
          active={active}
        />
      </div>

      {toggle && (
        <div className="md:hidden flex">
          <div
            className={`fixed top-0 left-0 w-80 bg-white shadow-black h-full transform transition-transform dark:bg-[#020817] ${
              toggle ? "translate-x-0" : "-translate-x-full"
            }  ease-in-out duration-1000 flex flex-col px-5`}
          >
            <div className="flex flex-col justify-between">
              <button
                onClick={handleClose}
                className="cursor-pointer text-[#2D60FF] flex justify-end mt-5"
              >
                <FaTimes className="text-3xl" />
              </button>
              <div className="px-3 mt-3 mb-4">
                <Image src="/Logo.png" width={183} height={36} alt="Logo" />
              </div>
            </div>
            <SidebarElements
              handleNav={handleNav}
              handleActive={handleActive}
              elements={elements}
              active={active}
            />
          </div>
        </div>
      )}
    </>
  );
};

export const SidebarLoading = () => {
  return (
    <div className="hidden md:flex md:flex-col md:gap-5 py-7 border-r h-svh sticky top-0 animate-pulse dark:bg-[#050914] dark:border-[#333B69]">
      <div className="px-5 py-2">
        <div className="bg-gray-300 dark:bg-[#333B69] w-44 h-9 rounded"></div>
      </div>
      <div className="flex flex-col gap-4 px-5">
        <div className="bg-gray-300 dark:bg-[#333B69] w-40 h-20 rounded"></div>
        <div className="bg-gray-300 dark:bg-[#333B69] w-40 h-20 rounded"></div>
        <div className="bg-gray-300 dark:bg-[#333B69] w-40 h-20 rounded"></div>
        <div className="bg-gray-300 dark:bg-[#333B69] w-40 h-20 rounded"></div>
        <div className="bg-gray-300 dark:bg-[#333B69] w-40 h-20 rounded"></div>
      </div>
    </div>
  );
};

export default Sidebar;
