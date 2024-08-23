

import React from "react";
import ServiceProvided from "@/components/ServiceProvided";
import BankservicesList from "@/components/BankservicesList";

const Services: React.FC = () => {
  return (
    <div className="mx-auto  max-w-sm lg:ml-72 sm:max-w-[1110px] md:pr-5">
      <div className="pt-5 pb-2 overflow-x-auto max-w-screen">
        <ServiceProvided />
      </div>
      <div className="py-12">
        <BankservicesList />
      </div>
    </div>
  );
};

export default Services;
