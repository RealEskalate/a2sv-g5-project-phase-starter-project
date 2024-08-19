import { Skeleton } from "@/components/ui/skeleton";
import React from "react";

const SkeletonCard = () => {
  return (
    <div className="flex w-full bg-white rounded-[20px] p-2 gap-2">
      <div className="w-1/4 items-center">
        <Skeleton className="h-12 w-12 rounded-[20px]" />
      </div>
      <div className="flex flex-col w-2/4 py-2 gap-1">
        <Skeleton className="h-4 w-full" />
        <Skeleton className="h-4 w-full" />
      </div>
      <div className="flex flex-col w-2/4 py-2 gap-1">
        <Skeleton className="h-4 w-full" />
        <Skeleton className="h-4 w-full" />
      </div>
      <div className="flex flex-col w-2/4 py-2 gap-1">
        <Skeleton className="h-4 w-full" />
        <Skeleton className="h-4 w-full" />
      </div>
    </div>
  );
};

export default SkeletonCard;
