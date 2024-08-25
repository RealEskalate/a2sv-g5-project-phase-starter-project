import { Skeleton } from "@/components/ui/skeleton";
import React from "react";

const CardlistcardSkeleton = () => {
  return (
    <div className="flex  w-full bg-white rounded-[20px] p-2 px-3 mb-2">
      <div className="flex-shrink-0 mr-4">
        <Skeleton className="lg:w-[55px] lg:h-[55px] w-[40px] h-[40px] sm:rounded-[20px]" />
      </div>
      <div className="flex-1 min-w-0 py-2 ">
        <Skeleton className="h-5 w-1/2" />
        <Skeleton className="h-5 w-1/2" />
      </div>
      <div className="flex-1 min-w-0 py-2">
        <Skeleton className="h-5 w-1/2" />
        <Skeleton className="h-5 w-1/2" />
      </div>
      <div className="hidden  lg:block flex-1 min-w-0 py-2">
        <p className="text-sm text-[#232323]">Card Number</p>
        <p className="text-xs text-[#718EBF]">{cardNumber}</p>
      </div>
      <div className="hidden  lg:block flex-1 min-w-0 py-2">
        <p className="text-sm text-[#232323]">Namain Card</p>
        <p className="text-xs text-[#718EBF]">{namainCard}</p>
      </div>
      <div className="flex items-center">
        <a className="text-sm text-[#1814F3] " href="#">
          View Details
        </a>
      </div>
    </div>
  );
};

export default CardlistcardSkeleton;
