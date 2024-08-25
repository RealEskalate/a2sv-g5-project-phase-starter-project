import React from "react";
import { CircleArrowDown, CircleArrowUp } from "lucide-react";
interface RecentTransactionDescriptionProps {
  amount: string;
  description: string;
}

const RecentTransactionDescription = ({
  amount,
  description,
}: RecentTransactionDescriptionProps) => {
  return (
    <div className="flex flex-row flex-wrap gap-3 items-center">
      <div>
        {amount === 'deposit' ? (
          <CircleArrowDown className="w-6 h-6 text-blue-steel" data-testid="downArrow"/>
        ) : (
          <CircleArrowUp className="w-6 h-6 text-blue-steel" data-testid="upArrow"/>
        )}
      </div>
        <p >{description}</p>
    </div>
  );
};

export default RecentTransactionDescription;
