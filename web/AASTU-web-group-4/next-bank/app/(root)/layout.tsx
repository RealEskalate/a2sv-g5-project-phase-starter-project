'use client'
import { FC, useState, useEffect } from 'react';
import { Sidebar } from '@/components/Sidebar';
import { Navbar } from '@/components/Navbar';
import { usePathname } from 'next/navigation';
import { sidebarLinks } from '@/constants';
import { NotificationProvider } from '@/services/NotificationContext';
import {  message } from 'antd';

const RootLayout: FC<{ children: React.ReactNode }> = ({ children }) => {
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);
  const toggleSidebar = () => setIsSidebarOpen(!isSidebarOpen);

  const pathname = usePathname();
  const pageTitle = pathname === "/"
    ? "Overview"
    : sidebarLinks.find(link => link.route === pathname)?.label || "";

  useEffect(() => {
    document.body.style.overflow = isSidebarOpen ? 'hidden' : 'auto';
  }, [isSidebarOpen]);
  const [messageApi, contextHolder] = message.useMessage();

  return (
    <>
    {contextHolder}
      <NotificationProvider>  
        <div className="flex overflow-x-hidden">
          <Sidebar isOpen={isSidebarOpen} toggleSidebar={toggleSidebar} />
          <div className={`flex-1 flex flex-col transition-transform duration-300 ${isSidebarOpen ? 'ml-0' : 'ml-0'} overflow-x-hidden`}>
            <Navbar pageTitle={pageTitle} toggleSidebar={toggleSidebar} />
            <main className="flex-grow pt-[120px] lg:pt-20">{children}</main>
          </div>
        </div>
      </NotificationProvider>
    </>
  );
};

export default RootLayout;
