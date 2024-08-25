"use client";
import React, { useEffect } from "react";
import Hamburger from "../../../public/assets/icons/hamburger-icon.svg";
import Settings from "../../../public/assets/icons/setting-icon.svg";
import Notification from "../../../public/assets/icons/notification-icon.svg";
import { usePathname } from "next/navigation";
import { useAppDispatch, useAppSelector } from "@/hooks/hoooks";
import { toggleHamburgerMenu } from "@/lib/redux/slices/uiSlice";
import Link from "next/link";

import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import AvatarSkeleton from "../AllSkeletons/Navigation/AvatarSkeleton";
import { useGetProfileQuery } from "@/lib/redux/api/profileAPI";
import { setProfile } from "@/lib/redux/slices/profileSlice";
import { LogOut } from "lucide-react";

export default function NavBar() {
  const path: { [key: string]: string } = {
    "/bank-dash": "Overview",
    "/bank-dash/transactions": "Transactions",
    "/bank-dash/accounts": "Accounts",
    "/bank-dash/investments": "Investments",
    "/bank-dash/credit-card": "Credit Card",
    "/bank-dash/loans": "Loans",
    "/bank-dash/services": "Services",
    "/bank-dash/settings": "Settings",
  };

  const getData = useAppSelector((state) => state.profile);
  const dispatch = useAppDispatch();
  const { data, isSuccess } = useGetProfileQuery();

  useEffect(() => {
    if (isSuccess && data) {
      dispatch(setProfile(data?.data));
      console.log(data.data);
    }
  }, [data, dispatch]);

  const pathname = usePathname();
  console.log(pathname);

  const handleClick = () => {
    console.log("toggle");
    dispatch(toggleHamburgerMenu());
  };

  return (
    <div className="w-full p-5 bg-white h-[70px] flex justify-between px-5 items-center">
      <button
        type="button"
        className="items-center rounded-lg sm:hidden font-bold"
        onClick={handleClick}
      >
        <Hamburger className="w-7 h-7 object-cover" />
      </button>
      <p className="font-semibold text-xl text-navy">
        {path[pathname as string]}
      </p>
      <div className="flex space-x-2 items-center">
        <div className="hidden md:flex ">
          <input
            type="text"
            placeholder="Search for something"
            className="bg-slate-100 px-3 pl-5 rounded-full py-2 outline-none"
          />
        </div>
        <div className=" hidden md:flex w-11 h-11 bg-slate-100 rounded-full justify-center items-center text-blue-steel">
          <Link href="/bank-dash/settings">
            <Settings className="w-6 h-6" />
          </Link>
        </div>
        <div className="hidden md:flex w-11 h-11 bg-slate-100 rounded-full justify-center items-center text-red-400">
          <Notification className="w-6 h-6" />
        </div>
        <div className="w-11 h-11 object-contain object-center rounded-full overflow-hidden">
          <Popover>
            <PopoverTrigger>
              {getData?.profilePicture ? (
                <img
                  src={getData.profilePicture || "/assets/default-user.png"}
                  alt=""
                  className="object-cover rounded-full w-[50px] h-[50px] md:w-[50px] md:h-[50px]"
                />
              ) : (
                <AvatarSkeleton />
              )}
            </PopoverTrigger>
            <PopoverContent className="divide-y divide-blue-200">
              <Link href={"/api/auth/signout"} className="w-full gap-4 flex">
              <LogOut />
                Logout
              </Link>
            </PopoverContent>
          </Popover>
        </div>
      </div>
    </div>
  );
}
