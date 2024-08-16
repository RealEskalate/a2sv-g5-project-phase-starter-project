"use client";
import React from "react";
import { MdOutlineSearch } from "react-icons/md";
import { GiHamburgerMenu } from "react-icons/gi";
import { IoSettingsOutline } from "react-icons/io5";
import { IoMdNotificationsOutline } from "react-icons/io";

interface Props {
  handleClick: () => void;
}
import Image from "next/image";
const Navbar = ({ handleClick }: Props) => {
  return (
    <div className="flex flex-col gap-5 py-5 border-b px-10">
      <div className="flex gap-5 justify-between items-center">
        <div className="text-2xl text-[#343C6A] md:hidden ">
          <button onClick={handleClick}>
            <GiHamburgerMenu />
          </button>
        </div>
        <div className="font-bold text-2xl text-[#343C6A]">Overview</div>

        <div className="flex gap-20">
          <div className="rounded-full hidden md:flex md:gap-2 bg-[#F5F7FA] text-[#8BA3CB] text-sm font-normal py-3 px-8 ml-2 items-center">
            <MdOutlineSearch className="text-xl" />
            <input
              type="text"
              placeholder="Search for Something"
              className="bg-transparent border-none outline-none text-[#8BA3CB] placeholder-[#8BA3CB] text-sm flex-grow"
            />
          </div>

          <div className="hidden md:flex gap-5 text-xl md:items-center">
            <div className="cursor-pointer text-xl bg-[#F5F7FA] rounded-full px-2 py-2">
              <IoSettingsOutline />
            </div>
            <div className="cursor-pointer text-xl bg-[#F5F7FA] rounded-full px-2 py-2">
              <IoMdNotificationsOutline />
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

      <div className="flex md:hidden rounded-full bg-[#F5F7FA] text-[#8BA3CB] text-sm font-normal gap-2 items-center py-3 px-4 ml-2">
        <span className="text-xl">
          <MdOutlineSearch />
        </span>
        Search for Something
      </div>
    </div>
  );
};

export default Navbar;
