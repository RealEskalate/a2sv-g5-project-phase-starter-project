import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faBars,
  faSearch,
  faMoon,
  faSun,
} from "@fortawesome/free-solid-svg-icons";
import { useAppDispatch, useAppSelector } from "@/app/Redux/store/store";
import { toggleDarkMode } from "@/app/Redux/slices/darkModeSlice";
import { usePathname } from "next/navigation";
import Image from "next/image";

const NavBar = ({ openSidebar }: { openSidebar: () => void }) => {
  const isDarkMode = useAppSelector((state) => state.darkMode.darkMode);
  const dispatch = useAppDispatch();
  const pathname = usePathname();
  const userData = useAppSelector((state) => state.user);
  console.log(userData, '--')
  const onDarkMode = () => {
    dispatch(toggleDarkMode());
  };

  const pageTitles: { [key: string]: string } = {
    "/": "Overview",
    "/transaction": "Transactions",
    "/account": "Accounts",
    "/investment": "Investments",
    "/credit-cards": "Credit Cards",
    "/loan": "Loans",
    "/service": "Services",
    "/privilages": "My Privileges",
    "/settings/editprofile": "Settings",
    "/login": "LogOut",
  };

  const currentPageTitle = pageTitles[pathname] || "Overview";

  return (
    <div className="w-full font-inter fixed left-0 z-10 flex flex-row justify-center items-center bg-white dark:bg-[#232328] xxs:px-[3%] xss:mb-2 sm:px-[4%] sm:gap-[6%] md:mb-0 lg:pl-[240px] pr-[3%] py-4">
      <button
        onClick={openSidebar}
        className="bg-[#F5F7FA] dark:bg-gray-600 rounded-[12px] p-3 py-2 flex items-center hover:bg-[#d0e6f6] dark:hover:bg-gray-600 lg:hidden"
      >
        <FontAwesomeIcon
          icon={faBars}
          className="text-2xl text-gray-700 dark:text-gray-300"
        />
      </button>
      <h1 className="xxs:text-2xl md:text-3xl font-semibold text-[#343C6A] dark:text-white xxs:flex xxs:grow xxs:justify-center md:justify-start">
        {currentPageTitle}
      </h1>
      <div className="flex justify-end gap-5  ">
        <div className="search-box relative flex gap-2 items-center text-base text-[#8BA3CB] dark:text-gray-400 xxs:bg-white xxs:dark:bg-[#232328] xxs:fixed xxs:top-16 xxs:w-full xxs:left-0 xxs:px-2 xxs:py-2 xxs:pb-4 xxs:shadow-sm xxs:grow sm:px-[4%] md:static md:top-0 md:left-0 md:shadow-none md:bg-transparent md:p-0">
          <FontAwesomeIcon
            icon={faSearch}
            className="relative xxs:left-10 md:left-12 text-xl"
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
          className="flex items-center bg-[#F5F7FA] dark:bg-gray-700 rounded-full p-3 hover:bg-gray-200 dark:hover:bg-gray-600 xxs:hidden md:flex"
        >
          <FontAwesomeIcon
            icon={isDarkMode ? faSun : faMoon}
            className="text-[#718EBF] dark:text-yellow-400 text-xl px-1"
          />
        </button>

        <button className="bg-[#F5F7FA] dark:bg-gray-700 rounded-full p-3 xxs:hidden md:flex">
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
            src={userData.user?.profilePicture || "/assets/profile-1.png"}
            alt="user image"
            width={64}
            height={64}
          />
        </button>
      </div>
    </div>
  );
};

export default NavBar;
