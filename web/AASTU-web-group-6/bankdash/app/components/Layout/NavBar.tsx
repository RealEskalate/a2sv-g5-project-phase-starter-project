import React from "react";
import Image from "next/image";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSearch } from "@fortawesome/free-solid-svg-icons";
import { faBell } from "@fortawesome/free-regular-svg-icons";

const NavBar = () => {
  return (
    <div className="w-[80%] fixed z-10 flex flex-row justify-center bg-white pl-[2%] pr-[5%] py-4">
      <h1 className="text-3xl font-semibold text-[#343C6A]">Overview</h1>
      <div className="flex justify-end gap-5 grow">
        {/* Search */}
        <div className="relative  flex gap-2 items-center text-base text-[#8BA3CB]">
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
