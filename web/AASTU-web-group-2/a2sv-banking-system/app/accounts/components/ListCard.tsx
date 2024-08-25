import React from "react";
import { IconType } from "react-icons";
type DataItem = {
  heading: string;
  text: string;
  headingStyle: string;
  dataStyle: string;
};

type Column = {
  icon: IconType;
  iconStyle: string;
  data: DataItem[]; // Updated to an array of DataItem objects
};

interface Props {
  column: Column;
  width: string;
  darkMode: string
}

const ListCard = ({ column, width, darkMode }: Props) => {
  return (
    <div
      className={`flex gap-3 items-center rounded-2xl px-5 py-4 bg-white ${width} dark:bg-[#020817] ${darkMode}`}
    >
      <div className={`text-3xl px-2 py-2 rounded-full ${column.iconStyle}`}>
        <column.icon />
      </div>
      <div className="flex justify-between w-full">
        {column.data.map((item, index) => (
          <div key={index}>
            <div className={`${item.headingStyle} dark:text-[#9faaeb]`}>{item.heading}</div>
            <div className={item.dataStyle}>{item.text}</div>
          </div>
        ))}
      </div>
    </div>
  );
};

export const ListCardLoading = () => {
  return (
    <div className="flex gap-3 items-center rounded-2xl px-5 py-4 bg-white dark:bg-[#020817] w-[48%] md:w-[23%] animate-pulse">
      <div className="text-3xl px-2 py-2 rounded-xl bg-gray-200 dark:bg-[#333B69]">
        <div className="w-8 h-8 bg-gray-300 dark:bg-[#555B85] rounded-full"></div>
      </div>
      <div className="flex justify-between w-full flex-col">
        <div>
          <div className="h-4 bg-gray-300 dark:bg-[#555B85] rounded w-full mb-2"></div>
          <div className="h-4 bg-gray-300 dark:bg-[#555B85] rounded w-full"></div>
        </div>
      </div>
    </div>
  );
};

export default ListCard;
