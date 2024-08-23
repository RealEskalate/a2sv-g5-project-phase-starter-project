import React from "react";

import { IconType } from "react-icons";

interface ServiceListProps {
  icon: IconType;
}
const ServiceList = ({ icon: Icon }: ServiceListProps) => {
  return (
    <div className="flex justify-between bg-white items-center  p-3 md:p-7  rounded-[20px] mb-6 ">
      <div className="flex items-center gap-3">
        <Icon className=" w-[45px] h-[45px] md:w-[60px] md:h-[60px]" />
        <div>
          <p className="service_list_title whitespace-nowrap">Business Loan</p>
          <p className="service_list_subtitle text-[#718EBF] whitespace-nowrap">
            It is a long established
          </p>
        </div>
      </div>
      <div className="hidden md:block">
        <p className="service_list_title">lorem</p>
        <p className="service_list_subtitle text-[#718EBF]">
          publishing something
        </p>
      </div>
      <div className="hidden md:block">
        <p className="service_list_title">lorem</p>
        <p className="service_list_subtitle text-[#718EBF]">
          publishing something
        </p>
      </div>
      <div className="hidden md:block">
        <p className="service_list_title">lorem</p>
        <p className="service_list_subtitle text-[#718EBF]">
          publishing something
        </p>
      </div>
      <button className="service_list_Button">
        View Details
      </button>
    </div>
  );
};

export default ServiceList;
