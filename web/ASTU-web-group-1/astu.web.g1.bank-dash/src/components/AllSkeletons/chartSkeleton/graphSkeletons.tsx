import { Skeleton } from "@/components/ui/skeleton";
import React from "react";

const GraphSkeletons = () => {
  return (
    <div className="w-full md:w-1/2 space-y-2 p-5">
      <Skeleton className="h-4 w-40 py-2 bg-slate-300 font-semibold mb-3"></Skeleton>
      <div className="border w-full bg-white space-y-5 items-end p-9 rounded-xl">
        <Skeleton className="border-2"></Skeleton>
        <Skeleton className="border-2"></Skeleton>
        <Skeleton className="border-2"></Skeleton>
        <Skeleton className="border-2"></Skeleton>
        <Skeleton className="border-2"></Skeleton>
        <Skeleton className="border-2"></Skeleton>
        <Skeleton className="border-2"></Skeleton>
        <Skeleton className="border-2"></Skeleton>
        <Skeleton className="border-2"></Skeleton>
        <Skeleton className="border-2"></Skeleton>
      </div>
    </div>
  );
};

export default GraphSkeletons;
