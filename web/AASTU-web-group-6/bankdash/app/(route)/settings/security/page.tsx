import SecurityForm from "@/app/components/Forms/SecurityForm";
import Navigation from "@/app/components/Settings/Navigation";
import React from "react";

const SecurityPage: React.FC = () => {
  return (

    <div className="w-full mt-2 px-10 py-7 flex flex-col ">
      <div className="bg-white rounded-xl">
        <Navigation />
        <SecurityForm />
      </div>
    </div>

      
  );
};

export default SecurityPage;
