import { FC } from 'react';
import Image from 'next/image';
import { FaSearch, FaCog, FaBell } from 'react-icons/fa';

interface NavbarProps {
  pageTitle: string;
  userProfileImage: string;
}

export const Navbar: FC<NavbarProps> = ({ pageTitle, userProfileImage }) => {
  return (
    <nav className="flex flex-col md:flex-row md:items-center p-8 bg-white md:shadow-lg">
      {/* Mobile View: Hamburger Menu */}
      <div className="flex justify-between md:hidden items-center">
        <button aria-label="Toggle Sidebar" className="cursor-pointer focus:outline-none">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="2" stroke="currentColor" className="w-6 h-6">
            <path strokeLinecap="round" strokeLinejoin="round" d="M4 6h16M4 12h16m-7 6h7" />
          </svg>
        </button>
        <div className="text-xl font-bold">{pageTitle}</div>
        <Image src={userProfileImage} alt="User Profile" width={32} height={32} className="rounded-full" />
      </div>

      {/* Mobile: Search Bar */}
      <div className="md:hidden mt-2">
        <div className="relative w-full">
          <FaSearch className="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-400" size={20} />
          <input
            type="text"
            placeholder="Search"
            className="w-full px-10 py-1 text-left bg-gray-100 border-gray-300 rounded-full focus:outline-none focus:ring-1 focus:ring-blue-500"
          />
        </div>
      </div>

      {/* Larger Screens: Full Navbar */}
      <div className="hidden md:flex flex-1 items-center justify-between">
        <div className="text-2xl font-extrabold">{pageTitle}</div>
        <div className="relative w-1/2">
          <FaSearch className="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-400" size={20} />
          <input
            type="text"
            placeholder="Search"
            className="w-full px-10 py-1 text-left bg-gray-100 border-gray-300 rounded-full focus:outline-none focus:ring-1 focus:ring-blue-500"
          />
        </div>
        <div className="flex items-center gap-4">
          <div className="flex items-center rounded-full bg-gray-100 p-2">
            <FaCog size={20} />
          </div>
          <div className="flex items-center rounded-full bg-gray-100 p-2">
            <FaBell size={20} />
          </div>
          <Image src={userProfileImage} alt="User Profile" width={32} height={32} className="rounded-full" />
        </div>
      </div>
    </nav>
  );
};
