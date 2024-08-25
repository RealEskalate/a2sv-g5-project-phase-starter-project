import Balancequick from "@/components/Balancequick";
import BalanceChart from "@/components/charts/Balance";
import BarNPie from "@/components/charts/barnpie";
import GradientStackedAreaChart from "@/components/charts/line";
import CardS from "@/components/CreditCards/CardS";
import QuickTransfer from "@/components/QuickTransfer";
import RecentTransactionTable from "@/components/RecentTable/RecentTransactionTable";
import TransactionCard from "@/components/recentTranCard/recentTransactionCard";
import Top from "@/components/Top";
import Image from "next/image";

export default function Home() {
  return (
    <div className="overflow-hidden min-h-[750px]">
      <Top topicName="Overview" />
      <div className="flex flex-col bg-slate-100 h-full px-4 py-3">
        <div className="flex flex-row justify-between">
          <CardS />
          <TransactionCard />
        </div>
        <div>
          <BarNPie />
        </div>
        <div className="flex gap-3">
          <div className="w-2/5">
            <div className="flex flex-col gap-3 p-4 rounded-lg w-full h-full">
              <h2 className="text-lg font-semibold text-gray-700">
                Quick Transfer
              </h2>
              <QuickTransfer />
            </div>
          </div>
          <div className="w-3/5">
            <GradientStackedAreaChart />
          </div>
        </div>
      </div>
    </div>
  );
}
