

import React from "react";
import ServiceProvided from "@/components/ServiceProvided";
import BankservicesList from "@/components/BankservicesList";

const Services: React.FC = () => {
  return (
    <div className="mx-auto max-w-sm lg:ml-80 sm:max-w-[1110px]">
      <div className="pt-5">
        <ServiceProvided />
      </div>
      <div className="py-10">
        <BankservicesList />
      </div>
    </div>
  );
};

export default Services;
