// src/components/TransactionTable/TransactionTableRow.tsx

import React from 'react';
import { Transaction } from '@/lib/utils';
import { FaArrowUp, FaArrowDown } from 'react-icons/fa';

interface TransactionTableRowProps {
  transaction: Transaction;
}

const TransactionTableRow: React.FC<TransactionTableRowProps> = ({ transaction }) => {
  const { description, id, type, category, card, date, amount } = transaction;

  return (
    <tr className="border-b ">
      <td className="px-4 py-2 flex items-center">
        {type === 'income' ? (
          <FaArrowUp className="text-green-500 mr-2" />
        ) : (
          <FaArrowDown className="text-red-500 mr-2" />
        )}
        {description}
      </td>
      <td className="px-4 py-2">{id}</td>
      <td className="px-4 py-2">{category}</td>
      <td className="px-4 py-2">{card}</td>
      <td className="px-4 py-2">{date}</td>
      <td className={`px-4 py-2 ${amount < 0 ? 'text-red-500' : 'text-green-500'}`}>
        {amount < 0 ? `-$${Math.abs(amount)}` : `+$${amount}`}
      </td>
      <td className="px-4 py-2">
        <button className="text-blue-500 hover:underline">Download</button>
      </td>
    </tr>
  );
};

export default TransactionTableRow;
