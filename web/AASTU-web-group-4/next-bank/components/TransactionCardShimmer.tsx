import React from 'react';

const TransactionCardShimmer: React.FC = () => {
  return (
    <div className="flex dark:bg-dark flex-col pt-4 md:flex-row md:justify-evenly gap-4 gap-y-6 w-auto rounded-2xl shadow-none border-none animate-pulse">
      
      {/* Desktop View */}
      <div className="hidden md:flex items-center">
        <div className="w-8 h-8 rounded-full bg-gray-300 mr-4"></div>
        <div className="w-40 bg-gray-300 h-6 rounded-md"></div>
      </div>
      <div className="hidden md:flex items-center w-24">
        <div className="w-full bg-gray-300 h-6 rounded-md"></div>
      </div>
      <div className="hidden md:flex items-center w-28">
        <div className="w-full bg-gray-300 h-6 rounded-md"></div>
      </div>
      <div className="hidden md:flex items-center w-20">
        <div className="w-full bg-gray-300 h-6 rounded-md"></div>
      </div>
      <div className="hidden md:flex items-center justify-end w-24">
        <div className="w-full bg-gray-300 h-6 rounded-md"></div>
      </div>
      
      {/* Mobile View */}
      <div className="md:hidden flex justify-between flex-row w-full">
        <div className="flex items-center">
          <div className="w-8 h-8 rounded-full bg-gray-300 mr-4"></div>
          <div className="flex flex-col">
            <div className="w-24 bg-gray-300 h-6 rounded-md mb-2"></div>
            <div className="w-16 bg-gray-300 h-4 rounded-md"></div>
          </div>
        </div>
        <div className="flex items-center">
          <div className="w-16 bg-gray-300 h-6 rounded-md"></div>
        </div>
      </div>
      
    </div>
  );
};

export default TransactionCardShimmer;
