import React from "react";
import services from "../../constants/services_constants/services.json"; // Assuming services.json contains an array of ServiceData
import Rectangle from "./rectangle";

interface ServiceData {
  id: string;
  name: string;
  details: string;
  numberOfUsers: number;
  status: string;
  type: string;
  icon: string;
}

const colors = [
  "255,224,235,1",
  "255,245,217,1",
  "255,224,235,1",
  "231,237,255,1",
  "220,250,248,1",
  "255,224,235,1",
];
export default function Services() {
  return (
    <div className="">
      {services.map((service: ServiceData, index: number) => (
        <div
          key={service.id}
          className="flex my-5 space-x-5 bg-white rounded-3xl p-4 w-full"
          style={{ backgroundColor: colors[index] }}
        >
          <Rectangle backgroundColor={colors[index]} src={service.icon} />
          <div className="flex flex-1 items-center justify-between">
            <div className="flex flex-col items-start">
              <p className="text-[rgba(35,35,35,1)] text-sm font-medium">
                {service.name}
              </p>
              <p className="text-[rgba(113,142,191,1)] text-xs font-normal">
                {service.details}
              </p>
            </div>
            <div className="hidden md:flex flex-col">
              <p className="text-[rgba(35,35,35,1)] text-sm font-medium">
                {service.status}
              </p>
              <p className="text-[rgba(113,142,191,1)] text-xs font-normal">
                {service.status}
              </p>
            </div>
            <div className="hidden md:flex flex-col">
              <p className="text-[rgba(35,35,35,1)] text-sm font-medium">
                {service.status}
              </p>
              <p className="text-[rgba(113,142,191,1)] text-xs font-normal">
                {service.status}
              </p>
            </div>
            <button className="text-xs font-medium text-[rgba(24,20,243,1)] py-1 px-3 w-fit rounded-full m-2  hover:text-[rgba(24,20,243,1)] md:text-[rgba(113,142,191,1)] md:border md:border-[rgba(113,142,191,1)] md:hover:border-[rgba(24,20,243,1)]">
              View Details
            </button>
          </div>
        </div>
      ))}
    </div>
  );
}
