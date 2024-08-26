"use client";
import LoanTable from "@/app/components/Loan/LoanTable";
import React, { useEffect, useState } from "react";
import { ApiResponse, LoanDetail, LoanType } from "@/types/LoanValue";
import Card from "../../components/Accounts/account";
import loanApi from "@/app/Services/api/loanApi";
import { useSession } from "next-auth/react";
import ApplyLoan from "@/app/components/Loan/ApplyForLoan";

const Loan = () => {
  const { data: session } = useSession();
  const accessToken = session?.accessToken as string;
  const [loading, setLoading] = useState(true);
  const [getLoan, setGetLoan] = useState<LoanDetail | null>(null);

  const fetchLoan = async () => {
    while (!accessToken) {
      await new Promise((resolve) => setTimeout(resolve, 150));
    }
    const loanData = await loanApi.detailData(accessToken);
    // console.log(loanData);
    setGetLoan(loanData.content);

    setLoading(false);
  };

  useEffect(() => {
    setLoading(true);
    fetchLoan();
    console.log(getLoan, "--");
  }, [accessToken]);
  return (
    <div className="px-5 space-y-4 mt-4 w-full h-screen">
      <div className="flex flex-col xxs:overflow-x-auto md:overflow-hidden lg:flex-row gap-6 xl:gap-7">
        <div className="flex scrollbar-hide overflow-x-scroll lg:w-[100%] gap-4 xl:gap-7">
          <Card
            title="Personal Loans"
            amount={getLoan?.personalLoan ?? null}
            color="#FFF5D9"
            icon="/assets/money-tag 1.svg"
            width="w-full"
          />
          <Card
            title="Corporate Loan"
            amount={getLoan?.corporateLoan ?? null}
            color="#E7EDFF"
            icon="/assets/expense.svg"
            width="w-full"
          />
          <Card
            title="Buissness Loan"
            amount={getLoan?.businessLoan ?? null}
            color="#FFE0EB"
            icon="/assets/growth.svg"
            width="w-full"
          />
          <Card
            title="My Balance"
            amount="$12,750"
            color="#DCFAF8"
            icon="/assets/custom.svg"
            width="w-full"
          />
        </div>
      </div>

      <LoanTable />
    </div>
  );
};

export default Loan;
