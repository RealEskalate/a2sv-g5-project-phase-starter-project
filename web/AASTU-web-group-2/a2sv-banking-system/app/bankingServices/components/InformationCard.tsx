import React from 'react';

interface InformationCardProps {
  logoBgColor?: string;
  logo: React.ReactNode;
  title: string;
  description: string;
  cardBgColor?: string;
}

const InformationCard: React.FC<InformationCardProps> = ({
  logoBgColor = '#E7EDFF',
  logo,
  title,
  description,
  cardBgColor = 'bg-green-200',
}) => {
  return (
    <div className="w-full dark:bg-[#020817] dark:text-[#9faaeb]">
      {/* Mobile Layout */}
      <div className={`dark:bg-[#020817] dark:border dark:border-[#333B69] flex items-center gap-4 py-6 px-8 ${cardBgColor} rounded-3xl text-nowrap md:hidden dark:border-[#333B69]`}>
        <div className="flex items-center rounded-full min-h-14 min-w-14 justify-center" style={{ backgroundColor: logoBgColor }}>
          {logo}
        </div>
        <div className="flex-col">
          <h1 className="text-[#232323] font-bold text-lg dark:text-white">{title}</h1>
          <h1 className="text-[#718EBF] text-sm dark:[#9faaeb]">{description}</h1>
        </div>
      </div>

      {/* Web Layout */}
      <div className={` dark:bg-[#020817] dark:border dark:border-[#333B69] hidden md:flex justify-center w-full items-center gap-4 py-4 px-4 ${cardBgColor} rounded-3xl text-nowrap  `}>
        <div className="flex items-center rounded-full min-h-16 min-w-16 justify-center " style={{ backgroundColor: logoBgColor }}>
          {logo}
        </div>
        <div className="flex-col justify-start ">
          <h1 className="text-[#232323] font-bold text-lg dark:text-white">{title}</h1>
          <h1 className="text-[#718EBF] text-sm dark:text-[#9faaeb]">{description}</h1>
        </div>
      </div>
    </div>
  );
};

export default InformationCard;
