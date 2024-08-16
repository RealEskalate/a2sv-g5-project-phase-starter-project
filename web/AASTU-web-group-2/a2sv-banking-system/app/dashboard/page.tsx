import Image from "next/image";
import ImageComponent from "./components/ImageComponent";
import Reviving from "./components/QuickTransfer"
import { BalanceHistory } from "./components/BalanceHistory";
import {WeeklyActivity} from "./components/WeeklyActivity";
import { ExpenseStatistics } from "./components/ExpenseStatistics"
import RecentTransaction from "./components/RecentTransaction";
// import {RecentTransaction} from "@/components/RecentTransaction"
export default function Home() {
  return (
    <div className=" flex flex-col ">
      <RecentTransaction/>
      <WeeklyActivity />
      <ExpenseStatistics/>

      <Reviving />
      {/* <p className="my-4 mx-4">fast spook</p> */}
      <BalanceHistory />
      {/* <BalanceHistory/> */}
    </div>
  );
}
