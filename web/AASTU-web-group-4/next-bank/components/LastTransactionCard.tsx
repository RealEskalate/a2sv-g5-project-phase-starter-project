import { useEffect, useState } from "react";
import { getAllTransactions } from "@/services/transactionfetch";
import { currentUser } from "@/services/userupdate";
import TransactionCard from "./TransactionCard";

// App Component
const App: React.FC = () => {
  const [transactions, setTransactions] = useState<[]>([]);
  const [currentuser,setCurrentuser] = useState("")

  useEffect(() => {
    const fetchTransactions = async () => {
      try {
        const transactionData = await getAllTransactions(0, 5); 
        const current = await currentUser();
        setCurrentuser(current.data.name)

        console.log("on the card user:" , currentUser); 
        
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
        <TransactionCard key={index} transaction={transaction} currentname={currentuser} />
      ))}
    </div>
  );
};

export default App;