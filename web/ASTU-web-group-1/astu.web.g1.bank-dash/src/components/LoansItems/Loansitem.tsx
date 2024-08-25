"use client";
import React from "react";
import LoanCard from "./LoanCard";
import { useGetDetailActiveLoansQuery } from "@/lib/redux/slices/activeLoanSlice";
import LoanSkeletons from "../AllSkeletons/loanSkeleton/loanSkeletons";

const Loansitem = () => {
  const { data, isLoading } = useGetDetailActiveLoansQuery();
  const allData = data?.data;
  if (isLoading) {
    return <LoanSkeletons />;
  }
  return (
    <div
      className="flex overflow-x-auto justify-around overflow-clip whitespace-nowrap w-full"
      style={{
        scrollbarWidth: "none",
        msOverflowStyle: "none",
      }}
    >
      <LoanCard
        image={"/assets/icons/personal.svg"}
        name={"Personal Loans"}
        amount={`$${allData?.personalLoan.toLocaleString()}`}
      />
      <LoanCard
        image={"/assets/icons/bag.svg"}
        name={"Corporate Loans"}
        amount={`$${allData?.corporateLoan.toLocaleString()}`}
      />
      <LoanCard
        image={"/assets/icons/businesstrack.svg"}
        name={"Business Loans"}
        amount={`$${allData?.businessLoan.toLocaleString()}`}
      />
      <LoanCard
        image={"/assets/icons/customLoan.svg"}
        name={"Custom Loans"}
        amount={`Choose Money`}
      />
    </div>
  );
};

export default Loansitem;
