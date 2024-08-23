"use client";
import "./globals.css";
import Provider from "@/lib/redux/Provider";
import Sidebar from "./components/SideBar"; // Update path as needed
import Navbar from "./components/NavBar"; // Update path as needed

import ClientLayout from '../lib/ClientLayout';

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <>
      <Provider>
        <html lang="en">
          <body className="flex h-screen overflow-hidden">
            <ClientLayout>
              {children}
            </ClientLayout>
            {/* <Sidebar />
            <div className="flex flex-col flex-grow h-full overflow-hidden md:w-4/5 lg:w-4/5">
              <Navbar />
              <main className="flex-grow overflow-y-auto bg-[#F5F7FA] p-1">
                {children}
              </main>
            </div> */}
          </body>
        </html>
      </Provider>
    </>
  );
}
