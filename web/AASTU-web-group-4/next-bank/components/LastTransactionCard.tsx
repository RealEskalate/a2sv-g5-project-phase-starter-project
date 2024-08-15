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
      className="w-4 h-4 text-teal-600"
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
      className="w-4 h-4 text-indigo-600"
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
      className="w-4 h-4 text-pink-600"
    >
      <path
        fillRule="evenodd"
        d="M8 8a3 3 0 100-6 3 3 0 000 6zm3.5 3c-1.38 0-2.629.56-3.5 1.463A5.978 5.978 0 004.5 11C2.57 11 1 12.57 1 14.5V15h14v-.5c0-1.93-1.57-3.5-3.5-3.5z"
      />
    </svg>
  </div>
);

// Transaction Card Component
const TransactionCard: React.FC<{ transaction: Transaction }> = ({
  transaction,
}) => {
  return (
    <div >
      <div className="hidden md:block flex items-center w-full md:w-auto">
        <div className="w-12 h-12 rounded-full bg-gray-100 flex items-center justify-center mr-4">
          {transaction.icon}
        </div>
        <div className="hidden md:block">
          <h3 className="text-base font-semibold">{transaction.title}</h3>
          <p className="text-sm text-gray-500">{transaction.date}</p>
        </div>
      </div>
      <div className="hidden md:block">
        <p className="text-sm font-medium text-gray-600">{transaction.type}</p>
      </div>
      <div className="hidden md:block">
        <p className="text-sm text-gray-500">{transaction.card}</p>
      </div>
      <div className="hidden md:block">
        <p>{transaction.status}</p>
      </div>
      <div className="hidden md:block">
        <p
          className={`text-lg ${
            transaction.amount.startsWith("+")
              ? "text-green-500"
              : "text-red-500"
          }`}
        >
          {transaction.amount}
        </p>
      </div>
      {/* Mobile view - only show the first and last divs */}
      <div className="md:hidden flex space-x-3 justify-between">
        <div className="flex ">
          <div className="w-12 h-12 rounded-full bg-gray-100 flex items-center justify-center mr-4">
            {transaction.icon}
          </div>
          <div>
            <h3 className="text-base font-semibold">{transaction.title}</h3>
            <p className="text-sm text-gray-500">{transaction.date}</p>
          </div>
        </div>
        <div>
          <p
            className={`text-lg ${
              transaction.amount.startsWith("+")
                ? "text-green-500"
                : "text-red-500"
            }`}
          >
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
    type: "Service",
    card: "1234 ****",
    status: "Completed",
    amount: "-$340",
    date: "25 Jan 2021",
  },
  {
    icon: <UserIcon />,
    title: "Emilly Wilson",
    type: "Transfer",
    card: "1234 ****",
    status: "Completed",
    amount: "+$780",
    date: "25 Jan 2021",
  },
];

// App Component
const App: React.FC = () => {
  return (
    <div className="p-6 bg-gray-50 min-h-screen">
      {transactions.map((transaction, index) => (
        <TransactionCard key={index} transaction={transaction} />
      ))}
    </div>
  );
};

export default App;
