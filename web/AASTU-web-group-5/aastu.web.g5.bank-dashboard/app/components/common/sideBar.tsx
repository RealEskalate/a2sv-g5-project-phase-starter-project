'use client'

import React from "react";
import Image from "next/image";
import { FaTimes } from "react-icons/fa";
import creditCard from "/public/assets/icons/credit-card 1.svg";
import econometrics from "/public/assets/icons/econometrics 1.svg";
import economicInvestment from "/public/assets/icons/economic-investment 1.svg";
import home from "/public/assets/icons/home 2.svg";
import mainIcon from "/public/assets/icons/iconfinder_vector_65_09_473792 1.png";
import loan from "/public/assets/icons/loan 1.svg";
import service from "/public/assets/icons/service 1.svg";
import settingsSolid from "/public/assets/icons/settings solid 1.svg";
import transfer from "/public/assets/icons/transfer 1.svg";
import user from "/public/assets/icons/user 3 1.svg";

const primary_2 = 'rgba(52, 60, 106, 1)';
const primary_3 = 'rgba(45, 96, 255, 1)';
const sidecolor = '#B1B1B1';

const SideBar = ({ isSidebarVisible, toggleSidebar }: { isSidebarVisible: boolean, toggleSidebar: () => void }) => {
    return (
        <div className={`pl-[38px] ${isSidebarVisible ? 'block' : 'hidden'} sm:block flex items-center flex-col min-w-full sm:w-auto `   }  >
            <div className="flex gap-[9px] items-center relative" style={{ height: "101px" }}>
                <Image src={mainIcon} alt="BankDash Logo" className="h-[36px] w-[36px]" />
                <div className="font-bold" style={{ color: primary_2 }}>BankDash.</div>
                <div></div>
                <button className="sm:hidden ml-auto absolute -right-8" onClick={toggleSidebar}>
                    <FaTimes size={24} />
                </button>
            </div>

            <div className="flex flex-col gap-[42px]">
                <div className="flex gap-[23px]">
                    <Image src={home} alt="Home Icon" className="h-[25px] w-[25px]" />
                    <div> Dashboard</div>
                </div>
                <div className="flex gap-[23px]">
                    <Image src={transfer} alt="Transfer Icon" className="h-[25px] w-[25px]" />
                    <div> Transactions</div>
                </div>
                <div className="flex gap-[23px]">
                    <Image src={user} alt="User Icon" className="h-[25px] w-[25px]" />
                    <div> Accounts</div>
                </div>
                <div className="flex gap-[23px]">
                    <Image src={economicInvestment} alt="Investments Icon" className="h-[25px] w-[25px]" />
                    <div> Investments</div>
                </div>
                <div className="flex gap-[23px]">
                    <Image src={creditCard} alt="Credit Card Icon" className="h-[25px] w-[25px]" />
                    <div> Credit Cards</div>
                </div>
                <div className="flex gap-[23px]">
                    <Image src={loan} alt="Loan Icon" className="h-[25px] w-[25px]" />
                    <div> Loans</div>
                </div>
                <div className="flex gap-[23px]">
                    <Image src={service} alt="Services Icon" className="h-[25px] w-[25px]" />
                    <div> Services</div>
                </div>
                <div className="flex gap-[23px]">
                    <Image src={econometrics} alt="Privileges Icon" className="h-[25px] w-[25px]" />
                    <div> My Privileges</div>
                </div>
                <div className="flex gap-[23px]">
                    <Image src={settingsSolid} alt="Settings Icon" className="h-[25px] w-[25px]" />
                    <div> Settings</div>
                </div>
            </div>
        </div>
    );
};

export default SideBar;