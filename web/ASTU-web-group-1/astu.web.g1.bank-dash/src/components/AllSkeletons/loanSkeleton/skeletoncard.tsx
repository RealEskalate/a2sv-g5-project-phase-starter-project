import { Skeleton } from "@/components/ui/skeleton";
import React from "react";

const Skeletoncard = () => {
  return (
    <div className="flex bg-[#FFFFFF] rounded-2xl px-4 py-4 mr-2 md:pl-5 md:pr-9 md:py-5">
      <div className="flex items-center space-x-1 w-full">
        <Skeleton className="object-cover rounded-full w-11 h-11 md:w-12 md:h-12 lg:w-14 lg:h-14 bg-slate-200" />
        <div className="space-y-1 max-w-full">
          <Skeleton className="flex h-5 w-28 bg-slate-200" />
          <Skeleton className="flex h-5 w-28 bg-slate-200" />
        </div>
      </div>
    </div>
  );
};

export default Skeletoncard;
