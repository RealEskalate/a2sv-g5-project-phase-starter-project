import React from "react";
import { ChevronLeftIcon, ChevronRightIcon } from "@heroicons/react/24/outline";

const Pagination = () => {
  return (
    <div className="flex justify-end gap-6 text-[#1814F3] font-[15px] font-medium content-end text-right w-full text-[15px] font-medium">
      <button className="flex items-center gap-1">
        <ChevronLeftIcon className="w-4 h-4" />
        Prev
      </button>
      <button className="box-content border bg-[#1814F3] px-4 py-3 text-white rounded-[10px]">
        1
      </button>
      <button>2</button>
      <button>3</button>
      <button>4</button>
      <button className="flex items-center gap-1">
        Next
        <ChevronRightIcon className="w-4 h-4" />
      </button>
    </div>
  );
};

export default Pagination;
