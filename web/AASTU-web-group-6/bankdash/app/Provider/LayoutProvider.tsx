"use client";

import React, { ReactNode } from "react";
import { usePathname } from "next/navigation";
import Sidebar from "../components/Layout/Sidebar";
import NavBar from "../components/Layout/NavBar";

const LayoutProvider = ({ children }: { children: ReactNode }) => {
  const pathname = usePathname();

  const noLayoutPaths = ["/login", "/signup"];

  const shouldRenderLayout = !noLayoutPaths.includes(pathname);

  return (
    <>
      {shouldRenderLayout ? (
        <section className="w-full flex relative">
          <div className="sidebar-container fixed z-20 h-screen w-[20%] bg-white pr-1">
            <Sidebar />
          </div>
          <div className="relative flex flex-col w-full pl-[20%]">
            <NavBar />
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
