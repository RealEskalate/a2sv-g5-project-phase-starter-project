import type { Metadata } from "next";
import { Inter } from "next/font/google";
import Sidebar from "./components/Sidebar/Sidebar";
import Header from "./components/Header/Header";
import "./globals.css";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="flex h-screen w-screen">
        <Sidebar />
        <div className="w-full flex flex-col h-full">
          <Header />
          {children}
        </div>
      </body>
    </html>
  );
}