"use client";
import React, { useState } from "react";
import ToggleButton from "../Button/ToggleButton";
import AuthService from "@/app/Services/api/authService";

const SecurityForm = () => {
  const [currentPassword, setCurrentPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [twoFactorAuth, setTwoFactorAuth] = useState(true);

  const handleToggle = (checked: boolean) => {
    setTwoFactorAuth(checked);
    console.log(
      "Two-factor Authentication is now",
      checked ? "Enabled" : "Disabled"
    );
  };

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();

    try {
      const response = await AuthService.changePassword(
        {
          password: currentPassword,
          newPassword: newPassword,
        },
        "accessToken"
      );
      console.log("Password changed:", response);
    } catch (error) {
      console.error("Error changing password:", error);
    }
  };

  return (
    <form className="w-full mt-3 space-y-6 px-3 py-4" onSubmit={handleSubmit}>
      <div className="mt-6">
        <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
          Two-factor Authentication
        </label>
        <div className="flex items-center mt-4">
          <ToggleButton
            onToggle={handleToggle}
            initialChecked={twoFactorAuth}
          />
          <span className="ml-3 text-sm text-gray-700 dark:text-gray-300">
            Enable or disable two-factor authentication
          </span>
        </div>
      </div>

      <div className="mt-6">
        <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
          Change Password
        </label>
        <div className="mt-4 space-y-4">
          <div className="xxs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
              Current Password
            </label>
            <input
              type="password"
              value={currentPassword}
              onChange={(e) => setCurrentPassword(e.target.value)}
              className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
            />
          </div>

          <div className="xxs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
              New Password
            </label>
            <input
              type="password"
              value={newPassword}
              onChange={(e) => setNewPassword(e.target.value)}
              className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
            />
          </div>
        </div>
      </div>

      <div className="flex justify-end mt-3 w-full ">
        <button
          className="px-4 py-2 bg-[#1814F3] text-white rounded-md hover:bg-[#0702db] transition-all duration-300"
          type="submit"
        >
          Save Changes
        </button>
      </div>
    </form>
  );
};

export default SecurityForm;
