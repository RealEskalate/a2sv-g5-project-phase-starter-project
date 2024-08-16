"use client";
import React, { useState } from "react";

const ProfileSecurity = () => {
  const [isToggled, setIsToggled] = useState(false);

  const handleToggle = () => {
    setIsToggled(!isToggled);
  };
  return (
    <div className="bg-white p-6 w-full flex flex-col">
      <div className="w-full mb-5">
        <p className="text-[#333B69] font-medium sm:text-[14px] md:text-[17px]">
          Two-factor Authentication
        </p>

        <div className="flex items-center mt-3 mb-5">
          {/* toggle button */}
          <div
            onClick={handleToggle}
            className={`mr-5 relative w-14 h-8 flex items-center rounded-full p-1 cursor-pointer ${
              isToggled ? "bg-[#16DBCC]" : "bg-gray-300"
            }`}
          >
            <div
              className={`bg-white w-6 h-6 rounded-full shadow-md transform duration-300 ease-in-out ${
                isToggled ? "translate-x-6" : ""
              }`}
            ></div>
          </div>
          {/* toggle button */}

          <span className="sm:text-[13px] md:text-[16px] text-[#232323] font-normal">
            Enable or disable two factor authentication
          </span>
        </div>

        <p className="text-[#333B69] font-medium sm:text-[14px] md:text-[17px] mb-3">
          Change Password
        </p>

        <div className="flex flex-col w-full sm:w-[80%] lg:w-[40%]">
          <p className="text-[#232323] sm:text-[13px] md:text-[16px] font-normal mb-2">
            Current Password
          </p>

          <input
            type="password"
            placeholder="************"
            className="border-2 border-[#DFEAF2] text-sm p-3 rounded-xl"
          />
        </div>

        <div className="flex flex-col mt-4 w-full sm:w-[80%] lg:w-[40%]">
          <p className="text-[#232323] sm:text-[13px] md:text-[16px] font-normal mb-2">
            New Password
          </p>

          <input
            type="password"
            placeholder="************"
            className="border-2 border-[#DFEAF2] text-sm p-3 rounded-xl"
          />
        </div>
      </div>

      <button className="sm:w-[150px] w-full bg-[#1814F3] text-md text-white font-medium flex justify-center rounded-lg py-3 sm:mt-auto sm:self-end">
        Save
      </button>
    </div>
  );
};

export default ProfileSecurity;
