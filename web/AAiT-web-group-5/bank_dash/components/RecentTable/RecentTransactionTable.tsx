import { Table } from "@mui/material";
import { TableDemo } from "./TableDemo";

export default function RecentTransactionTable() {
  return (
    <div className="pl-5">
      <p className="font-bold mb-3">Recent Transactions</p>
      <div className="flex gap-20 mb-5">
        <p className="text-violet-700 font-bold">All Transactions</p>
        <p className="font-thin">Income</p>
        <p className="font-thin">Expenses</p>
      </div>
      <TableDemo />
    </div>
  );
}
