import React from "react";
import { Skeleton } from "../../ui/skeleton";

const RecentTransactionDashSkeleton: React.FC = () => {
  return (
    <div className="w-full">
      <h1 className="text-[#333B69] pb-3 font-semibold">Recent Transaction</h1>

      <div className="max-w-md bg-white rounded-[15px] px-4 py-3">
        <div className="flow-root">
          <ul role="list" className="space-y-2">
            {[...Array(3)].map((_, index) => (
              <li key={index} className="py-1">
                <div className="flex items-center">
                  <Skeleton className="w-[40px] h-[40px] rounded-full bg-slate-200" />
                  <div className="flex-1 min-w-0 ms-4">
                    <Skeleton className="h-4 w-2/3 bg-slate-200 mb-1" />
                    <Skeleton className="h-4 w-2/3 bg-slate-200" />
                  </div>
                  <Skeleton className="h-4 w-16 bg-slate-200 ml-auto" />
                </div>
              </li>
            ))}
          </ul>
        </div>
      </div>
    </div>
  );
};

export default RecentTransactionDashSkeleton;
