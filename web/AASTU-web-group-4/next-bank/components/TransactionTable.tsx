'use client'

import React, { useState } from 'react';
import { transactions } from '@/lib/utils';
import { ITEMS_PER_PAGE, TABLE_HEADERS } from '@/constants/index';
import TransactionTableRow from './TransactionTableRow';
import Pagination from './Pagination';

const TransactionTable: React.FC = () => {
  const [currentPage, setCurrentPage] = useState(1);

  // Calculate the number of pages
  const totalPages = Math.ceil(transactions.length / ITEMS_PER_PAGE);

  // Get the transactions for the current page
  const paginatedTransactions = transactions.slice(
    (currentPage - 1) * ITEMS_PER_PAGE,
    currentPage * ITEMS_PER_PAGE
  );

  return (
    <div className="overflow-x-auto ">
      <table className="min-w-full bg-white shadow-md rounded-lg ">
        <thead>
          <tr>
            {TABLE_HEADERS.map((header) => (
              <th key={header} className="px-4 py-2 bg-gray-200">
                {header}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {paginatedTransactions.map((transaction) => (
            <TransactionTableRow key={transaction.id} transaction={transaction} />
          ))}
        </tbody>
      </table>

      <Pagination
        currentPage={currentPage}
        totalPages={totalPages}
        onPageChange={(page) => setCurrentPage(page)}
      />
    </div>
  );
};

export default TransactionTable;
