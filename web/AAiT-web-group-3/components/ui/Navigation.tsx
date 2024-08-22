"use client";
import React from "react";
import Navbar from "./Navbar";
import Sidebar from "./Sidebar/Sidebar";
import { Props } from "@/types";
import { useAppSelector } from "@/lib/store";

const Navigation: React.FC<Props> = ({ children }) => {
  const toggle = useAppSelector(
    (state) => state.navigationReducer.value.toggle
  );
  const session = true;
  // TODO: Handle fetching session and waiting for provider.

  return (
    <div className="flex w-full">
      {session ? (
        <div className="z-50 bg-white">
          <Sidebar />
        </div>
      ) : (
        <>
          <div className="hidden md:flex md:flex-col md:gap-5 py-7 border-r h-svh sticky top-0 animate-pulse">
            <div className="px-5 py-2">
              <div className="bg-gray-300 w-44 h-9 rounded"></div>
            </div>
            <div className="flex flex-col gap-4 px-5">
              <div className="bg-gray-300 w-40 h-20 rounded"></div>
              <div className="bg-gray-300 w-40 h-20 rounded"></div>
              <div className="bg-gray-300 w-40 h-20 rounded"></div>
              <div className="bg-gray-300 w-40 h-20 rounded"></div>
              <div className="bg-gray-300 w-40 h-20 rounded"></div>
            </div>
          </div>

          {toggle && (
            <div className="md:hidden flex animate-pulse">
              <div
                className={`fixed top-0 left-0 w-80 bg-white shadow-black h-full transform transition-transform ${
                  toggle ? "translate-x-0" : "-translate-x-full"
                } ease-in-out duration-1000 flex flex-col px-5`}
              >
                <div className="flex flex-col justify-between">
                  <button className="cursor-pointer text-primary-color-500 flex justify-end mt-5">
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
        </>
      )}

      <div className="flex flex-col w-full">
        {session ? (
          <div className="w-full bg-white md:sticky md:top-0 md:z-20">
            <Navbar />
          </div>
        ) : (
          <div className="flex flex-col gap-5 py-5 border-b px-10 animate-pulse">
            <div className="flex gap-5 justify-between items-center">
              <div className="text-2xl text-primary-color-800 md:hidden">
                <button>
                  <div className="bg-gray-300 w-8 h-8 rounded-full"></div>
                </button>
              </div>
              <div className="font-bold text-2xl text-primary-color-800 bg-gray-300 rounded w-32 h-8"></div>

              <div className="flex gap-20">
                <div className="rounded-full hidden md:flex md:gap-2 bg-primary-color-50 text-primary-color-200 text-sm font-normal py-3 px-8 ml-2 items-center">
                  <div className="bg-gray-300 w-5 h-5 rounded-full"></div>
                  <div className="bg-gray-300 h-6 w-full rounded"></div>
                </div>

                <div className="hidden md:flex gap-5 text-xl md:items-center">
                  <div className="cursor-pointer text-xl bg-primary-color-50 rounded-full px-2 py-2">
                    <div className="bg-gray-300 w-5 h-5 rounded-full"></div>
                  </div>
                  <div className="cursor-pointer text-xl bg-primary-color-50 rounded-full px-2 py-2">
                    <div className="bg-gray-300 w-5 h-5 rounded-full"></div>
                  </div>
                </div>
                <div className="items-center">
                  <div className="bg-gray-300 w-9 h-9 rounded-full"></div>
                </div>
              </div>
            </div>

            <div className="flex md:hidden rounded-full bg-primary-color-50 text-primary-color-200 text-sm font-normal gap-2 items-center py-3 px-4 ml-2">
              <div className="bg-gray-300 w-5 h-5 rounded-full"></div>
              <div className="bg-gray-300 h-6 w-full rounded"></div>
            </div>
          </div>
        )}
        {children}
      </div>
    </div>
  );
};

export default Navigation;
