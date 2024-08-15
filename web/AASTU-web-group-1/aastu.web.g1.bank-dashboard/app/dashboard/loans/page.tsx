import React from "react";
import Card from "./LoanComponents/Card";
import { TableDemo } from "./LoanComponents/Table/columns";
const Loans = () => {
  return (
    <div className="bg-slate-200 ">
      <div className="flex gap-2 overflow-x-scroll scrollbar-hidden">
        {([1, 2, 3, 4, 5, 6]).map(() => (
          <div className="my-2">
            <Card />
          </div>
          ))}
      </div>
      <div className="mt-5 p-2">
        <h1 className="text-2xl font-[600] text-[#333B69]">Active Loans Overview</h1>
        <TableDemo />
      </div>

    </div>

  );
};

export default Loans;
