'use client'
import { useState, Suspense } from 'react';
import { Inter } from "next/font/google";
import "./globals.css";
import NavBar from "./components/common/navBar";
import Sidebar from "./components/common/sideBar"; 
import { metadata } from "./layoutMetadata"; 

const inter = Inter({ subsets: ["latin"] });

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const [isSidebarVisible, setIsSidebarVisible] = useState(false);

  const toggleSidebar = () => {
    setIsSidebarVisible(!isSidebarVisible);
  };

  return (
    <html lang="en">
      <body className={inter.className}>
        <div className="min-h-screen bg-white sm:grid sm:grid-cols-[200px_1fr] md:grid-cols-[250px_1fr]">
          <div className={`fixed inset-0 bg-white z-50 sm:static sm:block ${isSidebarVisible ? 'block' : 'hidden'}`}>
            <Suspense fallback={<div>Loading...</div>}>
              <Sidebar isSidebarVisible={isSidebarVisible} toggleSidebar={toggleSidebar} />
            </Suspense>
          </div>
          <div className="flex flex-col w-full">
            <NavBar toggleSidebar={toggleSidebar} isSidebarVisible={isSidebarVisible} />
            <main>{children}</main>
          </div>
        </div>
      </body>
    </html>
  );
}