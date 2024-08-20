import { FC } from 'react';
import Image from 'next/image';
import { FaSearch, FaCog, FaBell } from 'react-icons/fa';

type NavbarProps = {
  pageTitle: string;
  toggleSidebar: () => void;
};

export const Navbar: FC<NavbarProps> = ({ pageTitle, toggleSidebar }) => {
  return (
    <nav className="flex flex-col p-4 bg-white shadow-md lg:pl-64">
      {/* Mobile View: Hamburger Menu */}
      <div className="lg:hidden flex justify-between items-center mb-2">
        <button onClick={toggleSidebar} aria-label="Toggle Sidebar" className="cursor-pointer focus:outline-none">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="2" stroke="currentColor" className="w-6 h-6">
            <path strokeLinecap="round" strokeLinejoin="round" d="M4 6h16M4 12h16m-7 6h7" />
          </svg>
        </button>
        <div className="text-xl font-bold flex-1 text-center">{pageTitle}</div>
        <Image src="/Images/profilepic.jpeg" alt="User Profile" width={32} height={32} className="rounded-full aspect-square object-cover" />
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
      <div className="hidden lg:flex items-center justify-around w-full">
        <div className="text-2xl font-extrabold">{pageTitle}</div>
        <div className="relative w-64"> {/* Fixed width for search bar */}
          <FaSearch className="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-400" size={20} />
          <input
            type="text"
            placeholder="Search for something"
            className="w-full px-10 py-2 text-left bg-gray-100 border-gray-300 rounded-full focus:outline-none focus:ring-1 focus:ring-blue-500"
          />
        </div>
        <div className="flex items-center gap-4">
          <div className="flex items-center rounded-full bg-gray-100 p-2">
            <FaCog className="text-blue-800" size={20} />
          </div>
          <div className="flex items-center rounded-full bg-gray-100 p-2">
            <FaBell className="text-red-600" size={20} />
          </div>
          <Image src="/Images/profilepic.jpeg" alt="User Profile" width={30} height={30} className="rounded-full aspect-square object-cover" />
        </div>
      </div>
    </nav>
  );
};
