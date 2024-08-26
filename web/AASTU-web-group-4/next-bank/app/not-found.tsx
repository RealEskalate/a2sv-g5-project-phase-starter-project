import React from "react";
import { colors } from "@/constants/index";
import Link from "next/link";
import Image from "next/image";

export default function NotFound() {
  return (
    <div className="w-[100%] h-screen bg-gray-100 dark:bg-dark grid grid-cols-1 lg:grid-cols-2 items-center">
      <div className="flex flex-col gap-5 p-10 lg:p-20">
        <div className="flex flex-col gap-2">
          <h2
            className={`text-[32px] font-bold ${colors.textblack} dark:text-blue-500`}
          >
            Ooops...
          </h2>
          <p
            className={`text-[20px] font-semibold ${colors.textblack} dark:text-blue-500`}
          >
            Page not found
          </p>
          <p
            className={`text-[14px] font-light ${colors.textblack} dark:text-blue-500`}
          >
            to access our services please go back to the dashboard page
          </p>
        </div>
        <div className="flex">
          <Link href="/" passHref>
            <button
              className={`px-6 py-3 rounded-lg font-semibold text-white bg-red-500 hover:bg-red-600 transition-colors`}
            >
              Go Back →
            </button>
          </Link>
        </div>
      </div>
      <div className="hidden lg:flex justify-center lg:justify-end p-10 lg:p-20">
        <Image
          src="/Images/404.png"
          width={500}
          height={500}
          alt="404 page not found"
        />
      </div>
    </div>
  );
}
