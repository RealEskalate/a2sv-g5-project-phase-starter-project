"use client";
import Card1 from "./components/Card1";
import { getServerSession } from "next-auth";
import { options } from "../api/auth/[...nextauth]/options";
import React, { useEffect, useState } from "react";
import { getSession } from "next-auth/react";
import { useRouter } from "next/navigation";
import Refresh from "../api/auth/[...nextauth]/token/RefreshToken";
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

// const loanid = "66c3054e80b7cf4a6c2f7709";
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
  personalLoan: any;
  businessLoan: any;
  corporateLoan: any;
}
type Data = {
  access_token: string;
  data: string;
  refresh_token: string;
};

type SessionDataType = {
  user: Data;
};
export default function Home() {
    const [session, setSession] = useState<Data | null>(null);
    const [access_token, setAccess_token] = useState("");
    const router = useRouter();
    const [loading, setloading] = useState(true);
    const [Loading, setLoading] = useState(true);
    const [f, setf] = useState<invoices[]>();
    const [data, setdata] = useState<loantype>();

    // Getting the session from the server and Access Token From Refresh
    useEffect(() => {
      const fetchSession = async () => {
        try {
          const sessionData = (await getSession()) as SessionDataType | null;
          setAccess_token(await Refresh());
          if (sessionData && sessionData.user) {
            setSession(sessionData.user);
          } else {
            router.push(
              `./api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`
            );
          }
        } catch (error) {
          console.error("Error fetching session:", error);
        } finally {
          setloading(false);
        }
      };

      fetchSession();
    }, [router]);

    // Combined fetching data to reduce multiple useEffect hooks
    useEffect(() => {
      const fetchData = async () => {
        if (!access_token) return;

        try {
          // Fetch f
          const data1 = await activeloansall(access_token);
          setf(data1);

          // Fetch data
          const d: loantype = await activeloansdetaildata(access_token);
          setdata(d);
        } catch (error) {
          console.error("Error fetching data:", error);
        } finally {
          setLoading(false);
        }
      };

      fetchData();
    }, [access_token]);

    if (loading||Loading) return <div>Loading...</div>;

  return (
    <div className="bg-gray-100 p-6 ">
      <div className="flex justify-between gap-8 overflow-x-auto [&::-webkit-scrollbar]:hidden">
        <Card1
          text="Personal Loans"
          img="/personal.png"
          num={data?.personalLoan}
        />
        <Card1
          text="Corporate Loans"
          img="/corporate.png"
          num={data?.corporateLoan}
        />
        <Card1
          text="Business Loans"
          img="/business.png"
          num={data?.businessLoan}
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
          {f?.map((invoice: invoices) => (
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
  );
}
