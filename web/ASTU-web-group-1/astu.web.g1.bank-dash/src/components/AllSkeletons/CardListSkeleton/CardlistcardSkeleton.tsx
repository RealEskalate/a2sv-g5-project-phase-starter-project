import { Skeleton } from "@/components/ui/skeleton";
import React from "react";

const CardlistcardSkeleton = () => {
  return (
    <div className="flex  w-full bg-white rounded-[20px] p-2 px-3 mb-2 space-x-1">
      <div className="flex-shrink-0 mr-4">
        <Skeleton className="lg:w-[55px] lg:h-[55px] w-[40px] h-[40px] sm:rounded-[20px]" />
      </div>
      <div className="flex-1 min-w-0 py-2 space-y-1">
        <Skeleton className="h-5 w-full" />
        <Skeleton className="h-5 w-full" />
      </div>
      <div className="flex-1 min-w-0 py-2 space-y-1">
        <Skeleton className="h-5 w-full" />
        <Skeleton className="h-5 w-full" />
      </div>
      <div className="hidden  lg:block flex-1 min-w-0 py-2 space-y-1">
        <Skeleton className="h-5 w-full" />
        <Skeleton className="h-5 w-full" />
      </div>
      <div className="hidden  lg:block flex-1 min-w-0 py-2 space-y-1">
        <Skeleton className="h-5 w-full" />
        <Skeleton className="h-5 w-full" />
      </div>
      <div className="flex items-center">
        <Skeleton className="h-5 w-full" />
      </div>
    </div>
  );
};

export default CardlistcardSkeleton;
