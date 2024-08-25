import React from "react";
import ListSkeleton from "@/components/AllSkeletons/Myinvestmentlistskeleton/Skeleton";
import SkeletonList from "@/components/AllSkeletons/trendingSkeleton/skeletonList";
import ChartSkeleton from "@/components/AllSkeletons/chartSkeleton/chartSkeleton";
import Skeletoncard from "@/components/AllSkeletons/loanSkeleton/skeletoncard";
<<<<<<< HEAD
import GraphSkeletons from "@/components/AllSkeletons/chartSkeleton/graphSkeletons";
=======
>>>>>>> 87ba7e7340a82fdda6571e0853e8742443d641a7

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
