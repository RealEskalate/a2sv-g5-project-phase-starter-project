import { Skeleton } from "@/components/ui/skeleton";
import React from "react";

const BankServiceSkeleton = () => {
  const BankServices = [1, 2, 3, 4, 5, 6];
  return (
    <div className="flex flex-col gap-4">
      <p className="text-[#333B69] pb-3 font-semibold">Bank Services List</p>
      <div className="flex flex-col gap-3">
        {BankServices.map((index) => (
          <div
            key={index}
            className="grid grid-cols-7 md:grid-cols-7 xl:grid-cols-11 items-center bg-white p-3 rounded-xl gap-2"
          >
            <div className="flex col-span-5 md:col-span-3 items-center gap-3 lg:gap-4">
              <Skeleton className="flex relative w-9 h-9 lg:w-12 lg:h-12 items-center"></Skeleton>
              <div className="flex flex-col w-full gap-1">
                <Skeleton className="w-full h-5"></Skeleton>
                <Skeleton className="w-full h-5"></Skeleton>
              </div>
            </div>
            <div className="flex-col  xl:flex col-span-2 hidden w-full gap-1">
              <Skeleton className="w-full h-5"></Skeleton>
              <Skeleton className="w-full h-5"></Skeleton>
            </div>
            <div className="flex-col  xl:flex col-span-2 hidden w-full gap-1">
              <Skeleton className="w-full h-5"></Skeleton>
              <Skeleton className="w-full h-5"></Skeleton>
            </div>
            <div className="flex-col  md:flex col-span-2 hidden gap-1">
              <Skeleton className="w-full h-5"></Skeleton>
              <Skeleton className="w-full h-5"></Skeleton>
            </div>
            <div className="flex-col  flex col-span-2 items-end md:items-start">
              <Skeleton className="flex items-center rounded-3xl px-7 w-32 h-12">
                <Skeleton className="bg-slate-300 w-full h-3" />
              </Skeleton>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default BankServiceSkeleton;
