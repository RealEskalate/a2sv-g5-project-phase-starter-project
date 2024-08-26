"use client";
import Card1 from "./components/Card1";
import { getServerSession } from "next-auth";
import { options } from "../api/auth/[...nextauth]/options";
import React, { useEffect, useState } from "react";
import { getSession } from "next-auth/react";
import { useRouter } from "next/navigation";
import Refresh from "../api/auth/[...nextauth]/token/RefreshToken";
import {
  activeloans,
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
import { SheetDemo } from "./components/Createloan";

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
  const [page, setpage] = useState<number>(0);
  const [total, settotal] = useState<number>(1);
  const [toggle, settoggle] = useState<boolean>(false);
  const numbers = [];

  // Loop from 0 to 5

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
        const data1 = await activeloansall(access_token, 9, page);
        // console.log(data1);
        setf(data1.content);
        settotal(data1.totalPages);

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
  }, [access_token, page, toggle]);

  if (loading || Loading)
    return (
      <div className="bg-gray-100 dark:bg-gray-900 p-6 animate-pulse">
        <div className="flex justify-between flex-wrap lg:flex-nowrap gap-4">
          <div className="bg-gray-300 dark:bg-gray-700 rounded-lg w-1/3 h-20"></div>
          <div className="bg-gray-300 dark:bg-gray-700 rounded-lg w-1/3 h-20"></div>
          <div className="bg-gray-300 dark:bg-gray-700 rounded-lg w-1/3 h-20"></div>
          <div className="bg-gray-300 dark:bg-gray-700 rounded-lg w-1/3 h-20"></div>
        </div>
        <div className=" ">
          <div className="col-span-2 lg:col-span-1">
            <div className="my-4 bg-gray-300 dark:bg-gray-700 rounded-lg h-8 w-1/3"></div>
            <div className="bg-gray-300 dark:bg-gray-700 rounded-lg h-10 my-3"></div>
            <div className="bg-gray-300 dark:bg-gray-700 rounded-lg h-10 my-3"></div>
            <div className="bg-gray-300 dark:bg-gray-700 rounded-lg h-10 my-3"></div>
            <div className="bg-gray-300 dark:bg-gray-700 rounded-lg h-10 my-3"></div>
            <div className="bg-gray-300 dark:bg-gray-700 rounded-lg h-10 my-3"></div>
            <div className="bg-gray-300 dark:bg-gray-700 rounded-lg h-10 my-3"></div>
            <div className="bg-gray-300 dark:bg-gray-700 rounded-lg h-10 my-3"></div>
            <div className="bg-gray-300 dark:bg-gray-700 rounded-lg h-10 my-3"></div>
          </div>
        </div>
      </div>
    );
  for (let i = 1; i <= total; i++) {
    numbers.push(i);
  }
  const handlePage = (page: number) => {
    setpage(page);
  };
  type Form = {
    loanAmount: number;
    duration: number;
    interestRate: number;
    type: string;
  };
  const handleform = async (data: Form) => {
    // console.log('handle',data);
    const res = await activeloans(access_token, data);
    settoggle(!toggle);
  }

  return (
    <div className="bg-gray-100 p-6 dark:bg-[#090b0e]">
      <div className="flex justify-between gap-8 overflow-x-auto [&::-webkit-scrollbar]:hidden ">
        <Card1
          text="Personal Loans"
          img="/personal.svg"
          num={data?.personalLoan}
          // num = {2}
        />
        <Card1
          text="Corporate Loans"
          img="/corporate.svg"
          num={data?.corporateLoan}
          // num={2}
        />
        <Card1
          text="Business Loans"
          img="/business.svg"
          num={data?.businessLoan}
          // num = {2}
        />
        <SheetDemo handleform={handleform} />
      </div>
      <div className="my-4 text-2xl font-bold text-[#333B69] dark:text-[#9faaeb]">
        Active Loans Overview
      </div>
      <Table className="bg-white shadow-1 rounded-3xl  dark:border dark:border-[#333B69] ">
        <TableHeader className="dark:bg-[#020817]">
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
        <TableBody className="dark:bg-[#050914] ">
          {f?.map((invoice: invoices, idx: number) => (
            <TableRow key={invoice.serialNumber}>
              <TableCell
                className={
                  invoice.serialNumber !== "Total"
                    ? "font-medium text-[#232323] hidden md:table-cell dark:text-[#9faaeb]"
                    : "font-medium text-[#FE5C73] hidden md:table-cell dark:text-[#9faaeb]"
                }
              >
                {page}
                {idx + 1}
              </TableCell>
              <TableCell
                className={
                  invoice.serialNumber !== "Total"
                    ? "font-medium text-[#232323] dark:text-[#9faaeb]"
                    : "font-medium text-[#FE5C73] dark:text-[#9faaeb]"
                }
              >
                {invoice.loanAmount}
              </TableCell>
              <TableCell
                className={
                  invoice.serialNumber !== "Total"
                    ? "font-medium text-[#232323] dark:text-[#9faaeb]"
                    : "font-medium text-[#FE5C73] dark:text-[#9faaeb]"
                }
              >
                {invoice.amountLeftToRepay}
              </TableCell>
              <TableCell
                className={
                  invoice.serialNumber !== "Total"
                    ? "font-medium text-[#232323] hidden md:table-cell dark:text-[#9faaeb]"
                    : "font-medium text-[#FE5C73] hidden md:table-cell dark:text-[#9faaeb]"
                }
              >
                {invoice.duration}
              </TableCell>
              <TableCell
                className={
                  invoice.serialNumber !== "Total"
                    ? "font-medium text-[#232323] hidden md:table-cell dark:text-[#9faaeb]"
                    : "font-medium text-[#FE5C73] hidden md:table-cell dark:text-[#9faaeb]"
                }
              >
                {invoice.interestRate}
              </TableCell>
              <TableCell
                className={
                  invoice.serialNumber !== "Total"
                    ? "font-medium text-[#232323] hidden md:table-cell dark:text-[#9faaeb]"
                    : "font-medium text-[#FE5C73] hidden md:table-cell dark:text-[#9faaeb]"
                }
              >
                {invoice.installment}
              </TableCell>
              <TableCell className="text-center ">
                <div
                  className={
                    invoice.serialNumber !== "01"
                      ? "border-2 rounded-full border-[#1814F3] md:border-[#232323] w-full h-full py-1 px-3 dark:text-[#9faaeb]"
                      : "border-2 rounded-full border-[#1814F3] w-full h-full py-1 px-3 dark:text-[#9faaeb]"
                  }
                >
                  <button
                    className={
                      invoice.serialNumber !== "01"
                        ? "text-[#1814F3] md:text-[#232323] font-bold dark:text-[#9faaeb]"
                        : "text-[#1814F3] font-semibold dark:text-[#9faaeb]"
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
      <div className="flex justify-center items-center pt-5">
        <div className="flex flex-wrap">
          <button
            onClick={() => {
              if (page > 0) {
                handlePage(page - 1);
              }
            }}
            className={
              page > 0
                ? "mx-3 text-slate-600 hover:text-slate-700 dark:text-[#9faaeb] font-bold text-xl"
                : "cursor-context-menu mx-3 text-slate-400  dark:text-[#9fabeb6b] font-bold text-xl"
            }
          >
            {"< "}Previous
          </button>
          {numbers.map((number) => (
            <button
              key=""
              onClick={() => handlePage(number - 1)}
              className={
                (page >= number && page - number < 1) ||
                (number >= page && number - page <= 2) ||
                (number === 3 && page < 2) ||
                (number === page - 1 && page === total - 1)
                  ? page != number - 1
                    ? "m-2 px-2 bg-blue-500 hover:bg-blue-700 rounded"
                    : "m-2 px-2 bg-red-500 hover:bg-red-700"
                  : "hidden"
              }
            >
              {" "}
              {number}
            </button>
          ))}
          <button
            onClick={() => handlePage(page + 1)}
            className={
              page < total - 1
                ? "mx-3 text-slate-400 hover:text-slate-700 dark:text-[#9faaeb] font-bold text-xl"
                : "cursor-context-menu mx-3 text-slate-400  dark:text-[#9fabeb6b] font-bold text-xl"
            }
          >
            Next{" >"}
          </button>
        </div>
      </div>
    </div>
  );
}
