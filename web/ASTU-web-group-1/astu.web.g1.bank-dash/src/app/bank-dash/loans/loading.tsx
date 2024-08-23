import ActiveLoanSkeleton from "@/components/AllSkeletons/ActiveSkeleton/ActiveLoanSkeleton";
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
