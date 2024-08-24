// "use client"

import React from 'react';
import Image from 'next/image';

import upCircled from '../../public/up-circled.svg'
import downCircled from '../../public/down-circled.svg'

const transactions = [
  {
    description: 'Spotify Subscription',
    transactionId: '#12548796',
    type: 'Shopping',
    card: '1234 ****',
    date: '28 Jan, 12.30 AM',
    amount: -2500,
    receipt: true,
  },
  {
    description: 'Freepik Sales',
    transactionId: '#12548796',
    type: 'Transfer',
    card: '1234 ****',
    date: '25 Jan, 10.40 PM',
    amount: 750,
    receipt: true,
  },
  {
    description: 'Mobile Service',
    transactionId: '#12548796',
    type: 'Service',
    card: '1234 ****',
    date: '20 Jan, 10.40 PM',
    amount: -150,
    receipt: true,
  },
  {
    description: 'Wilson',
    transactionId: '#12548796',
    type: 'Transfer',
    card: '1234 ****',
    date: '15 Jan, 03.29 PM',
    amount: -1050,
    receipt: true,
  },
  {
    description: 'Emilly',
    transactionId: '#12548796',
    type: 'Transfer',
    card: '1234 ****',
    date: '14 Jan, 10.40 PM',
    amount: 840,
    receipt: true,
  },
];

const RecentTransactions: React.FC = () => {
  return (
    <div className="overflow-x-auto bg-white rounded-3xl shadow-md p-4 max-sm:p-0 my-2 ">
      <table className="min-w-full  p-4 max-md:p-0  font-body max-md:text-[12px]  md:w-[743px]">
        <thead className=' border-b'>
          <tr className='font-medium text-base text-custom-light-purple max-sm:hidden'>
            <th className="px-6 py-4 text-left ">
              Description
            </th>
            <th className="px-6 py-4 text-left ">
              Transaction ID
            </th>
            <th className="px-6 py-4 text-left ">
              Type
            </th>
            <th className="px-6 py-4 text-left ">
              Card
            </th>
            <th className="px-6 py-4 text-left ">
              Date
            </th>
            <th className="px-6 py-4 text-left ">
              Amount
            </th>
            <th className="px-6 py-4 text-left ">
              Receipt
            </th>
          </tr>
        </thead>
        <tbody className='font-medium text-base font-body'>
          {transactions.map((transaction, index) => (
            <tr key={index} className="border-b last:border-none">
              <td className="px-6 py-4 text-gray-900 flex items-center gap-2 max-sm:py-1 max-sm:text-[15px]">
                
                 {transaction.type === 'Shopping' || transaction.type === 'Service' ? (
                  <Image src={upCircled} alt='arror'  />

                ) : (
                    <Image src={downCircled} alt='arror'  />


                )}{' '}
                {transaction.description}
              </td>
              <td className="px-6 py-4  text-gray-900 max-sm:hidden">{transaction.transactionId}</td>
              <td className="px-6 py-4  text-gray-900 max-sm:hidden">{transaction.type}</td>
              <td className="px-6 py-4  text-gray-900 max-sm:hidden">{transaction.card}</td>
              <td className="px-6 py-4  text-gray-900 max-sm:block max-sm:text-[#718EBF] max-sm:py-0  max-sm:w-[200px] max-sm:text-[12px] max-sm:text-center">{transaction.date}</td>
              <td className="px-6 py-4  font-medium">
                <span className={transaction.amount < 0 ? 'text-custom-pink-red' : 'text-custom-greenish'}>
                  {transaction.amount < 0 ? `-$${Math.abs(transaction.amount)}` : `+$${transaction.amount}`}
                </span>
              </td>
              <td className="px-6 py-4  text-blue-600">
                {transaction.receipt && (
                  <button className="px-2 py-1 font-normal text-sm  text-custom-purple hover:bg-blue-50 border border-custom-purple border-1 rounded-3xl max-sm:hidden">
                    Download
                  </button>
                )}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default RecentTransactions;

