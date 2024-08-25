import BalanceChart from "./charts/Balance";
import QuickTransfer from "./QuickTransfer";

export default function Balancequick() {
  return (
    <div className="flex gap-5 mt-3 w-full pl-4 pr-3">
      <div>
        <p className="font-bold text-xl mb-3">Quick Transfer</p>
        <QuickTransfer />
      </div>
      <div className="w-full">
        <p className="font-bold text-xl mb-3">Balance History</p>
        <BalanceChart />
      </div>
    </div>
  );
}
