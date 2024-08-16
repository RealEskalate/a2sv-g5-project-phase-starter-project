"use client";
import Image from "next/image";
import React, { useState } from "react";
import Sidebar from "../sidebar/Sidebar"; // Import Sidebar component

interface TitleProp {
  title: string;
}

const Navbar = ({ title }: TitleProp) => {
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);

  const toggleSidebar = () => {
    setIsSidebarOpen(!isSidebarOpen);
  };

  return (
    <>
      <header className="fixed top-0 left-0 right-0 bg-white shadow-md z-40">
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
            <Image
              src={"/assets/navbar/credit-card.svg"}
              width={36}
              height={36}
              alt="bankDash logo"
              className="mr-3 sm:block hidden"
            />
            <p className="font-black text-[25px] text-[#343C6A] sm:flex hidden">
              BankDash.
            </p>
          </div>

          <div className="flex-grow flex items-center justify-center">
            <p className="font-semibold text-[20px] sm:text-[25px] text-[#343C6A]">
              {title}
            </p>
          </div>

          <div className="hidden sm:flex gap-5 items-center">
            <div className="search-div bg-[#F5F7FA] items-center rounded-full h-[50px] pl-5 pr-5 pt-3 pb-3">
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
                className="text-[15px] bg-[#F5F7FA]"
              />
            </div>

            <div className="bg-[#F5F7FA] rounded-full flex justify-center items-center">
              <Image
                src={"/assets/navbar/settings.svg"}
                width={50}
                height={50}
                alt="settings"
                className="flex-shrink-0"
              />
            </div>

            <div className="bg-[#F5F7FA] rounded-full flex justify-center items-center">
              <Image
                src={"/assets/navbar/notification.svg"}
                width={50}
                height={50}
                alt="notification"
                className="flex-shrink-0"
              />
            </div>

            <Image
              src={"/assets/navbar/default-image.svg"}
              width={50}
              height={50}
              alt="profile-picture"
              className="object-fill rounded-full"
            />
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
