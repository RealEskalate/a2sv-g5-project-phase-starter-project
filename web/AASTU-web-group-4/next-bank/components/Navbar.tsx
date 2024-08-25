"use client";
import { FC, useState } from 'react';
import Image from 'next/image';
import { FaSearch, FaCog, FaBell } from 'react-icons/fa';
import Link from 'next/link';
import { signOut } from 'next-auth/react';
import Cookies from 'js-cookie';
import { useRouter } from 'next/navigation';
import { DropdownMenuDemo } from './dropdown';
import ThemeSwitch from './ThemeSwitch';
import { IoMdNotificationsOutline } from "react-icons/io";
import NotificationBell from './NotificationBell';


type NavbarProps = {
  pageTitle: string;
  toggleSidebar: () => void;
};

export const Navbar: FC<NavbarProps> = ({ pageTitle, toggleSidebar }) => {
  const [isOpen, setIsOpen] = useState(false);
  const router = useRouter(); // Initialize router here

  const toggleDropdown = () => {
    setIsOpen(prev => !prev);
  };

  

  const handleSettings = () => {
    // Navigate to settings page or open settings modal
  };

  return (
    <nav className="flex flex-col p-4 bg-white shadow-md lg:pl-64 dark:bg-dark text-gray-900 dark:text-white fixed w-full z-10">
      {/* Mobile View: Hamburger Menu */}
      <div className="lg:hidden flex justify-between items-center">
        <button onClick={toggleSidebar} aria-label="Toggle Sidebar" className="cursor-pointer focus:outline-none dark:text-blue-500">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="2" stroke="currentColor" className="w-6 h-6">
            <path strokeLinecap="round" strokeLinejoin="round" d="M4 6h16M4 12h16m-7 6h7" />
          </svg>
        </button>
        <div className="text-xl font-bold flex-1 text-center dark:text-blue-500">{pageTitle}</div>
        <DropdownMenuDemo />
      </div>

      {/* Mobile Search Bar */}
      <div className="lg:hidden relative w-full mt-2">
        <FaSearch className="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-400" size={20} />
        <input
          type="text"
          placeholder="Search for something"
          className="w-full px-10 py-2 text-left bg-gray-100 border-gray-300 rounded-full focus:outline-none focus:ring-1 focus:ring-blue-500"
        />
      </div>

      {/* Larger Screens: Full Navbar */}
      <div className="hidden lg:flex items-center justify-between w-full">
        <div className="text-2xl ml-8 mr-[100px] font-extrabold dark:text-blue-500">{pageTitle}</div>
        <div className="flex items-center gap-5 mr-8">
        <div className="relative w-72">
          <FaSearch className="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-400" size={20} />
          <input
            type="text"
            placeholder="Search for something"
            className="w-full px-10 py-2 text-left bg-gray-100 border-gray-300 rounded-full focus:outline-none focus:ring-1 focus:ring-blue-500"
          />
        </div>
          <div>
            <ThemeSwitch/>
          </div>
          <div>
            <NotificationBell />
          </div>
          <div className="relative inline-block">
            <div >
                <DropdownMenuDemo />
            </div>
          </div>
        </div>
      </div>
    </nav>
  );
};