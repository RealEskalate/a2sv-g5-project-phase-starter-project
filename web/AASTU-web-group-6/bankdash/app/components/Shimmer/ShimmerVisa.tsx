import React from "react";

export const ShimmerVisaCard: React.FC<{
  isBlack?: boolean;
  className?: string;
}> = ({ isBlack, className }) => {
  return (
    <div
      className={`w-full max-h-[242p xxs:min-w-[300px] font-Lato flex flex-col gap-2 grow rounded-3xl animate-pulse bg-gradient-to-r from-gray-200 via-gray-300 to-gray-200 dark:bg-gradient-to-r dark:from-[#222229] dark:via-[#353535] dark:to-[#232328]"
      } ${className}`}
    >
      <div className="flex flex-col gap-6 p-6">
        <div className="flex items-center justify-between">
          {/* Placeholder for Balance */}
          <div className="balance-box w-32 h-6 bg-gray-300 dark:bg-gray-600 rounded"></div>

          {/* Placeholder for Sim Icon */}
          <div className="w-12 h-8 bg-gray-300 dark:bg-gray-600 rounded"></div>
        </div>

        <div className="flex justify-between">
          {/* Placeholder for Card Holder */}
          <div className="flex flex-col gap-2 w-1/2">
            <div className="w-20 h-4 bg-gray-300 dark:bg-gray-600 rounded"></div>
            <div className="w-28 h-5 bg-gray-300 dark:bg-gray-600 rounded"></div>
          </div>

          {/* Placeholder for Valid Thru */}
          <div className="flex flex-col gap-2 w-1/4">
            <div className="w-16 h-4 bg-gray-300 dark:bg-gray-600 rounded"></div>
            <div className="w-16 h-5 bg-gray-300 dark:bg-gray-600 rounded"></div>
          </div>
        </div>
      </div>

      <div className="flex justify-between items-center gap-2 card-box rounded-b-3xl p-4 bg-gray-200 dark:bg-[#22272c]">
        {/* Placeholder for Card Number */}
        <div className="w-40 h-6 bg-gray-300 dark:bg-gray-600 rounded"></div>

        {/* Placeholder for Card Circles */}
        <div className="flex gap-1">
          <div className="w-7 h-7 bg-gray-300 dark:bg-gray-600 rounded-full"></div>
          <div className="w-7 h-7 bg-gray-300 dark:bg-gray-600 rounded-full -ml-4"></div>
        </div>
      </div>
    </div>
  );
};
