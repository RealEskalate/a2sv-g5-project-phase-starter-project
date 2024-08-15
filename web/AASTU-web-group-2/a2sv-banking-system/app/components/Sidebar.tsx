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
  MdOutlineCancel,
} from "react-icons/md";
import { RiHandCoinLine } from "react-icons/ri";
import { useRouter } from "next/navigation";
import SidebarElements from "./SidebarElements";

interface SidebarProps {
  toggle: boolean;
  handleClose: () => void;
}

const Sidebar = ({ toggle, handleClose }: SidebarProps) => {
  const router = useRouter();
  const [active, setActive] = useState("Dashboard");

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
      text: "Privileges",
      destination: "./privileges",
      icon: MdStar,
    },
    {
      id: 9,
      text: "Settings",
      destination: "./bankingSettings",
      icon: MdSettings,
    },
  ];

  const handleNav = (destination: string) => {
    router.push(destination);
  };
  const handleActive = (element:string) => {
    setActive(element);

  }

  return (
    <>
      <div className="hidden md:flex md:flex-col md:gap-10 py-7 border-r h-svh">
        <div className="px-5 py-2">
          <Image src="/Logo.png" width={183} height={36} alt="Logo" />
        </div>
        <SidebarElements
          handleNav={handleNav}
          handleActive = {handleActive}
          elements={elements}
          active={active}
        />
      </div>

      {toggle && (
        <div
          className={`fixed top-0 left-0 w-64 bg-white shadow-black h-full transform ${
            toggle ? "translate-x-0" : "-translate-x-full"
          } transition-transform ease-in-out duration-1000 flex flex-col gap-5 py-5 px-5`}
        >
          <div className="flex justify-end">
            <button onClick={handleClose} className="cursor-pointer">
              <MdOutlineCancel className="text-2xl" />
            </button>
          </div>
          <div className="px-5 py-2">
            <Image src="/Logo.png" width={183} height={36} alt="Logo" />
          </div>
          <SidebarElements
            handleNav={handleNav}
            elements={elements}
            active={active}
          />
        </div>
      )}
    </>
  );
};

export default Sidebar;
