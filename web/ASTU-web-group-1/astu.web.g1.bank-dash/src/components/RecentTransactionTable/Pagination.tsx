import React from 'react';
import { ChevronLeft, ChevronRight } from 'lucide-react';
import PageNumbers from './PageNumber';

interface paginationProps {
    totalPages:number;
    currentPage: number;
    setCurrentPage:any;
}
const Pagination = ({ totalPages, currentPage, setCurrentPage }:paginationProps) => {
    const handlePreviousPage = () => {
        if (currentPage > 1) {
          setCurrentPage(currentPage - 1);
        }
      };
    
      const handleNextPage = () => {
        if (currentPage < totalPages) {
          setCurrentPage(currentPage + 1);
        }
      };
    
      const renderPageNumbers = () => {
        const pages = [];
        const maxVisiblePages = 4; // Adjust this to change the number of visible pages before the ellipsis
    
        const startPage = Math.min(currentPage, totalPages - maxVisiblePages + 1);
        const endPage = Math.min(startPage + maxVisiblePages - 1, totalPages);
    
        // Render the range of pages
        for (let i = startPage; i <= endPage; i++) {
          pages.push(
            <li key={i}>
              <button
                onClick={() => setCurrentPage(i)}
                className={`flex items-center justify-center px-3 h-8 ${
                  currentPage === i && 'bg-blue-bright text-white rounded-xl'
                }`}
              >
                {i}
              </button>
            </li>
          );
        }
    
        // Ellipsis after the visible range if needed
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
    
      return (
        <nav className="flex items-center justify-end pt-4 text-blue-bright">
          <ul className="inline-flex items-center -space-x-px text-sm h-8 gap-[10px]">
            <li className="flex flex-row flex-wrap gap-3 items-center">
              <ChevronLeft />
              <button
                onClick={handlePreviousPage}
                disabled={currentPage === 1}
                className="flex items-center justify-center px-3 h-8"
              >
                Previous
              </button>
            </li>
    
            <PageNumbers totalPages={totalPages} currentPage={currentPage} setCurrentPage={setCurrentPage}/>
    
            <li className="flex flex-row flex-wrap gap-3 items-center">
              <button
                onClick={handleNextPage}
                disabled={currentPage === totalPages}
                className="flex items-center justify-center px-3 h-8"
              >
                Next
              </button>
              <ChevronRight />
            </li>
          </ul>
        </nav>
      );
  };
  
  export default Pagination;