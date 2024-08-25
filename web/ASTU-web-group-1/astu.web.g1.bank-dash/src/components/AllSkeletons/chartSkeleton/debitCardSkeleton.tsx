import { Skeleton } from "@/components/ui/skeleton";
import React from "react";

const DebitCardSkeleton = () => {
  return (
    <div className="w-full  md:w-8/12 mb-5">
      <Skeleton className="text-[#333B69] pb-2 font-semibold"></Skeleton>
      <div className="h-[400px] bg-white p-6 rounded-3xl">
        <div className="flex justify-between align-middle mb-6">
          <h3 className="text-blue-steel space-y-1">
            <Skeleton className="h-3 w-20"></Skeleton>
            <Skeleton className="h-3 w-20"></Skeleton>
          </h3>
          <div className="flex space-x-5">
            <div className="text-sm flex items-center text-blue-steel">
              <Skeleton className="h-4 w-4 rounded-sm inline-block me-2"></Skeleton>
            </div>
            <div className="text-sm flex items-center text-blue-steel">
              <Skeleton className="h-4 w-4 rounded-sm inline-block me-2"></Skeleton>
            </div>
          </div>
        </div>
        <div className="h-3/4 lg:h-[90%]">
          <div className="bg-slate-50 space-y-3 rounded-3xl md:px-10 md:py-10 p-4">
            <div className="flex items-end justify-between">
              <div className="flex items-baseline space-x-1">
                <Skeleton className="h-20 w-6 rounded-lg bg-slate-300" />
                <Skeleton className="h-40 w-6 rounded-lg bg-slate-300" />
              </div>
              <div className="flex items-baseline space-x-1">
                <Skeleton className="h-20 w-6 rounded-lg bg-slate-300" />
                <Skeleton className="h-32 w-6 rounded-lg bg-slate-300" />
              </div>
              <div className="flex items-baseline space-x-1">
                <Skeleton className="h-24 w-6 rounded-lg bg-slate-300" />
                <Skeleton className="h-10 w-6 rounded-lg bg-slate-300" />
              </div>
              <div className="flex items-baseline space-x-7">
                <Skeleton className="h-36 w-6 rounded-lg bg-slate-300" />
                <Skeleton className="h-28 w-6 rounded-lg bg-slate-300" />
              </div>
              <div className="flex items-baseline space-x-1">
                <Skeleton className="h-36 w-6 rounded-lg bg-slate-300" />
                <Skeleton className="h-28 w-6 rounded-lg bg-slate-300" />
              </div>
              <div className="flex items-baseline space-x-1">
                <Skeleton className="h-32 w-6 rounded-lg bg-slate-300" />
                <Skeleton className="h-40 w-6 rounded-lg bg-slate-300" />
              </div>
            </div>
            <div className="flex justify-between">
              <Skeleton className="h-3 w-20 bg-slate-200" />
              <Skeleton className="h-3 w-20 bg-slate-200" />
              <Skeleton className="h-3 w-20 bg-slate-200" />
              <Skeleton className="h-3 w-20 bg-slate-200" />
              <Skeleton className="h-3 w-20 bg-slate-200" />
              <Skeleton className="h-3 w-20 bg-slate-200" />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default DebitCardSkeleton;
