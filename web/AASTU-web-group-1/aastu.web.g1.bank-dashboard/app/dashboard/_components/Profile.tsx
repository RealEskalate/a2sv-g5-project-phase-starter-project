import React from "react";
import Image from "next/image";
import { useUser } from "@/contexts/UserContext";

interface Props {
  image: string;
  name: string;
  job: string;
  isSelected: boolean;
  onClick: () => void;
}

export const Profile = ({ image, name, job, isSelected, onClick }: Props) => {
  const { isDarkMode } = useUser();

  return (
    <div
      className={`
        min-w-[55px]
        cursor-pointer
        p-1
        rounded-xl
        transition-all
        duration-300
        ${
          isSelected
            ? `${
                isDarkMode
                  ? "bg-blue-800 border-blue-500"
                  : "bg-blue-100 border-blue-500"
              } shadow-lg`
            : `${
                isDarkMode
                  ? "bg-gray-700 border-transparent"
                  : "bg-transparent border-transparent"
              }`
        }
      `}
      border-2
      onClick={onClick}
    >
      <div className="flex flex-col items-center w-20">
        <Image
          src={image}
          alt={`profile picture`}
          className="!rounded-full object-cover object-center"
          width={50}
          height={50}
        />
        <div className="mt-2 w-12 truncate overflow-hidden group rel">
          <div className="  ">
            {/* Wrapping text elements in a div */}

            <h4
              className={`font-inter font-normal text-xs truncate   ${
                isDarkMode ? "text-white" : "text-black"
              }`}
            >
              {name}
            </h4>
            <h4
              className={`font-inter font-normal text-xs truncate overflow-x-hidden ${
                isDarkMode ? "text-[#9AA1B4]" : "text-[#718EBF]"
              }`}
            >
              {job}
            </h4>

            {/* Full text on hover */}
            {/* <div className="absolute inset-0 p-2 mt-1 text-sm text-white bg-black rounded-lg opacity-0 group-hover:opacity-100 z-10 max-w-xs">
              <p className="whitespace-normal break-words">{name}</p>

            </div> */}
          </div>
        </div>
      </div>
    </div>
  );
};
