import React, { useState } from "react";
import { FaArrowUp, FaArrowDown,FaArrowLeft,FaArrowRight } from "react-icons/fa";
import dummyData from "../components/dummyData"; // Adjust the path as needed

const itemsPerPage = 5;

const TableCard: React.FC = () => {
  const [currentPage, setCurrentPage] = useState<number>(0);

  // Paginate data
  const paginatedData = dummyData.slice(
    currentPage * itemsPerPage,
    (currentPage + 1) * itemsPerPage
  );

  const totalPages = Math.ceil(dummyData.length / itemsPerPage);

  const handlePageChange = (newPage: number) => {
    setCurrentPage(newPage);
  };

  // Pagination Controls
  const Pagination: React.FC<{ currentPage: number; totalPages: number; onPageChange: (page: number) => void; }> = ({ currentPage, totalPages, onPageChange }) => {
    const pageNumbers = Array.from({ length: totalPages }, (_, i) => i);

    return (
      <div className="flex justify-between items-center mt-4 space-x-2">
        <button
          className="flex items-center text-gray-600 p-2 hover:text-blue-500 disabled:opacity-50"
          onClick={() => onPageChange(currentPage - 1)}
          disabled={currentPage === 0}
        >
          <FaArrowLeft size={20} />
          <span className="ml-1">Prev</span>
        </button>

        <div className="flex items-center space-x-1">
          {pageNumbers.map((page) => (
            <button
              key={page}
              onClick={() => onPageChange(page)}
              className={`p-2 text-sm ${page === currentPage ? 'bg-blue-500 text-white' : 'bg-gray-200 text-gray-700'} hover:bg-blue-600 hover:text-white transition-colors`}
            >
              {page + 1}
            </button>
          ))}
        </div>

        <button
          className="flex items-center text-gray-600 p-2 hover:text-blue-500 disabled:opacity-50"
          onClick={() => onPageChange(currentPage + 1)}
          disabled={currentPage === totalPages - 1}
        >
          <span className="mr-1">Next</span>
          <FaArrowRight size={20} />
        </button>
      </div>
    );
  };

  return (
    <div className="flex flex-col gap-4">
      {paginatedData.map((data, index) => {
        const amount = data.column6;
        const isPositive = amount.startsWith('+') || parseFloat(amount.replace(/[^0-9.-]/g, '')) > 0;
        const amountColor = isPositive ? 'text-green-500' : 'text-red-500';
        const icon = isPositive ? <FaArrowUp className={`text-xl ${amountColor}`} /> : <FaArrowDown className={`text-xl ${amountColor}`} />;

        return (
          <div key={index} className="flex items-center p-4 border rounded-lg shadow-md">
            <div className="mr-4">
              {icon}
            </div>
            <div className="flex-1 flex justify-between items-center">
              <div className="flex-1">
                <p className="text-lg font-semibold">{data.column1}</p>
                <p className="text-sm text-gray-600">{data.column5}</p> {/* Date displayed below description */}
              </div>
              <p className={`text-lg ${amountColor} ml-4`}>{amount}</p> {/* Amount displayed in the same row */}
            </div>
          </div>
        );
      })}
      {/* Pagination Controls */}
      <Pagination
        currentPage={currentPage}
        totalPages={totalPages}
        onPageChange={handlePageChange}
      />
    </div>
  );
};

export default TableCard;
