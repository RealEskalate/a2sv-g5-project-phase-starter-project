import React, { useState, useEffect } from "react";
import { ChevronLeftIcon, ChevronRightIcon } from "@heroicons/react/24/outline";

const Pagination = ({
  updatePage,
  start,
}: {
  updatePage: (newPage: number) => void;
  start: boolean;
}) => {
  const [page, setPage] = useState(0);
  useEffect(() => {
    setPage(0); // Reset the page to 0 when start is true
  }, [start]);

  const handleNext = () => {
    const nextPage = page + 1;
    setPage(nextPage);
    updatePage(nextPage);
  };

  const handlePrev = () => {
    if (page > 0) {
      const prevPage = page - 1;
      setPage(prevPage);
      updatePage(prevPage);
    }
  };

  return (
    <div className="flex justify-end gap-6 text-[#1814F3] p-6 content-end text-right w-full text-[15px] font-medium">
      <button
        className={`flex items-center gap-1 ${
          page === 0 ? "opacity-50 cursor-not-allowed" : ""
        }`}
        onClick={handlePrev}
        disabled={page === 0}
      >
        <ChevronLeftIcon className="w-4 h-4" />
        Prev
      </button>
      {/* Render page numbers */}
      {[...Array(4)].map((_, index) => (
        <button
          key={index}
          className={`${
            page === index
              ? "box-content border bg-[#1814F3] px-4 py-3 text-white rounded-[10px]"
              : "dark:text-gray-300"
          }`}
          onClick={() => updatePage(index + 1)}
        >
          {index + 1}
        </button>
      ))}
      <button className="flex items-center gap-1" onClick={handleNext}>
        Next
        <ChevronRightIcon className="w-4 h-4" />
      </button>
    </div>
  );
};

export default Pagination;
