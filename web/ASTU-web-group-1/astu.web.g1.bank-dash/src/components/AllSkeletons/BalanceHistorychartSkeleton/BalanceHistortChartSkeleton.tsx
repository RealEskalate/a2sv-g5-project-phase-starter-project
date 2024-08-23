import React from "react";
import { Skeleton } from "../../ui/skeleton";

const BalanceHistoryChartSkeleton = () => {
  return (
    <div className="w-full md:w-7/12">
      <h1 className="text-[#333B69] pb-3 font-semibold">Balance History</h1>
      <div className="bg-white px-3 py-5 rounded-3xl">
        <Skeleton className="h-[240px] w-full bg-slate-200 rounded-2xl" />
      </div>
    </div>
  );
};

export default BalanceHistoryChartSkeleton;
