import { Skeleton } from "@/components/ui/skeleton";
import React from "react";

const InputSkeleton = () => {
  return (
    <div className="w-full space-y-1 my-3">
      <Skeleton className="gray-dark w-24 h-5 bg-slate-200">
        <br />
      </Skeleton>
      <Skeleton className="w-full border-2 border-[#DFEAF2] p-5 py-6 rounded-xl bg-slate-200" />
    </div>
  );
};

export default InputSkeleton;
