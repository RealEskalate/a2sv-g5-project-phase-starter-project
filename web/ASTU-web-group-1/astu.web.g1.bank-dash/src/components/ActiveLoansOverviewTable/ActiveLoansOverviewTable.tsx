"use client";
import { ActiveLoanDataType } from "@/types/active-loan.types";
import TableButton from "../TableButton/TableButton";
import { useGetAllActiveLoansQuery } from "@/lib/redux/slices/activeLoanSlice";

const ActiveLoansOverviewTable = () => {
  const { data, isLoading } = useGetAllActiveLoansQuery();

  const allData: ActiveLoanDataType[] | null = data?.data?.filter(
    (data) => data.type === "approved"
  );
  console.log("the table data is ", allData);
  return (
    <div className="flex flex-col gap-4">
      <h1 className="text-16px md:text-15px xl:text-18px text-[#333B69] font-semibold">
        Active Loans Overview
      </h1>
      <div className="relative overflow-x-auto bg-white px-4 md:px-6 pt-5 md:pt-6 rounded-2xl md:rounded-2xl">
        <table className="bg-white px-5 lg:px-11 w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
          <thead className=" text-12px md:text-15px font-Lato font-normal text-blue-steel bg-white border-b">
            <tr className="">
              <th scope="col" className="hidden md:table-cell pb-2">
                SL No
              </th>
              <th scope="col" className=" pb-2">
                Loan Money
              </th>
              <th scope="col" className=" pb-2">
                Left to replay
              </th>
              <th scope="col" className="hidden lg:table-cell pb-2">
                Duration
              </th>
              <th scope="col" className="hidden min-[900px]:table-cell pb-2">
                Interest rate
              </th>
              <th scope="col" className="hidden min-[900px]:table-cell pb-2">
                Installment
              </th>
              <th scope="col" className=" pb-2 w-fit">
                Repay
              </th>
            </tr>
          </thead>
          <tbody className="text-12px xl:text-15px text-gray-dark cursor-pointer  hover:bg-gray-100 dark:hover:bg-gray-700">
            {allData?.map((data, index) => (
              <tr
                key={index}
                className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-100 dark:hover:bg-gray-700"
              >
                <td className="hidden md:table-cell py-3">
                  {data.serialNumber}
                </td>
                <td className="py-3">{data.loanAmount}</td>
                <td className="py-3">{data.amountLeftToRepay}</td>
                <td className="hidden lg:table-cell py-3">{data.duration}</td>
                <td className="hidden min-[900px]:table-cell py-3">
                  {data.interestRate}
                </td>
                <td className="hidden min-[900px]:table-cell py-3">
                  {data.installment}
                </td>
                <td className="py-3 w-24 md:w-32 ">
                  <TableButton text="Repay" classname="px-6" />
                </td>
              </tr>
            ))}
            <tr className="bg-white align-bottom text-candyPink font-medium dark:bg-gray-800 dark:border-gray-700">
              <td className="hidden md:table-cell py-3 md:py-6">Total</td>
              <td className="py-3 md:py-6 flex flex-col">
                <span className="md:hidden">Total</span>
                $125,000
              </td>
              <td className="py-3 md:py-6">$750,000</td>
              <td className="hidden md:table-cell py-3 md:py-6"></td>
              <td className="hidden md:table-cell py-3 md:py-6"></td>
              <td className="hidden min-[900px]:table-cell py-3 md:py-6">
                $50,000 / month
              </td>
              <td className="py-3 md:py-6 whitespace-nowrap"></td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default ActiveLoansOverviewTable;
