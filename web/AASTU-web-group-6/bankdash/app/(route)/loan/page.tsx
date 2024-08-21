import LoanTable from "@/app/components/Loan/LoanTable";
import React from "react";
import { LoanType } from "@/types/LoanValue";
import Card from "../../components/Accounts/account";
// import loanApi from "@/app/Services/api/loanApi";
import { getServerSession } from "next-auth";
import { options } from "@/app/api/auth/[...nextauth]/options";

const Loan = async () => {
  const session = await getServerSession(options);
  const accessToken = session?.accessToken as string;
  // console.log(accessToken, "Server Token");
  // const loanData = await loanApi.getLoan(accessToken);
  const loanData = [];
  return (
    <div className="px-5 space-y-4 mt-4 w-full h-screen">
      <div className="flex flex-col xxs:overflow-x-auto md:overflow-hidden lg:flex-row gap-6 xl:gap-7">
        <div className="flex scrollbar-hide overflow-x-scroll lg:w-[100%] gap-4 xl:gap-7">
          <Card
            title="Personal Loans"
            amount="$12,750"
            color="#FFF5D9"
            icon="/assets/money-tag 1.svg"
            width="w-full"
          />
          <Card
            title="My Balance"
            amount="$5,600"
            color="#E7EDFF"
            icon="/assets/expense.svg"
            width="w-full"
          />
          <Card
            title="My Balance"
            amount="$500,000"
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
      <LoanTable loans={loanData} />
    </div>
  );
};

export default Loan;
