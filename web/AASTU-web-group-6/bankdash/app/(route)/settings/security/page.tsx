import SecurityForm from "@/app/components/Forms/SecurityForm";
import Navigation from "@/app/components/Settings/Navigation";
import React from "react";

const SecurityPage: React.FC = () => {
  return (
    // <div className="w-full flex justify-center">
    //   <div className="w-full max-w-[1110px] bg-white mx-4 mt-2 flex flex-col justify-center items-center rounded-xl ">
    <div className="w-full flex justify-center mt-15  md:max-h-[717px] ">
      <div className="w-full max-w-[1110px] bg-white px-3 mt-2 flex flex-col justify-center items-center rounded-xl  ">
        <Navigation />
        <SecurityForm />
      </div>
    </div>
  );
};

export default SecurityPage;
