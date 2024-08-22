'use client';

import { useState, useEffect } from "react";
import { Inter } from "next/font/google";
import { SessionProvider } from "next-auth/react";
import { Provider } from 'react-redux';
import store from './redux/store';
import "./globals.css";
import NavBar from "./components/common/navBar";
import SideBar from "./components/common/sideBar";
import { useSession } from "next-auth/react";

const inter = Inter({ subsets: ["latin"] });

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const [isSidebarVisible, setIsSidebarVisible] = useState(false);
  const [darkmode, setDarkmode] = useState(false);

  useEffect(() => {
    document.documentElement.setAttribute('data-theme', darkmode ? 'dark' : 'light');
  }, [darkmode]);

  const toggleSidebar = () => {
    setIsSidebarVisible(!isSidebarVisible);
  };

  const toggleDarkMode = () => {
    setDarkmode(!darkmode);
  };

  return (
    <html lang="en">
      <body className={inter.className}>
        <SessionProvider>
          <Provider store={store}>
            <SessionWrapper>
              <div className={`min-h-screen flex ${darkmode ? 'dark' : ''}`}>
                <SidebarWrapper
                  isSidebarVisible={isSidebarVisible}
                  toggleSidebar={toggleSidebar}
                />
                <div className="flex flex-col flex-1 transition-all duration-300">
                  <NavBar
                    toggleSidebar={toggleSidebar}
                    isSidebarVisible={isSidebarVisible}
                    toggleDarkMode={toggleDarkMode}
                    darkmode={darkmode}
                  />
                  <main >
                    {children}
                  </main>
                </div>
              </div>
            </SessionWrapper>
          </Provider>
        </SessionProvider>
      </body>
    </html>
  );
}

function SessionWrapper({ children }: { children: React.ReactNode }) {
  const { status } = useSession();

  if (status === "loading") {
    return <div>Loading...</div>;
  }

  return <>{children}</>;
}

function SidebarWrapper({ isSidebarVisible, toggleSidebar }: { isSidebarVisible: boolean, toggleSidebar: () => void }) {
  const { status } = useSession();

  return (
    status === "authenticated" && (
      <div className={`fixed inset-0 bg-white z-50 sm:static sm:block ${isSidebarVisible ? 'block' : 'hidden'}`}>
        <SideBar
          isSidebarVisible={isSidebarVisible}
          toggleSidebar={toggleSidebar}
        />
      </div>
    )
  );
}