'use client';

import React, { useEffect, useState } from 'react';
import Image from 'next/image';
import axios from 'axios';

interface Transaction {
  id: string;
  name: string;
  type: string;
  date: string;
  amount: string;
}

const RecentTransactionCard = () => {
  const [data, setData] = useState<Transaction[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  // Replace this with the actual way you retrieve the access token
  const accessToken = 'YOUR_ACCESS_TOKEN_HERE';

  useEffect(() => {
    const fetchRecentTransactions = async () => {
      try {
        const response = await axios.get<{ data: { content: Transaction[] } }>(
          'https://bank-dashboard-o9tl.onrender.com/transactions?page=0&size=3',
          {
            headers: {
              Authorization: `Bearer ${accessToken}`,
            },
          }
        );
        setData(response.data.data.content); // Adjusted to match your response format
      } catch (err) {
        setError('Failed to fetch data. Please check the console for more details.');
        console.error('Error fetching data:', err);
      } finally {
        setLoading(false);
      }
    };

    fetchRecentTransactions();
  }, [accessToken]);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;

  // Helper function to get icon path based on transaction type
  const getIconPath = (type: string) => {
    switch (type) {
      case 'transfer':
        return '/assets/icons/transfer.jpg';
      case 'shopping':
        return '/assets/icons/paypal.jpg';
      case 'deposit':
        return '/assets/icons/deposit.jpg';
      default:
        return '/assets/icons/transfer.jpg';
    }
  };

  return (
    <div>
      <p className='text-[#343C6A] text-lg font-semibold mb-4'>Recent Transactions</p>
      <div className='bg-white rounded-[25px] p-7'>
        {data.slice(0, 3).map((transaction) => (
          <div key={transaction.id} className="flex items-center justify-between space-x-8 mb-4">
            <div className="flex items-center space-x-4">
              <Image
                height={44}
                width={44}
                src={getIconPath(transaction.type)} // Icon based on transaction type
                alt={transaction.type}
                className="object-cover rounded-full"
              />
              <div>
                <p className="font-semibold text-sm md:text-base">{transaction.name}</p>
                <p className="text-xs md:text-sm text-gray-500">{transaction.date}</p>
              </div>
            </div>
            <p className={`font-semibold text-sm md:text-base ${transaction.amount[0] === '+' ? 'text-green-600' : 'text-red-700'}`}>
              {transaction.amount}
            </p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default RecentTransactionCard;
