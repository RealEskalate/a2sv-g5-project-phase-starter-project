import { AppWindowMacIcon } from 'lucide-react';
import Image from 'next/image';
// import { TbMoneybag } from "react-icons/tb";

interface BalanceCardProps {
  imageSrc: string;
  altText: string;
  title: string;
  balance: string;
}

const BalanceCard: React.FC<{ accountData: BalanceCardProps }> = ({ accountData }) => {
  return (
    <div className="bg-white rounded-lg shadow-md p-3 flex items-center space-x-3 w-full">
      {/* Icon */}
      <div className="bg-yellow-100 rounded-full p-2">
        <Image 
          src={accountData.imageSrc} // Replace with your icon path
          alt={accountData.altText}
          width={20}
          height={20}
        />
      </div>

      {/* Balance Details */}
      <div>
        <p className="text-blue-500 text-xs">{accountData.title}</p>
        <p className="text-black text-lg font-bold">{accountData.balance}</p>
      </div>
    </div>
  );
}

const accountDatas: BalanceCardProps[] = [
  {
    imageSrc: "/icons/money-bag-icon.svg",
    altText: "Money Bag Icon",
    title: "My Balance",
    balance: "$12,750",
  },
  {
    imageSrc: "/icons/income-icon.svg",
    altText: "Income Icon",
    title: "Income",
    balance: "$5,600",
  },
  {
    imageSrc: "/icons/expense-icon.svg",
    altText: "Expense Icon",
    title: "Expense",
    balance: "$3,460",
  },
  {
    imageSrc: "/icons/savings-icon.svg",
    altText: "Savings Icon",
    title: "Total Saving",
    balance: "$7,920",
  }
];

const App: React.FC = () => {
  return (
    <div className="flex flex-row gap-4 w-full justify-between">
      {accountDatas.map((accountData, index) => (
        <div key={index} className="w-full sm:w-1/2 lg:w-1/4">
          <BalanceCard accountData={accountData} />
        </div>
      ))}
    </div>
  );
};

export default App;
