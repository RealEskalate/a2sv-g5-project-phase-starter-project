import CardS from "@/components/CreditCards/CardS";
import QuickTransfer from "@/components/QuickTransfer";
import RecentTransactionTable from "@/components/RecentTable/RecentTransactionTable";
import TransactionCard from "@/components/recentTranCard/recentTransactionCard";
import Top from "@/components/Top";
import Image from "next/image";

export default function Home() {
  return (
    <div className="overflow-hidden ">
      <Top topicName="Overview"/>
      <div className="flex flex-col bg-slate-100">
        <div className="flex flex-row justify-evenly"> 
          <CardS />
          <TransactionCard />
        </div>
        <QuickTransfer />
      </div>
    </div>
  );
}
