"use client";

import React, { ReactNode, useEffect, useState } from "react";
import { usePathname } from "next/navigation";
import Sidebar from "../components/Layout/Sidebar";
import NavBar from "../components/Layout/NavBar";
import { useAppSelector } from "../Redux/store/store";

const LayoutProvider = ({ children }: { children: ReactNode }) => {
  const pathname = usePathname();

  const noLayoutPaths = ["/login", "/signup"];
  const shouldRenderLayout = !noLayoutPaths.includes(pathname);

  const [isOpen, setIsOpen] = useState(false);

  const toggle = () => {
    setIsOpen((prev) => !prev);
  };
  const isDarkMode = useAppSelector((state) => state.darkMode.darkMode);

  // Apply dark mode class directly
  const darkModeClass = isDarkMode ? "dark" : "";

  return (
    <main className={`bg-[#f5f7fa] ${darkModeClass} min-h-screen`}>
      {shouldRenderLayout ? (
        <section className={`w-full flex relative dark:bg-[#2D2E36] `}>
          {/* Sidebar */}
          <div
            className={`fixed z-20 h-screen sm:w-[50%] lg:w-[22%] lg:max-w-[220px] bg-white dark:bg-[#242428] pr-1 top-0 transition-transform duration-300 ease-in-out transform ${
              isOpen
                ? "translate-x-0 opacity-100"
                : "-translate-x-full opacity-0 sm:opacity-100"
            } lg:translate-x-0`}
          >
            <Sidebar isOpen={isOpen} closeSidebar={toggle} />
          </div>

          {/* Overlay for small screens */}
          {isOpen && (
            <div
              className="fixed inset-0 z-10 bg-black opacity-50 sm:block lg:hidden"
              onClick={toggle}
            ></div>
          )}

          {/* Main Content */}
          <div className="relative flex flex-col w-full sm:pl-0 lg:pl-[216px] dark:bg-[#2D2E36]">
            <NavBar openSidebar={toggle} />
            <section className="content flex w-full z-0 pt-20">
              {children}
            </section>
          </div>
        </section>
      ) : (
        <div>{children}</div>
      )}
    </main>
  );
};

export default LayoutProvider;
