import { Skeleton } from "@/components/ui/skeleton";
import React from "react";

const CardAndExpenceStatisticsSkeleton: React.FC = () => {
  return (
    <div>
      <Skeleton className="w-48 h-6 pb-2 bg-slate-200 "></Skeleton>
      <div className="w-full  bg-white p-6 rounded-3xl py-10 mt-2">
        <div className="flex justify-center">
          <Skeleton className=" w-56 h-32 bg-slate-200 " />
        </div>
        <div className="flex flex-col justify-items-center items-center mt-8 space-y-3">
          <Skeleton className="flex items-center w-3/4 h-5 bg-slate-200" />
          <Skeleton className="flex items-center w-3/4 h-5 bg-slate-200" />
        </div>
      </div>
    </div>
  );
};

export default CardAndExpenceStatisticsSkeleton;
