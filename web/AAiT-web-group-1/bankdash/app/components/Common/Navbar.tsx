'use client';
import React from "react";
const logo = require("../../../public/images/Logo.svg");
const bell = require("../../../public/images/bell.svg");
const settings = require("../../../public/images/settings.svg");
const hamburger = require("../../../public/images/ham.jpg");
const search = require("../../../public/images/search.svg");
const profile = require("../../../public/images/Mask Group.svg");
import Image from "next/image";
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../../../lib/redux/store';
import { toggleSidebar } from '../../../lib/redux/slices/menuSlice';

interface Props {
  title: string;
  profilepic: string;
}

const Header = () => {
  const title = "Overview"
  const profilepic = profile
  const dispatch: AppDispatch = useDispatch();

  const handleBurgerClick = () => {
    dispatch(toggleSidebar());
  };



  return (
    <div className="flex bg-white md:pr-10 pb-1 md:px-0 items-center lg:justify-between w-full top-0 fixed md:block z-10">
      {/* <div className="w-1/6 md:block hidden">
        <Image src={logo} className="ml-1" alt="LOGO" />
      </div> */}

      <div className=" w-full md:pl-4 flex">
        <div className="flex flex-row flex-wrap justify-center md:items-center align-middle w-full space-y-4 md:space-y-0 md:ml-60 ">
          <div className="md:hidden  flex justify-start align-middle pt-3 w-1/3 z-40" onClick={handleBurgerClick} >
            <Image
              className="w-8 h-8 order-1 bg-white"
              src={hamburger}
              alt="hamburger"
            />
          </div>

          <div className="md:text-2xl md:order-none text-xl order-2 md:mr-10 lg:mr-24 xl:mr-72 w-fit  md:mb-0 text-[#343C6A] font-bold ">
            <h1 className="ml-2 font-bold">{title}</h1>
          </div>

          <div className="flex lg:ml-32 order-4 md:order-none bg-[#F5F7FA] h-10 md:w-60 w-3/4 items-center rounded-full  ">
            <Image src={search} className="w-3 h-3 ml-2" alt="search" />
            <input
              className="bg-[#F5F7FA] md:w-38 w-10/12 px-4 outline-none placeholder:bg-[#F5F7FA] placeholder:text-xs md:placeholder:text-base"
              type="text"
              placeholder={`search for something`}
            />
          </div>

          <Image
            className="w-10 ml-7 h-20 md:block hidden"
            src={settings}
            alt="settings"
          />
          <Image
            className="w-10 h-20 ml-7 self-end md:block hidden"
            src={bell}
            alt="bell"
          />
          <div className="flex justify-end md:w-fit w-1/3 md:mb-5 order-3 md:order-3 md:ml-7">
            <Image
              className="w-10 h-10"
              src={profilepic}
              alt="profile_picture"
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Header;
