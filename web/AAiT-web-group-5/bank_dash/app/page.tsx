import BarNPie from "@/components/charts/barnpie";
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
        <div className="flex flex-row justify-evenly">
          <CardS />
          <TransactionCard />
        </div>
        <div>
          <BarNPie />
        </div>
        <div>
          <QuickTransfer />
        </div>
      </div>
    </div>
  );
}
