import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { Toaster } from "sonner";
import { UserProvider } from "@/contexts/UserContext";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "BankDash",
  description: "The best banking platform",
  icons: "./icons/logo.png",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
     

      <Toaster />
      <body className={`inter.className `}> 
        <UserProvider>{children}  </UserProvider></body>
    
    </html>
  );
}
