import {Table,TableBody,TableCaption, TableCell, TableFooter, TableHead,TableHeader,TableRow,} from "@/components/ui/table"
import { Invoices } from "./data-tabel"

export function TableDemo() {
    return (
      <div className="-z-50 p-3 bg-slate-600">
      <Table className="bg-slate-50 w-[90%] mx-auto rounded-lg mt-4 p-10 ">
        <TableHeader className="p-10">
          <TableRow className="text-[#243a61] font-[600] p-10">
            <TableHead className="hidden md:table-cell" >SL No</TableHead>
            <TableHead>Loan Money</TableHead>
            <TableHead>Left to repay</TableHead>
            <TableHead className="hidden md:table-cell">Duration</TableHead>
            <TableHead className="hidden md:table-cell">Interest rate</TableHead>
            <TableHead className="hidden md:table-cell">Installment</TableHead>
            <TableHead>Repay</TableHead>

          </TableRow>
        </TableHeader>
        <TableBody>
          {Invoices.map((Invoice) => (
            <TableRow key={Invoice.SLNo} className="font-[500] ">
              <TableCell className="hidden md:table-cell">{Invoice.SLNo}</TableCell>
              <TableCell>{Invoice.LoanMoney}</TableCell>
              <TableCell>{Invoice.LeftToPay}</TableCell>
              <TableCell className="hidden md:table-cell">{Invoice.Duration}</TableCell>
              <TableCell className="hidden md:table-cell">{Invoice.InterestRate}</TableCell>
              <TableCell className="hidden md:table-cell">{Invoice.Installment}</TableCell>
              <TableCell>
                <button className="border border-1 border-gray-800 rounded-full m-auto w-2/3 hover:text-blue-700 hover:border-blue-700">
                  Repay
                </button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
        <TableFooter>
          <TableRow>
            <TableCell colSpan={3}>Total</TableCell>
            <TableCell className="text-right">$2,500.00</TableCell>
          </TableRow>
        </TableFooter>
      </Table>
      </div>
    )
}
