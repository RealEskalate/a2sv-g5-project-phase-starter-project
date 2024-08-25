"use client"
import { Provider } from "react-redux";
import store from "./store";
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
      <Provider store={store}>
        <body className="flex w-full h-full">
          <Sidebar />
          <div className="w-full flex flex-col h-full">
            <Header />
            {children}
          </div>
        </body>
      </Provider>
    </html>
  );
}
