"use client";
import React from "react";
import { MdOutlineSearch } from "react-icons/md";
import { GiHamburgerMenu } from "react-icons/gi";
import { IoMoonOutline, IoSettingsOutline } from "react-icons/io5";
import { IoMdNotificationsOutline } from "react-icons/io";
import { useRouter } from "next/navigation";
interface Props {
  handleClick: () => void;
  toggleDarkMode: () => void;
}
import Image from "next/image";
const Navbar = ({ handleClick, toggleDarkMode }: Props) => {
  const route = useRouter();
  return (
    <div className="flex flex-col gap-5 py-5 border-b px-10">
      <div className="flex gap-5 justify-between items-center">
        <div className="text-2xl text-[#343C6A] md:hidden dark:text-white">
          <button onClick={handleClick}>
            <GiHamburgerMenu />
          </button>
        </div>
        <div className="font-bold text-2xl text-[#343C6A] dark:text-[#9faaeb]">
          Overview
        </div>

        <div className="flex gap-20">
          <div className="rounded-full hidden md:flex md:gap-2 bg-[#F5F7FA] text-[#8BA3CB] text-sm font-normal py-3 px-8 ml-2 items-center dark:bg-[#050914] dark:border dark:border-[#333B69]">
            <MdOutlineSearch className="text-xl" />
            <input
              type="text"
              placeholder="Search for Something"
              className="bg-transparent border-none outline-none text-[#8BA3CB] placeholder-[#8BA3CB] text-sm flex-grow"
            />
          </div>

          <div className="hidden md:flex gap-5 text-xl md:items-center">
            <div
              className="cursor-pointer text-xl bg-[#F5F7FA] rounded-full px-2 py-2 dark:bg-[#050914] dark:border dark:border-[#333B69]"
              onClick={() => route.push("./bankingSettings")}
            >
              <IoSettingsOutline />
            </div>
            <div className="cursor-pointer text-xl bg-[#F5F7FA] rounded-full px-2 py-2 dark:bg-[#050914] dark:border dark:border-[#333B69]">
              <IoMdNotificationsOutline />
            </div>
            <div
              className="cursor-pointer text-xl bg-[#F5F7FA] rounded-full px-2 py-2 dark:bg-[#050914] dark:border dark:border-[#333B69]"
              onClick={toggleDarkMode}
            >
              <IoMoonOutline />
            </div>
          </div>
          <div className="items-center">
            <Image
              src="/profile.png"
              alt="Profile"
              width={35}
              height={35}
            ></Image>
          </div>
        </div>
      </div>

      <div className="md:hidden rounded-full flex md:gap-2 bg-[#F5F7FA] text-[#8BA3CB] text-sm font-normal py-3 px-8 ml-2 items-center dark:bg-[#050914] dark:border dark:border-[#333B69]">
        <MdOutlineSearch className="text-xl" />
        <input
          type="text"
          placeholder="Search for Something"
          className="bg-transparent border-none outline-none text-[#8BA3CB] placeholder-[#8BA3CB] text-sm flex-grow"
        />
      </div>
    </div>
  );
};

export const NavBarLoading = () => {
  return (
    <div className="flex flex-col gap-5 py-5 border-b px-10 animate-pulse justify-between w-full dark:bg-[#050914] dark:border-[#333B69]">
      <div className="flex gap-5 justify-between items-center">
        <div className="text-2xl md:hidden">
          <button>
            <div className="bg-gray-300 dark:bg-[#333B69] w-8 h-8 rounded-full"></div>
          </button>
        </div>
        <div className="font-bold text-2xl bg-gray-300 dark:bg-[#333B69] rounded w-32 h-8"></div>

        <div className="flex gap-20">
          <div className="w-72 rounded-full hidden md:flex md:gap-2 bg-[#F5F7FA] dark:bg-[#050914] text-sm font-normal py-3 px-8 ml-2 items-center dark:border dark:border-[#333B69]">
            <div className="bg-gray-300 dark:bg-[#333B69] w-5 h-5 rounded-full"></div>
            <div className="bg-gray-300 dark:bg-[#333B69] h-6 w-full rounded-lg"></div>
          </div>

          <div className="hidden md:flex gap-5 text-xl md:items-center">
            <div className="cursor-pointer text-xl rounded-full px-2 py-2">
              <div className="bg-gray-200 dark:bg-[#333B69] w-5 h-5 rounded-full"></div>
            </div>
            <div className="cursor-pointer text-xl rounded-full px-2 py-2">
              <div className="w-5 h-5 rounded-full bg-gray-200 dark:bg-[#333B69]"></div>
            </div>
          </div>
          <div className="items-center">
            <div className="bg-gray-300 dark:bg-[#333B69] w-9 h-9 rounded-full"></div>
          </div>
        </div>
      </div>

      <div className="flex md:hidden rounded-full bg-[#F5F7FA] dark:bg-[#050914] text-[#8BA3CB] text-sm font-normal gap-2 items-center py-3 px-4 ml-2 dark:border dark:border-[#333B69]">
        <div className="bg-gray-300 dark:bg-[#333B69] w-5 h-5 rounded-lg"></div>
        <div className="bg-gray-300 dark:bg-[#333B69] h-6 w-full rounded"></div>
      </div>
    </div>
  );
};

export default Navbar;
