"use client";
import { useState } from "react";
import { FaBars } from "react-icons/fa";
import Link from "next/link";

const Nav = () => {
  const [isOpen, setIsOpen] = useState(false);

  const toggleMenu = () => {
    setIsOpen(!isOpen);
  };

  return (
    <div>
      <header className="flex justify-between items-center p-4 bg-gradient-to-r from-blue-500 to-purple-600 shadow-lg">
        <div className="text-white font-bold text-xl">A2SV Blogger</div>
        <div className="md:hidden">
          <FaBars
            className="text-white text-2xl cursor-pointer"
            onClick={toggleMenu}
          />
        </div>
        <div className={`md:flex space-x-8 text-white font-medium hidden `}>
          <ul className="flex flex-col md:flex-row md:space-x-8">
            <Link
              href={"/"}
              className="hover:text-gray-300 cursor-pointer"
              onClick={toggleMenu}
            >
              Home
            </Link>
            <Link
              href={"/Users"}
              className="hover:text-gray-300 cursor-pointer"
              onClick={toggleMenu}
            >
              User
            </Link>
            <Link
              href={"/auth/login"}
              className="hover:text-gray-300 cursor-pointer"
              onClick={toggleMenu}
            >
              Login
            </Link>
            <Link
              href={"/auth/signup"}
              className="hover:text-gray-300 cursor-pointer"
              onClick={toggleMenu}
            >
              Register
            </Link>
          </ul>
        </div>
      </header>

      <div>
        <div
          className={`md:hidden ${
            isOpen ? "block" : "hidden"
          } bg-gradient-to-r from-blue-500 to-purple-600 text-white`}
        >
          <ul className="flex flex-col space-y-4 p-4">
            <Link href={"/"} className="hover:text-gray-300 cursor-pointer">
              Home
            </Link>
            <Link
              href={"/Users"}
              className="hover:text-gray-300 cursor-pointer"
            >
              User
            </Link>
            <Link
              href={"/auth/login"}
              className="hover:text-gray-300 cursor-pointer"
            >
              Login
            </Link>
            <Link
              href={"/auth/signup"}
              className="hover:text-gray-300 cursor-pointer"
            >
              Register
            </Link>
          </ul>
        </div>
      </div>
    </div>
  );
};

export default Nav;
