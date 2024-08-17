import React from "react";
import ServiceList from "./serviceComponenet/ServiceList";
import BenefitComp from "./serviceComponenet/BenefitComp";

const Services = () => {
  return (
    <div className="p-4 flex  flex-col border-2 bg-gray-200 w-full h-full gap-5 lg:p-8">
      <div className=" flex gap-3 justify-start w-[100vw] md:w-full overflow-x-scroll scrollbar-hidden md:overflow-hidden lg:justify-between">
        <BenefitComp />
        <BenefitComp />
        <BenefitComp />
      </div>

      <div className="flex flex-col gap-[7px]">
        <h1 className="font-semibold text-lg text-[#343C6A]">Bank Services List</h1>
        <ServiceList />
        <ServiceList />
        <ServiceList />
        <ServiceList />
        <ServiceList />
        <ServiceList />
      </div>
    </div>
  );
};

export default Services;
