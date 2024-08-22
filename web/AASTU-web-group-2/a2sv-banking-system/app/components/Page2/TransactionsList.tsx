import React, { useState } from 'react';

interface Transaction {
  description: string;
  transactionId: string;
  type: string;
  receiverUserName?: string;
  date: string;
  amount?: string;
  receipt?: string;
  senderUserName?: string
}

interface TransactionsListProps {
  transactions: Transaction[];
}

type Column = {
  Header: string;
  accessor: keyof Transaction; 
  Cell?: (props: { value: any; row: Transaction }) => React.ReactNode; 
};

const TransactionsList: React.FC<TransactionsListProps> = ({ transactions }) => {
  const [currentPage, setCurrentPage] = useState<number>(1);
  const itemsPerPage = 5;
  const totalPages = Math.ceil(transactions.length / itemsPerPage);

  const paginatedTransactions = transactions.slice(
    (currentPage - 1) * itemsPerPage,
    currentPage * itemsPerPage
  );

  const columns: Column[] = [
    {
      Header: 'Description',
      accessor: 'description',
      Cell: ({ value, row }: { value: string; row: Transaction }) => {
        const amountValue = row.amount ? parseFloat(row.amount.replace(/[^0-9.-]/g, '')) : 0;
        const isPositive = amountValue > 0;

        return (
          <div className="flex items-center">
            <div className="border border-solid border-[#718EBF] rounded-full flex justify-center items-center h-6 w-6 -mr-4">
              <img
                src={isPositive ? '/downArrow.svg' : '/upArrow.svg'}
                alt={isPositive ? 'down arrow' : 'up arrow'}
                className="h-3 w-3"
              />
            </div>
            <span className="ml-6">{value}</span>
          </div>
        );
      },
    },
    {
      Header: 'Transaction ID',
      accessor: 'transactionId',
    },
    {
      Header: 'Type',
      accessor: 'type',
    },
    {
      Header: 'Card',
      accessor: 'receiverUserName',
    },
    {
      Header: 'Date',
      accessor: 'date',
    },
    {
      Header: 'Amount',
      accessor: 'amount',
      Cell: ({ value }: { value: string | undefined }) => {
        const amountValue = value ? parseFloat(value.replace(/[^0-9.-]/g, '')) : 0;
        const isPositive = amountValue > 0;
        return (
          <div className="flex items-center">
            <span className={isPositive ? 'text-green-500 ml-2' : 'text-red-500 ml-2'}>
              {value || 'N/A'}
            </span>
          </div>
        );
      },
    },
    {
      Header: 'Receipt',
      accessor: 'receipt',
      Cell: () => (
        <button className="text-[#123288] border border-[#123288] rounded-full px-2 py-1 hover:text-blue-600 hover:border-blue-600 transition-colors duration-300">
          Download
        </button>
      ),
    },
  ];

  return (
    <div className="p-4">
      {transactions.length === 0 ? (
        <p>No transactions available</p>
      ) : (
        <div className="p-4">
          {/* Desktop and Tablet View */}
          <div className="hidden md:block rounded-3xl shadow-md p-4 bg-white">
            <div className="overflow-x-auto [&::-webkit-scrollbar]:hidden">
              <table className="min-w-full divide-y divide-[#E6EFF5] bg-white">
                <thead className="bg-white font-inter font-medium">
                  <tr>
                    {columns.map((column, index) => (
                      <th
                        key={index}
                        className="px-6 py-3 text-left text-xs sm:text-sm font-medium text-[#718EBF]"
                      >
                        {column.Header}
                      </th>
                    ))}
                  </tr>
                </thead>
                <tbody className="bg-white divide-y divide-[#E6EFF5]">
                  {paginatedTransactions.map((row, rowIndex) => (
                    <tr key={rowIndex}>
                      {columns.map((column, colIndex) => (
                        <td
                          key={colIndex}
                          className="px-6 py-4 whitespace-nowrap font-normal text-xs sm:text-sm"
                        >
                          {column.Cell
                            ? column.Cell({ value: row[column.accessor], row }) 
                            : row[column.accessor] as React.ReactNode} 
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
            {paginatedTransactions.map((transaction, index) => {
              const amount = transaction.amount || '0';
              const isPositive = parseFloat(amount.replace(/[^0-9.-]/g, '')) > 0;
              return (
                <div key={index} className="border-b border-gray-200 py-2">
                  <div className="flex justify-between items-center">
                    <div className="flex items-center">
                      <div className="border border-solid border-[#718EBF] rounded-full flex justify-center items-center h-10 w-10 -mr-4">
                        <img
                          src={isPositive ? '/downArrow.svg' : '/upArrow.svg'}
                          alt={isPositive ? 'down arrow' : 'up arrow'}
                          className="h-5 w-5"
                        />
                      </div>
                      <div className="ml-3"> {/* Adjusted margin here */}
                        <span>{transaction.description}</span>
                        <div className="text-sm text-gray-500">{transaction.date}</div>
                      </div>
                    </div>
                    <span className={isPositive ? 'text-green-500' : 'text-red-500'}>
                      {amount}
                    </span>
                  </div>
                </div>
              );
            })}
          </div>

          {/* Pagination Controls */}
          <div className="flex items-center mt-4 space-x-2">
            <div className="ml-auto flex items-center space-x-1">
              <button
                onClick={() => setCurrentPage((prev) => Math.max(prev - 1, 1))}
                className="text-blue-600 rounded-full px-2 py-1"
              >
                Previous
              </button>
              <img src="/back.svg" alt="back" className="h-4 w-4 text-blue-600" />
            </div>
            <div className="flex space-x-1">
              {Array.from({ length: totalPages }, (_, index) => (
                <button
                  key={index}
                  onClick={() => setCurrentPage(index + 1)}
                  className={`px-3 py-1 rounded ${
                    currentPage === index + 1 ? 'bg-[#1814F3] text-white' : 'text-[#1814F3]'
                  }`}
                >
                  {index + 1}
                </button>
              ))}
            </div>
            <div className="flex items-center space-x-1">
              <button
                onClick={() => setCurrentPage((prev) => Math.min(prev + 1, totalPages))}
                className="text-blue-600 rounded-full px-1 py-1"
              >
                Next
              </button>
              <img src="/forward.svg" alt="forward" className="h-4 w-4 text-blue-600" />
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default TransactionsList;
