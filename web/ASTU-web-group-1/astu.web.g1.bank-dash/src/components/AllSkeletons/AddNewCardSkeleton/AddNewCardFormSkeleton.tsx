"use client";
import { Skeleton } from "@/components/ui/skeleton";
import React from "react";

const AddNewCardFormSkeleton = () => {
  return (
    <div className="space-y-6">
      <div className="flex flex-col md:flex-row md:gap-6 ">
        <Skeleton className="w-3/4 h-16 bg-slate-200" />
        <Skeleton className="w-3/4 h-16 bg-slate-200" />
      </div>
      <div className="flex flex-col md:flex-row md:gap-6 ">
        <Skeleton className="w-3/4 h-16 bg-slate-200" />
        <Skeleton className="w-3/4 h-16 bg-slate-200" />
      </div>
      <Skeleton className="bg-slate-200 px-10 py-3 rounded-lg w-full md:w-1/4 h-8 mt-4" />
    </div>
  );
};

export default AddNewCardFormSkeleton;
