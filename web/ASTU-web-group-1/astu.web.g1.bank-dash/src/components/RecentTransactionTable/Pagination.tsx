import React from "react";
import { ChevronLeft, ChevronRight } from "lucide-react";
import PageNumbers from "./PageNumber";

interface paginationProps {
  totalPages: number;
  currentPage: number;
  setCurrentPage: any;
}
const Pagination = ({
  totalPages,
  currentPage,
  setCurrentPage,
}: paginationProps) => {
  const handlePreviousPage = () => {
    if (currentPage > 0) {
      setCurrentPage(currentPage - 1);
    }
  };

  const handleNextPage = () => {
    if (currentPage < totalPages) {
      setCurrentPage(currentPage + 1);
    }
  };

  return (
    <nav className="flex items-center justify-end pt-4 text-blue-bright">
      <ul className="inline-flex items-center -space-x-px text-sm h-8 gap-[10px]">
        <li className="flex flex-row flex-wrap gap-3 items-center">
          <button
            onClick={handlePreviousPage}
            disabled={currentPage === 0}
            className="flex items-center justify-center px-3 h-8"
          >
            <ChevronLeft />
            Previous
          </button>
        </li>

        <PageNumbers
          totalPages={totalPages}
          currentPage={currentPage}
          setCurrentPage={setCurrentPage}
        />

        <li className="flex flex-row flex-wrap gap-3 items-center">
          <button
            onClick={handleNextPage}
            disabled={currentPage === totalPages - 1}
            className="flex items-center justify-center px-3 h-8"
          >
            Next
            <ChevronRight />
          </button>
        </li>
      </ul>
    </nav>
  );
};

export default Pagination;
