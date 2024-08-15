import React from "react";
import { CircleArrowDown, CircleArrowUp } from "lucide-react";
interface RecentTransactionDescriptionProps {
  amount: number;
  description: string;
}

const RecentTransactionDescription = ({
  amount,
  description,
}: RecentTransactionDescriptionProps) => {
  return (
    <div className="flex flex-row flex-wrap gap-3 items-center">
      <div>
        {amount < 0 ? (
          <CircleArrowDown className="w-8 h-8 text-blue-steel" />
        ) : (
          <CircleArrowUp className="w-8 h-8 text-blue-steel"/>
        )}
      </div>
      <div>
        <p>{description}</p>
      </div>
    </div>
  );
};

export default RecentTransactionDescription;
