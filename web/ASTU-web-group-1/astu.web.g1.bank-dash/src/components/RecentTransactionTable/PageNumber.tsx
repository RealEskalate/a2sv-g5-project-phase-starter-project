import React from 'react';

interface paginationProps {
    totalPages:number;
    currentPage: number;
    setCurrentPage:any;
}
const PageNumbers = ({ totalPages, currentPage, setCurrentPage }:paginationProps) => {
  const renderPageNumbers = () => {
    const pages = [];
    const maxVisiblePages = 3;
    const startPage = Math.max(Math.min(currentPage, totalPages - maxVisiblePages + 1),0);
    const endPage = Math.min(startPage + maxVisiblePages - 1, totalPages);
    console.log("the total pages are",totalPages,endPage,startPage)
    for (let i = startPage; i < endPage; i++) {
      pages.push(
        <li key={i}>
          <button
            onClick={() => setCurrentPage(i)}
            className={`flex items-center justify-center px-3 h-8 ${
              currentPage === i && 'bg-blue-bright text-white rounded-xl'
            }`}
          >
            {i+1}
          </button>
        </li>
      );
    }

    if (endPage < totalPages) {
      pages.push(
        <li key="end-ellipsis">
          <span className="px-3">...</span>
        </li>
      );
      pages.push(
        <li key={totalPages}>
          <button
            onClick={() => setCurrentPage(totalPages)}
            className={`flex items-center justify-center px-3 h-8 ${
              currentPage === totalPages && 'bg-blue-bright text-white rounded-xl'
            }`}
          >
            {totalPages}
          </button>
        </li>
      );
    }

    return pages;
  };

  return <ul className="inline-flex items-center -space-x-px text-sm h-8 gap-[10px]">{renderPageNumbers()}</ul>;
};

export default PageNumbers;
