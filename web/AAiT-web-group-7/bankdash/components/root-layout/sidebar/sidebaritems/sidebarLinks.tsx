import Image from "next/image";
import Link from "next/link";
import React from "react";

interface Props {
  title: string;
  Icon: string; 
  link: string;
  isActive: boolean;
  collapsed: boolean;
}

const SidebarLink: React.FC<Props> = ({ title, Icon, link, isActive, collapsed }) => {
  return (
    <Link href={link}>
      <div
        className={`text-sm flex items-center gap-3 p-4 mx-5 rounded-xl ${isActive ? "bg-blue-700 text-white" : ""}`}
      >
        <Image src={Icon} alt="title" className={`w-5 h-5 ${isActive ? '':""} `} />
        <p className={`${collapsed ? "hidden" : "block"}`}>{title}</p>
      </div>
    </Link>
  );
};

export default SidebarLink;
