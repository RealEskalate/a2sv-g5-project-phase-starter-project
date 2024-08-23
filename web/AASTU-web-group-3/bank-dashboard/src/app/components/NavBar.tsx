"use client";
import Image from "next/image";
import { GiHamburgerMenu } from "react-icons/gi";
import {
  IoSettingsOutline,
  IoSearchOutline,
  IoNotificationsOutline,
} from "react-icons/io5";
import { useSelector, useDispatch } from 'react-redux';
import { RootState } from '@/lib/redux/store'; 
import { toggleSidebar } from '@/lib/redux/slices/layoutSlice'; 
import { profilepic } from "@/../../public/Icons";

const Navbar = () => {
  const dispatch = useDispatch();
  const { ishidden, activeItem } = useSelector((state: RootState) => state.layout);

  return (
    <nav className="relative flex py-4 px-6 items-center gap-6 w-full bg-white shadow-md md:h-16">
      {!ishidden && (
        <GiHamburgerMenu
          className="md:hidden absolute top-5 left-5 text-3xl"
          onClick={() => dispatch(toggleSidebar())}
        />
      )}
      <div className="w-full flex flex-col md:flex-row gap-4 items-center justify-between md:w-[95%]">
        <div className="ml-[25%] md:ml-2 font-semibold text-[25px] text-[#343C6A]">
          {activeItem}
        </div>
        <div className="w-full md:w-auto flex items-center justify-between gap-4">
          <div className="w-full md:w-auto flex gap-2 items-center pl-5 py-3 bg-[#F5F7FA] rounded-full justify-start text-lg overflow-hidden">
            <IoSearchOutline className="text-[#718EBF] text-xl" />
            <input
              type="text"
              placeholder="search for something"
              className="outline-none text-md bg-[#F5F7FA]"
            />
          </div>
          <div className="hidden lg:block p-3 rounded-full text-xl text-[#718EBF] bg-[#F5F7FA]">
            <IoSettingsOutline />
          </div>
          <div className="hidden md:block p-3 rounded-full text-xl text-[#FE5C73] bg-[#F5F7FA]">
            <IoNotificationsOutline />
          </div>
        </div>
      </div>
      <div className="m-2 absolute top-0 right-0 rounded-full overflow-hidden w-12">
        <Image src={profilepic} alt="Profile Picture" width={48} height={48} />
      </div>
    </nav>
  );
};

export default Navbar;
