import React from "react";
import Image from "next/image";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faBars, faSearch } from "@fortawesome/free-solid-svg-icons";
import { faBell } from "@fortawesome/free-regular-svg-icons";

const NavBar = ({ openSidebar }: { openSidebar: () => void }) => {
  return (
    <div className="w-full fixed left-0 z-10 flex flex-row justify-center items-center bg-white sm:px-[4%] sm:gap-[6%] lg:pl-[240px] pr-[3%] py-4">
      <button
        onClick={openSidebar}
        className="bg-[#F5F7FA] rounded-[12px] p-3 py-2 flex items-center hover:bg-[#d0e6f6] lg:hidden"
      >
        <FontAwesomeIcon icon={faBars} className="text-2xl text-gray-700" />
      </button>
      <h1 className="text-3xl font-semibold text-[#343C6A] sm:hidden lg:block">
        Overview
      </h1>
      <div className="flex justify-end gap-5 grow">
        {/* Search */}
        <div className="relative  flex gap-2 items-center text-base text-[#8BA3CB] sm:grow  lg:grow-0">
          <FontAwesomeIcon
            icon={faSearch}
            className="absolute left-5 text-xl"
          />
          <input
            className="w-full  bg-[#F5F7FA] py-2 pl-12 pr-6 rounded-3xl border-solid border-2 border-blue-50 placeholder:text-base placeholder:text-[#8BA3CB] focus:outline-none focus:border-blue-200"
            type="text"
            name="search"
            id="search"
            placeholder="Search for something"
          />
        </div>

        {/* User tool */}
        <button className="bg-[#F5F7FA] rounded-full p-3">
          <Image
            src="/assets/settings 1.svg"
            alt="setting"
            width={24}
            height={24}
          />
        </button>
        <button className="bg-[#F5F7FA] rounded-full p-3">
          <Image
            src="/assets/bell-icon.svg"
            alt="notification"
            width={24}
            height={24}
          />
        </button>

        <button>
          <Image
            className="rounded-full"
            src="/assets/profile-1.png"
            alt="user image"
            width={32}
            height={32}
          />
        </button>
      </div>
    </div>
  );
};

export default NavBar;
