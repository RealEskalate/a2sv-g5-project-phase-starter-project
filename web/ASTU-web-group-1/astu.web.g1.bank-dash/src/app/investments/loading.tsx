import React from "react";
import { SkeletonCard } from "@/components/AllSkeletons/itemSkeleton/SkeletonCard";
import ListSkeleton from "@/components/AllSkeletons/listskeleton/Skeleton";
import SkeletonList from "@/components/AllSkeletons/trendingSkeleton/skeletonList";
import ChartSkeleton from "@/components/AllSkeletons/chartSkeleton/chartSkeleton";

export default function page() {
  return (
    <div>
      <SkeletonCard />
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
