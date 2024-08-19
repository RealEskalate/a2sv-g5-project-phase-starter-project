"use client";

import React, { useEffect, useState } from 'react';
import Image from 'next/image';
import axios from 'axios';

const RecentTransactionCard = () => {
  const [data, setData] = useState<any[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchRecentTransactions = async () => {
      try {
        const response = await axios.get(
          'https://bank-dashboard-6acc.onrender.com/transactions?page=0&size=3',
          {
            headers: {
              Authorization: `Bearer ${process.env.NEXT_PUBLIC_ACCESS_TOKEN}`,
            },
          }
        );
        setData(response.data.data);
      } catch (err: any) {
        setError('Failed to fetch data. Please check the console for more details.');
        console.error('Error fetching data:', err);
      } finally {
        setLoading(false);
      }
    };

    fetchRecentTransactions();
  }, []);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;

  const slicedTransactions = data.slice(0, 3);

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
      <p className='text-[#343C6A]'>Recent Transactions</p>
      <div>
        {slicedTransactions.map((transaction, index) => (
          <div key={index} className="flex items-center justify-between space-x-8 mb-4">
            <div className="flex items-center space-x-4">
              <Image
                height={44}
                width={44}
                src={getIconPath(transaction.type)} // Icon based on transaction type
                alt={transaction.type}
                className="object-cover"
              />
              <div>
                <p className="font-semibold text-sm md:text-base">{transaction.description}</p>
                <p className="text-xs md:text-sm text-gray-500">{new Date(transaction.date).toLocaleDateString()}</p>
              </div>
            </div>
            <p className={`font-semibold text-sm md:text-base ${transaction.amount >= 0 ? 'text-green-600' : 'text-red-700'}`}>
              {transaction.amount >= 0 ? `+$${transaction.amount}` : `-$${Math.abs(transaction.amount)}`}
            </p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default RecentTransactionCard;
