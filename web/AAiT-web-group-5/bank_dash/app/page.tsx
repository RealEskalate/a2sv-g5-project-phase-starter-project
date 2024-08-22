import CardS from "@/components/CardS";
import QuickTransfer from "@/components/QuickTransfer";
import RecentTransactionTable from "@/components/RecentTable/RecentTransactionTable";
import Top from "@/components/Top";
import Image from "next/image";

export default function Home() {
  return (
    <div className="overflow-hidden">
      <Top />
      <CardS />
      <RecentTransactionTable />{" "}
      {/* I used this as a place holder remove it when needed */}
      <QuickTransfer />
    </div>
  );
}
