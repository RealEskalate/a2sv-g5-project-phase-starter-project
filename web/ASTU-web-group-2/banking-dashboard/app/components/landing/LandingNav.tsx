import { useSession } from "next-auth/react";
import Image from "next/image";
import Link from "next/link";
import { motion as m, AnimatePresence } from "framer-motion";
import React, { useState } from "react";
import { AiOutlineClose } from "react-icons/ai"; // Import the close icon from React Icons
import SignIn from "../signIn/SignIn";

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
    <m.header className={`bg-[#083E9E] text-white flex justify-center items-center sm:justify-between p-2 sm:p-4 sm:pr-6 relative`}>
      <Link href="/" className="font-extrabold text-[25px] cursor-pointer">
        BankDash
      </Link>

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
      </div>

      {/* Overlay for Sidebar */}
      <AnimatePresence>
        {isMenuVisible && (
          <m.div
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.3 }}
            className="fixed inset-0 bg-black bg-opacity-50 z-[700]"
            onClick={toggleMenu}
          />
        )}
      </AnimatePresence>

      {/* Sidebar Menu */}
      <AnimatePresence>
        {isMenuVisible && (
          <m.div
            initial={{ x: "100%" }}
            animate={{ x: 0 }}
            exit={{ x: "100%" }}
            transition={{ duration: 0.3 }}
            className="fixed right-0 top-0 h-full z-[800] flex flex-col items-center gap-4 p-4 w-64 bg-[#002970ce] sm:hidden"
          >
            <div className="flex justify-end self-end">
              <AiOutlineClose
                size={25}
                className="cursor-pointer text-white"
                onClick={toggleMenu}
              />
            </div>
            <div className="cursor-pointer text-white" onClick={() => {scrollToRef(homeRef); toggleMenu();}}>
              Home
            </div>
            <div className="cursor-pointer text-white" onClick={() => {scrollToRef(servicesRef); toggleMenu();}}>
              Services
            </div>
            <div className="cursor-pointer text-white" onClick={() => {scrollToRef(aboutRef); toggleMenu();}}>
              About Us
            </div>
            {status === "authenticated" ? (
              <Link href="/dashboard">
                <div className="cursor-pointer text-white">Dashboard</div>
              </Link>
            ) : (
              <div onClick={() => {toggleSignInModal(); toggleMenu();}} className="cursor-pointer text-white">
                Login
              </div>
            )}
          </m.div>
        )}
      </AnimatePresence>

      {isSignInModalVisible && <SignIn onClose={toggleSignInModal} />}
    </m.header>
  );
};

export default LandingNav;
