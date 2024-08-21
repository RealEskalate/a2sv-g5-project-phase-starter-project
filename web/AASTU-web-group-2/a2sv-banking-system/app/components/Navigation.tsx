"use client";
import React, { useState, useEffect } from "react";
import Navbar from "./Navbar";
import Sidebar from "./Sidebar";
import { getSession } from "next-auth/react";
interface Props {
  children: React.ReactNode;
}

const Navigation: React.FC<Props> = ({ children }) => {
  const [toggle, setToggle] = useState(false);
  const [session, setSession] = useState(false);

  useEffect(() => {
    const fetchSession = async () => {
      const sessionData = await getSession();
      if (sessionData?.user) {
        console.log("Session Available")
        setSession(true);
      }
    };

    fetchSession();
  }, []);

  return (
    <div className="flex w-full">
      {session && (
        <div className="z-50 bg-white">
          <Sidebar
            toggle={toggle}
            handleClose={() => {
              setToggle(!toggle);
            }}
          />
        </div>
      )}

      <div className="flex flex-col w-full">
        {session && (
          <div className="w-full bg-white md:sticky md:top-0 md:z-10">
            <Navbar
              handleClick={() => {
                setToggle(!toggle);
              }}
            />
          </div>
        )}
        {children}
      </div>
    </div>
  );
};

export default Navigation;
