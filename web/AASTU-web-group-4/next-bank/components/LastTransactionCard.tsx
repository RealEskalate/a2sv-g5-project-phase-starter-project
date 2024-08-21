import React from "react";

// Interface for the Transaction data
interface Transaction {
  icon: JSX.Element;
  title: string;
  type: string;
  card: string;
  status: string;
  amount: string;
  date: string;
}

// Sample Icons (Using React Icons library or SVGs)
const SpotifyIcon = () => (
  <div className="w-6 h-6 rounded-full bg-teal-100 flex items-center justify-center">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="currentColor"
      viewBox="0 0 16 16"
      className="w-full h-full text-teal-600"
    >
      <path d="M8 0a8 8 0 100 16A8 8 0 008 0zm3.749 11.45a.62.62 0 01-.852.198c-2.34-1.396-5.292-1.71-8.756-.932a.623.623 0 11-.278-1.218c3.75-.86 6.993-.498 9.552 1.092a.623.623 0 01.198.86zm1.02-2.306a.778.778 0 01-1.067.247c-2.692-1.605-6.812-2.07-10.008-1.123a.777.777 0 11-.432-1.492c3.527-.956 8.068-.423 11.087 1.298.367.218.484.69.247 1.07zm.084-2.26c-3.009-1.782-7.802-2.27-10.603-1.235a.937.937 0 01-.585-1.78c3.27-1.076 8.608-.539 12.163 1.455a.938.938 0 01-.975 1.607z" />
    </svg>
  </div>
);

const MobileIcon = () => (
  <div className="w-6 h-6 rounded-full bg-indigo-100 flex items-center justify-center">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="currentColor"
      viewBox="0 0 16 16"
      className="w-full h-full text-indigo-600"
    >
      <path d="M11 1H5a2 2 0 00-2 2v10a2 2 0 002 2h6a2 2 0 002-2V3a2 2 0 00-2-2zm-3 13a1 1 0 110-2 1 1 0 010 2zm3-3H5V4h6v7z" />
    </svg>
  </div>
);

const UserIcon = () => (
  <div className="w-6 h-6 rounded-full bg-pink-100 flex items-center justify-center">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="currentColor"
      viewBox="0 0 16 16"
      className="w-full h-full text-pink-600"
    >
      <path
        fillRule="evenodd"
        d="M8 8a3 3 0 100-6 3 3 0 000 6zm3.5 3c-1.38 0-2.629.56-3.5 1.463A5.978 5.978 0 004.5 11C2.57 11 1 12.57 1 14.5V15h14v-.5c0-1.93-1.57-3.5-3.5-3.5z"
      />
    </svg>
  </div>
);

// Transaction Card Component
const TransactionCard: React.FC<{ transaction: Transaction }> = ({ transaction }) => {
  return (
    <div className="flex flex-col pt-4 md:flex-row gap-4 gap-y-6 w-auto rounded-2xl shadow-none border-none">
      {/* Desktop View */}
      <div className="hidden md:flex items-center">
        <div className="w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center mr-4">
          {transaction.icon}
        </div>
        <div className="w-40 truncate"> {/* Fixed width with truncation */}
          <h3 className="text-base font-semibold truncate">{transaction.title}</h3>
          <p className="text-sm text-gray-500 truncate">{transaction.date}</p>
        </div>
      </div>
      <div className="hidden md:flex items-center w-24 truncate">
        <p className="text-sm font-medium text-gray-600 truncate">{transaction.type}</p>
      </div>
      <div className="hidden md:flex items-center w-32 truncate"> 
        <p className="text-sm text-gray-500 truncate">{transaction.card}</p>
      </div>
      <div className="hidden md:flex items-center w-20 truncate">
        <p className="truncate">{transaction.status}</p>
      </div>
      <div className="hidden md:flex items-center justify-end w-24 truncate">
        <p className={`text-lg ${transaction.amount.startsWith('+') ? 'text-green-500' : 'text-red-500'} truncate`}>
          {transaction.amount}
        </p>
      </div>
      
      {/* Mobile View */}
      <div className="md:hidden flex justify-between flex-row w-full">
        <div className="flex items-center">
          <div className="w-12 h-12 rounded-full bg-gray-100 flex items-center justify-center mr-4">
            {transaction.icon}
          </div>
          <div>
            <h3 className="text-base font-semibold truncate">{transaction.title}</h3>
            <p className="text-sm text-gray-500 truncate">{transaction.date}</p>
          </div>
        </div>
        <div className="flex items-center">
          <p className={`text-lg ${transaction.amount.startsWith("+") ? "text-green-500" : "text-red-500"} truncate`}>
            {transaction.amount}
          </p>
        </div>
      </div>
    </div>
  );
};

// Sample Transaction Data
const transactions: Transaction[] = [
  {
    icon: <SpotifyIcon />,
    title: "Spotify Subscription",
    type: "Shopping",
    card: "1234 ****",
    status: "Pending",
    amount: "-$150",
    date: "25 Jan 2021",
  },
  {
    icon: <MobileIcon />,
    title: "Mobile Service",
    type: "Shopping",
    card: "1234 ****",
    status: "Completed",
    amount: "-$340",
    date: "25 Jan 2021",
  },
  {
    icon: <UserIcon />,
    title: "Emilly Wilson",
    type: "Shopping",
    card: "1234 ****",
    status: "Completed",
    amount: "+$780",
    date: "25 Jan 2021",
  },
];

// App Component
const App: React.FC = () => {
  return (
    <div className="p-3 gap-4 flex-1 h-auto bg-gray-50">
      {transactions.map((transaction, index) => (
        <TransactionCard key={index} transaction={transaction} />
      ))}
    </div>
  );
};

export default App;
