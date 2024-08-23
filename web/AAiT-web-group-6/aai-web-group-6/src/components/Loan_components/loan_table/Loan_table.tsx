import React from "react";
import { headers, data } from "../../../constants/loan_constants/table_data";

const Loan_table = () => {
  return (
    <div className="overflow-x-auto">
      <div className="bg-white rounded-[15px] w-full max-w-full h-max px-4">
        <table className="min-w-full">
          <thead className="">
            <tr>
              {headers.map((header, index) => (
                <td
                  key={index}
                  className="py-3 pb-2 px-4 md:px-8 border-b border-gray-400 text-left text-[16px] sm:text-[16px] text-[#718EBF]"
                >
                  {header}
                </td>
              ))}
            </tr>
          </thead>
          <tbody>
            {data.map((row, rowIndex) => (
              <tr key={rowIndex} className="text-[8pt] sm:text-[16px]">
                <td className="py-2 px-4 md:py-4 md:px-8 border-b border-gray-200">{row.id}</td>
                <td className="py-2 px-4 md:py-2 md:px-8 border-b border-gray-200">{row.loanMoney}</td>
                <td className="py-2 px-4 md:py-2 md:px-8 border-b border-gray-200">{row.leftToRepay}</td>
                <td className="py-2 px-4 md:py-2 md:px-8 border-b border-gray-200">{row.duration}</td>
                <td className="py-2 px-4 md:py-2 md:px-8 border-b border-gray-200">{row.interestRate}</td>
                <td className="py-2 px-4 md:py-2 md:px-8 border-b border-gray-200">{row.installment}</td>
                <td className="py-2 px-4 md:py-2 md:px-8 border-b border-gray-200">
                  <button className="border border-black px-3 sm:px-5 py-1 rounded-[30px] hover:border-blue-800">
                    {row.repay}
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
          <tfoot>
            <tr className="text-[8pt] sm:text-[16px] text-red-500">
              <td className="py-2 px-4 md:py-4 md:px-8 border-t border-gray-200">Total</td>
              <td className="py-2 px-4 md:py-4 md:px-8 border-t border-gray-200">$33,000</td> 
              <td className="py-2 px-4 md:py-4 md:px-8 border-t border-gray-200">$17,000</td> 
              <td className="py-2 px-4 md:py-4 md:px-8 border-t border-gray-200"></td>
              <td className="py-2 px-4 md:py-2 md:px-8 border-t border-gray-200"></td>
              <td className="py-2 px-4 md:py-4 md:px-8 border-t border-gray-200">$40,000</td>
              <td className="py-2 px-4 md:py-2 md:px-8 border-t border-gray-200"></td>
            </tr>
          </tfoot>
        </table>
      </div>
    </div>
  );
};

export default Loan_table;
