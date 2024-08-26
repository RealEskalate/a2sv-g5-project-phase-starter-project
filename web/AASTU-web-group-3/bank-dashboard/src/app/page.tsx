"use client";
import React from "react";
import Image from "next/image";
import { logo } from "@/../../public/Icons";
import Link from "next/link";
import mobileImage from "@/../../public/Images/landingmobileimage.png";
import personWithCreditCard from "@/../../public/Images/person-paying-with-its-credit-card.png";

const HomePage: React.FC = () => {
  return (
    <>
      <div className="flex flex-col w-full min-h-full bg-[#CBD6CC] justify-center">
        <div className="p-4 w-[95%]">
          <div className="flex items-center justify-between">
            <div className="flex items-center">
              <Image src={logo} alt="Logo" width={36} height={36} />
              <div className="text-[#343C6A] dark:text-darkText pl-2 md:text-xl md:pl-1 lg:pl-2 lg:text-2xl text-base xl:text-4xl md:text-[21px] font-[800] font-mont">
                BankDash.
              </div>
            </div>
            <div className="flex items-center gap-3">
              <Link
                href="/auth/signin"
                className="text-lg border border-[#343C6A] hover:bg-[#343C6A] hover:text-white px-4 py-2 rounded-3xl text-center"
              >
                Login
              </Link>
              <Link
                href="/auth/signup"
                className="text-lg border hover:border-[#343C6A] bg-[#343C6A] hover:bg-[#CBD6CC] text-white px-4 py-2 rounded-3xl text-center hover:text-black"
              >
                signUp
              </Link>
            </div>
          </div>
          <div className="flex flex-col-reverse md:flex-row my-16 w-full  ">
            <div className="w-full md:w-2/5">
              <div className="font-bold text-7xl py-6">
                Easy way to manage your money
              </div>
              <div className="py-6 w-3/4">
                A new way to make the payments easy reliable and secure. You can
                manage all your transactions from your mobile phone.
              </div>
              <Link
                href="/auth/signup"
                className="text-lg border hover:border-[#343C6A] bg-[#343C6A] hover:bg-[#CBD6CC] text-white px-4 py-2 rounded-3xl text-center hover:text-black"
              >
                Get Started
              </Link>
            </div>

            <div className="w-full md:w-3/5 flex items-center justify-center">
              <Image
                src={mobileImage}
                width={200}
                height={200}
                alt="sample moble of project"
                className="w-full md:w-3/5"
              />
            </div>
          </div>
        </div>
        <div className=" md:mt-16  bg-[#e1dddd] flex flex-col justify-center items-center  dark:bg-darkComponent md:rounded-t-full rounded-t-3xl p-6 pt-12">
          <div className="md:w-4/5 flex flex-wrap justify-between items-center gap-5">
            <div className="w-[45%] md:w-1/5 flex flex-col items-center justify-center text-center gap-1 bg-white shadow-xl hover:scale-105  h-24 ">
              <div className="font-bold text-2xl  ">dark Mode</div>
              <div className="font-thin text-sm">customize your theme</div>
            </div>
            <div className="w-[45%] md:w-1/5 flex flex-col items-center justify-center text-center gap-1 bg-white shadow-xl hover:scale-105 h-24 ">
              <div className="font-bold text-2xl  ">100% safe</div>
              <div className="font-thin text-sm">your money is safe</div>
            </div>
            <div className="w-[45%] md:w-1/5 flex flex-col items-center justify-center text-center gap-1 bg-white shadow-xl hover:scale-105 h-24 ">
              <div className="font-bold text-2xl  ">Quick Send</div>
              <div className="font-thin text-sm">transfer money in 1 click</div>
            </div>
            <div className="w-[45%] md:w-1/5 flex flex-col items-center text-center gap-1 justify-center bg-white shadow-xl hover:scale-105 h-24 ">
              <div className="font-bold text-2xl  ">Notification</div>
              <div className="font-thin text-sm">
                stay updated about your transactions
              </div>
            </div>
          </div>

          {/* <div className="w-full flex items-center justify-center">
              <div className="w-full md:w-3/5 flex items-center justify-center">
              <Image
                src={personWithCreditCard}
                width={300}
                height={300}
                alt="sample moble of project"
              />
            </div>
          </div> */}
        </div>
      </div>
    </>
  );
};

export default HomePage;
