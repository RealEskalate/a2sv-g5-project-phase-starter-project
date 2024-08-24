import axios from "axios";
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

const DebitCredit = async (accessToken : string) => {
  try {
    const expense_count = await axios.get(
      "https://bank-dashboard-rsf1.onrender.com/transactions/expenses?page=0&size=1",
      {
        headers: {
          Authorization:
          accessToken ? `Bearer ${accessToken}` : undefined,
        },
      }
    );
    const total_expense = expense_count.data.data.totalPages;
    const expense = await axios.get(
      `https://bank-dashboard-rsf1.onrender.com/transactions/expenses?page=0&size=${total_expense}`,
      {
        headers: {
          Authorization:
            accessToken ? `Bearer ${accessToken}` : undefined,
        },
      }
    );
    console.log(expense, "debit result");
    const expense_transactions = expense.data.data.content.map(
      (transaction: TransactionType) => ({
        ...transaction,
        type: "expense",
      })
    );
    const income_count = await axios.get(
      "https://bank-dashboard-rsf1.onrender.com/transactions/incomes?page=0&size=1",
      {
        headers: {
          Authorization:
            accessToken ? `Bearer ${accessToken}` : undefined,
        },
      }
    );
    const total_income = income_count.data.data.totalPages;
    const income = await axios.get(
      `https://bank-dashboard-rsf1.onrender.com/transactions/incomes?page=0&size=${total_income}`,
      {
        headers: {
          Authorization:
            accessToken ? `Bearer ${accessToken}` : undefined,
        },
      }
    );
    console.log(income, "income result");
    const income_transactions = income.data.data.content.map(
      (transaction: TransactionType) => ({
        ...transaction,
        type: "income",
      })
    );

    const combinedTransactions = [
      ...expense_transactions,
      ...income_transactions,
    ];
    const ans = groupTransactionsByWeek(combinedTransactions);
    return ans;
  } catch (error) {
    console.log(error, "error");
  }
};

export default DebitCredit;
