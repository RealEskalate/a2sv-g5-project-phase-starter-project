import { Skeleton } from "@/components/ui/skeleton";
import React from "react";
import CardSkeleton from "../CardSkeleton/CardSkeleton";

const CardListCardSkeleton = () => {
  const data = [1, 2];
  return (
    <div className="flex space-x-2">
      {data?.map((index: number) => (
        <CardSkeleton key={index} />
      ))}
    </div>
  );
};

export default CardListCardSkeleton;
