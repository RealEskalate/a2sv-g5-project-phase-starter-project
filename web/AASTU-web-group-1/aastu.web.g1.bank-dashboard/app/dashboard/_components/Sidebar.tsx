import React from "react";
import { sidebarLinks } from "@/constants";
import Link from "next/link";
const Sidebar = () => {
  return (
    <div className="bg-white h-[100vh] p-5">
      {/* <h1>BankDash</h1> */}
      {sidebarLinks.map((link, index) => {
        return (
          <Link
            href={link.route}
            className="flex items-center space-x-3 p-3 hover:bg-gray-100 cursor-pointer rounded-lg"
          >
            {/* <img src={link.icon} alt={link.title} className="w-6 h-6" /> */}
            <p>{link.title}</p>
          </Link>
        );
      })}
    </div>
  );
};

export default Sidebar;
