'use client'

import React, { useState } from "react";
import { FaBars, FaTimes } from "react-icons/fa";
import Image from "next/image";
import { signOut, useSession } from "next-auth/react";
import { useRouter } from "next/navigation";
import seetings from "/public/assets/icons/Group417.png";
import notification from "/public/assets/icons/Group418.png";
import person from "/public/assets/icons/MaskGroup.png";
import magnifying from "/public/assets/icons/magnifying-glass.png";

const NavBar = ({ toggleSidebar, isSidebarVisible }) => {
  const { data: session, status } = useSession();
  const router = useRouter();
  const [searchText, setSearchText] = useState("");

  return (
    <div className="bg-white shadow-md">
      {/* Mobile view */}
      <div className="flex justify-between items-center p-6 sm:hidden">
        <button onClick={toggleSidebar} aria-label="Toggle sidebar">
          {isSidebarVisible ? <FaTimes size={24} /> : <FaBars size={24} />}
        </button>
        <div className="font-semibold text-lg text-primary-2">Overview</div>
        {status === "authenticated" ? (
          <Image src={person} alt="User Icon" className="h-12 w-12" />
        ) : (
          <button
            onClick={() => router.push("/auth/signin")}
            className="bg-slate-500 text-white rounded-lg px-4 py-2"
          >
            Login
          </button>
        )}
      </div>

      {/* Desktop view */}
      <div className="hidden sm:flex justify-between items-center h-24 px-12">
        <div className="font-semibold text-xl text-primary-2">Overview</div>
        <div className="flex gap-5 items-center">
          <div className="relative w-full max-w-xs">
            <input
              type="text"
              placeholder="Search..."
              value={searchText}
              onChange={(e) => setSearchText(e.target.value)}
              className="pl-12 pr-12 py-2 w-full bg-gray-100 rounded-full border border-gray-300"
            />
            <div className="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-500">
              <Image src={magnifying} alt="Search Icon" width={16} height={16} />
            </div>
            {searchText && (
              <button
                onClick={() => setSearchText("")}
                className="absolute right-4 top-1/2 transform -translate-y-1/2 text-gray-500"
              >
                <FaTimes />
              </button>
            )}
          </div>
          <Image src={seetings} alt="Settings Icon" className="h-10 w-10" />
          <Image src={notification} alt="Notification Icon" className="h-10 w-10" />
          {status === "authenticated" ? (
            <>
              <Image src={person} alt="User Icon" className="h-12 w-12" />
              <button
                onClick={() => signOut()}
                className="bg-slate-500 text-white rounded-full px-4 py-2 ml-4"
              >
                Logout
              </button>
            </>
          ) : (
            <button
              onClick={() => router.push("/auth/signin")}
              className="bg-slate-500 text-white rounded-full px-4 py-2"
            >
              Login
            </button>
          )}
        </div>
      </div>
    </div>
  );
};

export default NavBar;
