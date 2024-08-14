import React from 'react';
import Image from 'next/image';

interface CardProps {
  image: string;
  title: string;
  description: string;
  color:string
}

const ServicesCard: React.FC<CardProps> = ({  image, title, description,color }) => {
  return (
    <div style={{ backgroundColor: color }}  className="flex min-w-fit lg:min-w-1/4  rounded-3xl p-5 mx-2 lg:w-1/4">
        <div  className="flex justify-center px-8 gap-1 items-center min-w-fit whitespace-nowrap">
      <Image
            src={image}
            alt=""
            width={70}
            height={70}
            className="object-cover w-13 h-13 md:w-14 md:h-14 lg:w-20 lg:h-20"
          />
      
      <div className="p-4">
        <h2 className="flex mb-1 text-gray-dark text-15px lg:text-18px font-[600]">{title}</h2>
        <p className="flex text-12px lg:text-14px text-blue-steel overflow-clip"
              style={{
                display: "-webkit-box",
                WebkitLineClamp: 2,
                WebkitBoxOrient: "vertical",
                overflow: "hidden",
                textOverflow: "ellipsis",
              }}>{description}</p>
      </div>
      </div>
    </div>
  );
}

export default ServicesCard;
