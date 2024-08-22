import { Table } from "@mui/material";
import { TableDemo } from "./TableDemo";

export default function RecentTransactionTable() {
  return (
    <div className="pl-5">
      <p className="font-bold mb-3">Recent Transactions</p>
      <div className="flex gap-20 mb-5">
        <button className="text-violet-700 font-bold">All Transactions</button>
        <div className="font-semibold text-slate-600">Income</div>
        <div className="font-semibold text-slate-600">Expenses</div>
      </div>
      <TableDemo />
    </div>
  );
}
