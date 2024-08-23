import RecentTransactionDashSkeleton from "@/components/AllSkeletons/RecentTransaction-DashSkelton/RecentTransactionDashSkelton";
import React from "react";
import WeeklyActivityChartSkeleton from "@/components/AllSkeletons/weeklyActivityChartSkeleton/WeeklyActvityChartSkeleton";
import ExpenseStatisticsSkeleton from "@/components/AllSkeletons/ExpenseStatistics/ExpenseStatisticsSkeleton";
import QuickTransferSkeleton from "@/components/AllSkeletons/QuickTransferSkeleton/QuickTransferSkeleton";
import BalanceHistoryChartSkeleton from "@/components/AllSkeletons/BalanceHistorychartSkeleton/BalanceHistortChartSkeleton";
import MyCardSkeleton from "@/components/AllSkeletons/MyCard/MyCardSkelelton";

const loading = () => {
  return (
    <>
      <div className="w-full lg:flex ">
        <div className="lg:w-2/3 md:pr-3 xl:pr-5 flex-shrink">
          <div className="w-full">
            <div className="flex justify-between">
              <p className="text-[#333B69] pb-3 font-semibold">My Card</p>
              <p className="text-[#333B69] pb-3 font-semibold">See All</p>
            </div>
            <div className="flex  overflow-x-auto space-x-2">
              <MyCardSkeleton />
              <MyCardSkeleton />
            </div>
          </div>
        </div>
        <div className="lg:w-1/3 w-full">
          <RecentTransactionDashSkeleton />
        </div>
      </div>
      <div className="md:flex my-5">
        <WeeklyActivityChartSkeleton />
        <ExpenseStatisticsSkeleton />
      </div>
      <div className="md:flex justify-between">
        <div className="w-5/12 pe-6">
          <QuickTransferSkeleton />
        </div>
        <BalanceHistoryChartSkeleton />
      </div>
    </>
  );
};

export default loading;
