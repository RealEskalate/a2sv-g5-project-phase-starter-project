import React, { useState } from 'react';

interface Transaction {
  description: string;
  transactionId: string;
  type: string;
  receiverUserName?: string;
  date: string;
  amount?: string;
  receipt?: string;
  senderUserName?: string;
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

  const columns: Column[] = [
    {
      Header: 'Description',
      accessor: 'description',
      Cell: ({ value, row }: { value: string; row: Transaction }) => {
        const amountValue = row.amount ? parseFloat(row.amount.replace(/[^0-9.-]/g, '')) : 0;
        const isPositive = amountValue > 0;

        return (
          <div className="flex items-center">
            <div className="border border-solid border-[#718EBF] dark:border-[#333B69] rounded-full flex justify-center items-center h-6 w-6 -mr-4">
              <img
                src={isPositive ? '/downArrow.svg' : '/upArrow.svg'}
                alt={isPositive ? 'down arrow' : 'up arrow'}
                className="h-3 w-3"
              />
            </div>
            <span className="ml-6 dark:text-[#9faaeb]">{value}</span>
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
            <span className={isPositive ? 'text-green-500 dark:text-green-400 ml-2' : 'text-red-500 dark:text-red-400 ml-2'}>
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
        <button className="text-[#123288] dark:text-[#9faaeb] border border-[#123288] dark:border-[#333B69] rounded-full px-2 py-1 hover:text-blue-600 dark:hover:text-blue-400 hover:border-blue-600 dark:hover:border-blue-400 transition-colors duration-300">
          Download
        </button>
      ),
    },
  ];

  return (
    <div className="p-4 dark:bg-[#020817]">
      {transactions.length === 0 ? (
        <p className="dark:text-[#9faaeb]">No transactions available</p>
      ) : (
        <div className="p-4">
          {/* Desktop and Tablet View */}
          <div className="hidden md:block rounded-3xl shadow-md p-4 bg-white dark:bg-[#050914]">
            <div className="overflow-x-auto [&::-webkit-scrollbar]:hidden">
              <table className="min-w-full divide-y divide-[#E6EFF5] dark:divide-[#333B69] bg-white dark:bg-[#050914]">
                <thead className="bg-white dark:bg-[#050914] font-inter font-medium">
                  <tr>
                    {columns.map((column, index) => (
                      <th
                        key={index}
                        className="px-6 py-3 text-left text-xs sm:text-sm font-medium text-[#718EBF] dark:text-[#9faaeb]"
                      >
                        {column.Header}
                      </th>
                    ))}
                  </tr>
                </thead>
                <tbody className="bg-white dark:bg-[#050914] divide-y divide-[#E6EFF5] dark:divide-[#333B69]">
                  {paginatedTransactions.map((row, rowIndex) => (
                    <tr key={rowIndex}>
                      {columns.map((column, colIndex) => (
                        <td
                          key={colIndex}
                          className="px-6 py-4 whitespace-nowrap font-normal text-xs sm:text-sm dark:text-[#9faaeb]"
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
                <div key={index} className="border-b border-gray-200 dark:border-[#333B69] py-2">
                  <div className="flex justify-between items-center">
                    <div className="flex items-center">
                      <div className="border border-solid border-[#718EBF] dark:border-[#333B69] rounded-full flex justify-center items-center h-10 w-10 -mr-4">
                        <img
                          src={isPositive ? '/downArrow.svg' : '/upArrow.svg'}
                          alt={isPositive ? 'down arrow' : 'up arrow'}
                          className="h-5 w-5"
                        />
                      </div>
                      <div className="ml-3"> {/* Adjusted margin here */}
                        <span className="dark:text-[#9faaeb]">{transaction.description}</span>
                        <div className="text-sm text-gray-500 dark:text-[#9faaeb]">{transaction.date}</div>
                      </div>
                    </div>
                    <span className={isPositive ? 'text-green-500 dark:text-green-400' : 'text-red-500 dark:text-red-400'}>
                      {amount}
                    </span>
                  </div>
                </div>
              );
            })}
          </div>

          <div className="flex items-center mt-4 space-x-2 justify-center">
            <button
              onClick={handlePreviousPage}
              disabled={currentPage === 1}
              className="text-blue-600 dark:text-[#9faaeb] rounded-full px-2 py-1 disabled:opacity-50 flex items-center"
            >
              <svg
                width="8"
                height="12"
                viewBox="0 0 8 12"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
                className="mr-1"
              >
                <path d="M7 1L2 6L7 11" stroke="#1814F3" strokeWidth="1.5" />
              </svg>
              Previous
            </button>
            <span className="px-3 py-1 rounded bg-[#1814F3] text-white">
              {currentPage}
            </span>
            <button
              onClick={handleNextPage}
              disabled={currentPage === totalPages}
              className="text-blue-600 dark:text-[#9faaeb] rounded-full px-1 py-1 disabled:opacity-50 flex items-center"
            >
              Next
              <svg
                width="8"
                height="12"
                viewBox="0 0 8 12"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
                className="ml-1"
              >
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






// //               </table>
// //             </div>
// //           </div>
// //           <div className="block md:hidden">
// //             {paginatedTransactions.map((transaction, index) => (
// //               <div key={index} className="border-b border-gray-200 py-2">
// //                 <div className="flex justify-between items-center">
// //                   <div className="flex items-center">
// //                     <div className="ml-3">
// //                       <span>{transaction.description}</span>
// //                       <div className="text-sm text-gray-500">{transaction.date}</div>
// //                     </div>
// //                   </div>
// //                   <span className="ml-2">
// //                     {transaction.amount || 'N/A'}
// //                   </span>
// //                 </div>
// //               </div>
// //             ))}
// //           </div>
// //           <div className="flex items-center mt-4 space-x-2">
// //             <div className="ml-auto flex items-center space-x-1">
// //               <button
// //                 onClick={() => setCurrentPage((prev) => Math.max(prev - 1, 1))}
// //                 className="text-blue-600 rounded-full px-2 py-1"
// //               >
// //                 Previous
// //               </button>
// //             </div>
// //             <div className="flex space-x-1">
// //               {Array.from({ length: totalPages }, (_, index) => (
// //                 <button
// //                   key={index}
// //                   onClick={() => setCurrentPage(index + 1)}
// //                   className={`px-3 py-1 rounded ${
// //                     currentPage === index + 1 ? 'bg-[#1814F3] text-white' : 'text-[#1814F3]'
// //                   }`}
// //                 >
// //                   {index + 1}
// //                 </button>
// //               ))}
// //             </div>
// //             <div className="flex items-center space-x-1">
// //               <button
// //                 onClick={() => setCurrentPage((prev) => Math.min(prev + 1, totalPages))}
// //                 className="text-blue-600 rounded-full px-1 py-1"
// //               >
// //                 Next
// //               </button>
// //             </div>
// //           </div>
// //         </div>
// //       )}
// //     </div>
// //   );
// // };

// // export default TransactionsList;
// import React, { useState } from 'react';

// interface Transaction {
//   description: string;
//   transactionId: string;
//   type: string;
//   receiverUserName?: string;
//   date: string;
//   amount?: string;
//   receipt?: string;
//   senderUserName?: string;
// }

// interface TransactionsListProps {
//   transactions: Transaction[];
// }

// const TransactionsList: React.FC<TransactionsListProps> = ({ transactions }) => {
//   const [currentPage, setCurrentPage] = useState<number>(1);
//   const itemsPerPage = 5;
//   const totalPages = Math.ceil(transactions.length / itemsPerPage);

//   const paginatedTransactions = transactions.slice(
//     (currentPage - 1) * itemsPerPage,
//     currentPage * itemsPerPage
//   );

//   const handleNextPage = () => {
//     if (currentPage < totalPages) {
//       setCurrentPage(currentPage + 1);
//     }
//   };

//   const handlePreviousPage = () => {
//     if (currentPage > 1) {
//       setCurrentPage(currentPage - 1);
//     }
//   };

//   return (
//     <div className="p-4 bg-[#f5f7fa] dark:bg-[#020817]">
//       {transactions.length === 0 ? (
//         <p className="text-[#343C6A] dark:text-[#9faaeb]">
//           No transactions available
//         </p>
//       ) : (
//         <div className="p-4">
//           {/* Responsive Table */}
//           <div className="rounded-3xl shadow-md p-4 bg-white dark:bg-[#050914] w-full border dark:border dark:border-[#333B69]">
//             <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-7 gap-4 text-xs sm:text-sm">
//               <div className="font-medium text-[#718EBF] dark:text-[#9faaeb]">
//                 Description
//               </div>
//               <div className="font-medium text-[#718EBF] dark:text-[#9faaeb]">
//                 Transaction ID
//               </div>
//               <div className="font-medium text-[#718EBF] dark:text-[#9faaeb]">
//                 Type
//               </div>
//               <div className="font-medium text-[#718EBF] dark:text-[#9faaeb]">
//                 Card
//               </div>
//               <div className="font-medium text-[#718EBF] dark:text-[#9faaeb]">
//                 Date
//               </div>
//               <div className="font-medium text-[#718EBF] dark:text-[#9faaeb]">
//                 Amount
//               </div>
//               <div className="font-medium text-[#718EBF] dark:text-[#9faaeb]">
//                 Receipt
//               </div>
//             </div>
//             {paginatedTransactions.map((transaction, index) => (
//               <div
//                 key={index}
//                 className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-7 gap-4 py-2 border-b dark:border-[#333B69]"
//               >
//                 <div className="text-[#343C6A] dark:text-[#9faaeb]">
//                   {transaction.description}
//                 </div>
//                 <div className="text-[#343C6A] dark:text-[#9faaeb]">
//                   {transaction.transactionId}
//                 </div>
//                 <div className="text-[#343C6A] dark:text-[#9faaeb]">
//                   {transaction.type}
//                 </div>
//                 <div className="text-[#343C6A] dark:text-[#9faaeb]">
//                   {transaction.receiverUserName || 'N/A'}
//                 </div>
//                 <div className="text-[#343C6A] dark:text-[#9faaeb]">
//                   {transaction.date}
//                 </div>
//                 <div className="text-[#343C6A] dark:text-[#9faaeb]">
//                   {transaction.amount || 'N/A'}
//                 </div>
//                 <div>
//                   <button className="text-[#123288] border border-[#123288] rounded-full px-2 py-1 hover:text-blue-600 hover:border-blue-600 transition-colors duration-300 dark:text-[#9faaeb] dark:border-[#333B69] dark:hover:text-blue-400 dark:hover:border-blue-400">
//                     Download
//                   </button>
//                 </div>
//               </div>
//             ))}
//           </div>

//           {/* Pagination Controls */}
//           <div className="flex items-center mt-4 space-x-2 justify-center">
//             <button
//               onClick={handlePreviousPage}
//               disabled={currentPage === 1}
//               className="text-blue-600 dark:text-[#9faaeb] rounded-full px-2 py-1 disabled:opacity-50 flex items-center"
//             >
//               <svg
//                 width="8"
//                 height="12"
//                 viewBox="0 0 8 12"
//                 fill="none"
//                 xmlns="http://www.w3.org/2000/svg"
//                 className="mr-1"
//               >
//                 <path d="M7 1L2 6L7 11" stroke="#1814F3" strokeWidth="1.5" />
//               </svg>
//               Previous
//             </button>
//             <span className="px-3 py-1 rounded bg-[#1814F3] text-white">
//               {currentPage}
//             </span>
//             <button
//               onClick={handleNextPage}
//               disabled={currentPage === totalPages}
//               className="text-blue-600 dark:text-[#9faaeb] rounded-full px-1 py-1 disabled:opacity-50 flex items-center"
//             >
//               Next
//               <svg
//                 width="8"
//                 height="12"
//                 viewBox="0 0 8 12"
//                 fill="none"
//                 xmlns="http://www.w3.org/2000/svg"
//                 className="ml-1"
//               >
//                 <path d="M1 11L6 6L1 1" stroke="#1814F3" strokeWidth="1.5" />
//               </svg>
//             </button>
//           </div>
//         </div>
//       )}
//     </div>
//   );
// };

// export default TransactionsList;
