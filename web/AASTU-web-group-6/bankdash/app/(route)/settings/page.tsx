"use client";
import React, { useState } from "react";
import EditProfileForm from "@/app/components/Forms/EditProfileForm";
import PreferencesForm from "@/app/components/Forms/PreferencesForm";
import SecurityForm from "@/app/components/Forms/SecurityForm";
import Navigation from "@/app/components/Settings/Navigation";

const SettingsPage: React.FC = () => {
  // State to manage the current active section
  const [activeSection, setActiveSection] = useState("editprofile");

  // Render the form based on the active section
  const renderForm = () => {
    switch (activeSection) {
      case "editprofile":
        return <EditProfileForm />;
      case "preference":
        return <PreferencesForm />;
      case "security":
        return <SecurityForm />;
      default:
        return <EditProfileForm />;
    }
  };

  return (
    <div className="w-full flex justify-center mt-15 md:max-h-[717px]">
      <div className="w-full max-w-[1110px] bg-white px-3 mt-2 flex flex-col justify-center items-center rounded-xl">
        {/* Navigation Component */}
        <Navigation
          activeSection={activeSection}
          setActiveSection={setActiveSection}
        />

        {/* Render the appropriate form */}
        <div className="w-full mt-4">{renderForm()}</div>
      </div>
    </div>
  );
};

export default SettingsPage;
