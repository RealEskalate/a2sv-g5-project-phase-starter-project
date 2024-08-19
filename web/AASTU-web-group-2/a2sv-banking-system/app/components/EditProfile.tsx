"use client";
import React from "react";
import Image from "next/image";
import { FaPencilAlt } from "react-icons/fa";
import { useState } from "react";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import { FaCaretDown } from "react-icons/fa";

const EditProfile = () => {
  const [startDate, setStartDate] = useState<Date | null>(null);

  return (
    <>
      <div className="flex flex-col items-center py-10 w-full">
        <div className="relative">
          <Image
            src="/ProfilePicture.png"
            alt="Profile Picture"
            width={170}
            height={170}
            className="rounded-full"
          />
          <span className="absolute bottom-5 right-0 bg-[#1814F3] rounded-full w-10 h-10 flex items-center justify-center text-white">
            <FaPencilAlt></FaPencilAlt>
          </span>
        </div>
      </div>

      <div>
        <form action="">
          <div className="flex flex-col gap-3">
            <div className="flex flex-col md:flex md:flex-row">
              {/* Name Goes In Here */}
              <div className="flex flex-col gap-3 px-6 py-3 md:w-[48%]">
                <label
                  htmlFor="name"
                  className="text-[#232323] font-semibold px-1"
                >
                  Your Name
                </label>
                <input
                  type="text"
                  name="name"
                  id="name"
                  className="border border-[#DFEAF2] focus:outline-[#DFEAF2] focus:border-[#DFEAF2] rounded-xl py-3 px-6 placeholder:text-[#718EBF]"
                  placeholder="Rebuma Tadele"
                />
              </div>

              {/* UserName Goes In Here */}
              <div className="flex flex-col gap-3 px-6 py-3 md:w-[48%]">
                <label
                  htmlFor="username"
                  className="text-[#232323] font-semibold px-1"
                >
                  Username
                </label>
                <input
                  type="text"
                  name="username"
                  id="username"
                  className="border border-[#DFEAF2] focus:outline-[#DFEAF2] focus:border-[#DFEAF2]  rounded-xl py-3 px-6 placeholder:text-[#718EBF]"
                  placeholder="rebuma"
                />
              </div>
            </div>

            <div className="flex flex-col md:flex md:flex-row">
              {/* Email Goes In Here */}

              <div className="flex flex-col gap-3 px-6 py-3 md:w-[48%]">
                <label
                  htmlFor="email"
                  className="text-[#232323] font-semibold px-1"
                >
                  Email
                </label>
                <input
                  type="email"
                  name="email"
                  id="email"
                  className="border border-[#DFEAF2] focus:outline-[#DFEAF2] focus:border-[#DFEAF2] rounded-xl py-3 px-6 placeholder:text-[#718EBF]"
                  placeholder="john@example.com"
                />
              </div>
              {/* Password Goes In Here */}

              <div className="flex flex-col gap-3 px-6 py-3 md:w-[48%]">
                <label
                  htmlFor="password"
                  className="text-[#232323] font-semibold px-1"
                >
                  Password
                </label>
                <input
                  type="password"
                  name="password"
                  id="password"
                  className="border border-[#DFEAF2] focus:outline-[#DFEAF2] focus:border-[#DFEAF2] rounded-xl py-3 px-6 placeholder:text-[#718EBF]"
                  placeholder="********"
                />
              </div>
            </div>

            <div className="flex flex-col md:flex md:flex-row">
              {/* Date Of Birth Goes In Here */}
              <div className="flex flex-col gap-3 px-6 py-3 md:w-[48%]">
                <label
                  htmlFor="datePicker"
                  className="text-[#232323] font-semibold px-1"
                >
                  Date Of Birth
                </label>
                <div className="relative w-full">
                  <DatePicker
                    selected={startDate}
                    onChange={(date: Date | null) => setStartDate(date)}
                    placeholderText="Date Of Birth"
                    className="w-full border border-[#DFEAF2] focus:outline-[#DFEAF2] focus:border-[#DFEAF2] rounded-xl py-3 px-6 placeholder:text-[#718EBF] bg-white"
                    dateFormat="MMMM d, yyyy"
                    id="datePicker"
                  />
                  <FaCaretDown className="absolute right-3 top-1/2 transform -translate-y-1/2 text-[#718EBF]" />
                </div>
              </div>
              {/* Present Address Goes In Here */}
              <div className="flex flex-col gap-3 px-6 py-3 md:w-[48%]">
                <label
                  htmlFor="presentAddress"
                  className="text-[#232323] font-semibold px-1"
                >
                  Present Address
                </label>
                <input
                  type="text"
                  name="presentAddress"
                  id="presentAddress"
                  className="border border-[#DFEAF2] focus:outline-[#DFEAF2] focus:border-[#DFEAF2] rounded-xl py-3 px-6 placeholder:text-[#718EBF]"
                  placeholder="Addis Ababa Ethiopia"
                />
              </div>
            </div>

            <div className="flex flex-col md:flex md:flex-row">
              {/* Permanent Address Goes In Here */}
              <div className="flex flex-col gap-3 px-6 py-3 md:w-[48%]">
                <label
                  htmlFor="permanentAddress"
                  className="text-[#232323] font-semibold px-1"
                >
                  Permanent Address
                </label>
                <input
                  type="text"
                  name="permanentAddress"
                  id="permanentAddress"
                  className="border border-[#DFEAF2] focus:outline-[#DFEAF2] focus:border-[#DFEAF2] rounded-xl py-3 px-6 placeholder:text-[#718EBF]"
                  placeholder="Addis Ababa, Ethiopia"
                />
              </div>

              {/* City Goes In Here */}
              <div className="flex flex-col gap-3 px-6 py-3 md:w-[48%]">
                <label
                  htmlFor="city"
                  className="text-[#232323] font-semibold px-1"
                >
                  City
                </label>
                <input
                  type="text"
                  name="city"
                  id="city"
                  className="border border-[#DFEAF2] focus:outline-[#DFEAF2] focus:border-[#DFEAF2] rounded-xl py-3 px-6 placeholder:text-[#718EBF]"
                  placeholder="Addis Ababa, Ethiopia"
                />
              </div>
            </div>

            <div className="flex flex-col md:flex md:flex-row">
              {/* Postal Code Goes In Here */}
              <div className="flex flex-col gap-3 px-6 py-3 md:w-[48%]">
                <label
                  htmlFor="postal"
                  className="text-[#232323] font-semibold px-1"
                >
                  Postal Code
                </label>
                <input
                  type="text"
                  name="name"
                  id="name"
                  className="border border-[#DFEAF2] focus:outline-[#DFEAF2] focus:border-[#DFEAF2] rounded-xl py-3 px-6 placeholder:text-[#718EBF]"
                  placeholder="10000"
                />
              </div>

              {/* Country Goes in here */}
              <div className="flex flex-col gap-3 px-6 py-3 md:w-[48%]">
                <label
                  htmlFor="country"
                  className="text-[#232323] font-semibold px-1"
                >
                  Country
                </label>
                <input
                  type="text"
                  name="name"
                  id="name"
                  className="border border-[#DFEAF2] focus:outline-[#DFEAF2] focus:border-[#DFEAF2] rounded-xl py-3 px-6 placeholder:text-[#718EBF]"
                  placeholder="Ethiopia"
                />
              </div>
            </div>

            <div className="flex flex-col gap-3 md:px-16 px-6 py-3 md:items-end">
              <button
                type="submit" 
                className="bg-[#1814F3] border border-[#1814F3] rounded-xl text-white px-6 py-3 font-semibold text-xl md:w-1/4"
              >
                Save
              </button>
            </div>
          </div>
        </form>
      </div>
    </>
  );
};

export default EditProfile;
