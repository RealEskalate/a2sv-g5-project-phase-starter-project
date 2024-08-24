import React from "react";
import Image from "next/image";

interface ServiceCardProps {
  iconSrc: string;
  title: string;
  subtitle: string;
  bgColor: string;
}

const ServiceCard: React.FC<ServiceCardProps> = ({
  iconSrc,
  title,
  subtitle,
  bgColor,
}) => {
  return (
    <div className="flex items-center p-8 bg-white rounded-lg shadow-lg">
      <div
        className={`flex items-center justify-center w-20 h-20 rounded-full mr-8 ${bgColor}`}
      >
        <Image src={iconSrc} alt={`${title} Icon`} width={50} height={50} />
      </div>
      <div>
        <h3 className="text-2xl font-semibold text-black">{title}</h3>
        <p className="text-lg text-gray-600">{subtitle}</p>
      </div>
    </div>
  );
};

export default ServiceCard;
