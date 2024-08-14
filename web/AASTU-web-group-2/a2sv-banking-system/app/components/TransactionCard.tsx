import React, { ReactNode } from "react";

interface Props {
  transactionIcon: ReactNode;
  amount: number;
  date: string;
  depositor: string;
}
const TransactionCard = ({
  transactionIcon,
  amount,
  date,
  depositor,
}: Props) => {
  return (
    <div className="flex p-4 w-[300px] gap-3 items-center">
      {transactionIcon}
      <div>
        <p className="text-[#232323] font-medium">{depositor}</p>
        <p className="text-[#8297c0]">{date}</p>
      </div>
      {amount >= 0 ? (
        <p className={`ml-auto text-[#41D4A8]`}>{"+$" + amount}</p>
      ) : (
        <p className={`ml-auto text-[#FF4B4A]`}>{"-$" + Math.abs(amount)}</p>
      )}
    </div>
  );
};

export default TransactionCard;
