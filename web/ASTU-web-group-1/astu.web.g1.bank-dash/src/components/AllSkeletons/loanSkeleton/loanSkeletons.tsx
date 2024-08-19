import React from "react";
import Skeletoncard from "./skeletoncard";

const LoanSkeletons = () => {
  return (
    <div
      className="flex overflow-x-auto justify-around overflow-clip whitespace-nowrap w-full"
      style={{
        scrollbarWidth: "none",
        msOverflowStyle: "none",
      }}
    >
      <Skeletoncard />
      <Skeletoncard />
      <Skeletoncard />
      <Skeletoncard />
    </div>
  );
};

export default LoanSkeletons;
