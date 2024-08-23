import { Skeleton } from "@/components/ui/skeleton";
import SkeletonCard from "./skeletonCard";

const ListSkeleton = () => {
  return (
    <div className="flex flex-col items-start px-2 w-full md:w-3/5 space-y-4">
      <Skeleton className="mt-2 bg-slate-200 w-52 h-5" />
      <SkeletonCard />
      <SkeletonCard />
      <SkeletonCard />
    </div>
  );
};

export default ListSkeleton;
