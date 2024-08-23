"use client";
import Image from "next/image";
import React, { useEffect, useState, useRef } from "react";
import Sidebar from "../sidebar/Sidebar"; // Import Sidebar component
import Link from "next/link";
import { usePathname } from "next/navigation";
import { useSelector } from "react-redux";
import { RootState } from "@/lib/store";
import { signOut } from "next-auth/react";

const Navbar = () => {
  // Access the user data from Redux store
  const user = useSelector((state: RootState) => state.user.user);

  // State for sidebar visibility
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);

  // State for dropdown visibility
  const [isDropdownOpen, setIsDropdownOpen] = useState(false);

  // Reference for the dropdown
  const dropdownRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        dropdownRef.current &&
        !dropdownRef.current.contains(event.target as Node)
      ) {
        setIsDropdownOpen(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);

    // Cleanup the event listener
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [dropdownRef]);

  // Toggle sidebar visibility
  const toggleSidebar = () => {
    setIsSidebarOpen((prev) => !prev);
  };

  // Toggle dropdown visibility
  const toggleDropdown = () => {
    setIsDropdownOpen((prev) => !prev);
  };

  // Get the current pathname from next/navigation
  const pathname = usePathname();

  // Function to capitalize the first letter of the page title
  const capitalizeFirstLetter = (text: string) => {
    if (!text || text === "/") return "Dashboard";
    text = text.replace("-", " ");
    return text.charAt(1).toUpperCase() + text.slice(2).toLowerCase();
  };

  const title = capitalizeFirstLetter(pathname);

  // Render loading state or user information

  if (!user) {
    return <p>Loading user data...</p>;
  }

  return (
    <>
      <header className="fixed top-0 md:left-[240px] max-md:left-0 right-0 bg-white shadow-md z-40 max-md:h-[140px]">
        <div className="flex w-full items-center justify-between px-4 py-2">
          <div className="flex items-center">
            <Image
              src="/assets/navbar/hamburger.svg"
              width={25}
              height={25}
              alt="hamburger"
              className="sm:hidden block cursor-pointer"
              onClick={toggleSidebar}
            />
          </div>

          <div className="flex-grow flex justify-start max-md:justify-center">
            <p className="font-semibold text-[20px] sm:text-[25px] text-[#343C6A]">
              {title.slice(0)}
            </p>
          </div>

          <div className=" sm:flex gap-5 items-center relative">
            <div className="hidden sm:flex search-div bg-[#F5F7FA]  items-center rounded-full h-[50px] px-4 py-2 w-full max-w-[400px] mx-auto">
              <Image
                src="/assets/navbar/magnifying-glass.svg"
                width={20}
                height={20}
                alt="magnifying-glass"
                className="mr-3 "
              />
              <input
                type="text"
                placeholder="Search for something"
                className="text-[15px] bg-[#F5F7FA] border border-transparent rounded-full p-2 w-full placeholder:text-[#B1B1B1] focus:outline-none focus:ring-0 focus:border-transparent"
              />
            </div>

            <Link
              href="/settings"
              className="sm:flex hidden bg-[#F5F7FA] rounded-full justify-center items-center"
            >
              <Image
                src="/assets/navbar/settings.svg"
                width={50}
                height={50}
                alt="settings"
                className="flex-shrink-0 min-w-fit"
              />
            </Link>

            <Link
              href="/"
              className="sm:flex hidden bg-[#F5F7FA] rounded-full  justify-center items-center"
            >
              <Image
                src="/assets/navbar/notification.svg"
                width={50}
                height={50}
                alt="notification"
                className="flex-shrink-0 min-w-fit"
              />
            </Link>

            <div className="relative" ref={dropdownRef}>
              <div className="flex justify-center items-center w-12 h-12 rounded-full object-scale-down">
                <Image
                  src={
                    user.profilePicture === "assets/navbar/default-image.svg"
                      ? ""
                      : user.profilePicture
                  }
                  width={50}
                  height={50}
                  alt="profile-picture"
                  className="block mx-auto h-12 rounded-full cursor-pointer"
                  onClick={toggleDropdown}
                />
              </div>

              {isDropdownOpen && (
                <div className="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg py-2 z-50 border-[1px] border-[#afafaf]">
                  <div
                    onClick={() => signOut({ callbackUrl: "/" })}
                    className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                  >
                    Logout
                  </div>
                </div>
              )}
            </div>
          </div>
        </div>

        <div className="search-div flex w-4/5 sm:hidden bg-[#F5F7FA] items-center rounded-full pl-5 pr-5 pt-3 pb-3 mt-2 mx-auto">
          <Image
            src="/assets/navbar/magnifying-glass.svg"
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
