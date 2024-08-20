import React from "react";
import { ArrowUpCircleIcon } from "@heroicons/react/24/outline";
import { LoanType } from "@/types/LoanValue";

interface LoanTableProps {
  loans: LoanType[];
}

const LoanTable = ({ loans }: LoanTableProps) => {
  const totalLoanAmount = loans.reduce((sum, loan) => sum + loan.loanAmount, 0);
  const totalLeftToRepay = loans.reduce(
    (sum, loan) => sum + loan.amountLeftToRepay,
    0
  );
  const totalInstallment = loans.reduce(
    (sum, loan) => sum + loan.installment,
    0
  );
  return (
    <div className="space-y mb-8">
      <div className="w-full bg-white rounded-[25px] px-8 py-6 dark:bg-[#232328]">
        <h3 className="font-semibold text-[22px] text-[#343C6A] dark:text-gray-300">
          Active Loans Overview
        </h3>
        <div className="overflow-x-auto custom-scrollbar dark:bg-[#232328] w-full">
          <table className="border-separate border-spacing-y-4 font-[16px] w-full detail-table  min-w-[1000px] sm:min-w-full">
            <thead>
              <tr className="text-[#718EBF] text-left dark:text-gray-200">
                <th className="hidden lg:table-cell">SL No</th>
                <th className="">Loan Money</th>
                <th className="">Left to repay</th>
                <th className="hidden lg:table-cell">Duration</th>
                <th className="hidden md:table-cell">Interest rate</th>
                <th className="hidden md:table-cell">Installment</th>
                <th>Repay</th>
              </tr>
            </thead>
            <tbody className="text-[#232323] p-8 space-y-4 dark:text-gray-300">
              {loans.map((detail, index) => (
                <tr key={index}>
                  <td className="hidden lg:table-cell">
                    {detail.serialNumber}
                  </td>
                  <td>${detail.loanAmount}</td>
                  <td>${detail.amountLeftToRepay}</td>
                  <td className="hidden lg:table-cell">
                    {detail.duration} Months
                  </td>
                  <td className="hidden md:table-cell">
                    {detail.interestRate}%
                  </td>
                  <td className="hidden md:table-cell">
                    ${detail.installment}
                  </td>
                  <td>
                    <p className="table-button">Replay</p>
                  </td>
                </tr>
              ))}
              <tr className="font-medium text-[#FE5C73] text-[16px]">
                <td className="hidden lg:table-cell">Total</td>
                <td>${totalLoanAmount}</td>
                <td>${totalLeftToRepay}</td>
                <td className="hidden lg:table-cell">-</td>
                <td className="hidden md:table-cell">-</td>
                <td className="hidden md:table-cell">${totalInstallment}</td>
                <td className="hidden md:table-cell">-</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default LoanTable;
