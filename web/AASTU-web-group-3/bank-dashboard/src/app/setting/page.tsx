"use client";
import React, { useState } from "react";
import EditProfile from "../components/setting/EditProfile";
import Preference from "../components/setting/Preference";
import Security from "../components/setting/Security";


const HomePage: React.FC = () => {
  const [activeTab, setActiveTab] = useState("Edit Profile");

  const renderContent = () => {
    switch (activeTab) {
      case "Edit Profile":
        return <EditProfile />;
      case "Preference":
        return <Preference />;
      case "Security":
        return <Security />;
      default:
        return null;
    }
  };

  return (
    <div className="w-11/12 mt-6 h-screen bg-white rounded-3xl">
      <div className="border-[#718EBF] border-b flex justify-between md:justify-start md:gap-4">
        <div
          onClick={() => setActiveTab("Edit Profile")}
          className={`cursor-pointer text-xl py-4 px-3 ${
            activeTab === "Edit Profile"
              ? "text-blue-600 border-blue-600 border-b-4"
              : "text-[#718EBF]"
          }`}
        >
          Edit Profile
        </div>
        <div
          onClick={() => setActiveTab("Preference")}
          className={`cursor-pointer text-xl py-4 px-3 ${
            activeTab === "Preference"
              ? "text-blue-600 border-blue-600 border-b-4"
              : "text-[#718EBF]"
          }`}
        >
          Preference
        </div>
        <div
          onClick={() => setActiveTab("Security")}
          className={`cursor-pointer text-xl py-4 px-3 ${
            activeTab === "Security"
              ? "text-blue-600 border-blue-600 border-b-4"
              : "text-[#718EBF]"
          }`}
        >
          Security
        </div>
      </div>

      <div>{renderContent()}</div>
    </div>
  );
};

export default HomePage;
