import React from "react";
import Image from "next/image";

interface Props {
  image: string;
  name: string;
  job: string;
  isSelected: boolean; // New prop to indicate if the profile is selected
  onClick: () => void; // New prop to handle click events
}

export const Profile = ({ image, name, job, isSelected, onClick }: Props) => {
  return (
    <div
      className={`min-w-[55px] cursor-pointer p-2 rounded-xl transition-all duration-300 ${
        isSelected
          ? "bg-blue-100 border-blue-500 shadow-lg"
          : "bg-transparent border-transparent"
      } border-2`}
      onClick={onClick}
    >
      <div className="flex flex-col items-center">
        <Image
          src={image}
          alt={`profile picture`}
          className="!rounded-full object-cover object-center"
          width={50}
          height={50}
        />
        <div className="mt-2 text-center">
          <h4 className="font-inter font-normal text-[12px]">{name}</h4>
          <h4 className="font-inter font-normal text-[12px] text-[#718EBF]">
            {job}
          </h4>
        </div>
      </div>
    </div>
  );
};
