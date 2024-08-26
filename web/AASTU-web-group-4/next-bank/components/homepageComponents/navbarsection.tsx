'use client'
import React, { useState, useEffect } from "react";
import Link from "next/link";
import Image from "next/image";

const Navbar = () => {
  const [isScrolled, setIsScrolled] = useState(false);

  useEffect(() => {
    const handleScroll = () => {
      if (window.scrollY > 0) {
        setIsScrolled(true);
      } else {
        setIsScrolled(false);
      }
    };

    window.addEventListener("scroll", handleScroll);

    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);

  return (
    <nav
      className={`fixed top-0 w-full h-20 z-50 transition-all duration-300 ${
        isScrolled ? "backdrop-blur-md shadow-md text-blue-500" : "bg-white"
      }`}
    >
      <div className="container flex justify-between items-center p-4">
        {/* Logo */}
        <div className="flex items-center">
          <Image
            src="https://cdn.freelogovectors.net/wp-content/uploads/2024/03/chase_logo-freelogovectors.net_.png"
            alt="Next Bank Logo"
            width={60} // Adjust width as needed
            height={40} // Adjust height as needed
            className="h-auto w-auto"
          />
          <span className="ml-3 text-xl font-bold">Next Bank</span>
        </div>

        {/* Menu */}
        <div className="space-x-8">
          <Link href="/signin" className="font-bold">
            Home
          </Link>
          <Link href="/signin" className="font-bold">
            Services
          </Link>
          <Link href="/signin" className="font-bold">
            Contact
          </Link>
          <Link href="/signin" className="font-bold">
            Sign In
          </Link>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
