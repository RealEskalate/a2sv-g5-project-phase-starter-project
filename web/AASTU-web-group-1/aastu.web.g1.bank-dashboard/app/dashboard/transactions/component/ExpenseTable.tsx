import React from "react";
import { CiSaveDown1, CiSaveUp1 } from "react-icons/ci";
import { getCurrentUser } from "./getCurrentUser";
import { TransactionData, UserData } from "@/types";



    


export const ExpenseTable: React.FC<{ transactions: TransactionData[] }> = ({
  transactions,
}) => {
  const [currentUser, setCurrentUser] = React.useState<UserData | null>(null);

  React.useEffect(() => {
    getCurrentUser().then(user => {
      if (user) {
        setCurrentUser(user);
      }
    });
  }, []);
  return (
    <div className="w-[100%] rounded-3xl shadow-md">
      <table className="w-[100%] divide-y divide-gray-200  rounded-xl md:shadow-lg md:border md:border-gray-300  ">
        <thead className="bg-white">
          <tr className="border-b-2">
            <th
              scope="col"
              className="lg:px-5 py-3 md:px-2 text-left text-xs font-medium text-gray-500 uppercase truncate md:max-w-[6rem] lg:max-w-[10rem]"
            >
              Description
            </th>
            <th
              scope="col"
              className="lg:px-5 py-3 md:px-2 text-left text-xs font-medium text-gray-500 uppercase truncate hidden md:table-cell md:max-w-[6rem]"
            >
              Transaction ID
            </th>
            <th
              scope="col"
              className="lg:px-5 py-3 md:px-2 text-left text-xs font-medium text-gray-500 uppercase hidden md:table-cell md:max-w-[6rem]"
            >
              Type
            </th>
            <th
              scope="col"
              className="lg:px-5 py-3 md:px-2 text-left text-xs font-medium text-gray-500 uppercase hidden md:table-cell md:max-w-[6rem]"
            >
              Card
            </th>
            <th
              scope="col"
              className="lg:px-5 py-3 md:px-2 text-left text-xs font-medium text-gray-500 uppercase hidden md:table-cell md:max-w-[6rem]"
            >
              Date
            </th>
            <th
              scope="col"
              className="lg:px-5 py-3 md:px-2 text-left text-xs font-medium text-gray-500 uppercase md:max-w-[4rem]"
            >
              Amount
            </th>
            <th
              scope="col"
              className="lg:px-5 py-3 md:px-2 text-left text-xs font-medium text-gray-500 uppercase hidden md:table-cell md:max-w-[6rem]"
            >
              Receipt
            </th>
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {transactions.map((transaction) => (
            <tr key={transaction.transactionId}>
              {/* Description */}
              <td className="lg:px-5 py-4 px-2 items-center text-sm truncate lg:max-w-[10rem] md:max-w-[6rem]">
                <div className="space-x-2 flex md:max-w-[6rem] lg:max-w-[10rem]">
                  <span className="inline-block align-middle">
                    {transaction.senderUserName === currentUser?.username ? (
                      <CiSaveUp1 size={20} />
                    ) : (
                      <CiSaveDown1 size={20} />
                    )}
                  </span>
                  <span
                    className="inline-block align-middle truncate "
                    title={transaction.description}
                  >
                    {transaction.description}
                  </span>
                </div>
              </td>
              {/* ID */}
              <td className="lg:px-5 py-4 md:px-2 text-sm truncate hidden md:table-cell md:max-w-[4rem]">
                {transaction.transactionId}
              </td>
              {/* Type */}
              <td className="lg:px-5 py-4 md:px-2 text-sm truncate hidden md:table-cell md:max-w-[6rem]">
                {transaction.type}
              </td>
              {/* reciver*/}
              <td className="lg:px-5 py-4 md:px-2 text-sm truncate hidden md:table-cell md:max-w-[6rem]">
                {transaction.receiverUserName !== null
                  ? transaction.receiverUserName
                  : "unknown"}
              </td>
              {/* Date */}
              <td className="lg:px-5 py-4 md:px-2 text-sm truncate hidden md:table-cell md:max-w-[6rem]">
                {transaction.date}
              </td>
              {/* Amount */}
              <td className="lg:px-5 py-4 md:px-2 text-sm truncate md:max-w-[4rem]">
                <p
                  className={
                    transaction.senderUserName !== currentUser?.username
                      ? "text-green-500"
                      : "text-red-500"
                  }
                >
                  ${transaction.amount}
                </p>
              </td>
              {/* Receipt */}
              <td className="lg:px-5 py-4 md:px-2 text-xs truncate hidden md:table-cell md:max-w-[3rem]">
                <button className="bg-transparent hover:bg-blue-500 text-blue-700 font-semibold hover:text-white py-2 px-1 border border-blue-500 hover:border-transparent rounded-2xl text-xs">
                  Download
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};
