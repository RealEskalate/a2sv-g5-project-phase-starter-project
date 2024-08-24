import { Skeleton } from "@/components/ui/skeleton";
import React from "react";
import InputSkeleton from "./inputSkeleton";

const EditProfileSkeleton = () => {
  return (
    <div className="w-full">
      <form action="">
        <div className="flex flex-col md:flex-row md:space-x-5 bg-white">
          <InputSkeleton />
          <InputSkeleton />
        </div>
        <div className="flex flex-col md:flex-row md:space-x-5 bg-white">
          <InputSkeleton />
          <InputSkeleton />
        </div>
        <div className="flex flex-col md:flex-row md:space-x-5 bg-white">
          <InputSkeleton />
          <InputSkeleton />
        </div>
        <div className="flex flex-col md:flex-row md:space-x-5 bg-white">
          <InputSkeleton />
          <InputSkeleton />
        </div>
        <div className="flex flex-col md:flex-row md:space-x-5 bg-white">
          <InputSkeleton />
          <InputSkeleton />
        </div>
        <div className="flex justify-end bg-white">
          <Skeleton className="bg-slate-200 px-16 py-5 rounded-lg w-full md:w-auto mt-4" />
        </div>
      </form>
    </div>
  );
};

export default EditProfileSkeleton;
