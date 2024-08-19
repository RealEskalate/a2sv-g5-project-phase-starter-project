import BankServiceSkeleton from "@/components/AllSkeletons/bankServiceSkeleton/bankServiceSkeleton";
import ServiceSkeletonList from "@/components/AllSkeletons/serviceSkeleton/serviceSkeletonList";
import React from "react";

export default function page() {
  return (
    <div>
      <ServiceSkeletonList />
      <BankServiceSkeleton />
    </div>
  );
}
