import {
  Table,
  TableBody,
  TableCell,
  TableFooter,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Invoices } from "./data-tabel";
import { useUser } from "@/contexts/UserContext";

export function TableDemo() {
  const { isDarkMode } = useUser();

  return (
    <Table
      className={`w-[90%] mx-auto rounded-2xl my-4 p-10 ${
        isDarkMode ? "bg-gray-800 text-gray-200" : "bg-white text-black"
      }`}
    >
      <TableHeader className="p-10">
        <TableRow
          className={`font-[600] p-10 border-b ${
            isDarkMode
              ? "border-b-gray-700 text-gray-400"
              : "border-b-slate-200 text-[#243a61]"
          }`}
        >
          <TableHead className="hidden md:table-cell">SL No</TableHead>
          <TableHead>Loan Money</TableHead>
          <TableHead>Left to repay</TableHead>
          <TableHead className="hidden md:table-cell">Duration</TableHead>
          <TableHead className="hidden md:table-cell">Interest rate</TableHead>
          <TableHead className="hidden md:table-cell">Installment</TableHead>
          <TableHead className="text-center">Repay</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {Invoices.map((Invoice) => (
          <TableRow
            key={Invoice.SLNo}
            className={`font-[500] ${
              isDarkMode ? "text-gray-300" : "text-black"
            }`}
          >
            <TableCell className="hidden md:table-cell">
              {Invoice.SLNo}
            </TableCell>
            <TableCell>{Invoice.LoanMoney}</TableCell>
            <TableCell>{Invoice.LeftToPay}</TableCell>
            <TableCell className="hidden md:table-cell">
              {Invoice.Duration}
            </TableCell>
            <TableCell className="hidden md:table-cell">
              {Invoice.InterestRate}
            </TableCell>
            <TableCell className="hidden md:table-cell">
              {Invoice.Installment}
            </TableCell>
            <TableCell className="text-center">
              <button
                className={`border border-1 rounded-full m-auto text-[10px] md:text-[15px] p-2 w-[65px] md:w-[75px] ${
                  isDarkMode
                    ? "border-gray-400 text-gray-400 hover:text-blue-400 hover:border-blue-400"
                    : "border-gray-800 text-black hover:text-blue-700 hover:border-blue-700"
                }`}
              >
                Repay
              </button>
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
      <TableFooter className={isDarkMode ? "text-gray-500" : "text-red-500"}>
        <TableRow className="table-cell md:hidden">
          <TableCell className="hidden md:table-cell">Total</TableCell>
        </TableRow>
        <TableRow>
          <TableCell className="hidden md:table-cell">Total</TableCell>
          <TableCell>$2,500.00</TableCell>
          <TableCell colSpan={3} className="hidden md:table-cell">
            $2,500.00
          </TableCell>
          <TableCell>$2,500.00</TableCell>
        </TableRow>
      </TableFooter>
    </Table>
  );
}
