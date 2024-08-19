import { Skeleton } from "@/components/ui/skeleton";
import React from "react";

const ServiceSkeletoncard = () => {
  return (
    <div className="flex bg-[#FFFFFF] rounded-2xl px-4 py-4 mr-2 md:pl-5 md:pr-9 md:py-5">
      <div className="flex w-full items-center gap-3 lg:px-3">
        <Skeleton className="object-cover rounded-full w-11 h-11 md:w-12 md:h-12 lg:w-14 lg:h-14" />
        <div className="space-y-1 max-w-full">
          <Skeleton className="flex h-5 w-28" />
          <Skeleton className="flex h-5 w-28" />
        </div>
      </div>
    </div>
  );
};

export default ServiceSkeletoncard;
