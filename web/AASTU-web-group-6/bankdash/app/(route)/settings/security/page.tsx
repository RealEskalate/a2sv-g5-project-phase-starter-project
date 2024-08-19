import SecurityForm from "@/app/components/Forms/SecurityForm";
import Navigation from "@/app/components/Settings/Navigation";
import React from "react";

const SecurityPage: React.FC = () => {
  return (

    <div className="w-full mt-2 px-4 flex flex-col bg-white">
      <Navigation/>
      <SecurityForm/>
    </div>

      
  );
};

export default SecurityPage;
