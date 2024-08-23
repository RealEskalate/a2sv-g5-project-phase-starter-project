import type { Metadata } from "next";
import "./globals.css";
import Sidebar from "./components/Sidebar";
import Navbar from "./components/Navbar";
export const metadata: Metadata = {
  title: "A2SV Wallet",
  description: "Built for a2sv",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="flex flex-col gap-5 ">
        <div className="bg-gray-300 px-5">

        <Navbar />
        </div>
        <div className="flex gap-5">
          <div className="bg-gray-300">
          <Sidebar></Sidebar>
          </div>
          {children}
        </div>
      </body>
    </html>
  );
}
