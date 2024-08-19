import { Skeleton } from "@/components/ui/skeleton";

export function SkeletonCard() {
  return (
    <div className="flex flex-col md:flex-row gap-4 justify-center w-full">
      <div className="flex bg-[#FFFFFF] rounded-3xl md:w-1/3 p-5 mx-1">
        <div className="flex gap-3 items-center w-full">
          <Skeleton className="object-cover rounded-full w-11 h-11 md:w-12 md:h-12 lg:w-14 lg:h-14" />
          <div className="space-y-1 p-1 w-3/5">
            <Skeleton className="h-4 w-full" />
            <Skeleton className="h-4 w-full" />
          </div>
        </div>
      </div>
      <div className="flex bg-[#FFFFFF] rounded-3xl md:w-1/3 p-5 mx-1">
        <div className="flex gap-3 items-center w-full">
          <Skeleton className="rounded-full w-11 h-11 md:w-12 md:h-12 lg:w-14 lg:h-14" />
          <div className="space-y-1 p-1 w-3/5">
            <Skeleton className="h-4 w-full" />
            <Skeleton className="h-4 w-full" />
          </div>
        </div>
      </div>
      <div className="flex bg-[#FFFFFF] rounded-3xl md:w-1/3 p-5 mx-1">
        <div className="flex gap-3 items-center w-full">
          <Skeleton className="object-cover rounded-full w-11 h-11 md:w-12 md:h-12 lg:w-14 lg:h-14" />
          <div className="space-y-1 p-1 w-3/5">
            <Skeleton className="h-4 w-full" />
            <Skeleton className="h-4 w-full" />
          </div>
        </div>
      </div>
    </div>
  );
}
