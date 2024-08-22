import React from "react";

export const ShimmerRecent: React.FC = () => {
  return (
    <div className="flex flex-col w-full gap-4 bg-gray-100 border dark:bg-[#232328] rounded-3xl shadow-gray-50 p-6 animate-pulse">
      {[...Array(3)].map((_, key) => (
        <div
          key={key}
          className="recentTr w-full flex gap-4 items-center justify-center"
        >
          {/* Placeholder for Icon */}
          <div className="icon flex items-center rounded-full p-4 bg-gray-200 dark:bg-gray-600 w-12 h-12"></div>

          {/* Placeholder for Transaction Details */}
          <div className="flex flex-col gap-1 w-1/2">
            <div className="title h-4 bg-gray-300 dark:bg-gray-600 rounded w-3/4"></div>
            <div className="date h-3 bg-gray-300 dark:bg-gray-600 rounded w-1/2"></div>
          </div>

          {/* Placeholder for Price */}
          <div className="price flex grow justify-end">
            <div className="h-4 bg-gray-300 dark:bg-gray-600 rounded w-16"></div>
          </div>
        </div>
      ))}
    </div>
  );
};

export default ShimmerRecent;
