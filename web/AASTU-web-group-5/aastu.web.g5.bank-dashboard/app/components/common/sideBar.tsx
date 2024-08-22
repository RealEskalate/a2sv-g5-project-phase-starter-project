"use client";

import React, { useState } from "react";
import Image from "next/image";
import { usePathname, useRouter } from "next/navigation";
import { FaTimes } from "react-icons/fa";

import creditCard from "/public/assets/icons/credit-card 1.svg";
import econometrics from "/public/assets/icons/econometrics 1.svg";
import economicInvestment from "/public/assets/icons/economic-investment 1.svg";
import enabledHome from "/public/assets/icons/home 2.svg";
import mainIcon from "/public/assets/icons/iconfinder_vector_65_09_473792 1.png";
import loan from "/public/assets/icons/loan 1.svg";
import service from "/public/assets/icons/service 1.svg";
import settingsSolid from "/public/assets/icons/settings solid 1.svg";
import transfer from "/public/assets/icons/transfer 1.svg";
import user from "/public/assets/icons/user 3 1.svg";

import enabledCreditCard from "/public/assets/icons/enabled/credit-card 1.svg";
import enabledEconometrics from "/public/assets/icons/enabled/econometrics 1.svg";
import enabledEconomicInvestment from "/public/assets/icons/enabled/economic-investment 1.svg";
import home from "/public/assets/icons/enabled/Vector.svg";
import enabledLoan from "/public/assets/icons/enabled/loan 1.svg";
import enabledService from "/public/assets/icons/enabled/service 1.svg";
import enabledSettingsSolid from "/public/assets/icons/enabled/settings solid 1.svg";
import enabledTransfer from "/public/assets/icons/enabled/transfer 1.svg";
import enabledUser from "/public/assets/icons/enabled/user 3 1.svg";

const primary_2 = "rgba(52, 60, 106, 1)";
const primary_3 = "rgba(45, 96, 255, 1)";
const sidecolor = "#B1B1B1";

const SideBar = ({ isSidebarVisible, toggleSidebar }: { isSidebarVisible: boolean, toggleSidebar: () => void }) => {
    const pathname = usePathname() || 'Dashboard';
    const [enabled, setEnabled] = useState<string>(pathname);
    
    const router = useRouter();

    const handleIconClick = (option: string, path: string) => {
        setEnabled(option);
        router.push(path);
    };

    return (
        <div className={`pl-[38px]   ${isSidebarVisible ? 'block' : 'hidden'} sm:block flex items-center flex-col min-w-full sm:w-auto`}>
            <div className="flex gap-[9px] items-center relative" style={{ height: "101px" }}>
                <Image src={mainIcon} alt="BankDash Logo" className="h-[36px] w-[36px]" />
                <div className="font-bold" style={{ color: primary_2 }}>BankDash.</div>
                <button className="sm:hidden ml-auto absolute -right-8" onClick={toggleSidebar}>
                    <FaTimes size={24} />
                </button>
            </div>


            <div className="flex flex-col gap-[42px]">
                <SidebarItem
                    isEnabled={enabled === "home"}
                    onClick={() => handleIconClick("home", "/Dashboard")}
                    icon={home}
                    enabledIcon={enabledHome}
                    label="Dashboard"
                />
                <SidebarItem
                    isEnabled={enabled === "transfer"}
                    onClick={() => handleIconClick("transfer", "/Transaction")}
                    icon={transfer}
                    enabledIcon={enabledTransfer}
                    label="Transactions"
                />
                <SidebarItem
                    isEnabled={enabled === "user"}
                    onClick={() => handleIconClick("user", "/Accounts")}
                    icon={user}
                    enabledIcon={enabledUser}
                    label="Accounts"
                />
                <SidebarItem
                    isEnabled={enabled === "economicInvestment"}
                    onClick={() => handleIconClick("economicInvestment", "/Investments")}
                    icon={economicInvestment}
                    enabledIcon={enabledEconomicInvestment}
                    label="Investments"
                />
                <SidebarItem
                    isEnabled={enabled === "creditCard"}
                    onClick={() => handleIconClick("creditCard", "/CreditCards")}
                    icon={creditCard}
                    enabledIcon={enabledCreditCard}
                    label="Credit Cards"
                />
                <SidebarItem
                    isEnabled={enabled === "loan"}
                    onClick={() => handleIconClick("loan", "/Loan")}
                    icon={loan}
                    enabledIcon={enabledLoan}
                    label="Loans"
                />
                <SidebarItem
                    isEnabled={enabled === "service"}
                    onClick={() => handleIconClick("service", "/Services")}
                    icon={service}
                    enabledIcon={enabledService}
                    label="Services"
                />
                <SidebarItem
                    isEnabled={enabled === "econometrics"}
                    onClick={() => handleIconClick("econometrics", "/privileges")}
                    icon={econometrics}
                    enabledIcon={enabledEconometrics}
                    label="My Privileges"
                />
                <SidebarItem
                    isEnabled={enabled === "settings"}
                    onClick={() => handleIconClick("settings", "/settings")}
                    icon={settingsSolid}
                    enabledIcon={enabledSettingsSolid}
                    label="Settings"
                />
            </div>
        </div>
    );
};

const SidebarItem = ({ isEnabled, onClick, icon, enabledIcon, label }: { isEnabled: boolean, onClick: () => void, icon: string, enabledIcon: string, label: string }) => (
    <div className="flex gap-[23px]" onClick={onClick}>
        <Image src={isEnabled ? enabledIcon : icon} alt={`${label} Icon`} className="h-[25px] w-[25px]" />
        <div>{label}</div>
    </div>
);

export default SideBar;
