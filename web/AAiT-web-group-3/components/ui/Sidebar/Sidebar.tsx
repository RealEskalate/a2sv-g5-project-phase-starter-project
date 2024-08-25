import React from "react";
import { useDispatch } from "react-redux";
import Image from "next/image";

import { FaTimes } from "react-icons/fa";

import { useAppSelector, AppDispatch } from "@/lib/store";
import { updateToggle } from "@/lib/features/navigation/navigationSlice";

import SidebarElements from "./SidebarElements";

const Sidebar = () => {
  const toggle = useAppSelector(
    (state) => state.navigationReducer.value.toggle
  );
  const dispatch = useDispatch<AppDispatch>();
  return (
    <>
      <div className="hidden md:flex md:flex-col md:gap-5 py-7 border-r h-svh sticky top-0">
        <div className="px-5 py-2">
          <Image
            // TODO: Change Image
            src="https://res.cloudinary.com/dtt1wnvfb/image/upload/v1701954159/photo_2023-12-07%2016.02.23.jpeg.jpg"
            width={223}
            height={36}
            alt="Logo"
          />
        </div>
        <SidebarElements />
      </div>

      {toggle && (
        <div className="md:hidden flex opacity-90">
          <div
            className={`fixed top-0 left-0 w-80 bg-white shadow-black h-full transform transition-transform${
              toggle ? "translate-x-0" : "-translate-x-full"
            }  ease-in-out duration-1000 flex flex-col px-5`}
          >
            <div className="flex flex-col justify-between">
              <button
                onClick={() => {
                  dispatch(updateToggle(!toggle));
                }}
                className="cursor-pointer text-primary-color-500 flex justify-end mt-5"
              >
                <span className="text-3xl">
                  <FaTimes />
                </span>
                ;
              </button>
              <div className="px-3 mt-3 mb-4">
                <Image
                  src="https://res.cloudinary.com/dtt1wnvfb/image/upload/v1701954159/photo_2023-12-07%2016.02.23.jpeg.jpg"
                  width={183}
                  height={36}
                  alt="Logo"
                />
              </div>
            </div>
            <SidebarElements />
          </div>
        </div>
      )}
    </>
  );
};

export default Sidebar;
