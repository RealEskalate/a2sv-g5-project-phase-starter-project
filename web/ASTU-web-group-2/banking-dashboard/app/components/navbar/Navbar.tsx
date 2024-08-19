"use client";
import Image from "next/image";
import React, { useState } from "react";
import Sidebar from "../sidebar/Sidebar"; // Import Sidebar component
import Link from "next/link";
import { usePathname } from "next/navigation";
import { title } from "process";


const Navbar = () => {
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);

  const toggleSidebar = () => {
    setIsSidebarOpen(!isSidebarOpen);
  };
  const pathname = usePathname()
  const capitalizeFirstLetter = (text:string) => {
    if (!text || text == "/") return 'Dashboard';
    text = text.replace('-', ' ')
    return text.charAt(1).toUpperCase() + text.slice(2).toLowerCase();
  };
  const title = capitalizeFirstLetter(pathname);

  return (
    <>
      <header className="fixed top-0 md:left-[240px] max-md:left-0 right-0 bg-white shadow-md z-40 max-md:h-[140px]">
        <div className="flex w-full items-center justify-between px-4 py-2">
          <div className="flex items-center">
            <Image
              src={"/assets/navbar/hamburger.svg"}
              width={25}
              height={25}
              alt="hamburger"
              className="sm:hidden block cursor-pointer"
              onClick={toggleSidebar}
            />
          </div>

          <div className="flex-grow flex justify-start max-md:justify-center">
            <p className="font-semibold text-[20px] sm:text-[25px] text-[#343C6A]">
              {title}
            </p>
          </div>

          <div className="hidden sm:flex gap-5 items-center">
            <div className="search-div bg-[#F5F7FA] flex items-center rounded-full h-[50px] px-4 py-2 w-full max-w-[400px] mx-auto">
              <Image
                src={"/assets/navbar/magnifying-glass.svg"}
                width={20}
                height={20}
                alt="magnifying-glass"
                className="mr-3"
              />
              <input
                type="text"
                placeholder="Search for something"
                className="text-[15px] bg-[#F5F7FA] border border-transparent rounded-full px-4 py-2 w-full placeholder:text-[#B1B1B1] focus:outline-none focus:ring-0 focus:border-transparent"
              />
            </div>

            <Link
              href="/settings"
              className="bg-[#F5F7FA] rounded-full flex justify-center items-center"
            >
              <Image
                src={"/assets/navbar/settings.svg"}
                width={50}
                height={50}
                alt="settings"
                className="flex-shrink-0 min-w-fit"
              />
            </Link>

            <Link
              href="/"
              className="bg-[#F5F7FA] rounded-full flex justify-center items-center"
            >
              <Image
                src={"/assets/navbar/notification.svg"}
                width={50}
                height={50}
                alt="notification"
                className="flex-shrink-0 min-w-fit"
              />
            </Link>

            <Link
              href="/settings"
              className="flex items-center gap-3 bg-[#F5F7FA] rounded-full"
            >
              <Image
                src={"/assets/navbar/default-image.svg"}
                width={50}
                height={50}
                alt="profile-picture"
                className="object-fill rounded-full"
              />
            </Link>
          </div>
        </div>

        <div className="search-div flex w-4/5 sm:hidden bg-[#F5F7FA] items-center rounded-full pl-5 pr-5 pt-3 pb-3 mt-2 mx-auto">
          <Image
            src={"/assets/navbar/magnifying-glass.svg"}
            width={20}
            height={20}
            alt="magnifying-glass"
            className="mr-5"
          />
          <input
            type="text"
            placeholder="Search for something"
            className="text-md bg-[#F5F7FA] w-full"
          />
        </div>
      </header>

      <Sidebar isOpen={isSidebarOpen} toggleSidebar={toggleSidebar} />
    </>
  );
};

export default Navbar;
