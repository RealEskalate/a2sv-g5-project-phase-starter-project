import React from "react";
import { Skeleton } from "../../ui/skeleton";

const WeeklyActivitySkeleton = () => {
  return (
    <div className="w-full md:w-8/12 md:me-6">
      <h1 className="text-[#333B69] pb-3 font-semibold">Weekly Activity</h1>

      <div className="bg-white rounded-3xl p-10">
        <Skeleton className="h-[300px] w-full bg-slate-200 rounded-md" />
      </div>
    </div>
  );
};

export default WeeklyActivitySkeleton;
