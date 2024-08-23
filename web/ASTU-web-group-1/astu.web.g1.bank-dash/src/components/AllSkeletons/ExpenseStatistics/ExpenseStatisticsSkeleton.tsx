import React from "react";
import { Skeleton } from "../../ui/skeleton";

const ExpenseStatisticsSkeleton = () => {
  return (
    <div className="w-full md:w-4/12">
      <h1 className="text-[#333B69] pb-3 font-semibold">Expense Statistics</h1>

      <div className="bg-white rounded-3xl flex justify-center items-center h-[350px]">
        <Skeleton className="h-[300px] w-[300px] bg-slate-200 rounded-full" />
      </div>
    </div>
  );
};

export default ExpenseStatisticsSkeleton;
