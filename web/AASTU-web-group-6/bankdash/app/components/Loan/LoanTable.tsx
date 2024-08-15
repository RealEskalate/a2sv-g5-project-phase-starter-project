import React from "react";
import { ArrowUpCircleIcon } from "@heroicons/react/24/outline";
import LoanValue from "@/types/LoanValue";

interface LoanTableProps {
  loans: LoanValue[];
}

const LoanTable = ({ loans }: LoanTableProps) => {
  const totalLoanAmount = loans.reduce((sum, loan) => sum + loan.loanAmount, 0);
  const totalLeftToRepay = loans.reduce(
    (sum, loan) => sum + loan.leftToRepay,
    0
  );
  const totalInstallment = loans.reduce(
    (sum, loan) => sum + loan.Installment,
    0
  );
  return (
    <div className="space-y mb-8">
      <div className="w-full bg-white rounded-[25px] px-8 py-6">
        <h3 className="font-semibold text-[22px] text-[#343C6A]">
          Active Loans Overview
        </h3>
        <div className="overflow-x-auto custom-scrollbar">
          <table className="border-separate border-spacing-y-4 font-[16px] w-full detail-table  min-w-[900px]">
            <thead>
              <tr className="text-[#718EBF] text-left">
                <th>SL No</th>
                <th>Loan Money</th>
                <th>Left to repay</th>
                <th>Duration</th>
                <th>Interest rate</th>
                <th>Installment</th>
                <th>Repay</th>
              </tr>
            </thead>
            <tbody className="text-[#232323] p-8 space-y-4">
              {loans.map((detail, index) => (
                <tr key={index}>
                  <td>{detail.id}</td>
                  <td>${detail.loanAmount}</td>
                  <td>${detail.leftToRepay}</td>
                  <td>{detail.Duration} Months</td>
                  <td>{detail.interestRate}%</td>
                  <td>${detail.Installment}</td>
                  <td>
                    <p className="table-button">Replay</p>
                  </td>
                </tr>
              ))}
              <tr className="font-medium text-[#FE5C73] text-[16px]">
                <td>Total</td>
                <td>${totalLoanAmount}</td>
                <td>${totalLeftToRepay}</td>
                <td>-</td>
                <td>-</td>
                <td>${totalInstallment}</td>
                <td>-</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default LoanTable;
