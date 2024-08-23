"use client";
import React from "react";
import { useState } from "react";
import EditProfile from "@/components/settings/EditProfile/page";
import PrefPage from "@/components/settings/Preferences/page";
import SecuritySettings from "@/components/settings/Security/page";
const SettingPage = () => {
  const [activeButton, setActiveButton] = useState("edit");
  const handleOnClick = (button: string) => {
    setActiveButton(button);
  };
  return (
    <div className="py-8 px-10">
      <div className="flex flex-col rounded-[25px] w-fill  pt-[37px] px-[30px]  bg-white">
        <div className="flex flex-row border-b  w-fill h-[30px] text-[#718EBF] gap-12">
          <button
            className={`w-[114px] h-[30px]  items-center ${
              activeButton === "edit"
                ? "border-b-2  text-[#1814F3] border-[#1814F3]"
                : ""
            }  `}
            onClick={() => handleOnClick("edit")}
          >
            Edit Profile
          </button>
          <button
            className={`w-[114px] h-[30px]  items-center ${
              activeButton === "preferences"
                ? "border-b-2  text-[#1814F3] border-[#1814F3]"
                : ""
            }  `}
            onClick={() => handleOnClick("preferences")}
          >
            Preferences
          </button>
          <button
            className={`w-[114px] h-[30px]  items-center ${
              activeButton === "Security"
                ? "border-b-2  text-[#1814F3] border-[#1814F3]"
                : ""
            }  `}
            onClick={() => handleOnClick("Security")}
          >
            Security
          </button>
        </div>
        {activeButton === "edit" && <EditProfile />}
        {activeButton === "preferences" && <PrefPage />}
        {activeButton === "Security" && <SecuritySettings />}
      </div>
    </div>
  );
};

export default SettingPage;
