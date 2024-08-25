import ActiveLoanSkeleton from "@/components/AllSkeletons/ActiveLoansSkeleton/ActiveLoanSkeleton";
import LoanSkeletons from "@/components/AllSkeletons/loanSkeleton/loanSkeletons";
import React from "react";

export default function page() {
  return (
    <div className="flex flex-col gap-5">
      <LoanSkeletons />
      <ActiveLoanSkeleton />
    </div>
  );
}
