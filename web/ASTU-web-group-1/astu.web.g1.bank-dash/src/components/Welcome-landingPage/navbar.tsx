import Image from "next/image";
import Link from "next/link";
import React from "react";

const Navbar = () => {
  return (
    <>
      <div
        className="flex w-full z-20 bg-slate-50 px-10 md:px-20 justify-between items-center fixed space-x-3 lg:px-14 border-b shadow-lg border-gray-100"
        id="dashboard"
      >
        <div className="flex flex-col md:flex-row w-fit md:w-fit space-x-2 items-center py-2">
          <Image
            src="/assets/images/logo.jpg"
            alt=""
            width={100}
            height={100}
            className="object-cover rounded-full h-14 w-14 md:w-16 md:h-16"
          />
          <h1 className="font-Inter hidden sm:block font-extrabold text-xl md:text-2xl text-deepNavy">
            BANK DASHBOARD
          </h1>
        </div>
        <div className="flex md:space-x-5 space-x-6 items-center justify-end ">
          <Link
            href="#dashboard"
            className="px-2 text-deepNavy text-center font-semibold md:text-lg hover:text-indigo-600 cursor-pointer hidden md:block"
          >
            Dashboard
          </Link>
          <Link
            href="#why"
            className="text-deepNavy text-center font-semibold md:text-lg hover:text-indigo-600 cursor-pointer md:block"
          >
            Why
          </Link>
          <Link
            href="#contacts"
            className="text-deepNavy text-center font-semibold md:text-lg hover:text-indigo-600 cursor-pointer md:block"
          >
            Contacts
          </Link>
          <Link
            href="/api/auth/signin"
            className="text-deepNavy text-center font-semibold md:text-lg hover:text-indigo-600 cursor-pointer md:block"
          >
            SignIn
          </Link>
        </div>
      </div>
    </>
  );
};

export default Navbar;
