import { Skeleton } from "../../ui/skeleton";

const ChartSkeleton = () => {
  return (
    <div className="w-full mr-5">
      <div className="w-full">
        <Skeleton className="h-5 w-52 bg-slate-200 mb-6 py-2 rounded-lg" />
        <div className="bg-slate-50 space-y-3 rounded-3xl md:px-10 md:py-10 p-4 h-60 lg:h-72">
          <div className="flex items-end justify-between">
            <div className="flex items-baseline space-x-1">
              <Skeleton className="h-20 w-6 rounded-lg bg-slate-300" />
              <Skeleton className="h-40 w-6 rounded-lg bg-slate-300" />
            </div>
            <div className="flex items-baseline space-x-1">
              <Skeleton className="h-20 w-6 rounded-lg bg-slate-300" />
              <Skeleton className="h-32 w-6 rounded-lg bg-slate-300" />
            </div>
            <div className="flex items-baseline space-x-1">
              <Skeleton className="h-24 w-6 rounded-lg bg-slate-300" />
              <Skeleton className="h-10 w-6 rounded-lg bg-slate-300" />
            </div>
            <div className="flex items-baseline space-x-7">
              <Skeleton className="h-36 w-6 rounded-lg bg-slate-300" />
              <Skeleton className="h-28 w-6 rounded-lg bg-slate-300" />
            </div>
            <div className="flex items-baseline space-x-1">
              <Skeleton className="h-36 w-6 rounded-lg bg-slate-300" />
              <Skeleton className="h-28 w-6 rounded-lg bg-slate-300" />
            </div>
            <div className="flex items-baseline space-x-1">
              <Skeleton className="h-32 w-6 rounded-lg bg-slate-300" />
              <Skeleton className="h-40 w-6 rounded-lg bg-slate-300" />
            </div>
          </div>
          <div className="flex justify-between">
            <Skeleton className="h-3 w-20 bg-slate-200" />
            <Skeleton className="h-3 w-20 bg-slate-200" />
            <Skeleton className="h-3 w-20 bg-slate-200" />
            <Skeleton className="h-3 w-20 bg-slate-200" />
            <Skeleton className="h-3 w-20 bg-slate-200" />
            <Skeleton className="h-3 w-20 bg-slate-200" />
          </div>
        </div>
      </div>
    </div>
  );
};

export default ChartSkeleton;
