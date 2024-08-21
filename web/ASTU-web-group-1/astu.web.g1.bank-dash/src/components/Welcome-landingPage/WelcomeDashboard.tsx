import Image from "next/image";
import Link from "next/link";
import React from "react";

const WelcomeDashboard = () => {
  return (
    <>
      <div className="flex flex-col w-full lg:w-3/4 p-2 md:p-10 space-y-5 md:space-y-10 relative">
        <Image
          src="/assets/images/coin.jpg"
          alt="Bank Logo"
          className="object-cover z-15 opacity-30 rounded-lg"
          fill
        />
        <p className="z-10 text-deepNavy font-Inter text-5xl md:text-6xl font-extrabold text-center rounded-3xl p-2">
          Welcome to Your Trusted Financial Partner
        </p>
        <p className="z-10 text-gray-700 font-Inter text-sm">
          At Our, we understand that your financial journey is unique. That’s
          why we offer a wide range of services tailored to meet your needs.
          From secure savings and checking accounts to personalized loan options
          and investment advice, our goal is to help you achieve your financial
          aspirations with ease and confidence.
        </p>
        <div className="flex justify-center items-center">
          <Link
            href="/api/auth/signup"
            className="z-10 font-extrabold w-fit px-3 py-1 rounded-2xl text-deepNavy text-2xl hover:bg-green-200 border hover:shadow-[10px_10px_20px_rgba(135,206,235,0.3)]"
          >
            GET-START &rarr;
          </Link>
        </div>
      </div>
    </>
  );
};

export default WelcomeDashboard;
