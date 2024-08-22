"use client";
import React from "react";
import { useDispatch } from "react-redux";
import { useRouter } from "next/navigation";

import { pages } from "./SidebarData";

import { AppDispatch, useAppSelector } from "@/lib/store";
import {
  updateActivePage,
  updateToggle,
} from "@/lib/features/navigation/navigationSlice";

const SidebarElements = () => {
  const dispatch = useDispatch<AppDispatch>();
  const activePage = useAppSelector(
    (state) => state.navigationReducer.value.activePage
  );
  const toggle = useAppSelector(
    (state) => state.navigationReducer.value.toggle
  );
  const router = useRouter();

  const handleNav = async (destination: string) => {
    router.push(destination);
  };
  const handleActive = (element: string) => {
    dispatch(updateActivePage(element));
    dispatch(updateToggle(!toggle));
  };

  return (
    <div className="flex flex-col gap-5 mb-5">
      {pages.map((each_page, index) => (
        <div
          key={index}
          className={`${
            activePage === each_page.text
              ? "text-primary-color-500 border-l-2"
              : "text-primary-color-100"
          } flex gap-3 items-center font-semibold text-l`}
        >
          <button
            onClick={() => {
              handleActive(each_page.text);
              handleNav(each_page.destination);
            }}
            className={`flex items-center w-full`}
          >
            <span
              className={`${
                activePage === each_page.text
                  ? "bg-primary-color-500"
                  : "hidden"
              } rounded-r-lg w-1 h-10`}
            ></span>
            <div className="px-5 flex items-center gap-6">
              <span className="text-2xl">
                <each_page.icon />
              </span>
              {each_page.text}
            </div>
          </button>
        </div>
      ))}
    </div>
  );
};

export default SidebarElements;
