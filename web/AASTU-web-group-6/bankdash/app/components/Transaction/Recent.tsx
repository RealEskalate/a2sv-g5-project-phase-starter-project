import React from "react";
import { ArrowUpCircleIcon } from "@heroicons/react/24/outline";

const Recent = () => {
  return (
    <div className="space-y-7 mb-8">
      <h3 className="font-semibold text-[22px] text-[#343C6A]">
        Recent Transactions
      </h3>
      <div className="flex gap-16 px-1 border-b text-[16px] text-[#718EBF] font-semibold">
        <p className=" border-[#1814F3] border-b-[3px] pb-2 text-[#1814F3] ">
          All Transactions
        </p>
        <p> Income </p>
        <p> Expense </p>
      </div>
      <div className=" w-full bg-white rounded-[25px] px-8 py-6">
        <table className="border-separate border-spacing-y-4 font-[16px] w-full transaction-table ">
          <thead>
            <tr className="text-[#718EBF] text-left">
              <th> Description </th>
              <th> Transaction ID</th>
              <th> Type </th>
              <th> Card </th>
              <th> Date </th>
              <th> Amount </th>
              <th> Receipt </th>
            </tr>
          </thead>
          <tbody className="text-[#232323] p-8 space-y-4">
            <tr>
              {/* Their is goin to be an icon if it is deposit or  */}
              <td className="flex gap-2 items-center">
                <ArrowUpCircleIcon className="transaction-icon" />
                Spotify Subscription{" "}
              </td>
              <td> #123456789 </td>
              <td> Shopping </td>
              <td> 123**** </td>
              <td>28 Jan, 12:36 PM</td>
              <td>
                <p className="text-[#FE5C73]">-25.99$</p>
              </td>
              <td>
                <p className="table-button">Download</p>
              </td>
            </tr>

            <tr>
              {/* Their is goin to be an icon if it is deposit or  */}
              <td className="flex gap-2 items-center">
                <ArrowUpCircleIcon className="transaction-icon" />
                Spotify Subscription{" "}
              </td>
              <td> #123456789 </td>
              <td> Shopping </td>
              <td> 123**** </td>
              <td>28 Jan, 12:36 PM</td>
              <td>
                <p className="text-[#FE5C73]">-25.99$</p>
              </td>
              <td>
                <p className="table-button">Download</p>
              </td>
            </tr>

            <tr>
              {/* Their is goin to be an icon if it is deposit or  */}
              <td className="flex gap-2 items-center">
                <ArrowUpCircleIcon className="transaction-icon" />
                Spotify Subscription{" "}
              </td>
              <td> #123456789 </td>
              <td> Shopping </td>
              <td> 123**** </td>
              <td>28 Jan, 12:36 PM</td>
              <td>
                <p className="text-[#FE5C73]">-25.99$</p>
              </td>
              <td>
                <p className="table-button">Download</p>
              </td>
            </tr>

            <tr>
              {/* Their is goin to be an icon if it is deposit or  */}
              <td className="flex gap-2 items-center">
                <ArrowUpCircleIcon className="transaction-icon" /> Spotify
                Subscription{" "}
              </td>
              <td> #123456789 </td>
              <td> Shopping </td>
              <td> 123**** </td>
              <td>28 Jan, 12:36 PM</td>
              <td>
                <p className="text-[#FE5C73]">-25.99$</p>
              </td>
              <td>
                <p className="table-button">Download</p>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default Recent;
