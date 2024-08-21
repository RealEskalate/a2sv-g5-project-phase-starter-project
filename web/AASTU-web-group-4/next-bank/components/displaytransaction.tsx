import React from 'react'
import { FaArrowCircleUp, FaArrowCircleDown } from 'react-icons/fa';

import Pagination from '@/components/Pagination';
const displaytransaction = (alltransaction : any ,type:any ) => {
  return (
    <div>
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
            {alltransaction.map((transaction : any, index:any) => (
              <tr
                key={transaction.transactionId}
                className={`border-b border-gray-200 ${index % 2 === 0 ? 'bg-gray-50' : 'bg-white'} hover:bg-gray-100 transition-colors duration-300`}
              >
                <td className="p-4 flex items-center text-sm text-gray-700 truncate">
                  {(transaction.type === 'deposit' || type === 'income')? (
                    <FaArrowCircleDown className="text-green-500 text-lg mr-2" />
                  ) : (
                    <FaArrowCircleUp className="text-red-500 text-lg mr-2" />
                  )}
                  {transaction.description}
                </td>
                <td className="p-4 text-sm text-gray-600">{transaction.transactionId}</td>
                <td className="p-4 text-sm text-gray-600">{transaction.type}</td>
                <td className="p-4 text-sm text-gray-600">Card Name</td>
                <td className="p-4 text-sm text-gray-600">{transaction.date}</td>
                <td className={`p-4 text-sm ${transaction.type === 'deposit' ? 'text-green-500' : 'text-red-500'}`}>
                  {transaction.type === 'deposit' ? (
                    <div> 
                      +
                      {transaction.amount}
                      $</div>
                    ) : (
                      <div> 
                      -
                      {transaction.amount}
                      $
                      </div>
                    )}
                  
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
        {alltransaction.map((transaction:any) => (
          <div key={transaction.transactionId} className="flex justify-between bg-white p-4 mb-2 rounded-lg shadow-sm border border-gray-200 items-center">
            <div>
              <div className="flex items-center mb-2">
                {transaction.type === 'deposit' ? (
                  <FaArrowCircleDown className="text-green-500 text-xl mr-2" />
                ) : (
                  <FaArrowCircleUp className="text-red-500 text-xl mr-2" />
                )}
                <span className="font-semibold">{transaction.description}</span>
              </div>
              <div className="text-[12px] text-gray-400 mb-1 pl-5">{transaction.date}</div>
            </div>
            <div>
              <div className={`font-bold ${(transaction.type === 'deposit')? 'text-green-500' : 'text-red-500'}`}>
              {transaction.type=== 'deposit'? (
                    <div> 
                      +
                      {transaction.amount}
                      $</div>
                    ) : (
                      <div> 
                      -
                      {transaction.amount}
                      $
                      </div>
                    )}
              </div>
            </div>
          </div>
        ))}
      </div>

      {/* Pagination */}
     
    </div>
    
  )
}

export default displaytransaction
