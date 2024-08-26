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
  additionalServices?: Service[];
  viewDetailsLink?: string;
}

const BankServiceList: React.FC<BankServiceListProps> = ({
  logoBgColor = "bg-pink-100",
  logoSvg = (
    <svg className="w-8 h-8 text-pink-500 dark:text-[#9faaeb]" fill="currentColor " viewBox="0 0 24 24">
      {/* Default SVG content */}
    </svg>
  ),
  serviceName = "Business loans",
  serviceDescription = "It is a long established",
  additionalServices = [],
  viewDetailsLink = "#",
}) => {
  return (
    <div className="mx-5 my-5 bg-white p-4 rounded-xl dark:bg-[#020817] dark:text-[#9faaeb] dark:border dark:border-[#333B69]">
      {/* Mobile Layout */}
      <div className="flex text-nowrap text-sm justify-between md:hidden dark:text-[#9faaeb]">
        <div className="flex items-center">
          <div className={`${logoBgColor} p-2 rounded-lg`}>
            roller
            {logoSvg}
          </div>
          <div className="ml-4">
            <h2 className="text-lg font-medium text-gray-900 dark:text-[#9faaeb]">{serviceName}</h2>
            <p className="text-sm text-gray-500 dark:text-[#9faaeb]">{serviceDescription}</p>
          </div>
        </div>
        <a href={viewDetailsLink} className="text-blue-500 font-semibold mt-2">View Details</a>
      </div>

      {/* Web Layout */}
      <div className="hidden md:flex items-center justify-between dark:text-[#9faaeb]">
        <div className="flex items-center w-1/4">
          <div className={`${logoBgColor} p-2 rounded-xl`}>
            {logoSvg}
          </div>
          <div className="ml-4">
            <h2 className="text-lg font-medium text-gray-900 dark:text-[#9faaeb]">{serviceName}</h2>
            <p className="text-sm text-gray-500 dark:text-[#9faaeb]">{serviceDescription}</p>
          </div>
        </div>
        <div className="flex flex-1 justify-around items-center">
          {additionalServices.map((service, index) => (
            <div key={index} className="text-center w-1/4">
              <h2 className="text-lg font-medium text-gray-900 dark:text-[#9faaeb]">{service.name}</h2>
              <p className="text-sm text-gray-500 dark:text-[#9faaeb]">{service.description}</p>
            </div>
          ))}
        </div>
        <a href={viewDetailsLink} className="text-[#718EBF] border-[#718EBF] font-semibold hover:border-[#1814F3] hover:text-[#1814F3] border rounded-3xl px-8 py-2">View Details</a>
      </div>
    </div>
  );
};

export default BankServiceList;
