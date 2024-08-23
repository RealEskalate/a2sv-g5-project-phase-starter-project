"use client";
import { useState, useEffect } from "react";
import NavBar from "@/components/root-layout/Navbar/NavBar";
import Sidebar from "@/components/root-layout/sidebar/sidebar";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>): JSX.Element {
  const [open, setOpen] = useState<boolean>(true);
  const [collapse, setCollapse] = useState<boolean>(false);

  useEffect(() => {
    window.addEventListener("resize", () => {
      window.innerWidth < 767 ? setOpen(false) : setOpen(true);
    });
  }, []);
  return (
    <>
      <div className="flex">
        <Sidebar
          open={open}
          onClose={() => {
            setOpen(false);
          }}
          collapse={collapse}
          onCollapse={() => {
            setCollapse(!collapse);
          }}
        />
        <div className={`w-full h-full border ${collapse ? "md:ml-[110px]" : "md:ml-[260px]"}`}>
          <NavBar
            openSidebar={() => {
              setOpen(true);
            }}
          />
          <div className="max-w-[1300px] mx-auto bg-gray-50">{children}</div>
        </div>
      </div>
    </>
  );
}
