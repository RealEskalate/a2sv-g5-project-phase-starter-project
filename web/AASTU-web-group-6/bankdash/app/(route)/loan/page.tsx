import LoanTable from "@/app/components/Loan/LoanTable";
import React from "react";
import LoanValue from "@/types/LoanValue";

const loans: LoanValue[] = [
  {
    description: "Personal Loan",
    id: 1,
    loanAmount: 20000,
    leftToRepay: 14500,
    Duration: 8,
    interestRate: 12,
    Installment: 2500,
  },
  {
    description: "Personal Loan",
    id: 2,
    loanAmount: 20000,
    leftToRepay: 14500,
    Duration: 8,
    interestRate: 12,
    Installment: 2500,
  },
];

const Loan = () => {
  return (
    <div>
      <LoanTable loans={loans} />
    </div>
  );
};

export default Loan;
