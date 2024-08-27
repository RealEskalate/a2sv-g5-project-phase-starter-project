import Image from "next/image";
import { GiHamburgerMenu } from "react-icons/gi";
import {
  IoSettingsOutline,
  IoSearchOutline,
  IoNotificationsOutline,
  IoMoon,
  IoSunny,
} from "react-icons/io5";
import { useSelector, useDispatch } from 'react-redux';
import { FaUserEdit } from "react-icons/fa";
import { AiOutlineLogout } from "react-icons/ai";
import { RootState } from '@/lib/redux/store'; 
import { toggleSidebar } from '@/lib/redux/slices/layoutSlice'; 
import { profilepic } from "@/../../public/Icons";
import { useGetCurrentUserQuery } from "@/lib/redux/api/settingApi";
import { useEffect, useState } from "react";
import { setError, setLoading, setSetting } from "@/lib/redux/slices/settingSlice";
import { signOut } from "next-auth/react";
import { useTheme } from "@/contexts/Theme";
import React from 'react';
import Link from "next/link";

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  name:string
}

const ProfileModal = ({ isOpen, onClose,name }: ModalProps) => {
  if (!isOpen) return null;

  return (
    <div className="bg-white dark:bg-darkBackground rounded-xl shadow-lg w-52 h-52 md:w-56 md:h-52 p-6 absolute top-14 right-0 z-50">
      <div className="w-full flex items-center justify-between border-b-2 pb-3 dark:border-darkPage">
        <div className="font-semibold">
          {name}
        </div>
        <button className=" px-3 py-1 text-2xl bg-[#eceef0] hover:bg-[#bcbdc0] dark:bg-gray-700 dark:hover:bg-gray-600 rounded-full" onClick={onClose}>
          &times;
        </button>
      </div>
      <div className="flex flex-col mt-2 gap-4">
        <Link href={"/setting"} className="flex gap-3 items-center rounded-xl text-xl font-semibold hover:bg-[#F5F7FA] dark:hover:bg-gray-700 p-2" onClick={onClose}>
          <FaUserEdit />
          Edit Profile
        </Link>
        <button
          className="flex gap-3 text-red-600 items-center rounded-xl text-xl font-semibold hover:bg-[#F5F7FA] dark:hover:bg-gray-700 p-2"
          onClick={() => {
            signOut({ callbackUrl: '/auth/signin' });
            onClose();
          }}
        >
          <AiOutlineLogout />
          Logout
        </button>
      </div>
    </div>
  );
};

const Navbar = () => {
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false);
  const dispatch = useDispatch();
  const { ishidden, activeItem } = useSelector((state: RootState) => state.layout);
  const { setting, loading, error } = useSelector((state: RootState) => state.setting);
  const { data, isLoading, isError } = useGetCurrentUserQuery();

  useEffect(() => {
    dispatch(setLoading(isLoading));
  
    if (data) {
      dispatch(setSetting([data]));
    }
  
    if (isError) {
      dispatch(setError("Error loading transactions"));
    }
  }, [data, isLoading, isError, dispatch]);

  const { theme, toggleTheme } = useTheme(); // Updated to use `toggleTheme`

  return (
    <>
      <nav className="relative flex py-4 px-6 items-center gap-6 w-full bg-white dark:bg-darkBackground shadow-md md:h-16">
        {!ishidden && (
          <GiHamburgerMenu
            className="md:hidden absolute top-5 left-5 text-3xl text-[#343C6A] dark:text-darkText"
            onClick={() => dispatch(toggleSidebar())}
          />
        )}
        <div className="w-full flex flex-col md:flex-row gap-4 items-center justify-between md:w-[95%]">
          <div className="ml-[25%] md:ml-6 font-semibold text-[25px] text-[#343C6A] dark:text-darkText">
            {activeItem}
          </div>
          <div className="w-full md:w-auto flex items-center justify-between gap-4">
            <div className="w-full md:w-auto flex gap-2 items-center pl-5 py-3 bg-[#F5F7FA] dark:bg-gray-700 rounded-full justify-start text-lg overflow-hidden">
              <IoSearchOutline className="text-[#718EBF] dark:text-darkAccent text-xl" />
              <input
                type="text"
                placeholder="search for something"
                className="outline-none text-md bg-[#F5F7FA] dark:bg-gray-700 dark:text-darkText"
              />
            </div>
            <div className="lg:block p-3 rounded-full text-xl text-[#718EBF] dark:text-yellow-300 bg-[#F5F7FA] dark:bg-gray-700" onClick={toggleTheme}>
              {theme === 'light' ? <IoMoon /> : <IoSunny />}
            </div>
            <div className="hidden md:block p-3 rounded-full text-xl text-[#FE5C73] bg-[#F5F7FA] dark:bg-gray-700 dark:text-darkText">
              <IoSettingsOutline />
            </div>
            <div className="hidden md:block p-3 rounded-full text-xl text-[#FE5C73] bg-[#F5F7FA] dark:bg-gray-700 dark:text-darkText">
              <IoNotificationsOutline />
            </div>
          </div>
        </div>
        <div className="m-2 absolute top-0 right-0 rounded-full overflow-hidden w-14 h-14 cursor-pointer">
          <Image 
            src={setting[0]?.data?.profilePicture} 
            alt="Profile Picture" 
            width={56} 
            height={48} 
            onClick={() => setIsModalOpen(true)} // Open the modal on click
          />
        </div>
        <ProfileModal 
          isOpen={isModalOpen} 
          onClose={() => setIsModalOpen(false)}
          name = {setting[0]?.data?.name} // Close the modal
        />
      </nav>
    </>
  );
};

export default Navbar;
