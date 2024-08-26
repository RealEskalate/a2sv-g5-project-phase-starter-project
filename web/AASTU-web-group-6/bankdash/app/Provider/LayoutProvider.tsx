"use client";

import React, { ReactNode, useEffect, useState } from "react";
import { usePathname } from "next/navigation";
import Sidebar from "../components/Layout/Sidebar";
import NavBar from "../components/Layout/NavBar";
import { useAppSelector } from "../Redux/store/store";
import { useSession } from "next-auth/react";
import useCardDispatch from "../Redux/Dispacher/useCardDispatch";
import useTranDispatch from "../Redux/Dispacher/useTranDispatch";
// import useUserDataDispatch from "../Redux/Dispacher/useUserDataDispatch";

const LayoutProvider = ({ children }: { children: ReactNode }) => {
  const pathname = usePathname();

  const noLayoutPaths = ["/login", "/signup"];
  const shouldRenderLayout = !noLayoutPaths.includes(pathname);

  const [isOpen, setIsOpen] = useState(false);

  const toggle = () => {
    setIsOpen((prev) => !prev);
  };

  const { data: session } = useSession();
  const accessToken = session?.accessToken as string;

  console.log(session?.accessToken, "token");
  // Update initial card and tran data using the custom hook
  // useUserDataDispatch(accessToken);
  useCardDispatch(accessToken);
  useTranDispatch(accessToken);
  const isDarkMode = useAppSelector((state) => state.darkMode.darkMode);

  // Apply dark mode class directly
  const darkModeClass = isDarkMode ? "dark" : "";

  return (
    <main
      className={`bg-[#f5f7fa] ${darkModeClass} min-h-screen overflow-x-hidden`}
    >
      {shouldRenderLayout ? (
        <section className={`w-full flex relative dark:bg-[#2D2E36] `}>
          {/* Sidebar */}
          <div
            className={`fixed z-20 h-screen xxs:w-[264px] lg:w-[22%] lg:max-w-[220px] bg-white dark:bg-[#242428] top-0 transition-transform duration-300 ease-in-out transform ${
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
          <div className="relative w-full sm:pl-0 lg:pl-[216px] dark:bg-[#2D2E36]">
            <NavBar openSidebar={toggle} />
            <section className="content flex w-full z-0 xxs:pt-28 md:pt-20">
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
