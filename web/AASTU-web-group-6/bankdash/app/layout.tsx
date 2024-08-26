import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { usePathname } from "next/navigation";
import LayoutProvider from "./Provider/LayoutProvider";
import ReduxProvider from "./Redux/ReduxProvider";
import SessionWrapper from "./Provider/SessionWrapper";
import { Toaster } from "@/components/ui/toaster";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "BankDash",
  description: "Banking Made Easy",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <SessionWrapper>
      <html lang="en">
        <body className={`${inter.className}`}>
          <ReduxProvider>
            <LayoutProvider>{children}</LayoutProvider>
            <Toaster />
          </ReduxProvider>
        </body>
      </html>
    </SessionWrapper>
  );
}
