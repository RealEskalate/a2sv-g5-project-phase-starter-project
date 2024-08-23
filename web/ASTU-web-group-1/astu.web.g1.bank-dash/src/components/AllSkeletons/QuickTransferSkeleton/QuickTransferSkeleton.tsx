import React from "react";
import { Skeleton } from "../../ui/skeleton";

const QuickTransferSkeleton = () => {
  return (
    <div className="flex flex-col">
      <Skeleton className="h-6 w-1/3 bg-slate-200 mb-3" />
      <div className="flex flex-col lg:gap-5 rounded-3xl md:w-[320px] lg:w-full w-[320px] py-5 px-6 bg-white">
        <div className="px-10 md:px-6">
          <Skeleton className="h-[120px] w-full bg-slate-200 rounded-2xl mb-4" />
        </div>
        <div className="flex flex-row w-full text-15px mt-4">
          <Skeleton className="h-5 w-1/2 bg-slate-200 mr-2" />
          <div className="flex relative flex-row items-center h-10 rounded-full w-full">
            <Skeleton className="flex-1 h-full bg-slate-200 rounded-full" />
            <Skeleton className="absolute right-0 h-full w-[80px] bg-blue-bright rounded-full" />
          </div>
        </div>
      </div>
    </div>
  );
};

export default QuickTransferSkeleton;
