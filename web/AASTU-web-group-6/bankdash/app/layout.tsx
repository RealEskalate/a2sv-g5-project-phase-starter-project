import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import NavBar from "./components/Layout/NavBar";
import Sidebar from "./components/Layout/Sidebar";
import { usePathname } from "next/navigation";
import LayoutProvider from "./Provider/LayoutProvider";
import ReduxProvider from "./Redux/ReduxProvider";
import SessionWrapper from "./Provider/SessionWrapper";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <SessionWrapper>
      <html lang="en">
        <body className={`${inter.className} bg-[#f5f7fa]`}>
          {/* i just add redux provider here */}
          <ReduxProvider>
            <LayoutProvider>{children}</LayoutProvider>
          </ReduxProvider>
        </body>
      </html>
    </SessionWrapper>
  );
}
