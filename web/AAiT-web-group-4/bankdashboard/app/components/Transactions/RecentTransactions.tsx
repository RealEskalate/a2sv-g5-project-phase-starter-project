
import React, { useState } from 'react';

const RecentTransactions: React.FC = () => {
  const transactionsPerPage = 5;
  const [currentPage, setCurrentPage] = useState(1);

  // Dummy data
  const transactions = [
    { id: 1, description: 'Spotify Subscription', transactionId: '#12548796', type: 'Shopping', card: '1234 ****', date: '28 Jan, 12:30 AM', amount: '-$2,500', receipt: 'Download', amountColor: 'text-red-500' },
    { id: 2, description: 'Freepik Sales', transactionId: '#12548796', type: 'Transfer', card: '1234 ****', date: '25 Jan, 10:40 PM', amount: '+$750', receipt: 'Download', amountColor: 'text-green-500' },
    { id: 3, description: 'Mobile Service', transactionId: '#12548796', type: 'Service', card: '1234 ****', date: '20 Jan, 10:40 PM', amount: '-$150', receipt: 'Download', amountColor: 'text-red-500' },
    { id: 4, description: 'Wilson', transactionId: '#12548796', type: 'Transfer', card: '1234 ****', date: '15 Jan, 03:29 PM', amount: '-$1,050', receipt: 'Download', amountColor: 'text-red-500' },
    { id: 5, description: 'Emilly', transactionId: '#12548796', type: 'Transfer', card: '1234 ****', date: '14 Jan, 10:40 PM', amount: '+$840', receipt: 'Download', amountColor: 'text-green-500' },
    { id: 6, description: 'Netflix Subscription', transactionId: '#12548796', type: 'Entertainment', card: '1234 ****', date: '13 Jan, 09:20 AM', amount: '-$15.99', receipt: 'Download', amountColor: 'text-red-500' },
    { id: 7, description: 'Uber Ride', transactionId: '#12548796', type: 'Transport', card: '1234 ****', date: '12 Jan, 08:15 PM', amount: '-$25', receipt: 'Download', amountColor: 'text-red-500' },
    { id: 8, description: 'Salary Deposit', transactionId: '#12548796', type: 'Income', card: '1234 ****', date: '11 Jan, 02:00 PM', amount: '+$3,500', receipt: 'Download', amountColor: 'text-green-500' },
    { id: 9, description: 'Amazon Purchase', transactionId: '#12548796', type: 'Shopping', card: '1234 ****', date: '10 Jan, 06:45 PM', amount: '-$99', receipt: 'Download', amountColor: 'text-red-500' },
    { id: 10, description: 'Starbucks', transactionId: '#12548796', type: 'Food', card: '1234 ****', date: '09 Jan, 07:00 AM', amount: '-$5.50', receipt: 'Download', amountColor: 'text-red-500' },
    { id: 11, description: 'Spotify Subscription', transactionId: '#12548797', type: 'Shopping', card: '5678 ****', date: '28 Jan, 12:30 AM', amount: '-$2,500', receipt: 'Download', amountColor: 'text-red-500' },
    { id: 12, description: 'Freepik Sales', transactionId: '#12548797', type: 'Transfer', card: '5678 ****', date: '25 Jan, 10:40 PM', amount: '+$750', receipt: 'Download', amountColor: 'text-green-500' },
    { id: 13, description: 'Mobile Service', transactionId: '#12548797', type: 'Service', card: '5678 ****', date: '20 Jan, 10:40 PM', amount: '-$150', receipt: 'Download', amountColor: 'text-red-500' },
    { id: 14, description: 'Wilson', transactionId: '#12548797', type: 'Transfer', card: '5678 ****', date: '15 Jan, 03:29 PM', amount: '-$1,050', receipt: 'Download', amountColor: 'text-red-500' },
    { id: 15, description: 'Emilly', transactionId: '#12548797', type: 'Transfer', card: '5678 ****', date: '14 Jan, 10:40 PM', amount: '+$840', receipt: 'Download', amountColor: 'text-green-500' },
    { id: 16, description: 'Netflix Subscription', transactionId: '#12548797', type: 'Entertainment', card: '5678 ****', date: '13 Jan, 09:20 AM', amount: '-$15.99', receipt: 'Download', amountColor: 'text-red-500' },
    { id: 17, description: 'Uber Ride', transactionId: '#12548797', type: 'Transport', card: '5678 ****', date: '12 Jan, 08:15 PM', amount: '-$25', receipt: 'Download', amountColor: 'text-red-500' },
    { id: 18, description: 'Salary Deposit', transactionId: '#12548797', type: 'Income', card: '5678 ****', date: '11 Jan, 02:00 PM', amount: '+$3,500', receipt: 'Download', amountColor: 'text-green-500' },
    { id: 19, description: 'Amazon Purchase', transactionId: '#12548797', type: 'Shopping', card: '5678 ****', date: '10 Jan, 06:45 PM', amount: '-$99', receipt: 'Download', amountColor: 'text-red-500' },
    { id: 20, description: 'Starbucks', transactionId: '#12548797', type: 'Food', card: '5678 ****', date: '09 Jan, 07:00 AM', amount: '-$5.50', receipt: 'Download', amountColor: 'text-red-500' },
  ];

  const totalPages = Math.ceil(transactions.length / transactionsPerPage);

  const handlePageChange = (newPage: number) => {
    setCurrentPage(newPage);
  };

  const paginatedTransactions = transactions.slice(
    (currentPage - 1) * transactionsPerPage,
    currentPage * transactionsPerPage
  );

  return (
    <div className="p-8">
      {/* Recent Transactions Header */}
      <h1 className="font-inter text-lg font-semibold text-[#343C6A]">Recent Transactions</h1>

      {/* Tabs */}
      <div className="flex mt-4 space-x-8">
        <button className="text-blue-600 font-medium">All Transactions</button>
        <button className="text-gray-400">Income</button>
        <button className="text-gray-400">Expense</button>
      </div>

      {/* Table */}
      <div className="mt-8 bg-white shadow-md rounded-lg overflow-hidden">
        <table className="min-w-full text-left">
          <thead>
            <tr className="text-sm font-medium text-gray-500">
              <th className="px-6 py-4">Description</th>
              <th className="px-6 py-4">Transaction ID</th>
              <th className="px-6 py-4">Type</th>
              <th className="px-6 py-4">Card</th>
              <th className="px-6 py-4">Date</th>
              <th className="px-6 py-4">Amount</th>
              <th className="px-6 py-4">Receipt</th>
            </tr>
          </thead>
          <tbody className="text-sm">
            {paginatedTransactions.map((transaction) => (
              <tr key={transaction.id} className="border-t border-gray-200">
                <td className="px-6 py-4">
                  <div className="flex items-center space-x-2">
                    <div className="flex justify-center items-center w-8 h-8 rounded-full border-2 border-[#718EBF]">
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        strokeWidth="2"
                        stroke="#718EBF"
                        className="w-3 h-3"
                      >
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          d={
                            transaction.amount.startsWith('-')
                              ? 'M5 10l7-7m0 0l7 7m-7-7v18' // Upward arrow for negative transactions (withdrawal)
                              : 'M5 14l7 7m0 0l7-7m-7 7V3' // Downward arrow for positive transactions (income)
                          }
                        />
                      </svg>
                    </div>
                    <p>{transaction.description}</p>
                  </div>
                </td>
                <td className="px-6 py-4">{transaction.transactionId}</td>
                <td className="px-6 py-4">{transaction.type}</td>
                <td className="px-6 py-4">{transaction.card}</td>
                <td className="px-6 py-4">{transaction.date}</td>
                <td className={`px-6 py-4 ${transaction.amountColor}`}>{transaction.amount}</td>
                <td className="px-6 py-4">
                  <button className="px-4 py-2 text-blue-600 border border-blue-600 rounded-full">
                    {transaction.receipt}
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {/* Pagination */}
      <div className="flex justify-end items-center mt-4">
        <nav className="inline-flex rounded-md shadow-sm" aria-label="Pagination">
          <button
            onClick={() => handlePageChange(currentPage - 1)}
            disabled={currentPage === 1}
            className="relative inline-flex items-center px-2 py-2 text-sm font-medium text-blue-500 bg-white border border-gray-300 rounded-l-md hover:bg-gray-50"
          >
            &lt; Previous
          </button>
          {[...Array(totalPages)].map((_, pageIndex) => (
            <button
              key={pageIndex + 1}
              onClick={() => handlePageChange(pageIndex + 1)}
              className={`relative inline-flex items-center px-4 py-2 text-sm font-medium ${
                currentPage === pageIndex + 1
                  ? 'text-white bg-blue-600'
                  : 'text-blue-600 bg-white'
              } border border-gray-300`}
            >
              {pageIndex + 1}
            </button>
          ))}
          <button
            onClick={() => handlePageChange(currentPage + 1)}
            disabled={currentPage === totalPages}
            className="relative inline-flex items-center px-2 py-2 text-sm font-medium text-blue-500 bg-white border border-gray-300 rounded-r-md hover:bg-gray-50"
          >
            Next &gt;
          </button>
        </nav>
      </div>
    </div>
  );
};

export default RecentTransactions;
