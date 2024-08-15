import DepateAndCredit from "@/components/Charts/DepateAndCredit";
import MyExpence from "@/components/Charts/MyExpence";
import dynamic from "next/dynamic";
const BalanceHistory = dynamic(
  () => import("@/components/Charts/BalanceHistory"),
  { ssr: false }
);

export default function Home() {
  return (
    <div className="md:flex md:space-x-6">
      <DepateAndCredit />
    </div>
  );
}
