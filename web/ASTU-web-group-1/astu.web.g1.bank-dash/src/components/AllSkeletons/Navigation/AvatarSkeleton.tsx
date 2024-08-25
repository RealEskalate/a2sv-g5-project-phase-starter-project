import { Skeleton } from "@/components/ui/skeleton";

const AvatarSkeleton = () => {
  return (
    <Skeleton className="object-cover rounded-full w-[50px] h-[50px] md:w-[50px] md:h-[50px] " />
  );
};

export default AvatarSkeleton;
