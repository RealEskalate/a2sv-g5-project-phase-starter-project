import LoanTable from "@/app/components/Loan/LoanTable";
import React from "react";
import { LoanType } from "@/types/LoanValue";
import Card from "../../components/Accounts/account";
import loanApi from "@/app/Services/api/loanApi";


const Loan = async () => {
  const loanData = await loanApi.getLoan(
    "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJ0bmFob20iLCJpYXQiOjE3MjQxNDYwNjgsImV4cCI6MTcyNDIzMjQ2OH0.Y00dc0ACMvkHK5ZYWVsBxK5lk2l5VB_6xnnFAMXrkMjRjl2jxEZHglllGcw_S61p"
  );

  return (
    <div className="px-5 space-y-4 mt-4 w-full h-screen">
      <div className="flex flex-col lg:flex-row gap-7">
        <div className="flex lg:w-[45%] gap-7">
          <Card
            title="Personal Loans"
            amount="$50,000"
            color="#E7EDFF"
            icon="/assets/user-blue.svg"
            width="w-[45%]"
          />
          <Card
            title="My Balance"
            amount="$100,000"
            color="#FFF5D9"
            icon="/assets/briefcase 1.svg"
            width="w-[45%]"
          />
        </div>
        <div className="flex lg:w-[45%] gap-7">
          <Card
            title="My Balance"
            amount="$500,000"
            color="#FFE0EB"
            icon="/assets/growth.svg"
            width="w-[45%]"
          />
          <Card
            title="My Balance"
            amount="$12,750"
            color="#DCFAF8"
            icon="/assets/custom.svg"
            width="w-[45%]"
          />
        </div>
      </div>
      <LoanTable loans={loanData} />
    </div>
  );
};

export default Loan;
