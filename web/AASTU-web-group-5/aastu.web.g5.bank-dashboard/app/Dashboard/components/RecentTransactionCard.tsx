'use client';

import React, { useEffect, useState } from 'react';
import Image from 'next/image';
import axios from 'axios';
import { useSession } from 'next-auth/react';

const TransactionDate = ({ date }: { date: string }) => {
  // Parse the date string into a Date object
  const parsedDate = new Date(date);
  
  // Format the date to "dd MMMM yyyy"
  const formattedDate = parsedDate.toLocaleDateString('en-GB', {
    day: '2-digit',
    month: 'long',
    year: 'numeric',
  });
  
  return formattedDate;
}

interface Transaction {
  transactionId: string;
  senderUserName: string;
  type: string;
  date: string;
  amount: number;
  description: string;
  receiverUserName: string | null;
}

const RecentTransactionCard = () => {
  const { data: session } = useSession();
  const user = session?.user as { accessToken?: string; refreshToken?: string };
  
  const [data, setData] = useState<Transaction[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchRecentTransactions = async () => {
      if (!user?.accessToken) {
        setError('No access token available');
        setLoading(false);
        return;
      }

      try {
        const response = await axios.get<{ data: { content: Transaction[] } }>(
          'https://bank-dashboard-rsf1.onrender.com/transactions?page=0&size=3',
          {
            headers: {
              Authorization: `Bearer ${user.accessToken}`,
            },
          }
        );
        setData(response.data.data.content);
      } catch (err) {
        if (err.response && err.response.status === 401 && user.refreshToken) {
          try {
            // Attempt to refresh the access token
            const refreshResponse = await axios.post('https://bank-dashboard-rsf1.onrender.com/auth/refresh_token', {}, {
              headers: {
                'Authorization': `Bearer ${user.refreshToken}`,
              },
            });

            const refreshedTokens = refreshResponse.data;
            const newAccessToken = refreshedTokens.data.access_token;

            // Retry fetching data with the new access token
            const retryResponse = await axios.get<{ data: { content: Transaction[] } }>(
              'https://bank-dashboard-rsf1.onrender.com/transactions?page=0&size=3',
              {
                headers: {
                  Authorization: `Bearer ${newAccessToken}`,
                },
              }
            );

            setData(retryResponse.data.data.content);
          } catch (refreshError) {
            setError('Failed to refresh access token or fetch data');
          }
        } else {
          setError('Failed to fetch data. Please check the console for more details.');
          console.error('Error fetching data:', err);
        }
      } finally {
        setLoading(false);
      }
    };

    fetchRecentTransactions();
  }, [user?.accessToken, user?.refreshToken]);

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
      <div className='bg-white rounded-[25px] p-7'>
        {data.slice(0, 3).map((transaction) => (
          <div key={transaction.transactionId} className="flex items-center justify-between space-x-8 mb-4">
            <div className="flex items-center space-x-4">
              <Image
                height={44}
                width={44}
                src={getIconPath(transaction.type)} // Icon based on transaction type
                alt={transaction.type}
                className="object-cover rounded-full"
              />
              <div>
              <p className="font-medium text-[16px] leading-[19.36px] text-left font-inter">
  {transaction.senderUserName || transaction.receiverUserName}
</p>
                <p className="text-xs md:text-sm text-gray-500">
                  <TransactionDate date={transaction.date} />
                </p>
              </div>
            </div>
            <p className={`font-medium text-sm md:text-base ${transaction.type === 'deposit' ? 'text-green-600' : 'text-red-700'}`}>
  ${transaction.amount}
</p>

          </div>
        ))}
      </div>
    </div>
  );
};

export default RecentTransactionCard;
