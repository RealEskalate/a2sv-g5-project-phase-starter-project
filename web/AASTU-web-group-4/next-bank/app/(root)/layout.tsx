'use client'

import { FC, useState } from 'react';
import { Sidebar } from '@/components/Sidebar';
import { Navbar } from '@/components/Navbar';
import { usePathname } from 'next/navigation';
import { sidebarLinks, user } from '@/constants';



const RootLayout: FC<{ children: React.ReactNode }> = ({ children }) => {
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);
  const toggleSidebar = () => setIsSidebarOpen(!isSidebarOpen);

  const pathname = usePathname();
  const pageTitle = pathname === "/"
    ? "Overview"
    : sidebarLinks.find(link => link.route === pathname)?.label || "";

  return (
    <div className="flex overflow-x-hidden">
      <Sidebar isOpen={isSidebarOpen} toggleSidebar={toggleSidebar} />
      <div className="flex-1 flex flex-col">
        <Navbar pageTitle={pageTitle} toggleSidebar={toggleSidebar} userProfileImage={user.profileImage} />
        <main className="flex-grow">{children}</main>
      </div>
    </div>
  );
};

export default RootLayout;
