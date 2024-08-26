"use client";
import React, { useState, useEffect } from "react";
import Navbar, { NavBarLoading } from "./Navbar";
import Sidebar, { SidebarLoading } from "./Sidebar";
import { getSession } from "next-auth/react";
import Loading from "../accounts/components/Loading";
import { useDarkMode } from "./Context/DarkModeContext";

interface Props {
  children: React.ReactNode;
}

const Navigation: React.FC<Props> = ({ children }) => {
  const [toggle, setToggle] = useState(false);
  const [session, setSession] = useState(false);
  const [loading, setLoading] = useState(true);
  const { darkMode, toggleDarkMode } = useDarkMode(); // Use dark mode context

  useEffect(() => {
    const fetchSession = async () => {
      const sessionData = await getSession();
      if (sessionData?.user) {
        setSession(true);
      }
      setLoading(false);
    };

    fetchSession();
  }, []);

  return (
    <>
      {loading ? (
        <div className="flex w-full">
          <SidebarLoading />
          <div className="flex flex-col w-full">
            <div className="flex w-full">
              <NavBarLoading />
            </div>
            <Loading />
          </div>
        </div>
      ) : (
        <div className="flex w-full">
          {session && (
            <div className="z-50 bg-white">
              <Sidebar
                toggle={toggle}
                handleClose={() => {
                  setToggle(!toggle);
                }}
              />
            </div>
          )}
          {toggle && (
            <div className="md:hidden flex animate-pulse">
              <div
                className={`fixed top-0 left-0 w-80 bg-white shadow-black h-full transform transition-transform ${
                  toggle ? "translate-x-0" : "-translate-x-full"
                } ease-in-out duration-1000 flex flex-col px-5`}
              >
                <div className="flex flex-col justify-between">
                  <button className="cursor-pointer text-[#2D60FF] flex justify-end mt-5">
                    <div className="bg-gray-300 w-8 h-8 rounded-full"></div>
                  </button>
                  <div className="px-3 mt-3 mb-4">
                    <div className="bg-gray-300 w-44 h-9 rounded"></div>
                  </div>
                </div>
                <div className="flex flex-col gap-4 px-5">
                  <div className="bg-gray-300 w-32 h-6 rounded"></div>
                  <div className="bg-gray-300 w-32 h-6 rounded"></div>
                  <div className="bg-gray-300 w-32 h-6 rounded"></div>
                  <div className="bg-gray-300 w-32 h-6 rounded"></div>
                </div>
              </div>
            </div>
          )}

          <div className="flex flex-col w-full">
            {session && (
              <div className="w-full bg-white md:sticky md:top-0 md:z-20 dark:bg-[#020817]">
                <Navbar
                  handleClick={() => {
                    setToggle(!toggle);
                  }}
                  toggleDarkMode={toggleDarkMode}
                />
              </div>
            )}
            {children}
          </div>
        </div>
      )}
    </>
  );
};

export default Navigation;
