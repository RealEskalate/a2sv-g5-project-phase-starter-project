import SecurityForm from "@/app/components/Forms/SecurityForm";
import Navigation from "@/app/components/Settings/Navigation";
import React from "react";

const SecurityPage: React.FC = () => {
  return (
<div className="w-full lg:max-w-[1000px] mt-2 flex flex-col justify-center ">
      <div className="bg-white  dark:bg-[#232328] rounded-xl min-h-screen">
        <Navigation />
        <SecurityForm />
      </div>
    </div>
  );
};

export default SecurityPage;
