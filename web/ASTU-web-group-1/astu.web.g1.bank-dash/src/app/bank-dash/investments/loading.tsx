import React from "react";
import ListSkeleton from "@/components/AllSkeletons/Myinvestmentlistskeleton/Skeleton";
import SkeletonList from "@/components/AllSkeletons/trendingSkeleton/skeletonList";
import ChartSkeleton from "@/components/AllSkeletons/chartSkeleton/chartSkeleton";
import Skeletoncard from "@/components/AllSkeletons/loanSkeleton/skeletoncard";

export default function page() {
  return (
    <div>
      <Skeletoncard />
      <div className="md:flex space-x-4">
        <ChartSkeleton />
        <ChartSkeleton />
      </div>
      <div className="flex flex-col md:flex-row ">
        <ListSkeleton />
        <SkeletonList />
      </div>
    </div>
  );
}
