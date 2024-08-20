"use client";
import React, { useState } from "react";
import ToggleButton from "../Button/ToggleButton";
import authService from "../../services/authService"; // Make sure to import authService

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
      const response = await authService.changePassword({
        password: currentPassword,
        newPassword: newPassword,
      });
      console.log("Password changed:", response);
    } catch (error) {
      console.error("Error changing password:", error);
    }
  };

  return (
    <form className="mt-8 space-y-6 px-3 py-4" onSubmit={handleSubmit}>
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
          <div>
            <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
              Current Password
            </label>
            <input
              type="password"
              value={currentPassword}
              onChange={(e) => setCurrentPassword(e.target.value)}
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
              New Password
            </label>
            <input
              type="password"
              value={newPassword}
              onChange={(e) => setNewPassword(e.target.value)}
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
            />
          </div>
        </div>
      </div>

      <div className="flex lg:justify-end mt-3 xs:w-full sm:justify-center">
        <button
          type="submit"
          className="w-[192px] py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700"
        >
          Save
        </button>
      </div>
    </form>
  );
};

export default SecurityForm;
