import Card1 from "./components/Card1";
import { getServerSession } from "next-auth";
import { options } from "../api/auth/[...nextauth]/options";
import {
  activeloansall,
  activeloansdetaildata,
} from "./back/ActiveLoanController";
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

const loanid = "66c3054e80b7cf4a6c2f7709";
// const token ="eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJsc2FqZGxzanNuIiwiaWF0IjoxNzI0MTU1NzkzLCJleHAiOjE3MjQyNDIxOTN9.wi7oRgF81zMp1v8tPzRPmAj4GOLaYy4bV_TMVvtWmzg2mjrTThiruT_Fswcyu1eq";

interface invoices {
  serialNumber: string;
  loanAmount: number;
  amountLeftToRepay: number;
  duration: number;
  interestRate: number;
  installment: number;
  type: string;
  activeLoneStatus: string;
  userId: string;
}
interface loantype {
  personalLoan: number;
  businessLoan: number;
  corporateLoan: number;
}
type ISODateString = string;

interface DefaultSession {
  user?: {
    name?: string | null;
    email?: string | null;
    image?: string | null;
    access_token?: string | any;
  };
  expires: ISODateString;
}
export default async function Home() {
  const seasion: DefaultSession | null = await getServerSession(options);
  const f = await activeloansall(seasion?.user?.access_token);
  // const f = await activeloansall(token);
  // console.log("asa1111", f);
  const data: loantype = await activeloansdetaildata(
    seasion?.user?.access_token
  );
  // console.log("s", data);

  return (
    // <main className="mt-16 ml-72">
    <div className="bg-gray-100 p-6 ">
      <div className="flex justify-between gap-8 overflow-x-auto [&::-webkit-scrollbar]:hidden">
        <Card1
          text="Personal Loans"
          img="/personal.png"
          num={data.personalLoan}
        />
        <Card1
          text="Corporate Loans"
          img="/corporate.png"
          num={data.corporateLoan}
        />
        <Card1
          text="Business Loans"
          img="/business.png"
          num={data.businessLoan}
        />
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
          {f.map((invoice: invoices) => (
            <TableRow key={invoice.serialNumber}>
              <TableCell
                className={
                  invoice.serialNumber !== "Total"
                    ? "font-medium text-[#232323] hidden md:table-cell"
                    : "font-medium text-[#FE5C73] hidden md:table-cell"
                }
              >
                {invoice.serialNumber}
              </TableCell>
              <TableCell
                className={
                  invoice.serialNumber !== "Total"
                    ? "font-medium text-[#232323]"
                    : "font-medium text-[#FE5C73]"
                }
              >
                {invoice.loanAmount}
              </TableCell>
              <TableCell
                className={
                  invoice.serialNumber !== "Total"
                    ? "font-medium text-[#232323]"
                    : "font-medium text-[#FE5C73]"
                }
              >
                {invoice.amountLeftToRepay}
              </TableCell>
              <TableCell
                className={
                  invoice.serialNumber !== "Total"
                    ? "font-medium text-[#232323] hidden md:table-cell"
                    : "font-medium text-[#FE5C73] hidden md:table-cell"
                }
              >
                {invoice.duration}
              </TableCell>
              <TableCell
                className={
                  invoice.serialNumber !== "Total"
                    ? "font-medium text-[#232323] hidden md:table-cell"
                    : "font-medium text-[#FE5C73] hidden md:table-cell"
                }
              >
                {invoice.interestRate}
              </TableCell>
              <TableCell
                className={
                  invoice.serialNumber !== "Total"
                    ? "font-medium text-[#232323] hidden md:table-cell"
                    : "font-medium text-[#FE5C73] hidden md:table-cell"
                }
              >
                {invoice.installment}
              </TableCell>
              <TableCell className="text-center ">
                <div
                  className={
                    invoice.serialNumber !== "01"
                      ? "border-2 rounded-full border-[#1814F3] md:border-[#232323] w-full h-full py-1 px-3"
                      : "border-2 rounded-full border-[#1814F3] w-full h-full py-1 px-3"
                  }
                >
                  <button
                    className={
                      invoice.serialNumber !== "01"
                        ? "text-[#1814F3] md:text-[#232323] font-bold"
                        : "text-[#1814F3] font-semibold"
                    }
                  >
                    repay
                  </button>{" "}
                </div>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
    // </main>
  );
}
