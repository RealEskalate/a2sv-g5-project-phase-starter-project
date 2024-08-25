import React from "react";
import Image from "next/image";

interface CardProps {
  title: string;
  description: string;
  color: string;
  children: React.ReactNode;
}
const ServicesCard: React.FC<CardProps> = ({
  children,
  title,
  description,
  color,
}) => {
  return (
    <div
      style={{ backgroundColor: color }}
      className="flex min-w-fit lg:min-w-1/4  rounded-3xl p-5 lg:w-1/4"
    >
      <div className="flex justify-center items-center  whitespace-nowrap lg:px-6">
        <div className="px-4">{children}</div>

        <div>
          <h2 className="flex mb-1 text-gray-dark text-15px lg:text-18px font-[600]">
            {title}
          </h2>
          <p
            className="flex text-12px lg:text-14px text-blue-steel overflow-clip"
            style={{
              display: "-webkit-box",
              WebkitLineClamp: 2,
              WebkitBoxOrient: "vertical",
              overflow: "hidden",
              textOverflow: "ellipsis",
            }}
          >
            {description}
          </p>
        </div>
      </div>
    </div>
  );
};

export default ServicesCard;
