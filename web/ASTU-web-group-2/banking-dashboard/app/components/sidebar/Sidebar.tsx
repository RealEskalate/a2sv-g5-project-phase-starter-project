import React from "react";
import { sidebarListItems } from "./sidebarListItems";
import Link from "next/link";
import Image from "next/image";

const Sidebar = () => {
  return (
    <div className="absolute top-0 left-0 bottom-0 border-r border-[#E6EFF5] w-[240px]">
      <div className="flex flex-col w-[189px] h-[420px] gap-2  pt-[114px]">
        {sidebarListItems.map((item, index) => (
          <div className="flex gap-6">
            <Link className="flex gap-5 " key={index} href={item.path}>
              {index === 0 ? (
                <div className="h-[50px] w-[5px] bg-[#2D60FF] rounded-r-lg"></div>
              ) : (
                <div className="h-[50px] w-[5px] active:bg-[#2D60FF] rounded-r-lg"></div>
              )}
              <div className="flex gap-5 pt-3">
                <div>
                  <Image
                    src={`${item.icon}`}
                    alt="icon"
                    width={20}
                    height={20}
                  />
                </div>
                {index === 0 ? (
                  <h1 className="font-medium text-[16px] text-[#2D60FF]  ">
                    {" "}
                    {item.name}
                  </h1>
                ) : (
                  <h1 className="font-medium text-[16px]  text-[#B1B1B1] active:text-[#2D60FF] ">
                    {item.name}
                  </h1>
                )}
              </div>
            </Link>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Sidebar;
