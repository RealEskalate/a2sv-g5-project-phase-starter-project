import PreferencesForm from "@/app/components/Forms/PreferencesForm";
import Navigation from "@/app/components/Settings/Navigation";
import React from "react";

const PreferencesPage: React.FC = () => {
  return (
  <div className="w-full lg:max-w-[1000px] mt-2 flex flex-col justify-center items-center ">
      <div className="bg-white  dark:bg-[#232328] rounded-xl min-h-screen">

        <Navigation />
        <PreferencesForm />
      </div>
    </div>
  );
};

export default PreferencesPage;
