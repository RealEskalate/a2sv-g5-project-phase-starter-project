import React, { useState } from "react";
import { FaArrowUp, FaArrowDown } from "react-icons/fa";

const itemsPerPage = 5;

interface TableCardProps {
  data: {
    column1: string; // description
    column2: string; // transactionId (not used here)
    column3: string; // type (transaction type)
    column4: string; // cardInfo (not used here)
    column5: string; // date
    column6: number; // amount
    column7: string; // receiptInfo (not used here)
  }[];
}

const TableCard: React.FC<TableCardProps> = ({ data }) => {
  const [currentPage, setCurrentPage] = useState<number>(0);

  // Paginate data
  const paginatedData = data.slice(
    currentPage * itemsPerPage,
    (currentPage + 1) * itemsPerPage
  );

  const totalPages = Math.ceil(data.length / itemsPerPage);

  const handlePageChange = (newPage: number) => {
    setCurrentPage(newPage);
  };

  // Pagination Controls
  const Pagination: React.FC<{ currentPage: number; totalPages: number; onPageChange: (page: number) => void; }> = ({ currentPage, totalPages, onPageChange }) => {
    const pageNumbers = Array.from({ length: totalPages }, (_, i) => i);

    return (
      <div className="flex justify-between items-center mt-4 space-x-2">
        <button
          className="flex items-center text-gray-400 dark:text-gray-600 p-2 hover:text-blue-500 dark:hover:text-blue-400 disabled:opacity-50"
          onClick={() => onPageChange(currentPage - 1)}
          disabled={currentPage === 0}
        >
          <FaArrowUp size={20} />
          <span className="ml-1">Prev</span>
        </button>

        <div className="flex items-center space-x-1">
          {pageNumbers.map((page) => (
            <button
              key={page}
              onClick={() => onPageChange(page)}
              className={`p-2 text-sm ${page === currentPage ? 'bg-blue-500 text-white dark:bg-blue-600 dark:text-white' : 'bg-gray-200 text-gray-700 dark:bg-gray-700 dark:text-gray-300'} hover:bg-blue-600 hover:text-white dark:hover:bg-blue-500 dark:hover:text-white transition-colors`}
            >
              {page + 1}
            </button>
          ))}
        </div>

        <button
          className="flex items-center text-gray-400 dark:text-gray-600 p-2 hover:text-blue-500 dark:hover:text-blue-400 disabled:opacity-50"
          onClick={() => onPageChange(currentPage + 1)}
          disabled={currentPage === totalPages - 1}
        >
          <span className="mr-1">Next</span>
          <FaArrowDown size={20} />
        </button>
      </div>
    );
  };

  return (
    <div className="flex flex-col gap-4 p-4 bg-white dark:bg-gray-900 shadow-md dark:shadow-gray-800 rounded-lg">
      {paginatedData.map((item, index) => {
        const amount = item.column6;
        const isPositive = item.column3 === "deposit" || amount > 0;
        const amountColor = isPositive ? 'text-green-500' : 'text-red-500';
        const icon = isPositive ? <FaArrowUp className={`text-xl ${amountColor}`} /> : <FaArrowDown className={`text-xl ${amountColor}`} />;

        return (
          <div key={index} className="flex items-center justify-between p-4 border rounded-lg bg-gray-100 dark:bg-gray-800 border-gray-300 dark:border-gray-700">
            <div className="flex-1 flex flex-col">
              <div className="text-sm font-semibold text-gray-800 dark:text-gray-200">{item.column1}</div>
              <div className="text-xs text-gray-500 dark:text-gray-400">{item.column5}</div>
            </div>
            <div className="flex items-center space-x-2">
              {icon}
              <span className={`text-lg font-bold ${amountColor}`}>
                ${amount}
              </span>
            </div>
          </div>
        );
      })}

      <Pagination
        currentPage={currentPage}
        totalPages={totalPages}
        onPageChange={handlePageChange}
      />
    </div>
  );
};

export default TableCard;
