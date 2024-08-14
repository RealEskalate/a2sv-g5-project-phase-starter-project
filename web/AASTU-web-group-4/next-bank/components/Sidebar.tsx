import { FC } from 'react';
import Link from 'next/link';
import { sidebarLinks, user } from "@/constants";

export const Sidebar: FC = () => {
  return (
    <div className="hidden md:flex flex-col w-64 h-full bg-white shadow-lg p-4 fixed left-0 top-0">
      <div className="space-y-4">
        {sidebarLinks.map((link) => (
          <Link key={link.route} href={link.route} className="flex items-center p-2 hover:underline hover:text-blue-600">
            <link.Icon className="mr-3" size={25} />
            <span>{link.label}</span>
          </Link>
        ))}
      </div>
    </div>
  );
};
