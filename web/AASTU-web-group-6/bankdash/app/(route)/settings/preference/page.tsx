import PreferencesForm from "@/app/components/Forms/PreferencesForm";
import Navigation from "@/app/components/Settings/Navigation";
import React from "react";

const PreferencesPage: React.FC = () => {
  return (

    <div className="w-full mt-2 px-4 flex flex-col bg-white">
      <Navigation/>
      <PreferencesForm/>
    </div>
  );
};

export default PreferencesPage;

