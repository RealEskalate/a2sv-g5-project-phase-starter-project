import React, { useState } from 'react';

interface Transaction {
  description: string;
  transactionId: string;
  type: string;
  receiverUserName?: string;
  senderUserName?: string;
  date: string;
  amount?: string;
  receipt?: string;
}

interface TransactionsListProps {
  transactions: Transaction[];
  loading: boolean;
  activeTab: string; 
}

const TransactionsList: React.FC<TransactionsListProps> = ({ transactions, loading, activeTab }) => {
  const [currentPage, setCurrentPage] = useState<number>(1);
  const itemsPerPage = 5;
  const totalPages = Math.ceil(transactions.length / itemsPerPage);

  const paginatedTransactions = transactions.slice(
    (currentPage - 1) * itemsPerPage,
    currentPage * itemsPerPage
  );

  const handlePageClick = (pageNumber: number) => {
    setCurrentPage(pageNumber);
  };

  const handleNextPage = () => {
    if (currentPage < totalPages) {
      setCurrentPage(currentPage + 1);
    }
  };

  const handlePreviousPage = () => {
    if (currentPage > 1) {
      setCurrentPage(currentPage - 1);
    }
  };

  const renderPagination = () => {
    const pageNumbers: (number | string)[] = [];
    const maxPageNumbersToShow = 3;
    

    if (totalPages <= maxPageNumbersToShow) {
      for (let i = 1; i <= totalPages; i++) {
        pageNumbers.push(i);
      }
    } else {
      let startPage = Math.max(1, currentPage - 2);
      let endPage = Math.min(totalPages, currentPage + 2);

      if (currentPage <= 3) {
        startPage = 1;
        endPage = maxPageNumbersToShow;
      } else if (currentPage >= totalPages - 2) {
        startPage = totalPages - maxPageNumbersToShow + 1;
        endPage = totalPages;
      }

      for (let i = startPage; i <= endPage; i++) {
        pageNumbers.push(i);
      }

      if (startPage > 1) {
        pageNumbers.unshift(1);
        if (startPage > 2) {
          pageNumbers.splice(1, 0, '...');
        }
      }

      if (endPage < totalPages) {
        pageNumbers.push(totalPages);
        if (endPage < totalPages - 1) {
          pageNumbers.splice(pageNumbers.length - 1, 0, '...');
        }
      }
    }

    return pageNumbers.map((pageNumber, index) =>
      typeof pageNumber === 'number' ? (
        <button
          key={index}
          onClick={() => handlePageClick(pageNumber)}
          className={`px-3 py-1 rounded ${pageNumber === currentPage ? 'bg-[#1814F3] text-white' : 'text-blue-600 dark:text-[#9faaeb]'}`}
        >
          {pageNumber}
        </button>
      ) : (
        <span key={index} className="px-3 py-1 text-blue-600 dark:text-[#9faaeb]">
          {pageNumber}
        </span>
      )
    );
  };

  const columns = ['Description', 'Transaction ID', 'Type', 'Card', 'Date', 'Amount', 'Receipt'];

  const getNoDataMessage = () => {
    switch (activeTab) {
      case 'Income':
        return 'No income available.';
      case 'Expense':
        return 'No expense available.';
      default:
        return 'No transactions available.';
    }
  };

  const shouldDisplayAsPositive = (transaction: Transaction) => {
    const { senderUserName, receiverUserName } = transaction;
    return senderUserName && receiverUserName && senderUserName !== receiverUserName;
  };

  return (
    <div className="dark:bg-[#090b0e]">
      {loading ? (
        <div className="p-4">
          {/* Shimmer effect for loading state */}
          <div className="hidden md:block rounded-3xl shadow-md p-4 bg-white dark:bg-[#050914]">
            <div className="overflow-x-auto [&::-webkit-scrollbar]:hidden">
              <table className="min-w-full divide-y divide-[#E6EFF5] dark:divide-[#333B69]">
                <thead className="bg-white dark:bg-[#020817] font-inter font-black">
                  <tr>
                    {columns.map((column, index) => (
                      <th
                        key={index}
                        className="px-6 py-3 text-left text-xs sm:text-sm font-black text-[#718EBF] dark:text-[#9faaeb]"
                      >
                        {column}
                      </th>
                    ))}
                  </tr>
                </thead>
                <tbody className="bg-white dark:bg-[#020817] divide-y divide-[#E6EFF5] dark:divide-[#333B69]">
                  {[...Array(5)].map((_, rowIndex) => (
                    <tr key={rowIndex} className="animate-pulse">
                      {columns.map((_, colIndex) => (
                        <td key={colIndex} className="px-6 py-4">
                          <div className="h-4 bg-gray-300 dark:bg-[#333B69] rounded"></div>
                        </td>
                      ))}
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>

          {/* Mobile View */}
          <div className="block md:hidden">
            {[...Array(5)].map((_, index) => (
              <div key={index} className="border-b border-gray-200 dark:border-[#333B69] py-4 animate-pulse">
                <div className="flex justify-between items-center">
                  <div className="flex items-center space-x-2">
                    <div className="w-10 h-10 bg-gray-300 dark:bg-[#020817] rounded-full"></div>
                    <div>
                      <div className="h-4 bg-gray-300 dark:bg-[#020817] rounded w-32 mb-2"></div>
                      <div className="h-4 bg-gray-300 dark:bg-[#020817] rounded w-20"></div>
                    </div>
                  </div>
                  <div className="h-4 bg-gray-300 dark:bg-[#020817] rounded w-16"></div>
                </div>
              </div>
            ))}
          </div>
        </div>
      ) : (
        <div className="p-4">
          {/* Table rendering */}
          <div className="hidden md:block rounded-3xl shadow-md p-4 bg-white dark:bg-[#050914] md:w-11/12">
            <div className="overflow-x-auto [&::-webkit-scrollbar]:hidden">
              <table className="min-w-full divide-y divide-[#E6EFF5] dark:divide-[#333B69]">
                <thead className="bg-white dark:bg-[#050914] font-inter font-bold">
                  <tr>
                    {columns.map((column, index) => (
                      <th
                        key={index}
                        className="px-6 py-3 text-left text-xs sm:text-sm font-black text-[#718EBF] dark:text-[#9faaeb]"
                      >
                        {column}
                      </th>
                    ))}
                  </tr>
                </thead>
                <tbody className="bg-white dark:bg-[#050914] divide-y divide-[#E6EFF5] dark:divide-[#333B69]">
                  {paginatedTransactions.map((row, rowIndex) => {
                    const amount = row.amount || '0';
                    const isPositive = shouldDisplayAsPositive(row);

                    return (
                      <tr key={rowIndex}>
                        <td className="px-6 py-4">
                          <div className="flex items-center space-x-6">
                            <div className="border border-solid rounded-full flex justify-center items-center h-7 w-7 -mr-4">
                              <img
                                src={isPositive ? '/downArrow.svg' : '/upArrow.svg'}
                                alt={isPositive ? 'down arrow' : 'up arrow'}
                                className="h-4 w-4"
                              />
                            </div>
                            <span>{row.description}</span>
                          </div>
                        </td>
                        <td className="px-6 py-4">{row.transactionId}</td>
                        <td className="px-6 py-4">{row.type}</td>
                        <td className="px-6 py-4">{row.receiverUserName || 'Unknown'}</td>
                        <td className="px-6 py-4">{row.date}</td>
                        <td className="px-6 py-4">
                          <span className={isPositive ? 'text-green-500 dark:text-green-400' : 'text-red-500 dark:text-red-400'}>
                            {isPositive ? `+${amount}` : `-${amount}`}
                          </span>
                        </td>
                        <td className="px-6 py-4">
                          <button className="text-[#123288] dark:text-[#9faaeb] border border-[#123288] dark:border-[#333B69] rounded-full px-2 py-1 hover:text-blue-600 dark:hover:text-blue-400 hover:border-blue-600 dark:hover:border-blue-400 transition-colors duration-300">Download</button>
                        </td>
                      </tr>
                    );
                  })}
                </tbody>
              </table>
            </div>
          </div>

              {/* Mobile View */}
              <div className="block md:hidden">
              {paginatedTransactions.map((transaction, index) => {
                const amount = transaction.amount || '0';
                const isSameUser = transaction.senderUserName === transaction.receiverUserName;
                const isPositive = !isSameUser;

                return (
                  <div key={index} className="border-b border-gray-200 dark:border-[#333B69] py-4">
                    <div className="flex justify-between items-center">
                      <div className="flex items-center space-x-6">
                        <div className="border border-solid border-[#718EBF] dark:border-[#333B69] rounded-full flex justify-center items-center h-7 w-7 -mr-4">
                          <img
                            src={isPositive ? '/downArrow.svg' : '/upArrow.svg'}
                            alt={isPositive ? 'down arrow' : 'up arrow'}
                            className="h-4 w-4"
                          />
                        </div>
                        <div>
                          <span className="dark:text-[#9faaeb]">{transaction.description}</span>
                          <div className="text-sm text-gray-500 dark:text-[#9faaeb]">{transaction.date}</div>
                        </div>
                      </div>
                      <span className={isPositive ? 'text-green-500 dark:text-green-400' : 'text-red-500 dark:text-red-400'}>
                        {isPositive ? `+${amount} `: `-${amount}`}
                      </span>
                    </div>
                  </div>
                );
              })}
            </div>

          {/* Pagination */}
          <div className="flex items-center mt-4 space-x-2 justify-center">
            <button
              onClick={handlePreviousPage}
              disabled={currentPage === 1}
              className="text-blue-600 dark:text-[#9faaeb] rounded-full px-2 py-1 disabled:opacity-50 flex items-center"
            >
              <svg width="8" height="12" viewBox="0 0 8 12" fill="none" xmlns="http://www.w3.org/2000/svg" className="mr-1">
                <path d="M7 1L2 6L7 11" stroke="#1814F3" strokeWidth="1.5" />
              </svg>
              Previous
            </button>
            {renderPagination()}
            <button
              onClick={handleNextPage}
              disabled={currentPage === totalPages}
              className="text-blue-600 dark:text-[#9faaeb] rounded-full px-1 py-1 disabled:opacity-50 flex items-center"
            >
              Next
              <svg width="8" height="12" viewBox="0 0 8 12" fill="none" xmlns="http://www.w3.org/2000/svg" className="ml-1">
                <path d="M1 11L6 6L1 1" stroke="#1814F3" strokeWidth="1.5" />
              </svg>
            </button>
          </div>
        </div>
      )}
    </div>
  );
};

export default TransactionsList;
