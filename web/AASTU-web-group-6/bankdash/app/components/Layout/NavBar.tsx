import React, { useEffect, useState } from "react";
import Image from "next/image";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faBars,
  faSearch,
  faMoon,
  faSun,
} from "@fortawesome/free-solid-svg-icons";
import { useAppDispatch, useAppSelector } from "@/app/Redux/store/store";
import { setDarkMode, toggleDarkMode } from "@/app/Redux/slices/darkModeSlice";

const NavBar = ({ openSidebar }: { openSidebar: () => void }) => {
  const isDarkMode = useAppSelector((state) => state.darkMode.darkMode);
  const dispatch = useAppDispatch();

  const onDarkMode = () => {
    dispatch(toggleDarkMode());
  };

  return (
    <div className="w-full fixed left-0 z-10 flex flex-row justify-center items-center bg-white dark:bg-[#232328] sm:px-[4%] sm:gap-[6%] lg:pl-[240px] pr-[3%] py-4">
      <button
        onClick={openSidebar}
        className="bg-[#F5F7FA] dark:bg-gray-600 rounded-[12px] p-3 py-2 flex items-center hover:bg-[#d0e6f6] dark:hover:bg-gray-600 lg:hidden"
      >
        <FontAwesomeIcon
          icon={faBars}
          className="text-2xl text-gray-700 dark:text-gray-300"
        />
      </button>
      <h1 className="text-3xl font-semibold text-[#343C6A] dark:text-white xs:hidden sm:hidden lg:block">
        Overview
      </h1>
      <div className="flex justify-end gap-5 grow">
        <div className="relative flex gap-2 items-center text-base text-[#8BA3CB] dark:text-gray-400 xs:grow sm:grow lg:grow-0">
          <FontAwesomeIcon
            icon={faSearch}
            className="absolute left-5 text-xl"
          />
          <input
            className="w-full bg-[#F5F7FA] dark:bg-gray-700 py-2 pl-12 pr-6 rounded-3xl border-solid border-2 border-blue-50 dark:border-gray-600 placeholder:text-base placeholder:text-[#8BA3CB] dark:placeholder:text-gray-400 focus:outline-none focus:border-blue-200 dark:focus:border-gray-500"
            type="text"
            name="search"
            id="search"
            placeholder="Search for something"
          />
        </div>

        <button
          onClick={onDarkMode}
          className="flex items-center bg-[#F5F7FA] dark:bg-gray-700 rounded-full p-3 hover:bg-gray-200 dark:hover:bg-gray-600"
        >
          <FontAwesomeIcon
            icon={isDarkMode ? faSun : faMoon}
            className="text-[#718EBF] dark:text-yellow-400 text-xl px-1"
          />
        </button>

        <button className="bg-[#F5F7FA] dark:bg-gray-700 rounded-full p-3">
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
