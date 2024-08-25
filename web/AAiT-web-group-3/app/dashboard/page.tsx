import CreditCard from "../components/credit-card";
import { PiFoldersBold } from "react-icons/pi";
import WeeklyActivityChart from "@/components/dashboard/weeklyActChart";
import ExpenseStatChart from "@/components/dashboard/expenseStatChart";
import { FiSend } from "react-icons/fi";
import BalanceHistoryChart from "@/components/dashboard/balanceHistoryChart";
import QuickTransferCard from "@/components/dashboard/quickTransferCard";

const page = () => {
  return (
    <div className="p-4 pr-0 w-full bg-primary-color-50">
      <div className="flex flex-col gap-y-3 w-full">
        <div className="grid grid-cols-1 lg:grid-cols-12 gap-4 w-full">
          <div className="flex flex-col lg:col-span-8">
            <div className="flex items-center justify-between mb-2 mr-4">
              <h2 className="text-primary-color-800 font-semibold md:text-lg">
                My Cards
              </h2>
              <h3 className="text-primary-color-800 font-semibold">See All</h3>
            </div>

            <div className="flex gap-4 lg:gap-x-10 items-center overflow-x-scroll no-scrollbar w-full">
              <CreditCard color="blue" />
              <CreditCard color="white" />
            </div>
          </div>

          <div className="flex flex-col lg:col-span-4 mr-4 lg:ml-5 lg:mr-2">
            <h2 className="text-primary-color-800 font-semibold md:text-lg mb-3">
              Recent Transaction
            </h2>
            <div className="bg-white rounded-xl flex flex-col gap-y-6 min-h-[170px] min-w[230px] p-3">
              <div className="flex items-center justify-between">
                <span className="bg-[#FFF5D9] w-10 h-10 flex items-center justify-center rounded-full">
                  <PiFoldersBold color="#FFBB38" fontSize={18} />
                </span>
                <div>
                  <h4 className="text-[13px]">Deposit from my</h4>
                  <p className="text-primary-color-200 text-[12px]">
                    28 January 2021
                  </p>
                </div>
                <span className="text-[11px] text-[#41D4A8] ml-2">-$2,569</span>
              </div>

              <div className="flex items-center justify-between w-full">
                <span className="bg-[#FFF5D9] w-10 h-10 flex items-center justify-center rounded-full">
                  <PiFoldersBold color="#FFBB38" fontSize={18} />
                </span>
                <div>
                  <h4 className="text-[13px]">Deposit from my</h4>
                  <p className="text-primary-color-200 text-[12px]">
                    28 January 2021
                  </p>
                </div>
                <span className="text-[11px] text-[#41D4A8] ml-2">-$5,369</span>
              </div>
              <div className="flex items-center justify-between w-full">
                <span className="bg-[#FFF5D9] w-10 h-10 flex items-center justify-center rounded-full">
                  <PiFoldersBold color="#FFBB38" fontSize={18} />
                </span>
                <div>
                  <h4 className="text-[13px]">Deposit from my</h4>
                  <p className="text-primary-color-200 text-[12px]">
                    28 January 2021
                  </p>
                </div>
                <span className="text-[11px] text-red-600 ml-2">-$569</span>
              </div>
            </div>
          </div>
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-12 gap-4 w-full pr-4">
          <div className="flex flex-col lg:col-span-8 h-[260px]">
            <h2 className="text-primary-color-800 font-semibold md:text-lg mb-2">
              Weekly Activity
            </h2>

            <WeeklyActivityChart />
          </div>

          <div className="flex flex-col lg:col-span-4  h-[260px]">
            <h2 className="text-primary-color-800 font-semibold md:text-lg mb-2">
              Expense Statistics
            </h2>
            <ExpenseStatChart />
          </div>
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-12 gap-4 pr-4">
          <div className="flex flex-col lg:col-span-4">
            <h2 className="text-primary-color-800 font-semibold md:text-lg mb-2">
              Quick Transfer
            </h2>
            <div className="bg-white rounded-xl flex flex-col items-center gap-y-6 min-h-[170px] min-w[230px] h-[220px] p-5">
              <QuickTransferCard />

              <div className="flex items-center justify-between h-full w-full">
                <p className="text-xs text-primary-color-200 font-medium">
                  Write Amount
                </p>
                <span className="text-xs bg-primary-color-50 rounded-full w-fit pl-2">
                  $25.00
                  <span className="bg-[#1814F3] rounded-full p-2 px-3 w-fit inline-flex items-center gap-2 ml-3 text-white">
                    Send
                    <FiSend color="#fff" fontSize={12} />
                  </span>
                </span>
              </div>
            </div>
          </div>

          <div className="flex flex-col lg:col-span-8">
            <h2 className="text-primary-color-800 font-semibold md:text-lg mb-2">
              Balance History
            </h2>
            <BalanceHistoryChart />
          </div>
        </div>
      </div>
    </div>
  );
};

export default page;
