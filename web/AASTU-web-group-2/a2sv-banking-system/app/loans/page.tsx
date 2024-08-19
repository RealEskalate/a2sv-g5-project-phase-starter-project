import Card1 from "./components/Card1";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableFooter,
  TableRow,
} from "@/app/loans/components/table";
const invoices = [
  {
    slno: "01",
    loanmoney: "$100,000",
    Lefttorepay: "$40,500",
    duration: "8 Months",
    interestrate: "12%",
    installment: "$2,000 / month",
    repay: "Repay",
  },
  {
    slno: "02",
    loanmoney: "$500,000",
    Lefttorepay: "$250,000",
    duration: "36 Months",
    interestrate: "10%",
    installment: "$8,000 / month",
    repay: "Repay",
  },
  {
    slno: "03",
    loanmoney: "$900,000",
    Lefttorepay: "$40,500",
    duration: "12 Months",
    interestrate: "12%",
    installment: "$5,000 / month%",
    repay: "Repay",
  },
  {
    slno: "04",
    loanmoney: "$50,000",
    Lefttorepay: "$40,500",
    duration: "25 Months",
    interestrate: "5%",
    installment: "$2,000 / month",
    repay: "Repay",
  },
  {
    slno: "05",
    loanmoney: "$50,000",
    Lefttorepay: "$40,500",
    duration: "5 Months",
    interestrate: "16%",
    installment: "$10,000 / month",
    repay: "Repay",
  },
  {
    slno: "06",
    loanmoney: "$80,000",
    Lefttorepay: "$25,500",
    duration: "14 Months",
    interestrate: "8%",
    installment: "$2,000 / month",
    repay: "Repay",
  },
  {
    slno: "07",
    loanmoney: "$12,000",
    Lefttorepay: "$5,500",
    duration: "9 Months",
    interestrate: "13%",
    installment: "$500 / month",
    repay: "Repay",
  },
  {
    slno: "08",
    loanmoney: "$160,000",
    Lefttorepay: "$100,800",
    duration: "3 Months",
    interestrate: "12%",
    installment: "$900 / month",
    repay: "Repay",
  },
  {
    slno: "Total",
    loanmoney: "$125,0000",
    Lefttorepay: "$750,000",
    duration: "",
    interestrate: "",
    installment: "$50,000 / month",
    repay: "",
  },
];

export default function Home() {
  return (
    <div className="bg-gray-100 p-6">
      <div className="flex justify-between gap-8 overflow-x-auto [&::-webkit-scrollbar]:hidden">
        <Card1 text="Personal Loans" img="/personal.png" num="$50,000" />
        <Card1 text="Corporate Loans" img="/corporate.png" num="$100,000" />
        <Card1 text="Business Loans" img="/business.png" num="$500,000" />
        <Card1 text="Custom Loans" img="/custom.png" num="Choose Money" />
      </div>
      <div className="my-4 text-2xl font-bold text-[#333B69]">
        Active Loans Overview
      </div>
      <Table className="bg-white shadow-1 rounded-3xl">
        <TableHeader>
          <TableRow className="text-[#718EBF]">
            <TableHead className="w-[100px] text-[#718EBF] hidden md:table-cell">
              SL No
            </TableHead>
            <TableHead className="text-[#718EBF]">Loan Money</TableHead>
            <TableHead className="text-[#718EBF]">Left to repay</TableHead>
            <TableHead className="text-[#718EBF] hidden md:table-cell">
              Duration
            </TableHead>
            <TableHead className="text-[#718EBF] hidden md:table-cell">
              Interest rate
            </TableHead>
            <TableHead className="text-[#718EBF] hidden md:table-cell">
              Installment
            </TableHead>
            <TableHead className=" text-[#718EBF]">Repay</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {invoices.map((invoice) => (
            <TableRow key={invoice.slno}>
              <TableCell
                className={
                  invoice.slno !== "Total"
                    ? "font-medium text-[#232323] hidden md:table-cell"
                    : "font-medium text-[#FE5C73] hidden md:table-cell"
                }
              >
                {invoice.slno}
              </TableCell>
              <TableCell
                className={
                  invoice.slno !== "Total"
                    ? "font-medium text-[#232323]"
                    : "font-medium text-[#FE5C73]"
                }
              >
                {invoice.loanmoney}
              </TableCell>
              <TableCell
                className={
                  invoice.slno !== "Total"
                    ? "font-medium text-[#232323]"
                    : "font-medium text-[#FE5C73]"
                }
              >
                {invoice.Lefttorepay}
              </TableCell>
              <TableCell
                className={
                  invoice.slno !== "Total"
                    ? "font-medium text-[#232323] hidden md:table-cell"
                    : "font-medium text-[#FE5C73] hidden md:table-cell"
                }
              >
                {invoice.duration}
              </TableCell>
              <TableCell
                className={
                  invoice.slno !== "Total"
                    ? "font-medium text-[#232323] hidden md:table-cell"
                    : "font-medium text-[#FE5C73] hidden md:table-cell"
                }
              >
                {invoice.interestrate}
              </TableCell>
              <TableCell
                className={
                  invoice.slno !== "Total"
                    ? "font-medium text-[#232323] hidden md:table-cell"
                    : "font-medium text-[#FE5C73] hidden md:table-cell"
                }
              >
                {invoice.installment}
              </TableCell>
              {invoice.repay !== "" && (
                <TableCell className="text-center ">
                  <div
                    className={
                      invoice.slno !== "01"
                        ? "border-2 rounded-full border-[#1814F3] md:border-[#232323] w-full h-full py-1 px-3"
                        : "border-2 rounded-full border-[#1814F3] w-full h-full py-1 px-3"
                    }
                  >
                    <button
                      className={
                        invoice.slno !== "01"
                          ? "text-[#1814F3] md:text-[#232323] font-bold"
                          : "text-[#1814F3] font-semibold"
                      }
                    >
                      {invoice.repay}
                    </button>{" "}
                  </div>
                </TableCell>
              )}
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
}
