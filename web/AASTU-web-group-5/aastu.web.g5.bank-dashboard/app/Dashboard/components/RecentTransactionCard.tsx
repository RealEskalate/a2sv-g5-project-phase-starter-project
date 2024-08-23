"use client"

import React from 'react';
import Image from 'next/image'; // Assuming you're using Next.js
import { RecentTransaction } from './mockData'; // Adjust the path as needed
import { Slice } from 'lucide-react';

<<<<<<< HEAD
const RecentTransactionCard = () => {
  const [data, setData] = useState<any[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchRecentTransactions = async () => {
      try {
        const response = await axios.get(
          'https://bank-dashboard-o9tl.onrender.com/transactions?page=0&size=3',
          {
            headers: {
              Authorization: `Bearer ${accessToken}`,
            },
          }
        );
        setData(response.data.data.content); // Adjusted to match your response format
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

  const transactions = data.slice(0, 3); // Ensure we're getting the first 3 transactions

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
      <div className='bg-white rounded-[25px] p-7'>
        {transactions.map((transaction, index) => (
          <div>
          <div key={index} className="flex items-center justify-between space-x-8 mb-4 ">
            <div className="flex items-center space-x-4">
              <Image
                height={44}
                width={44}
                src={getIconPath(transaction.type)} // Icon based on transaction type
                alt={transaction.type}
                className="object-cover"
              />
=======
export const RecentTransactionCard = () => {
  const slicedTransactions = RecentTransaction.slice(0, 3);

  return (
    <div>
      <p>Recent Transactions</p>
      <div>
        {slicedTransactions.map((transaction, index) => (
          <div key={index} className="flex items-center justify-between space-x-8 mb-4">
            <div className="flex items-center justify-between space-x-8 mb-4">
              <Image height={44} width={44} src={transaction.image} alt="invoice" className="rounded-full object-cover" />
>>>>>>> aastu.web.g5.yetnayet.transactions
              <div>
                <p className="font-semibold text-sm md:text-base">{transaction.name}</p>
                <p className="text-xs md:text-sm text-gray-500">{transaction.date}</p>
              </div>
            </div>
            {transaction.amount[0]=="+" ? 
            <p className="font-semibold text-green-600 text-sm md:text-base">{transaction.amount}</p>
            : <p className="font-semibold text-red-700 text-sm md:text-base">{transaction.amount}</p> }
          </div>
          </div>
        ))}
      </div>
    </div>
  );
};


export default RecentTransactionCard;