"use client";
import React from "react";
import { Skeleton } from "@/components/ui/skeleton";

const CardSkeleton = () => {
  return (
    <div>
      <div
        className={`bg-white w-[280px] h-[175px] rounded-3xl flex flex-col justify-between`}
      >
        <div className="flex flex-col  px-4  pt-4 h-full">
          <div className="flex justify-between pt-2">
            <div className="flex flex-col ">
              <Skeleton className="bg-slate-200 w-24 h-4 mb-1"></Skeleton>
              <Skeleton className="bg-slate-200 w-24 h-4"></Skeleton>
            </div>
            <Skeleton className="w-12 h-8 items-center  bg-slate-200" />
          </div>
          <div className="flex flex-1 flex-col justify-center pt-3">
            <Skeleton className="bg-slate-200  w-full h-4 mb-1"></Skeleton>
            <Skeleton className="bg-slate-200 w-full h-5 "></Skeleton>
          </div>

          <div className={` rounded-b-3xl flex justify-between py-3 `}>
            <Skeleton className="bg-slate-200 w-full h-8"></Skeleton>
          </div>
        </div>
      </div>
    </div>
  );
};

export default CardSkeleton;
