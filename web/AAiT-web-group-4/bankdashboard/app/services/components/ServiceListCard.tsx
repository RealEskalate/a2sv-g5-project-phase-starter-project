import React from "react";
import Image from "next/image";

interface DetailItem {
  title: string;
  subtitle: string;
}

interface ServiceListCardProps {
  iconSrc: string;
  mainTitle: string;
  mainSubtitle: string;
  details: DetailItem[];
  buttonText: string;
  bgColor: string;
}

const ServiceListCard: React.FC<ServiceListCardProps> = ({
  iconSrc,
  mainTitle,
  mainSubtitle,
  details,
  buttonText,
  bgColor,
}) => {
  return (
    <div className="flex items-center p-4 bg-white rounded-lg shadow-md space-x-50">
      <div
        className={`flex items-center justify-center w-16 h-16 rounded-full mr-2 ${bgColor}`}
      >
        <Image src={iconSrc} alt={`${mainTitle} Icon`} width={32} height={32} />
      </div>

      <div className="flex-grow mr-0">
        <h3 className="text-base font-semibold text-black">{mainTitle}</h3>
        <p className="text-sm text-gray-500">{mainSubtitle}</p>
      </div>
      <div className="flex items-center justify-center mr-20">
      <div className="hidden tablet:flex flex-1 space-x-20">
        {details.map((item, index) => (
          <div key={index} className="flex-1 text-center">
            <h3 className="text-base text-responsive font-semibold text-black">{item.title}</h3>
            <p className="text-sm text-responsive  text-gray-500">{item.subtitle}</p>
          </div>
        ))}
      </div>
      </div>
      

      <div>
      <button  
  className="px-4 py-1 rounded-full   
             bg-white text-customBlue   
             border-2 border-transparent   
             md:border-gray-500   
             text-gray-500   
             hover:text-blue-600  
             hover:border-transparent
             md:hover:border-blue-600
             transition duration-300 ease-in-out  
             ">  
  {buttonText}  
</button>






      </div>
    </div>
  );
};

export default ServiceListCard;
