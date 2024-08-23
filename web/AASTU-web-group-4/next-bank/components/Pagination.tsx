// src/components/TransactionTable/Pagination.tsx

import React from 'react';

interface PaginationProps {
  currentPage: number;
  totalPages: number;
  onPageChange: (page: number) => void;
}

const Pagination: React.FC<PaginationProps> = ({ currentPage, totalPages, onPageChange }) => {
  const pages = Array.from({ length: totalPages }, (_, index) => index + 1);

  return (
    <div className="flex justify-center items-center my-6 space-x-2">
      {/* Previous Button */}
      <button
        className={`px-4 py-2 border rounded-lg shadow-sm ${currentPage <= 1 ? ' text-gray-400' : 'bg-transparent text-blue-500 hover:bg-blue-100'} transition-colors duration-200`}
        onClick={() => onPageChange(currentPage - 1)}
        disabled={currentPage <= 0}
      >
        Previous
      </button>

      {/* Page Number Buttons */}
      {pages.map(page => (
        <button
          key={page}
          className={`px-4 py-2 border rounded-lg shadow-sm ${page-1 === currentPage ? 'bg-blue-600 text-white' : 'bg-transparent text-blue-500 border-blue-500 hover:bg-blue-100'} transition-colors duration-200`}
          onClick={() => onPageChange(page-1)}
        >
          {page}
        </button>
      ))}

      {/* Next Button */}
      <button
        className={`px-4 py-2 border rounded-lg shadow-sm ${currentPage >= totalPages ? ' text-gray-400' : 'bg-transparent text-blue-500  hover:bg-blue-100'} transition-colors duration-200`}
        onClick={() => onPageChange(currentPage + 1)}
        disabled={currentPage >= totalPages}
      >
        Next
      </button>
    </div>
  );
};

export default Pagination;
