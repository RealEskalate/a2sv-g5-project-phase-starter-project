'use client';
import React, { useEffect, useState } from "react";
import dashboard from "../../../public/images/Dashboard.svg";
import dashboardNone from "../../../public/images/Dashboard-none.svg";
import transaction from "../../../public/images/transaction.svg";
import transactionBlue from "../../../public/images/transaction-blue.svg";
import account from "../../../public/images/user 3 1.svg";
import accountBlue from "../../../public/images/user-blue.svg";
import investments from "../../../public/images/economic-investment 1.svg";
import investmentsBlue from "../../../public/images/investment-blue.svg";
import credit from "../../../public/images/credit-card 1.svg";
import creditBlue from "../../../public/images/credit-blue.svg";
import loans from "../../../public/images/loan 1.svg";
import loansBlue from "../../../public/images/loan-blue.svg";
import service from "../../../public/images/service.svg";
import serviceBlue from "../../../public/images/service-blue.svg";
import setting from "../../../public/images/setting.svg";
import settingBlue from "../../../public/images/setting-blue.svg";
import logo from "../../../public/images/Logo.svg";
import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation"; // Import the usePathname hook
import { useSelector } from 'react-redux';
import { RootState } from "@/lib/redux/store";
import { useRouter } from "next/navigation";

interface Props {
  selected: string | string[];
}

const Sidebar = () => {
  const isSidebarVisible = useSelector((state: RootState) => state.menu.isSidebarVisible);
  const pathname = usePathname();
  const [selected, setSelected] = useState("Dashboard");
  const router = useRouter();

  useEffect(() => {
    if (pathname) {
      const pathToItem: { [key: string]: string } = {
        "/": "Dashboard",
        "/transactions": "Transactions",
        "/accounts": "Accounts",
        "/investments": "Investments",
        "/credit-cards": "Credit Cards",
        "/loans": "Loans",
        "/services": "Services",
        "/setting": "Setting",
      };

      setSelected(pathToItem[pathname] || ""); 
    }
  }, [pathname]);

  const sidebarStyle = `w-1/6 bg-white left-0 top-0 w-fit h-full fixed pt-1 ${isSidebarVisible ? '': 'hidden'} md:block z-20 ` 

  
  const handleRoute = (route: string) => {
    // use next/router to navigate to the route
    router.push(route);
  }

  return (
    <div className={sidebarStyle}>
      <ul className="flex mt-5 justify-center space-y-7 flex-col mr-5">
          <div className="w-2/3 ml-10 md:block hidden">
            <Image src={logo} className="ml-1" alt="LOGO" />
          </div>

          <li
            className={
              selected === "Dashboard"
                ? `text-[#2D60FF] relative font-semibold ] flex space-x-5`
                : `text-[#B1B1B1] relative flex space-x-5  ` + `  cursor-pointer`
            }
            onClick={() => handleRoute("/")}
          >
            {selected === "Dashboard" ? (
              <div>
                <div className="w-1 h-10 bg-blue-800 absolute left-0 -top-1rounded-r-lg"></div>
                <Image className="ml-10" src={dashboard} alt="dash" />
              </div>
            ) : (
              <Image className="ml-10" src={dashboardNone} alt="dash" />
            )}
            <h2>Dashboard</h2>
          </li>



        <li
          className={
            selected === "Transactions"
              ? `text-[#2D60FF] relative font-semibold ] flex space-x-5`
              : `text-[#B1B1B1] relative flex space-x-5  ` + `  cursor-pointer`
          }
        >
          {selected === "Transactions" ? (
            <div>
              <div className="w-1 h-10 bg-blue-800 absolute left-0 -top-1 rounded-r-lg"></div>
              <Image className="ml-10" src={transactionBlue} alt="dash" />
            </div>
          ) : (
            <Image className="ml-10" src={transaction} alt="dash" />
          )}
          <h2>Transactions</h2>
        </li>



        <li
          className={
            selected === "Accounts"
              ? `text-[#2D60FF] relative font-semibold ] flex space-x-5`
              : `text-[#B1B1B1] relative flex space-x-5 ` + `  cursor-pointer`
          }
        >
          {selected === "Accounts" ? (
            <div>
              <div className="w-1 h-10 bg-blue-800 absolute left-0 -top-1 rounded-r-lg"></div>
              <Image className="ml-10" src={accountBlue} alt="dash" />
            </div>
          ) : (
            <Image className="ml-10" src={account} alt="dash" />
          )}
          <h2>Accounts</h2>
        </li>



        <li
          className={
            selected === "Investments"
              ? `text-[#2D60FF] relative font-semibold ] flex space-x-5`
              : `text-[#B1B1B1] relative flex space-x-5 ` + `  cursor-pointer`
          }
        >
          {selected === "Investments" ? (
            <div>
              <div className="w-1 h-10 bg-blue-800 absolute left-0 -top-1 rounded-r-lg"></div>
              <Image className="ml-10" src={investmentsBlue} alt="dash" />
            </div>
          ) : (
            <Image className="ml-10" src={investments} alt="dash" />
          )}
          <h2>Investments</h2>
        </li>



        <li
          className={
            selected === "Credit Cards"
              ? `text-[#2D60FF] relative font-semibold ] flex space-x-5`
              : `text-[#B1B1B1] relative flex space-x-5  ` + `  cursor-pointer`
          }
        >
          {selected === "Credit Cards" ? (
            <div>
              <div className="w-1 h-10 bg-blue-800 absolute left-0 -top-1 rounded-r-lg"></div>
              <Image className="ml-10" src={creditBlue} alt="dash" />
            </div>
          ) : (
            <Image className="ml-10" src={credit} alt="dash" />
          )}
          <h2>Credit Cards</h2>
        </li>



        <li
          className={
            selected === "Loans"
              ? `text-[#2D60FF] relative font-semibold ]  flex space-x-5`
              : `text-[#B1B1B1] relative flex space-x-5 ` + `  cursor-pointer `
          }
        >
          {selected === "Loans" ? (
            <div>
              <div className="w-1 h-10 bg-blue-800  absolute left-0 --top-1 rounded-r-lg"></div>
              <Image className="ml-10" src={loansBlue} alt="dash" />
            </div>
          ) : (
            <Image className="ml-10" src={loans} alt="dash" />
          )}
          <h2>Loans</h2>
        </li>



        <li
          className={
            selected === "Services"
              ? `text-[#2D60FF] relative font-semibold ] flex space-x-5`
              : `text-[#B1B1B1] relative flex space-x-5  ` + `  cursor-pointer`
          }
          onClick={() => handleRoute("/services")}
        >
          {selected === "Services" ? (
            <div>
              <div className="w-1 h-10 bg-blue-800 absolute left-0 --top-1rounded-r-lg"></div>
              <Image className="ml-10" src={serviceBlue} alt="dash" />
            </div>
          ) : (
            <Image className="ml-10" src={service} alt="dash" />
          )}
          <h2>Services</h2>
        </li>



        <li
          className={
            selected === "Setting"
              ? `text-[#2D60FF] relative font-semibold ] flex space-x-5`
              : `text-[#B1B1B1] relative flex space-x-5  ` + `  cursor-pointer`
          }
        >
          {selected === "Setting" ? (
            <Image className="ml-10" src={settingBlue} alt="dash" />
          ) : (
            <Image className="ml-10" src={setting} alt="dash" />
          )}
          <h2>Setting</h2>
        </li>



      </ul>
    </div>
  );
};

export default Sidebar;
