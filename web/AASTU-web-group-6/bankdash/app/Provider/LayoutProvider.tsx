"use client";

import React, { ReactNode, useState } from "react";
import { usePathname } from "next/navigation";
import Sidebar from "../components/Layout/Sidebar";
import NavBar from "../components/Layout/NavBar";

const LayoutProvider = ({ children }: { children: ReactNode }) => {
  const pathname = usePathname();

  const noLayoutPaths = ["/login", "/signup"];
  const shouldRenderLayout = !noLayoutPaths.includes(pathname);

  const [isOpen, setIsOpen] = useState(false);

  const toggle = () => {
    setIsOpen((prev) => !prev);
  };

  return (
    <>
      {shouldRenderLayout ? (
        <section className="w-full flex relative">
          {/* Sidebar */}
          <div
            className={`fixed z-20 h-screen sm:w-[50%] lg:w-[22%] lg:max-w-[220px] bg-white pr-1 top-0 transition-transform duration-300 ease-in-out transform ${
              isOpen ? "translate-x-0" : "-translate-x-full"
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
          <div className="relative flex flex-col w-full sm:pl-0 lg:pl-[216px]">
            <NavBar openSidebar={toggle} />
            <section className="content flex w-full z-0 pt-20">
              {children}
            </section>
          </div>
        </section>
      ) : (
        <div>{children}</div>
      )}
    </>
  );
};

export default LayoutProvider;
