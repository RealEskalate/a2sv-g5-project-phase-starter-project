import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Sidebar from "@/components/Layout/sidebar";
import Header from "@/components/Layout/header";


const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Bank Dash",
  description: "implemented by G54 G8",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${inter.className} flex `}>
        <Sidebar/>
        <div className="flex flex-col w-full">
          <Header />
          <main className="p-6 bg-gray-100 min-h-full">
            {children}
          </main>
        </div>
      </body>
    </html>
  );
}