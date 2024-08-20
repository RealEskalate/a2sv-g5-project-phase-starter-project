import PreferencesForm from "@/app/components/Forms/PreferencesForm";
import Navigation from "@/app/components/Settings/Navigation";
import React from "react";

const PreferencesPage: React.FC = () => {
  return (
    <div className="w-full mt-2 px-10 py-7 flex flex-col ">
      <div className="bg-white rounded-xl">
        <Navigation />
        <PreferencesForm />
      </div>
    </div>
  );
};

export default PreferencesPage;
