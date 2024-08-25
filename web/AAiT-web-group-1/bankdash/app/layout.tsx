'use client';
import "./globals.css";
import Header from "./components/Common/Navbar";
import Sidebar from "./components/Common/Sidebar";
import { Provider } from "react-redux";
import store from "@/lib/redux/store";


export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <Provider store={store}> 
            <body className="w-full h-full flex bg-[#f5f7fa]">
              <div className="">
                <Sidebar />
              </div>
              <div className="h-fit bg-red-300">
                <Header />
              </div>
              <main className="mt-16 lg:ml-64 sm:ml-60 w-full ">{children}</main>
            </body>
      </Provider>
    </html>
  );
}
