import React from "react";
import BarGraph from "@/app/components/Transaction/BarGraph";
import Recent from "@/app/components/Transaction/Recent";
import Pagination from "@/app/components/Transaction/Pagination";

const Transaction = () => {
  return (
    <div className="space-y-6 bg-[#F5F7FA] px-10">
      <div>
        <BarGraph />
      </div>
      <Recent />
      <Pagination />
    </div>
  );
};

export default Transaction;
