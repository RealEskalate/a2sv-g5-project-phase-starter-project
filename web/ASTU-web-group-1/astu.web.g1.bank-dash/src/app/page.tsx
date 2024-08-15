import CardExpenceStatistics from "@/components/Charts/CardExpenceStatistics";
import DebiteAndCredit from "@/components/Charts/DebiteAndCredit";
import DepateAndCredit from "@/components/Charts/DebiteAndCredit";
import ExpenseStatistics from "@/components/Charts/ExpenseStatistics";
import MonthlyRevenue from "@/components/Charts/MonthlyRevenue";
import MyExpence from "@/components/Charts/MyExpence";
import WeeklyActivity from "@/components/Charts/WeeklyActivity";
import YearlyTotalInvestment from "@/components/Charts/YearlyTotalInvestment";
import dynamic from "next/dynamic";
const BalanceHistory = dynamic(
  () => import("@/components/Charts/BalanceHistory"),
  { ssr: false }
);

export default function Home() {
  return (
    <div className="md:flex md:space-x-6">
      <DebiteAndCredit />
    </div>
  );
}
