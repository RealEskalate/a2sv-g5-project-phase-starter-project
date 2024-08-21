"use client";
import React, { useState } from "react";
import Profile from './profile';
import Preference from "./preference";
import Security from "./security";
import { useSession } from "next-auth/react";
import { useSelector } from "react-redux";
import User from '../../type/user'
const activeColor = "text-blue-700";
const disabledColor = "text-slate-400";

function Settings() {
  const user = useSelector((state: { user: User }) => state.user);
  const session = useSession()
  const [enabled, setEnabled] = useState("Edit Profile");

  const renderContent = () => {
    switch (enabled) {
      case 'Edit Profile':
        return <Profile/>;
      case "Preference":
        return <Preference />;
      case "Security":
        return <Security />;
      default:
        return null;
    }
  };

  return (
    <div className="bg-white rounded-2xl pt-4 text-neutral-800 w-[95%] text-sm md:text-base justify-center mt-6 px-6 sm:px-10 pb-5 mx-auto">
      <div className="relative mt-8 md:mt-10">
        <div className="flex gap-10 md:flex-row md:gap-16 lg:gap-20">
          <div
            className={`cursor-pointer ${
              enabled === "Edit Profile"
                ? `${activeColor} custom-underline`
                : disabledColor
            }`}
            onClick={() => setEnabled("Edit Profile")}
          >
            Edit Profile
          </div>
          <div
            className={`cursor-pointer ${
              enabled === "Preference"
                ? `${activeColor} custom-underline`
                : disabledColor
            }`}
            onClick={() => setEnabled("Preference")}
          >
            Preference
          </div>
          <div
            className={`cursor-pointer ${
              enabled === "Security"
                ? `${activeColor} custom-underline`
                : disabledColor
            }`}
            onClick={() => setEnabled("Security")}
          >
            Security
          </div>
        </div>
        <div className="mt-10">{renderContent()}</div>
      </div>
    </div>
  );
}

export default Settings;
