import React from 'react';

interface Service {
  name: string;
  description: string;
}

interface BankServiceListProps {
  logoBgColor?: string;
  logoSvg?: React.ReactNode;
  serviceName?: string;
  serviceDescription?: string;
  additionalServices?: Service[]; // Correctly define the array type
  viewDetailsLink?: string;
}

const BankServiceList: React.FC<BankServiceListProps> = ({
  logoBgColor = "bg-pink-100",
  logoSvg = (
    <svg className="w-8 h-8 text-pink-500" fill="currentColor" viewBox="0 0 24 24">
      {/* Default SVG content */}
    </svg>
  ),
  serviceName = "Business loans",
  serviceDescription = "It is a long established",
  additionalServices = [], // Default to an empty array
  viewDetailsLink = "#",
}) => {
  return (
    <div className="mx-5 my-5 bg-[#FFFFFF] p-4 rounded-xl">
      {/* Mobile Layout */}
      <div className="flex text-nowrap text-sm justify-between md:hidden">
        <div className="flex items-center">
          <div className={`${logoBgColor} p-2 rounded-lg`}>
            {logoSvg}
          </div>
          <div className="ml-4">
            <h2 className="text-lg font-medium text-[#232323]">{serviceName}</h2>
            <p className="text-sm text-[#718EBF]">{serviceDescription}</p>
          </div>
        </div>
        <a href={viewDetailsLink} className="text-[#1814F3] font-semibold mt-2">View Details</a>
      </div>

      {/* Web Layout */}
      <div className="hidden md:flex items-center justify-between">
        <div className="flex items-center">
          <div className={`${logoBgColor} p-2 rounded-xl`}>
            {logoSvg}
          </div>
          <div className="ml-4 ">
            <h2 className="text-lg font-medium">{serviceName}</h2>
            <p className="text-sm text-[#718EBF]">{serviceDescription}</p>
          </div>
        </div>
        <div className="flex flex-1 justify-around items-center text-left">
          {additionalServices.map((service, index) => (
            <div key={index}>
              <h2 className="text-lg font-medium">{service.name}</h2>
              <p className="text-sm text-[#718EBF]">{service.description}</p>
            </div>
          ))}
        </div>
        <a href={viewDetailsLink} className="text-[#718EBF] border-[#718EBF] font-semibold hover:border-[#1814F3] hover:text-[#1814F3] border rounded-3xl px-8 py-2">View Details</a>
      </div>
    </div>
  );
};

export default BankServiceList;
