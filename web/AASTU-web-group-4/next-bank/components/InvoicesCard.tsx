import React from 'react';

interface Transaction {
  id: number;
  icon: string;
  name: string;
  time: string;
  amount: string;
  iconBgColor: string;
}

const transactions: Transaction[] = [
  {
    id: 1,
    icon: '🍏', // Placeholder for Apple Store icon
    name: 'Apple Store',
    time: '5h ago',
    amount: '$450',
    iconBgColor: 'bg-teal-100'
  },
  {
    id: 2,
    icon: '👤', // Placeholder for person icon
    name: 'Michael',
    time: '2 days ago',
    amount: '$160',
    iconBgColor: 'bg-yellow-100'
  },
  {
    id: 3,
    icon: '🎮', // Placeholder for Playstation icon
    name: 'Playstation',
    time: '5 days ago',
    amount: '$1085',
    iconBgColor: 'bg-blue-100'
  },
  {
    id: 4,
    icon: '👤', // Placeholder for person icon
    name: 'William',
    time: '10 days ago',
    amount: '$90',
    iconBgColor: 'bg-pink-100'
  },
];

const TransactionList: React.FC = () => {
  return (
    <div className="flex-1 flex flex-col justify-between bg-white rounded-lg shadow-md p-4 space-y-4">
      {transactions.map(transaction => (
        <div key={transaction.id} className="flex items-center justify-between">
          <div className={`w-10 h-10 flex items-center justify-center rounded-full ${transaction.iconBgColor}`}>
            <span className="text-2xl">{transaction.icon}</span>
          </div>
          <div className="flex-1 px-4">
            <div className="text-gray-800 font-medium">{transaction.name}</div>
            <div className="text-gray-400 text-sm">{transaction.time}</div>
          </div>
          <div className="text-gray-800 font-semibold">{transaction.amount}</div>
        </div>
      ))}
    </div>
  );
};

export default TransactionList;

// const TransactionList: React.FC = () => {
//   return (
//     <div className="w-auto flex  flex-col justify-between bg-white rounded-lg shadow-md p-4 space-y-4">
//       {transactions.map(transaction => (
//         <div key={transaction.id} className="flex w-auto items-center justify-between">
//           <div className={`w-10 h-10 flex items-center justify-center rounded-full ${transaction.iconBgColor}`}>
//             <span className="text-2xl">{transaction.icon}</span>
//           </div>
//           <div className="flex-1">
//             <div className="text-gray-800 font-medium">{transaction.name}</div>
//             <div className="text-gray-400 text-sm">{transaction.time}</div>
//           </div>
//           <div className="text-gray-800 font-semibold">{transaction.amount}</div>
//         </div>
//       ))}
//     </div>
//   );
// };

// export default TransactionList;
