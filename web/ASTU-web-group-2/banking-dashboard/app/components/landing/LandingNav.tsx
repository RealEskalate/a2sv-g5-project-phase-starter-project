"use client";
import { useSession } from "next-auth/react";
import Image from "next/image";
import Link from "next/link";
import React, { useEffect, useState } from "react";

const LandingNav = ({bgWhite}:{bgWhite:boolean}) => {
  const [isMenuVisible, setIsMenuVisible] = useState<boolean>(false);
  const { data: session, status } = useSession();

  const toggle = () => {
    setIsMenuVisible(!isMenuVisible);
  };

  useEffect(() => {
    console.log(isMenuVisible);
  }, [isMenuVisible]);

  return (
    <header className={`bg-[#083E9E] text-white flex justify-center items-center sm:justify-between p-2 sm:p-4  relative `}>
      <div className="font-extrabold text-[25px]">BankDash</div>

      <div className="flex items-center">
        <Image
          src="assets/landing/hamburger.svg"
          width={25}
          height={25}
          alt="hamburger"
          className="sm:hidden absolute right-5"
          onClick={toggle}
        />

        <div
          className={`${
            isMenuVisible
              ? "absolute right-[-20px] bg-white mt-1 top-full text-[#083E9E] z-50 rounded-lg p-2"
              : "hidden"
          } sm:flex sm:flex-row gap-20 mr-10`}
        >
          <Link href={"#home"}>
            <div className="mb-2 sm:mb-0">Home</div>
          </Link>
          <Link href={"#services"}>
            <div className="mb-2 sm:mb-0">Services</div>
          </Link>

          <Link href={"#about"}>
            <div className="mb-2 sm:mb-0">About Us</div>
          </Link>
          {status === "authenticated" ? (
            <Link href={"/dashboard"}>
              <div className="mb-2 sm:mb-0">DashBoard</div>
            </Link>
          ) : (
            <Link href={"/login"}>
              <div className="mb-2 sm:mb-0">Login</div>
            </Link>
          )}
        </div>
      </div>
    </header>
  );
};

export default LandingNav;
