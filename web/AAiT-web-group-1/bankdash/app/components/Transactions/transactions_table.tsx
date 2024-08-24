"use client";
import React, { useState } from 'react';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';

interface Transaction {
  description: string;
  transactionId: string;
  type: string;
  card: string;
  date: string;
  amount: string;
  amountColor: string;
}

interface TransactionsTableProps {
  transactions: Transaction[];
}



const TransactionsTable = () => {

  const transactions = [
  {
    description: 'Spotify Subscription',
    transactionId: '#12548796',
    type: 'Shopping',
    card: '1234 ****',
    date: '28 Jan, 12.30 AM',
    amount: '-$9.99',
    amountColor: '',
  },
  {
    description: 'Amazon Purchase',
    transactionId: '#12548797',
    type: 'Shopping',
    card: '5678 ****',
    date: '27 Jan, 3.15 PM',
    amount: '-$45.00',
    amountColor: '',
  },
  {
    description: 'Salary Credit',
    transactionId: '#12548798',
    type: 'Transfer',
    card: '1234 ****',
    date: '25 Jan, 10.40 PM',
    amount: '+$3,000',
    amountColor: '',
  },
  {
    description: 'Electricity Bill',
    transactionId: '#12548799',
    type: 'Utilities',
    card: '5678 ****',
    date: '24 Jan, 8.00 AM',
    amount: '-$120.50',
    amountColor: '',
  },
  {
    description: 'Freelance Payment',
    transactionId: '#12548800',
    type: 'Transfer',
    card: '1234 ****',
    date: '23 Jan, 5.30 PM',
    amount: '+$500',
    amountColor: '',
  },
  {
    description: 'Netflix Subscription',
    transactionId: '#12548801',
    type: 'Shopping',
    card: '5678 ****',
    date: '22 Jan, 9.00 PM',
    amount: '-$15.99',
    amountColor: '',
  },
  {
    description: 'Gym Membership',
    transactionId: '#12548802',
    type: 'Shopping',
    card: '1234 ****',
    date: '21 Jan, 7.00 AM',
    amount: '-$50.00',
    amountColor: '',
  },
  {
    description: 'Grocery Shopping',
    transactionId: '#12548803',
    type: 'Shopping',
    card: '5678 ****',
    date: '20 Jan, 4.00 PM',
    amount: '-$200.75',
    amountColor: '',
  },
  {
    description: 'Book Sale',
    transactionId: '#12548804',
    type: 'Transfer',
    card: '1234 ****',
    date: '19 Jan, 11.30 AM',
    amount: '+$30.00',
    amountColor: '',
  },
  {
    description: 'Coffee Shop',
    transactionId: '#12548805',
    type: 'Shopping',
    card: '5678 ****',
    date: '18 Jan, 2.00 PM',
    amount: '-$5.50',
    amountColor: '',
  },
  {
    description: 'Spotify Subscription',
    transactionId: '#12548806',
    type: 'Shopping',
    card: '1234 ****',
    date: '17 Jan, 12.30 AM',
    amount: '-$9.99',
    amountColor: '',
  },
  {
    description: 'Freelance Payment',
    transactionId: '#12548807',
    type: 'Transfer',
    card: '5678 ****',
    date: '16 Jan, 10.40 PM',
    amount: '+$750',
    amountColor: '',
  },
  {
    description: 'Electricity Bill',
    transactionId: '#12548808',
    type: 'Utilities',
    card: '1234 ****',
    date: '15 Jan, 8.00 AM',
    amount: '-$120.50',
    amountColor: '',
  },
  {
    description: 'Amazon Purchase',
    transactionId: '#12548809',
    type: 'Shopping',
    card: '5678 ****',
    date: '14 Jan, 3.15 PM',
    amount: '-$45.00',
    amountColor: '',
  },
  {
    description: 'Salary Credit',
    transactionId: '#12548810',
    type: 'Transfer',
    card: '1234 ****',
    date: '13 Jan, 10.40 PM',
    amount: '+$3,000',
    amountColor: '',
  },
  {
    description: 'Spotify Subscription',
    transactionId: '#12548796',
    type: 'Shopping',
    card: '1234 ****',
    date: '28 Jan, 12.30 AM',
    amount: '-$9.99',
    amountColor: '',
  },
  {
    description: 'Amazon Purchase',
    transactionId: '#12548797',
    type: 'Shopping',
    card: '5678 ****',
    date: '27 Jan, 3.15 PM',
    amount: '-$45.00',
    amountColor: '',
  },
  // Add more transactions as needed
];
  // Pagination state
  const [currentPage, setCurrentPage] = useState(1);
  const rowsPerPage = 5;
  
  // Calculate start and end index for the current page
  const indexOfLastRow = currentPage * rowsPerPage;
  const indexOfFirstRow = indexOfLastRow - rowsPerPage;
  const currentRows = transactions.slice(indexOfFirstRow, indexOfLastRow);

  // Calculate total pages
  const totalPages = Math.ceil(transactions.length / rowsPerPage);

  // Handle page change
  const handlePageChange = (page: number) => {
    setCurrentPage(page);
  };

  return (
    <div className="p-5 bg-[#F5F7FA] relative min-h-screen">
      {/* Header */}
      <div>
        <p className="font-inter text-[18px] font-semibold leading-[21.78px] text-left text-[#343C6A] mb-5">
          Recent Transactions
        </p>
      </div>

      {/* Tabs */}
      <div className="font-inter text-[13px] font-medium leading-[15.73px] text-left text-[#718EBF] flex justify-between w-[400px] mb-4">
        <p className='text-[#1814F3] border-b-2 rounded-sm border-[#1814F3]'>All Transactions</p>
        <p>Income</p>
        <p>Expense</p>
      </div>

      {/* Transactions Table */}
      <div className="overflow-x-auto rounded-3xl bg-white text-[14px] px-5">
        <table className="min-w-full">
          <thead>
            <tr>
              <th className="text-left p-4 text-[#718EBF] font-medium">Description</th>
              <th className="text-left p-2 text-[#718EBF] font-medium">Transaction ID</th>
              <th className="text-left p-4 text-[#718EBF] font-medium">Type</th>
              <th className="text-left p-4 text-[#718EBF] font-medium">Card</th>
              <th className="text-center p-4 text-[#718EBF] font-medium">Date</th>
              <th className="text-right p-4 text-[#718EBF] font-medium">Amount</th>
              <th className="text-right p-4 text-[#718EBF] font-medium">Receipt</th>
            </tr>
          </thead>
          <tbody>
            {currentRows.map((transaction, index) => (
              <tr key={index} className="border-t">
                <td className="my-4 text-[#343C6A] flex">
                  {parseFloat(transaction.amount.replace('$', '')) >= 0 ? (
                    <img src="./images/Group 474.png" alt="upArrow" className='mr-2 w-[20px] h-[20px]' />
                  ) : (
                    <img src="./images/Group 474 (1).png" alt="downArrow" className='mr-2 w-[20px] h-[20px]' />
                  )} {transaction.description}
                </td>
                <td className="text-[#343C6A]">{transaction.transactionId}</td>
                <td className="text-[#343C6A]">{transaction.type}</td>
                <td className="text-[#343C6A]">{transaction.card}</td>
                <td className="text-[#343C6A] text-center">{transaction.date}</td>
                <td className={`font-medium text-center ${parseFloat(transaction.amount.replace('$', '')) >= 0 ? "text-green-500" : "text-red-500"}`}>{transaction.amount}</td>
                <td className="text-right">
                  <button className="px-3 py-1 border border-[#123288] text-[#123288] rounded-3xl">
                    Download
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {/* Pagination Controls
      ................................... */}

      {/* prev-button */}
      <div className="flex justify-end absolute right-8 bottom-10 text-[#1814F3] font-inter text-[15px] font-semibold leading-[18.15px]">
        <button
          onClick={() => handlePageChange(currentPage - 1)}
          disabled={currentPage === 1}
          className={`mx-1 ${currentPage === 1 ? 'cursor-not-allowed text-gray-400' : 'hover:font-black'}`}>
            {<ChevronLeftIcon/>}Previous
        </button>

        {/* page-index */}
        {Array.from({ length: totalPages }, (_, page_index) => (
          <button
              key={page_index}
              onClick={() => handlePageChange(page_index + 1)}
              className={`px-3 py-2 mx-1 rounded-xl hover:font-black ${currentPage === page_index + 1 ? 'bg-[#1814F3] text-white' : ''}`}>
              {page_index + 1}
          </button>
        ))}

        {/* next-button */}
        <button
          onClick={() => handlePageChange(currentPage + 1)}
          disabled={currentPage === totalPages}
          className={`mx-1 ${currentPage === totalPages ? 'cursor-not-allowed text-gray-400' : 'hover:font-black'}`}>
          Next{<ChevronRightIcon/>}
        </button>
      </div>

    </div>
  );
};

export default TransactionsTable;

