import { BanknotesIcon } from "@heroicons/react/24/outline";

// Interface for the Transaction data
interface Transaction {
  receiverUserName: string;
  type: string;
  amount: string;
  date: string;
}

interface Currentname {
    name:string,
   
}

// Transaction Card Component
const TransactionCard: React.FC<{ transaction: Transaction,currentname:string }> = ({ transaction,currentname}) => {
  const isPositive = currentname === transaction.receiverUserName;
  const amountColor = isPositive ? "text-green-500" : "text-red-500"
  const sign = isPositive ? "+" : "-";
  console.log("current user :",currentname)
  console.log("reciever :", transaction.receiverUserName)

  const formattedAmount = `${sign}${'$'}${Math.abs(Number(transaction.amount)).toLocaleString()}`;

  return (
    <div className="flex flex-col pt-4 md:flex-row md:justify-evenly gap-4 gap-y-6 w-auto rounded-2xl shadow-none border-none">
      {/* Desktop View */}
      <div className="hidden md:flex items-center">
        <div className="w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center mr-4">
          <BanknotesIcon className="h-8 w-8 text-green-500"/>
          {/* {transaction.icon} */}
        </div>
        <div className="w-40 truncate"> {/* Fixed width with truncation */}
          <h3 className="text-base font-semibold truncate">{transaction.receiverUserName}</h3>
          <p className="text-sm text-gray-500 truncate">{transaction.date}</p>
        </div>
      </div>
      <div className="hidden md:flex items-center w-24 truncate">
        <p className="text-sm font-medium text-gray-600 truncate">{transaction.type}</p>
      </div>
      <div className="hidden md:flex items-center w-28 truncate"> 
        <p className="text-sm text-gray-500 truncate">1234 ****</p>
      </div>
      <div className="hidden md:flex items-center w-20 truncate">
        <p className="truncate">Completed</p>
      </div>
      <div className="hidden md:flex items-center justify-end w-24 truncate">
        <p className={`text-lg  ${amountColor} truncate`}>
        {formattedAmount}
        </p>
      </div>
      
      {/* Mobile View */}
      <div className="md:hidden flex justify-between flex-row w-full">
        <div className="flex items-center">
          <div className="w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center mr-4">
          <BanknotesIcon className="h-8 w-8 text-green-500"/>
          </div>
          <div>
            <h3 className="text-base font-semibold truncate">{transaction.receiverUserName}</h3>
            <p className="text-sm text-gray-500 truncate">{transaction.date}</p>
          </div>
        </div>
        <div className="flex items-center">
          <p className={`text-lg ${amountColor} truncate`}>
            {formattedAmount}
          </p>
        </div>
      </div>
    </div>
  );
};
 export default TransactionCard