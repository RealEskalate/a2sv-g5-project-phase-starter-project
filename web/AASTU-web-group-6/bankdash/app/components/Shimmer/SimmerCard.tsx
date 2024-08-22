import React from "react";

const ShimmerCard = () => {
  return (
    <div className="mb-5 flex items-center grow justify-between h-24 bg-white dark:bg-[#232328] gap-5 border dark:border-gray-700 rounded-3xl animate-pulse xs:w-full md:w-full lg:w-full">
      <div className="icon rounded-full p-3 bg-gray-300 dark:bg-gray-700 ml-4 w-14 h-14"></div>
      <div className="flex flex-col gap-2 w-1/4">
        <div className="h-4 bg-gray-300 dark:bg-gray-700 rounded w-3/4"></div>
        <div className="h-3 bg-gray-200 dark:bg-gray-600 rounded w-1/2"></div>
      </div>
      <div className="flex flex-col gap-2 w-1/4">
        <div className="h-4 bg-gray-300 dark:bg-gray-700 rounded w-3/4"></div>
        <div className="h-3 bg-gray-200 dark:bg-gray-600 rounded w-1/2"></div>
      </div>
      <div className="flex flex-col gap-2 w-1/4">
        <div className="h-4 bg-gray-300 dark:bg-gray-700 rounded w-3/4"></div>
        <div className="h-3 bg-gray-200 dark:bg-gray-600 rounded w-1/2"></div>
      </div>
      <div className="flex flex-col gap-2 w-1/4">
        <div className="h-4 bg-gray-300 dark:bg-gray-700 rounded w-3/4"></div>
        <div className="h-3 bg-gray-200 dark:bg-gray-600 rounded w-1/2"></div>
      </div>
      <div className="w-36 h-8 bg-gray-300 dark:bg-gray-700 rounded-[50px] mr-8"></div>
    </div>
  );
};

export default ShimmerCard;
