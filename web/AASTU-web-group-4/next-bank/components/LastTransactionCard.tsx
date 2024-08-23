import { useEffect, useState } from "react";
import { getAllTransactions } from "@/services/transactionfetch";
import { currentuser } from "@/services/userupdate";
import TransactionCard from "./TransactionCard";

// App Component
const App: React.FC = () => {
  const [transactions, setTransactions] = useState<[]>([]);
  const [currentUser,setCurrentuser] = useState("")

  useEffect(() => {
    const fetchTransactions = async () => {
      try {
        const transactionData = await getAllTransactions(0, 5); 
        const current = await currentuser();
        setCurrentuser(current.data.name)

        console.log("on the card user:" , currentuser); 
        
        if (Array.isArray(transactionData.data.content)) {
          setTransactions(transactionData.data.content);
        } else {
          console.error("Transaction data is not an array");
        }
      } catch (error) {
        console.error("Failed to fetch transactions", error);
      }
    };

    fetchTransactions();
  }, []);

  return (
    <div className="p-3 gap-4 flex-1 h-auto bg-gray-50 dark:bg-dark text-gray-900 dark:text-white">
      {transactions.map((transaction, index) => (
        <TransactionCard key={index} transaction={transaction} currentname={currentUser} />
      ))}
    </div>
  );
};

export default App;