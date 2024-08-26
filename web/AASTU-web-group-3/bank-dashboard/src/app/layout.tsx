"use client";
import "./globals.css";
import Provider from "@/lib/redux/Provider";
import { ThemeProvider } from "@/contexts/Theme";

import ClientLayout from "../lib/ClientLayout";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <>
      <Provider>
        <ThemeProvider>
          <html lang="en" data-theme="light">
            <body className="flex h-screen overflow-hidden ">
              <ClientLayout>{children}</ClientLayout>
            </body>
          </html>
        </ThemeProvider>
      </Provider>
    </>
  );
}
