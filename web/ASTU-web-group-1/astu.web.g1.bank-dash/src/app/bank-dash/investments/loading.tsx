import React from "react";
import ListSkeleton from "@/components/AllSkeletons/Myinvestmentlistskeleton/Skeleton";
import SkeletonList from "@/components/AllSkeletons/trendingSkeleton/skeletonList";
import Skeletoncard from "@/components/AllSkeletons/loanSkeleton/skeletoncard";
import GraphSkeletons from "@/components/AllSkeletons/chartSkeleton/graphSkeletons";

export default function page() {
  return (
    <div>
      <Skeletoncard />
      <div className="md:flex space-x-4">
        <GraphSkeletons />
        <GraphSkeletons />
      </div>
      <div className="flex flex-col md:flex-row ">
        <ListSkeleton />
        <SkeletonList />
      </div>
    </div>
  );
}
