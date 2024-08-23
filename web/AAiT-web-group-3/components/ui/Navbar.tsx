"use client";
import React from "react";
import { useDispatch } from "react-redux";
import Image from "next/image";

//  Import Icons
import { MdOutlineSearch } from "react-icons/md";
import { GiHamburgerMenu } from "react-icons/gi";
import { IoSettingsOutline } from "react-icons/io5";
import { IoMdNotificationsOutline } from "react-icons/io";

// Import State Management Utils
import { AppDispatch, useAppSelector } from "@/lib/store";
import { updateToggle } from "@/lib/features/navigation/navigationSlice";

const Navbar = () => {
  const dispatch = useDispatch<AppDispatch>();
  const toggle = useAppSelector(
    (state) => state.navigationReducer.value.toggle
  );
  const activePage = useAppSelector(
    (state) => state.navigationReducer.value.activePage
  );
  return (
    <div className="flex flex-col gap-5 py-5 border-b px-10">
      <div className="flex gap-5 justify-between items-center">
        <div className="text-2xl text-primary-color-800 md:hidden ">
          <button
            onClick={() => {
              dispatch(updateToggle(!toggle));
            }}
          >
            <GiHamburgerMenu />
          </button>
        </div>
        <div className="font-bold text-2xl text-primary-color-800">
          {activePage}
        </div>

        <div className="flex gap-20">
          <div className="rounded-full hidden md:flex md:gap-2 bg-primary-color-50 text-primary-color-200 text-sm font-normal py-3 px-8 ml-2 items-center">
            <span className="text-xl">
              <MdOutlineSearch />
            </span>

            <input
              type="text"
              placeholder="Search for Something"
              className="bg-transparent border-none outline-none text-primary-color-200 placeholder-primary-color-200  text-sm flex-grow"
            />
          </div>

          <div className="hidden md:flex gap-5 text-xl md:items-center">
            <div className="cursor-pointer text-xl bg-primary-color-50 rounded-full px-2 py-2">
              <span className="text-2xl">
                <IoSettingsOutline />
              </span>
            </div>
            <div className="cursor-pointer text-xl bg-primary-color-50 text-red-400 rounded-full px-2 py-2">
              <span className="text-2xl">
                <IoMdNotificationsOutline />
              </span>
            </div>
          </div>
          <div className="items-center">
            <Image
              //TODO: CHANGE IMAGE
              src="https://res.cloudinary.com/dtt1wnvfb/image/upload/v1701954159/photo_2023-12-07%2016.02.23.jpeg.jpg"
              alt="Profile"
              width={35}
              height={35}
            ></Image>
          </div>
        </div>
      </div>

      <div className="flex md:hidden rounded-full bg-primary-color-50 text-primary-color-200  text-sm font-normal gap-2 items-center py-3 px-4 ml-2">
        <span className="text-xl">
          <MdOutlineSearch />
        </span>
        Search for Something
      </div>
    </div>
  );
};

export default Navbar;
