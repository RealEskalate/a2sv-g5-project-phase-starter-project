import { Inter } from "next/font/google";

export interface TransactionType {
  description: string;
  transactionId: string;
  type: string;
  card?: string;
  date: string;
  amount: number;
}

const inter = Inter({ subsets: ["latin"] });
const Transaction = ({
  description,
  transactionId,
  type,
  card = "",
  date,
  amount,
}: TransactionType) => {
  return (
    <tr className="border-t border-[#E6EFF5]">
      <td className="py-6 px-6 flex items-center gap-3">
        <img
          src={
            amount < 0
              ? "/assets/transaction/withdraw.svg"
              : "/assets/transaction/deposit.svg"
          }
          alt="icon"
          className="w-6 h-6"
        />
        {description}
      </td>
      <td className="py-3 px-6">{transactionId}</td>
      <td className="py-3 px-6">{type}</td>
      <td className="py-3 px-6">{card}</td>
      <td className="py-3 px-6">{date}</td>
      <td
        className={`py-3 px-6 ${
          type.toLowerCase() != "deposit" ? "text-red-500" : "text-green-500"
        }`}
      >
        {type.toLowerCase() != "deposit"
          ? `-$${Math.abs(amount)}`
          : `+$${amount}`}
      </td>
      <td className="py-3 px-6">
        <button
          className={`${inter.className} border-[1px] border-[#123288] text-[#123288] rounded-3xl px-4 py-2 hover:border-[#1814F3] hover:text-[#1814F3]`}
        >
          Download
        </button>
      </td>
    </tr>
  );
};

export default Transaction;
