// src/components/RecentTransactions.tsx

import { useState } from 'react';
import { FaArrowCircleUp, FaArrowCircleDown } from 'react-icons/fa';
import { transactionsData, ITEMS_PER_PAGE } from '@/constants';
import Pagination from '@/components/Pagination';

const RecentTransactions = () => {
  const [filter, setFilter] = useState<'all' | 'income' | 'expense'>('all');
  const [currentPage, setCurrentPage] = useState(1);

  // Filter transactions based on the selected tab
  const filteredTransactions = transactionsData.filter(transaction => {
    if (filter === 'all') return true;
    return transaction.type === filter;
  });

  // Pagination calculations
  const totalPages = Math.ceil(filteredTransactions.length / ITEMS_PER_PAGE);
  const startIndex = (currentPage - 1) * ITEMS_PER_PAGE;
  const currentTransactions = filteredTransactions.slice(startIndex, startIndex + ITEMS_PER_PAGE);

  return (
    <div className="p-4 md:ml-64">
      {/* Tabs */}
      <div className="flex flex-wrap mb-4">
        <button
          onClick={() => setFilter('all')}
          className={`font-bold px-4 py-2 rounded-t-lg ${filter === 'all' ? 'border-b-2 border-blue-500' : 'text-gray-600'}`}
        >
          All Transactions
        </button>
        <button
          onClick={() => setFilter('income')}
          className={`font-bold px-4 py-2 rounded-t-lg ${filter === 'income' ? 'border-b-2 border-blue-500' : 'text-gray-600'}`}
        >
          Income
        </button>
        <button
          onClick={() => setFilter('expense')}
          className={`font-bold px-4 py-2 rounded-t-lg ${filter === 'expense' ? 'border-b-2 border-blue-500' : 'text-gray-600'}`}
        >
          Expenses
        </button>
      </div>

      {/* Transactions Table for Desktop/Tablets */}
      <div className="hidden md:block overflow-x-auto">
        <table className="min-w-full bg-white rounded-lg shadow-md border border-gray-200">
          <thead className="bg-blue-50">
            <tr>
              <th className="p-4 text-left text-sm font-semibold text-gray-700">Description</th>
              <th className="p-4 text-left text-sm font-semibold text-gray-700">Transaction ID</th>
              <th className="p-4 text-left text-sm font-semibold text-gray-700">Type</th>
              <th className="p-4 text-left text-sm font-semibold text-gray-700">Card</th>
              <th className="p-4 text-left text-sm font-semibold text-gray-700">Date</th>
              <th className="p-4 text-left text-sm font-semibold text-gray-700">Amount</th>
              <th className="p-4 text-left text-sm font-semibold text-gray-700">Receipt</th>
            </tr>
          </thead>
          <tbody>
            {currentTransactions.map((transaction, index) => (
              <tr
                key={transaction.id}
                className={`border-b border-gray-200 ${index % 2 === 0 ? 'bg-gray-50' : 'bg-white'} hover:bg-gray-100 transition-colors duration-300`}
              >
                <td className="p-4 flex items-center text-sm text-gray-700 truncate">
                  {transaction.type === 'income' ? (
                    <FaArrowCircleUp className="text-green-500 text-lg mr-2" />
                  ) : (
                    <FaArrowCircleDown className="text-red-500 text-lg mr-2" />
                  )}
                  {transaction.description}
                </td>
                <td className="p-4 text-sm text-gray-600">{transaction.id}</td>
                <td className="p-4 text-sm text-gray-600">{transaction.type}</td>
                <td className="p-4 text-sm text-gray-600">Card Name</td>
                <td className="p-4 text-sm text-gray-600">{transaction.date}</td>
                <td className={`p-4 text-sm ${transaction.type === 'income' ? 'text-green-500' : 'text-red-500'}`}>
                  {transaction.amount}
                </td>
                <td className="p-4">
                  <button className="text-blue-500 text-sm hover:underline">Download</button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {/* Transactions for Mobile View */}
      <div className="block md:hidden">
        {currentTransactions.map((transaction) => (
          <div key={transaction.id} className="flex justify-between bg-white p-4 mb-2 rounded-lg shadow-sm border border-gray-200 items-center">
            <div>
              <div className="flex items-center mb-2">
                {transaction.type === 'income' ? (
                  <FaArrowCircleUp className="text-green-500 text-xl mr-2" />
                ) : (
                  <FaArrowCircleDown className="text-red-500 text-xl mr-2" />
                )}
                <span className="font-semibold">{transaction.description}</span>
              </div>
              <div className="text-[12px] text-gray-400 mb-1 pl-5">{transaction.date}</div>
            </div>
            <div>
              <div className={`font-bold ${transaction.type === 'income' ? 'text-green-500' : 'text-red-500'}`}>
                {transaction.amount}
              </div>
            </div>
          </div>
        ))}
      </div>

      {/* Pagination */}
      <Pagination
        currentPage={currentPage}
        totalPages={totalPages}
        onPageChange={setCurrentPage}
      />
    </div>
  );
};

export default RecentTransactions;
