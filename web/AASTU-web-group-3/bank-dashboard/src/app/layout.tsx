"use client";
import "./globals.css";
import Provider from "@/lib/redux/Provider";
import { ThemeProvider } from "@/contexts/Theme";
import { SessionProvider } from "next-auth/react";

import ClientLayout from "../lib/ClientLayout";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <>
      <Provider>
        <ThemeProvider>
          <SessionProvider>
          <html lang="en" data-theme="light">
            <body className="flex h-screen overflow-hidden ">
              <ClientLayout>{children}</ClientLayout>
            </body>
          </html>
          </SessionProvider>
        </ThemeProvider>
      </Provider>
    </>
  );
}
