"use client";
import { useSession } from "next-auth/react";
import Image from "next/image";
import Link from "next/link";
import React, { useState } from "react";
import { motion as m } from 'framer-motion';

interface LandingNavProps {
  bgWhite: boolean;
  homeRef: React.RefObject<HTMLDivElement>;
  servicesRef: React.RefObject<HTMLDivElement>;
  aboutRef: React.RefObject<HTMLDivElement>;
}

const LandingNav: React.FC<LandingNavProps> = ({ bgWhite, homeRef, servicesRef, aboutRef }) => {
  const [isMenuVisible, setIsMenuVisible] = useState<boolean>(false);
  const { data: session, status } = useSession();

  const toggle = () => {
    setIsMenuVisible(!isMenuVisible);
  };

  const scrollToRef = (ref: React.RefObject<HTMLDivElement>) => {
    if (ref.current) {
      window.scrollTo({
        top: ref.current.offsetTop,
        behavior: 'smooth'
      });
    }
  };

  return (
    <m.header 
     
  
      className={`bg-[#083E9E] text-white flex justify-center items-center sm:justify-between p-2 sm:p-4 relative`}
    >
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
          <div className="mb-2 sm:mb-0" onClick={() => scrollToRef(homeRef)}>Home</div>
          <div className="mb-2 sm:mb-0" onClick={() => scrollToRef(servicesRef)}>Services</div>
          <div className="mb-2 sm:mb-0" onClick={() => scrollToRef(aboutRef)}>About Us</div>
          {status === "authenticated" ? (
            <Link href="/dashboard">
              <div className="mb-2 sm:mb-0">Dashboard</div>
            </Link>
          ) : (
            <Link href="/login">
              <div className="mb-2 sm:mb-0">Login</div>
            </Link>
          )}
        </div>
      </div>
    </m.header>
  );
};

export default LandingNav;
