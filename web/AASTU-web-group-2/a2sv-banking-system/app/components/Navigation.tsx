"use client";
import React, { useState } from "react";
import Navbar from "./Navbar";
import Sidebar from "./Sidebar";
interface Props {
  children: React.ReactNode;
}
const Navigation = ({ children }: Props) => {
  const [toggle, setToggle] = useState(false);
  return (
    <>
      <div className="flex w-full">
        <div className="">
          <Sidebar
            toggle={toggle}
            handleClose={() => {
              setToggle(!toggle);
            }}
          ></Sidebar>
        </div>

        <div className="flex flex-col gap-5 w-full">
          <div className="w-full">
            <Navbar
              handleClick={() => {
                setToggle(!toggle);
              }}
            />
          </div>
          {children}
        </div>
      </div>
    </>
  );
};

export default Navigation;
