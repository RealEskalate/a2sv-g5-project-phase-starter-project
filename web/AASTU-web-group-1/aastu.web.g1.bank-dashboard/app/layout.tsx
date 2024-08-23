import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
// import { Toaster } from "sonner";
import { Toaster } from "@/components/ui/toaster";
import { UserProvider } from "@/contexts/UserContext";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "BankDash",
  description: "The best banking platform",
  icons: { icon: "/icons/logo.png" },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <head>
        <link rel="icon" href="/icons/logo.png" />
      </head>
      <body className={`inter.className`}>
        <Toaster />
        <UserProvider>{children} </UserProvider>
      </body>
    </html>
  );
}
