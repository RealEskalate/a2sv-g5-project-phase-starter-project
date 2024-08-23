import React from "react";

import { IconType } from "react-icons";

interface ServiceListProps {
  icon: IconType;
  icon_bg?: string;
  icon_color?: string;
}
const ServiceList = ({ icon: Icon,icon_bg,icon_color }: ServiceListProps) => {
  return (
    <div className="flex justify-between bg-white items-center  p-3 md:p-7  rounded-[20px] mb-6 ">
      <div className="flex items-center gap-3">
        <div className="w-[45px] h-[45px] rounded-xl flex items-center justify-center " style={{backgroundColor:icon_bg}}>
        <Icon className=" w-[20px] h-[20px] "  style={{color:icon_color}}/>

        </div>
        <div>
          <p className="service_list_title ">Business Loan</p>
          <p className="service_list_subtitle text-[#718EBF] ">
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
