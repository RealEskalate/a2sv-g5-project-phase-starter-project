"use client";
import React, { useState } from "react";
import { useForm } from "react-hook-form";
import EditProfile from "./EditProfile";
import Preferences from "./Preferences";
import Security from "./Security";

const SettingsPage: React.FC = () => {
  const [activeTab, setActiveTab] = useState("Edit Profile");
  const handleTabClick = (tabName: string) => {
    setActiveTab(tabName);
  };

  return (
    <div className="bg-white w-full rounded-3xl p-4 m-4">
      {/* Selection Buttons */}
      <div className="flex mb-6 container">
        {/* Edit Profile */}
        <button
          onClick={() => handleTabClick("Edit Profile")}
          className={`px-2 py-2 font-medium relative ${
            activeTab === "Edit Profile" ? "text-blue-500" : "text-gray-600"
          }`}
        >
          {activeTab === "Edit Profile" && (
            <span className="absolute bottom-0 left-0 w-full h-[3px] bg-blue-500"></span>
          )}
          Edit Profile
        </button>

        {/* Preferences */}
        <button
          onClick={() => handleTabClick("Preferences")}
          className={`px-4 py-2 font-medium relative ${
            activeTab === "Preferences" ? "text-blue-500" : "text-gray-600"
          }`}
        >
          {activeTab === "Preferences" && (
            <span className="absolute bottom-0 left-0 w-full h-[3px] bg-blue-500"></span>
          )}
          Preferences
        </button>

        {/* Security */}
        <button
          onClick={() => handleTabClick("Security")}
          className={`px-4 py-2 font-medium relative ${
            activeTab === "Security" ? "text-blue-500" : "text-gray-600"
          }`}
        >
          {activeTab === "Security" && (
            <span className="absolute bottom-0 left-0 w-full h-[3px] bg-blue-500"></span>
          )}
          Security
        </button>
      </div>

      {/* Rendered Page */}
      {activeTab === "Edit Profile" && <EditProfile />}
      {activeTab === "Preferences" && <Preferences />}
      {activeTab === "Security" && <Security />}
    </div>
  );
};

export default SettingsPage;
