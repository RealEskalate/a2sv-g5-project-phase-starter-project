import React from "react";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import Image from "next/image";

const Header = ({ title }: { title: string }) => {
  return (
    <div className="bg-white mb-5 max-md:hidden">
      <div className="flex justify-between px-10 py-4">
        <h1 className="text-3xl text-primaryBlack font-[600]">{title}</h1>

        <div className="flex gap-5 items-center">
          {/* Search */}
          <div className="flex gap-3 bg-[#F5F7FA] p-3 rounded-full">
            <Image
              src="/icons/Search.svg"
              width={20}
              height={20}
              alt="Search"
            />
            <input
              className="outline-none bg-[#F5F7FA]"
              type="text"
              placeholder="Search for something"
            />
          </div>

          {/* Settings */}
          <div className="bg-[#F5F7FA] p-2 rounded-full cursor-pointer">
            <Image
              src="/icons/Settings.svg"
              width={22}
              height={22}
              alt="Notif"
            />
          </div>

          <div className="bg-[#F5F7FA] p-2 rounded-full cursor-pointer">
            {/* Notif */}
            <Image
              src="/icons/Notification.svg"
              width={22}
              height={22}
              alt="Notif"
            />
          </div>
          {/* Avatar */}
          <Avatar>
            <AvatarImage src="https://github.com/shadcn.png" />
            <AvatarFallback>CN</AvatarFallback>
          </Avatar>
        </div>
      </div>
    </div>
  );
};

export default Header;
