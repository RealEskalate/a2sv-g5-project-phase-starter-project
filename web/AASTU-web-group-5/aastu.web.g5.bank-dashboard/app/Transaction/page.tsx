import React from "react";
import { BarChartComponent } from "./components/BarChartComponent";
import { TableComponent } from "./components/TableComponent";
import dummyData from "./components/dummyData"; // Adjust the path as needed
import columns from "./components/columns"; // Adjust the path as needed
// import { CardsComponent } from "./components/CardsComponent"; // Assuming you have a Cards component

function Transactions() {
  return (
    <div className="flex flex-col gap-6 p-6">
      {/* First Row: Cards and BarChart */}
      <div className="flex flex-col md:flex-row gap-6">
        <div className="flex-1">
          <div className="flex justify-between items-center mb-4">
            <h2 className="text-lg font-semibold font-Inter text-[#343C6A]">
              My Cards
            </h2>
            <button className="text-[#343C6A] font-Inter font-medium">
              + Add Card
            </button>
          </div>
          {/* <CardsComponent /> */}
          <p>card 1</p>
          <p>card 2</p>
        </div>
        <div className="flex-1 md:w-2/5">
          <h2 className="text-lg font-semibold mb-4 font-Inter text-[#343C6A]">
            My Expenses
          </h2>
          <BarChartComponent />
        </div>
      </div>

      {/* Second Row: Table */}
      <div className="flex-1 w-full">
        <h2 className="text-lg font-semibold mb-4 font-Inter text-[#343C6A]">
          Recent Transactions
        </h2>
        <TableComponent columns={columns} data={dummyData} />
      </div>
    </div>
  );
}

export default Transactions;
