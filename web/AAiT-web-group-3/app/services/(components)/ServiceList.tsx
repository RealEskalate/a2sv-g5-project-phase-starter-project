import React from "react";

import { IconType } from "react-icons";

interface ServiceListProps {
  icon: IconType;
}
const ServiceList = ({ icon: Icon }: ServiceListProps) => {
  return (
    <div className="flex justify-between bg-white items-center p-7  rounded-[20px]">
      <div className="flex items-center gap-3">
        <Icon className="w-[60px] h-[60px]" />
        <div>
          <p className="service_list_title">Business Loan</p>
          <p className="service_list_subtitle">It is a long established</p>
        </div>
      </div>
      <div>
        <p className="service_list_title">lorem</p>
        <p className="service_list_subtitle">publishing something</p>
      </div>
      <div>
        <p className="service_list_title">lorem</p>
        <p className="service_list_subtitle">publishing something</p>
      </div>
      <div>
        <p className="service_list_title">lorem</p>
        <p className="service_list_subtitle">publishing something</p>
      </div>
      <button className="service_list_subtitle border border-[#718EBF] rounded-[50px] px-5 py-2">View Details</button>
    </div>
  );
};

export default ServiceList;
