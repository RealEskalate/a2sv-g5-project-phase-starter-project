import {Table,TableBody,TableCaption, TableCell, TableFooter, TableHead,TableHeader,TableRow,} from "@/components/ui/table"
import { Invoices } from "./data-tabel"

export function TableDemo() {
    return (
      <Table className="bg-white w-[90%] mx-auto rounded-2xl my-4 p-10 ">
        <TableHeader className="p-10">
          <TableRow className="text-[#243a61] font-[600] p-10  border-b-slate-200">
            <TableHead className="hidden md:table-cell" >SL No</TableHead>
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
            <TableRow key={Invoice.SLNo} className="font-[500] ">
              <TableCell className="hidden md:table-cell">{Invoice.SLNo}</TableCell>
              <TableCell>{Invoice.LoanMoney}</TableCell>
              <TableCell>{Invoice.LeftToPay}</TableCell>
              <TableCell className="hidden md:table-cell">{Invoice.Duration}</TableCell>
              <TableCell className="hidden md:table-cell">{Invoice.InterestRate}</TableCell>
              <TableCell className="hidden md:table-cell">{Invoice.Installment}</TableCell>
              <TableCell className="text-center">
                <button className="border border-1 border-gray-800 rounded-full m-auto hover:text-blue-700 hover:border-blue-700 text-[10px] md:text-[15px] p-2 w-[65px] md:w-[75px]">
                  Repay
                </button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
        <TableFooter className="text-red-500">
          <TableRow className="table-cell md:hidden">
          <TableCell className="hidden md:table-cell">Total</TableCell>
          </TableRow>
          <TableRow >  
            <TableCell className="hidden md:table-cell">Total</TableCell> 
            <TableCell>$2,500.00</TableCell>
            <TableCell colSpan={3} className="hidden md:table-cell">$2,500.00</TableCell>
            <TableCell>$2,500.00</TableCell>
          </TableRow>
        </TableFooter>
      </Table>
    )
}
