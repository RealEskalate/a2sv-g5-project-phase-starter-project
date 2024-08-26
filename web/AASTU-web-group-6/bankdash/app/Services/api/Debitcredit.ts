
interface TransactionType {
  transactionId: string;
  type: string;
  senderUserName: string;
  description: string;
  date: string;
  amount: number;
  receiverUserName: string | null;
}

interface GroupedTransactions {
  [weekKey: string]: {
    income: { [dayKey: string]: number };
    expense: { [dayKey: string]: number };
  };
}
function groupTransactionsByWeek(
  transactions: TransactionType[]
): GroupedTransactions {
  const grouped: GroupedTransactions = {};

  transactions.forEach((transaction) => {
    const date = new Date(transaction.date);
    const day = date.getDay();
    const startOfWeek = new Date(date);
    startOfWeek.setDate(date.getDate() - day);

    const weekKey = startOfWeek.toISOString().split("T")[0];

    if (!grouped[weekKey]) {
      grouped[weekKey] = { income: {}, expense: {} };
    }

    const typeGroup = transaction.type === "income" ? "income" : "expense";
    const dayKey = date.toISOString().split("T")[0];

    if (!grouped[weekKey][typeGroup][dayKey]) {
      grouped[weekKey][typeGroup][dayKey] = 0;
    }

    grouped[weekKey][typeGroup][dayKey] += transaction.amount;
  });

  return grouped;
}

const DebitCredit =  (accessToken : string , income:any , expense:any) => {
   
    
    const income_transactions = income.map(
      (transaction: TransactionType) => ({
        ...transaction,
        type: "income",
      })
    );
    
    const expense_transactions = expense.map(
          (transaction: TransactionType) => ({
            ...transaction,
            type: "expense",
          })
        );
    const combinedTransactions = [
      ...expense_transactions,
      ...income_transactions,
    ];
    const ans = groupTransactionsByWeek(combinedTransactions);
    return ans;
  
};

export default DebitCredit;
