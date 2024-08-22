import {
    Table,
    TableBody,
    TableCaption,
    TableCell,
    TableFooter,
    TableHead,
    TableHeader,
    TableRow,
  } from "@/components/ui/table";
  
  const invoices = [
    {
      SL_no: "01.",
      LoanMoney: "$100,000",
      LeftToRepay: "$40,000",
      Duration: "6 Months",
      InterestRate: "5%",
      Installment: "$10,000/month",
    }
  ];
  
  export function LoanTable() {
    return (
      <Table className="mb-10">
        <TableHeader>
          <TableRow>
            <TableHead className="w-[100px]">SL No</TableHead>
            <TableHead>Loan Money</TableHead>
            <TableHead>Left To Pay</TableHead>
            <TableHead>Duration</TableHead>
            <TableHead>Interest Rate</TableHead>
            <TableHead>Installment</TableHead>
            <TableHead>Repay</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {invoices.map((invoice) => (
            <TableRow key={invoice.SL_no}>
            <TableCell className="font-medium">{invoice.SL_no}</TableCell>
              <TableCell className="font-medium">{invoice.LoanMoney}</TableCell>
              <TableCell>{invoice.LeftToRepay}</TableCell>
              <TableCell>{invoice.Duration}</TableCell>
              <TableCell>{invoice.InterestRate}</TableCell>
              <TableCell>{invoice.Installment}</TableCell>
              
              <TableCell >
                <button className="rounded-xl border-blue-600 p-2 text-blue-600 border flex items-center justify-center transition ease-in hover:bg-blue-600 hover:text-white">
                  Repay
                </button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
        <TableFooter>
          <TableRow>
            {/* <TableCell colSpan={1}>125,000</TableCell> */}
            <TableCell colSpan={1}>Total</TableCell>
            <TableCell className="text-left text-red-500">$2,500.00</TableCell>
            <TableCell className="text-left text-red-500">$2,500.00</TableCell>
            <TableCell className="text-left"></TableCell>
            <TableCell className="text-left"></TableCell>
            <TableCell className="text-left text-red-500">50000</TableCell>
          </TableRow>
        </TableFooter>
      </Table>
    );
  }
  