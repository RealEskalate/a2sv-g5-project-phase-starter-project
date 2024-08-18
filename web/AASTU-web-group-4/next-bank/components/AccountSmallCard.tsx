import Image from 'next/image';

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
      <div className="bg-yellow-100 rounded-xl p-2">
        <Image 
          src={accountData.imageSrc} // Replace with your icon path
          alt={accountData.altText}
          width={20}
          height={40}
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
    <div className="grid grid-cols-2 sm:grid-cols-2 lg:grid-cols-4 gap-4 w-full">
      {accountDatas.map((accountData, index) => (
        <BalanceCard key={index} accountData={accountData} />
      ))}
    </div>
  );
};


export default App;
