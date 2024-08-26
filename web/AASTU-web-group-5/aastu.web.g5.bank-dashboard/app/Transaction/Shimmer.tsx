import React from "react";

const Shimmer: React.FC = () => {
  return (
    <div className="relative overflow-hidden bg-gray-200 animate-pulse rounded-md h-full">
      <div className="absolute inset-0 bg-gradient-to-r from-gray-200 via-gray-300 to-gray-200 bg-[length:200%_auto] animate-shimmer"></div>
    </div>
  );
};

export default Shimmer;
