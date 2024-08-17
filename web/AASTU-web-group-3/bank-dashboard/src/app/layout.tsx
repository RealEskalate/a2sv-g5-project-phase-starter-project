"use client";
import Image from "next/image";
import { useEffect, useState } from "react";
import Link from "next/link";
import "./globals.css";
import { GiHamburgerMenu } from "react-icons/gi";
import {
  IoSettingsOutline,
  IoSearchOutline,
  IoNotificationsOutline,
} from "react-icons/io5";
import { menuItems, logo } from "./../../public/Icons";
import ProfilePic from "./../../public/profilepic.png";

export default function ClientSideComponent({
  children,
}: {
  children: React.ReactNode;
}) {
  const [activeItem, setActiveItem] = useState("Dashboard");
  const [ishidden, setIshidden] = useState(false);

  useEffect(() => {
    const currentPath = window.location.pathname;
    const activeMenuItem = menuItems.find((item) =>
      currentPath.startsWith(item.href)
    );
    setActiveItem(activeMenuItem ? activeMenuItem.title : "Dashboard");
  }, []);

  return (
    <>
      <html lang="en">
        <body className="flex">
          <aside
            className={`fixed min-h-full top-0 left-0 z-40 w-64  bg-white shadow-md transition-transform transform ${
              ishidden ? "translate-x-0" : "-translate-x-full"
            } md:translate-x-0 md:relative md:w-[30%] lg:w-[20%] md:flex md:flex-col md:gap-6 md:py-1 md:border md:border-[#E6EFF5] `}
          >
            <div className="ml-5 mt-3 flex items-center gap-8">
              <div className="flex justify-between items-center ">
                <Image src={logo} alt="" width={36} height={36} />
                <div className="text-[#343C6A] pl-2 md:text-xl md:pl-1 lg:pl-2 lg:text-2xl text-base xl:text-4xl md:text-[21px] font-[800] font-mont">
                  BankDash.
                </div>
              </div>
              <div
                onClick={() => setIshidden(false)}
                className="md:hidden rounded-full text-end py-3 px-4 text-blue-600 font-bold bg-[#F5F7FA]"
              >
                X
              </div>
            </div>
            <ul className="flex flex-col ">
              {menuItems.map((item) => (
                <Link href={item.href} key={item.title}>
                  <li
                    onClick={() => setActiveItem(item.title)}
                    className={`flex gap-3 items-center px-8 py-3 text-md md:text-lg ${
                      activeItem === item.title
                        ? "border-l-4 border-l-[#2D60FF] text-[#2D60FF] font-bold"
                        : "text-[#B1B1B1]"
                    }`}
                  >
                    <Image
                      src={item.icon}
                      alt={item.title}
                      width={24}
                      height={24}
                      className={`w-6 h-6 ${
                        activeItem === item.title
                          ? "filter-active"
                          : "filter-inactive"
                      }`}
                    />
                    {/* <img
                      src={item.icon}
                      alt={item.title}
                      className={`w-6 h-6 ${
                        activeItem === item.title
                          ? "filter-active"
                          : "filter-inactive"
                      }`}
                    /> */}
                    <div>{item.title}</div>
                  </li>
                </Link>
              ))}
            </ul>
          </aside>
          <div className="w-full md:w-4/5">
            <nav className="relative flex py-4 px-6 items-center gap-6 w-full  md:h-16">
              {!ishidden && (
                <GiHamburgerMenu
                  className={`md:hidden absolute top-5 left-5 text-3xl`}
                  onClick={() => setIshidden(true)}
                />
              )}
              <div className="w-full flex flex-col md:flex-row gap-4 items-center justify-between md:w-[95%]">
                <div className="ml-[25%] md:ml-6 font-semibold text-[25px] text-[#343C6A]">
                  {activeItem}
                </div>
                <div className="w-full md:w-auto flex items-center justify-between gap-4">
                  <div className="w-full md:w-auto flex gap-2 items-center pl-5 py-3 bg-[#F5F7FA] rounded-full justify-start text-lg overflow-hidden">
                    <IoSearchOutline className="text-[#718EBF] text-xl" />
                    <input
                      type="text"
                      placeholder="search for something"
                      className="outline-none text-md bg-[#F5F7FA]"
                    />
                  </div>
                  <div className="hidden lg:block p-3 rounded-full text-xl text-[#718EBF] bg-[#F5F7FA]">
                    <IoSettingsOutline />
                  </div>
                  <div className="hidden md:block p-3 rounded-full text-xl text-[#FE5C73] bg-[#F5F7FA]">
                    <IoNotificationsOutline />
                  </div>
                </div>
              </div>
              <div className="m-2 absolute top-0 right-0 rounded-full overflow-hidden w-12">
                <Image width={48} height={48} src={ProfilePic.src} alt=""  />
              </div>
            </nav>
            <main className="bg-[#F5F7FA] p-1">{children}</main>
          </div>
        </body>
      </html>
    </>
  );
}
