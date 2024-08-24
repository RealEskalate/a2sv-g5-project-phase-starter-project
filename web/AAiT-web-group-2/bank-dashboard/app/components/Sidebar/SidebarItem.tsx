'use client'
import React from "react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import Image from "next/image";

interface SidebarItemProps {
  item: {
    icon: string;
    label: string;
    route: string;
  };
  pageName: string;
  setPageName: (name: string) => void;
}

const SidebarItem: React.FC<SidebarItemProps> = ({
  item,
  pageName,
  setPageName,
}) => {
  const pathname = usePathname();

  const handleClick = () => {
    const updatedPageName =
      pageName !== item.label.toLowerCase() ? item.label.toLowerCase() : "";
    setPageName(updatedPageName);
  };

  const isActive = item.route === pathname;

  return (
    <li>
      <Link
        href={item.route}
        onClick={handleClick}
        className={`${
          isActive ? "text-blue-500" : "text-gray-500"
        } group relative flex items-center gap-2.5 rounded-sm px-4 py-2 font-medium duration-300 ease-in-out hover:text-blue-500`}
      >
        <Image src={item.icon} alt={item.label} width={24} height={24} />
        {item.label}
      </Link>
    </li>
  );
};

export default SidebarItem;
