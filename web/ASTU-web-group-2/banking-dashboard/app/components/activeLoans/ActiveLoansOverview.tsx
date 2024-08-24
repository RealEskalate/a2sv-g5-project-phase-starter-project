/* eslint-disable jsx-a11y/alt-text */
/* eslint-disable @next/next/no-img-element */
"use client";
import React, { useEffect } from "react";
import { defaultloans, loan } from "./ActiveLoansItems";
import { useSession } from "next-auth/react";
import { useGetMyLoanServiceQuery } from "@/lib/service/LoanService";
import LoanTableSkeleton from "./LoanTableSkeleton";
import ErrorImage from "../Error/ErrorImage";

const ActiveLoansOverview = () => {
  const { data: session } = useSession();
  const accessToken = session?.user.accessToken!;
  useEffect(() => {
    console.log("data", session);
  }, [session]);
  const { data, isLoading, isError, isSuccess, error } =
    useGetMyLoanServiceQuery(accessToken);
  let loans: loan[] = [];

  if (isLoading) {
    return (
      <div>
        <LoanTableSkeleton />
      </div>
    );
  }
  if (isError) {
    console.log("errorrrrs");
    console.log(error);
    return (
      <div>
        <ErrorImage />
      </div>
    );
  }
  if (isSuccess) {
    console.log("data", data);
    loans = data.data.content;
  }

  return (
    <div className="bg-white rounded-3xl w-full h-max-[466px] sm:h-max-[500px] md:h-max-[625px]  p-3">
      {loans.length === 0 ? (
        <div className="flex justify-center">
          <img
            src="/assets/bankService/empty-image.png"
            className="w-fit h-fit"
          />
        </div>
      ) : (
        <table className="min-w-full divide-y">
          <thead>
            <tr>
              <th
                scope="col"
                className="px-4 py-3 text-start text-[#718EBF] hidden sm:table-cell  sm:text-[12px] md:text-[16px] font-medium"
              >
                SL No
              </th>
              <th
                scope="col"
                className="p-3 text-start text-[#718EBF] sm:text-[12px] md:text-[16px] font-medium"
              >
                Loan Money
              </th>
              <th
                scope="col"
                className="p-3 text-start text-[#718EBF] sm:text-[12px] md:text-[16px] font-medium"
              >
                Left to repay
              </th>
              <th
                scope="col"
                className="p-3 text-start text-[#718EBF] hidden sm:table-cell  sm:text-[12px] md:text-[16px] font-medium"
              >
                Duration
              </th>
              <th
                scope="col"
                className="p-3 text-start text-[#718EBF] hidden sm:table-cell  sm:text-[12px] md:text-[16px] font-medium"
              >
                Interest rate
              </th>
              <th
                scope="col"
                className="p-3 text-start text-[#718EBF] hidden sm:table-cell  sm:text-[12px] md:text-[16px] font-medium"
              >
                Installment
              </th>
              <th
                scope="col"
                className="p-3 text-start text-[#718EBF] sm:text-[12px] md:text-[16px] font-medium"
              >
                Repay
              </th>
            </tr>
          </thead>
          <tbody className="divide-y divide-gray-200 dark:divide-neutral-700">
            {loans.map((loan, index) => (
              <tr key={index}>
                {index === loans.length - 1 ? (
                  <>
                    <td className="px-4 py-3 hidden sm:table-cell text-[#FE5C73] font-medium sm:text-[12px] md:text-[16px]">
                      Total.
                    </td>
                    <td className="p-3 text-[#FE5C73] font-medium sm:text-[12px] md:text-[16px]">
                      <span className="block sm:hidden">Total</span>
                      <span>${loan.loanAmount}</span>
                    </td>
                    <td className="p-3 text-[#FE5C73] font-medium sm:text-[12px] md:text-[16px]">
                      ${loan.amountLeftToRepay}
                    </td>
                    <td className="p-3 hidden sm:table-cell  sm:text-[12px] md:text-[16px]">
                      {loan.duration}
                    </td>
                    <td className="p-3 hidden sm:table-cell  sm:text-[12px] md:text-[16px]">
                      {loan.interestRate}
                    </td>
                    <td className="p-3 hidden sm:table-cell  text-[#FE5C73] font-medium sm:text-[12px] md:text-[16px]">
                      {loan.installment} / Month
                    </td>
                  </>
                ) : (
                  <>
                    <td className="px-4 py-2 hidden sm:table-cell   text-[#232323] align-middle font-normal sm:text-[12px] md:text-[16px]">
                      {index + 1 < 10 ? `0${index + 1}.` : index + 1}
                    </td>
                    <td className="p-3 text-[#232323] font-normal sm:text-[12px] md:text-[16px]">
                      ${loan.loanAmount}
                    </td>
                    <td className="p-3 text-[#232323] font-normal sm:text-[12px] md:text-[16px]">
                      ${loan.amountLeftToRepay}
                    </td>
                    <td className="p-3 hidden sm:table-cell  text-[#232323] font-normal sm:text-[12px] md:text-[16px]">
                      {loan.duration} Months
                    </td>
                    <td className="p-3 hidden sm:table-cell  text-[#232323] font-normal sm:text-[12px] md:text-[16px]">
                      {loan.interestRate} %
                    </td>
                    <td className="p-3 hidden sm:table-cell  text-[#232323] font-normal sm:text-[12px] md:text-[16px]">
                      {loan.installment} / Month
                    </td>
                    <td className="py-3 text-[#232323] font-normal sm:text-[12px] md:text-[16px]">
                      <button className="sm:font-normal md:font-medium border-[#1814F3] text-[#1814F3] sm:text-[rgb(35,35,35)] sm:border-[rgb(35,35,35)] border-2 rounded-full sm:hover:border-[#1814F3] sm:hover:text-[#1814F3] py-[5px] px-[15px]  sm:py-[8px] sm:px-[22px] md:px-[27px] md:py-[8px]">
                        Repay
                      </button>
                    </td>
                  </>
                )}
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
};

export default ActiveLoansOverview;
