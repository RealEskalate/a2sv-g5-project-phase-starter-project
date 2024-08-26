import { useSession } from "next-auth/react";
import Image from "next/image";
import Link from "next/link";
import { motion as m, AnimatePresence } from "framer-motion";
import React, { useState } from "react";
import { AiOutlineClose } from "react-icons/ai"; // Import the close icon from React Icons
import SignIn from "../authAndProfileManagement/signIn/SignIn";

interface LandingNavProps {
  bgWhite: boolean;
  homeRef: React.RefObject<HTMLDivElement>;
  servicesRef: React.RefObject<HTMLDivElement>;
  aboutRef: React.RefObject<HTMLDivElement>;
}

const LandingNav: React.FC<LandingNavProps> = ({
  bgWhite,
  homeRef,
  servicesRef,
  aboutRef,
}) => {
  const [isMenuVisible, setIsMenuVisible] = useState<boolean>(false);
  const [isSignInModalVisible, setIsSignInModalVisible] = useState<boolean>(false);
  const { data: session, status } = useSession();

  const toggleMenu = () => {
    setIsMenuVisible(!isMenuVisible);
  };

  const scrollToRef = (ref: React.RefObject<HTMLDivElement>) => {
    if (ref.current) {
      window.scrollTo({
        top: ref.current.offsetTop,
        behavior: "smooth",
      });
    }
  };

  const toggleSignInModal = () => {
    setIsSignInModalVisible(!isSignInModalVisible);
  };

  return (
    <m.header 
     


      className={`bg-[#083E9E] text-white flex justify-center items-center sm:justify-between p-2 sm:p-4 relative z-[800]`}
    >
      <Link href="/" className="font-extrabold text-[25px] cursor-pointer">BankDash</Link>

      {/* Desktop Navigation Links */}
      <div className="hidden sm:flex gap-10 items-center">
        <div className="cursor-pointer text-white" onClick={() => scrollToRef(homeRef)}>
          Home
        </div>
        <div className="cursor-pointer text-white" onClick={() => scrollToRef(servicesRef)}>
          Services
        </div>
        <div className="cursor-pointer text-white" onClick={() => scrollToRef(aboutRef)}>
          About Us
        </div>
        {status === "authenticated" ? (
          <Link href="/dashboard">
            <div className="cursor-pointer text-white">Dashboard</div>
          </Link>
        ) : (
          <div onClick={toggleSignInModal} className="cursor-pointer text-white">
            Login
          </div>
        )}
      </div>

      {/* Mobile Menu Toggle Button */}
      <div className="flex items-center sm:hidden">
        <Image
          src="assets/landing/hamburger.svg"
          width={25}
          height={25}
          alt="hamburger"
          className={`absolute right-5 ${isMenuVisible&&"hidden"}`}
          onClick={toggleMenu}
        />

        <div
          className={`${
            isMenuVisible
              ? "absolute right-[-20px] bg-white mt-1 top-full text-[#083E9E] rounded-lg p-2"
              : "hidden"
          } sm:flex sm:flex-row gap-20 mr-10`}
        >
          <div className="mb-2 sm:mb-0 cursor-pointer" onClick={() => scrollToRef(homeRef)}>Home</div>
          <div className="mb-2 sm:mb-0 cursor-pointer" onClick={() => scrollToRef(servicesRef)}>Services</div>
          <div className="mb-2 sm:mb-0 cursor-pointer" onClick={() => scrollToRef(aboutRef)}>About Us</div>
          {status === "authenticated" ? (
            <Link href="/dashboard">
              <div className="mb-2 sm:mb-0">Dashboard</div>
            </Link>
          ) : (
            <div onClick={toggleSignInModal} className="mb-2 sm:mb-0 cursor-pointer">
              Login
            </div>
          )}
        </div>
      </div>
      {isSignInModalVisible && <SignIn onClose={toggleSignInModal} />}
    </m.header>
  );
};

export default LandingNav;
