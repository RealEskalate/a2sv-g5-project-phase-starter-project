import { Skeleton } from "@/components/ui/skeleton";
import React from "react";

const CardListCardSkeleton = () => {
  return (
    <div className="flex  w-full bg-white rounded-[20px] p-2 px-3 mb-2">
      <div className="flex-shrink-0 mr-4">
        <Skeleton className="lg:w-[40px] lg:h-[40px] w-[30px] h-[30px] rounded-3xl bg-slate-200" />
      </div>
      <div className="flex-1 min-w-0 py-2 ">
        <Skeleton className="bg-slate-200 w-24 h-4 mb-1"></Skeleton>
        <Skeleton className="bg-slate-200 w-24 h-4"></Skeleton>
      </div>
      <div className="flex-1 min-w-0 py-2">
        <Skeleton className="bg-slate-200 w-24 h-4 mb-1"></Skeleton>
        <Skeleton className="bg-slate-200 w-24 h-4"></Skeleton>
      </div>
      <div className="hidden  lg:block flex-1 min-w-0 py-2">
        <Skeleton className="bg-slate-200 w-24 h-4 mb-1"></Skeleton>
        <Skeleton className="bg-slate-200 w-24 h-4"></Skeleton>
      </div>
      <div className="hidden  lg:block flex-1 min-w-0 py-2">
        <Skeleton className="bg-slate-200 w-24 h-4 mb-1"></Skeleton>
        <Skeleton className="bg-slate-200 w-24 h-4"></Skeleton>
      </div>
      <div className="flex items-center">
        <Skeleton className="bg-slate-200 w-24 h-6 " />
      </div>
    </div>
  );
};

export default CardListCardSkeleton;
