import { Skeleton } from "@/components/ui/skeleton";
import React from "react";

const SkeletonCard = () => {
  return (
    <tr>
      <td className="font-[400] font-Inter text-sm py-1.5 pl-4">
        <Skeleton className="h-4 w-10" />
      </td>
      <td className="font-[400] font-Inter text-sm py-1.5 pl-4 truncate">
        <Skeleton className="h-4 w-10" />
      </td>
      <td className="font-[400] font-Inter text-sm py-1.5 pl-4 truncate">
        <Skeleton className="h-4 w-10" />
      </td>
      <td>
        <Skeleton className="h-4 w-10" />
      </td>
    </tr>
  );
};

export default SkeletonCard;
