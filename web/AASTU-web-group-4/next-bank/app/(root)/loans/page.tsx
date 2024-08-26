"use client";

import React, { useEffect, useState } from "react";
import Cookies from "js-cookie";

import {
  Carousel,
  CarouselItem,
  CarouselContent,
} from "@/components/ui/carousel";
import { loanCardMapping } from "@/constants/index";
import Link from "next/link";
import { getMyLoans, getLoanDetailData } from "@/services/activeloan";
import CustomLoans from "@/public/icons/CustomLoans";
import { TbFileSad } from "react-icons/tb";
import { colors } from "@/constants";
import { UserData } from "@/types";
import { currentuser } from "@/services/userupdate";
import { createTransaction } from "@/services/transactionfetch";
import { access } from "fs";
import { message } from "antd";
import { ArrowPathIcon } from "@heroicons/react/24/outline";
import Pagination from "@/components/Pagination";
interface LoanCard {
  icon: (props: any) => JSX.Element;
  title: string;
  loanAmount: string;
}

const LoanCard: React.FC<{
  icon: React.FC<React.SVGProps<SVGSVGElement>>;
  title: string;
  description: string;
}> = ({ icon: Icon, title, description }) => (
  <div className="flex items-center space-x-4">
    <Icon className="pl-1 w-12 h-12" aria-hidden="true" />
    <div>
      <h3 className="text-sm text-gray-500">{title}</h3>
      <p className="text-lg font-extrabold">{description}</p>
    </div>
  </div>
);

const LoansPage: React.FC = () => {
  // State for loan cards
  const [loanCards, setLoanCards] = useState<LoanCard[]>([]);
  const [loanCardsLoading, setLoanCardsLoading] = useState(true);
  const [loanCardsError, setLoanCardsError] = useState<string | null>(null);

  // State for loans
  const [loans, setLoans] = useState([]);
  const [loansLoading, setLoansLoading] = useState(true);
  const [loansError, setLoansError] = useState<string | null>(null);
  const ITEMS_PER_PAGE = 5;
  const [totalPages, setTotalPages] = useState(0);
  const [currentPage, setCurrentPage] = useState(0);
  useEffect(() => {
    const fetchLoans = async () => {
      setLoansLoading(true);
      setLoansError(null);

      try {
        const response = await getMyLoans(currentPage, ITEMS_PER_PAGE);
        if (response.data.content.length === 0) {
          setLoansError("No active loans found.");
        } else {
          setLoans(response.data.content);
          setTotalPages(response.data.totalPages);
        }
      } catch (error) {
        setLoansError("Error fetching the loans.");
      } finally {
        setLoansLoading(false);
      }
    };
    fetchLoans();
  }, [currentPage]);
  const [messageApi, contextHolder] = message.useMessage();
  const success = (amount: string, username: string) => {
    messageApi.open({
      type: "success",
      content: `Successfully transferred ${amount} to ${username}`,
      duration: 4,
    });
  };
  const errormessage = () => {
    messageApi.open({
      type: "error",
      content: "Transaction was not successful",
      duration: 4,
    } as any);
  };

  const lowbalance = () => {
    messageApi.open({
      type: "error",
      content: "Insufficient funds",
      duration: 4,
    } as any);
  };

  const accessToken = Cookies.get("accessToken") || "";
  const selectedUser = "soll";
  const [isLoading, setIsLoading] = useState(false);
  const[selectedIndex , setselecteIndex] = useState(-1)
  const onSubmit = async (amount: string ,index:any) => {
    setselecteIndex(index)
    setIsLoading(true);
    const transactionData = {
      type: "transfer",
      description: `Transfer to [${selectedUser}]`,
      amount: amount,
      receiverUserName: selectedUser,
    };
    console.log("transactionData:", transactionData);

    try {
      const res = await createTransaction(transactionData, accessToken);
      if (res.success && parseInt(transactionData.amount) < accountBalance) {
        success(transactionData.amount, transactionData.receiverUserName);
        setLoans(loans.filter(( _ , i) => i !== index));
      } else if (parseInt(transactionData.amount) > accountBalance) {
        lowbalance();
        console.error(
          "Insufficient funds , typeof(accountBalance):",
          accountBalance,
          transactionData.amount
        );
      } else {
        errormessage();
        console.error("Failed to create transaction", res);
      }
    } catch (error) {
      errormessage();
      console.error("Error creating transaction:", error);
    }
    setIsLoading(false);
    setselecteIndex(-1)
  };
  
  useEffect(() => {
    const fetchLoanCard = async () => {
      setLoanCardsLoading(true);
      setLoanCardsError(null);

      try {
        const response = await getLoanDetailData();
        const { data } = response;

        const loanCard = loanCardMapping.map((loan) => ({
          icon: loan.icon,
          title: loan.title,
          loanAmount: `$${data[loan.descriptionKey]?.toLocaleString() || 0}`,
        }));

        loanCard.push({
          icon: CustomLoans,
          title: "Custom Loans",
          loanAmount: "Choose Loans",
        });

        if (loanCard.length === 0) {
          setLoanCardsError("No loan details available.");
        } else {
          setLoanCards(loanCard);
        }
      } catch (error) {
        setLoanCardsError("Error fetching the loans summary data.");
      } finally {
        setLoanCardsLoading(false);
      }
    };

    fetchLoanCard();
  }, []);

  const [accountBalance, setAccountBalance] = useState(0);
  const [info, setinfo] = useState<UserData>();
  const [visible, setvisible] = useState(false);
  useEffect(() => {
    const fetch = async () => {
      try {
        const data = await currentuser();
        setinfo(data.data || []);
        setAccountBalance(data.data.accountBalance);
      } catch (error) {
        console.error("Error:", error);
      }
    };
    fetch();
  }, []);

  const totalLoanMoney = loans.reduce(
    (total, loan: any) => total + parseInt(loan.loanAmount),
    0
  );
  const totalLeftToRepay = loans.reduce(
    (total, loan: any) => total + parseInt(loan.amountLeftToRepay),
    0
  );
  const totalInstallment = loans.reduce(
    (total, loan: any) => total + parseInt(loan.installment),
    0
  );

  return (
    <div className="flex lg:ml-64 lg:max-w-[100%] px-6 dark:text-blue-500 bg-gray-100 dark:bg-dark ">
      {/* Mobile and Tablet View */}

      {contextHolder}
      <div className="w-[100%] block lg:hidden">
        {loanCardsLoading ? (
          <div className="flex w-[100%] py-6  justify-center gap-5">
            <div className=" lg:hidden flex items-center space-x-4 animate-pulse p-4 bg-gray-200 rounded-md w-[240px] h-[85px]">
              <div className="pl-1 w-12 h-12 bg-gray-300 rounded-full" />
              <div>
                <div className="w-20 h-4 bg-gray-300 rounded mb-2" />
                <div className="w-32 h-6 bg-gray-300 rounded" />
              </div>
            </div>

            <div className=" lg:hidden flex items-center space-x-4 animate-pulse p-4 bg-gray-200 rounded-md w-[240px] h-[85px]">
              <div className="pl-1 w-12 h-12 bg-gray-300 rounded-full" />
              <div>
                <div className="w-20 h-4 bg-gray-300 rounded mb-2" />
                <div className="w-32 h-6 bg-gray-300 rounded" />
              </div>
            </div>

            <div className=" lg:hidden flex items-center space-x-4 animate-pulse p-4 bg-gray-200 rounded-md w-[240px] h-[85px]">
              <div className="pl-1 w-12 h-12 bg-gray-300 rounded-full" />
              <div>
                <div className="w-20 h-4 bg-gray-300 rounded mb-2" />
                <div className="w-32 h-6 bg-gray-300 rounded" />
              </div>
            </div>
          </div>
        ) : loanCardsError ? (
          <div className="flex flex-col items-center justify-center text-center space-y-4">
            <TbFileSad
          className={`text-gray-300 dark:text-[#993d4b] w-[400px] h-[70px] pb-2 block mx-auto`}
          strokeWidth={1}
        />
            <div className="text-red-500 ">{loanCardsError}</div>
          </div>
        ) : (
          <Carousel>
            <CarouselContent className="p-6">
              {loanCards.map((loanItem, index) => (
                <CarouselItem
                  key={index}
                  className="w-[240px] h-[85px] mx-auto mr-4 flex-none"
                >
                  <div className="shadow-lg p-4 rounded-md flex items-center h-full">
                    <LoanCard
                      icon={loanItem.icon}
                      title={loanItem.title}
                      description={loanItem.loanAmount}
                    />
                  </div>
                </CarouselItem>
              ))}
            </CarouselContent>
          </Carousel>
        )}

        <div className="w-[100%] flex flex-col items-center mt-8 text-sm">
          <h2 className="text-lg font-bold mb-4 ml-5 dark:text-blue-500">
            Active Loans Overview
          </h2>
          {loansLoading ? (
            <div className="w-[100%] flex justify-center">
              <table className="w-[70%]  h-[85px] bg-white rounded-lg shadow-md text-[12px] dark:bg-dark dark:text-white animate-pulse">
                <thead>
                  <tr>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400 bg-gray-300 dark:bg-gray-700 h-6">
                      Loan Money
                    </th>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400 bg-gray-300 dark:bg-gray-700 h-6">
                      Left to Repay
                    </th>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400 bg-gray-300 dark:bg-gray-700 h-6">
                      {" "}
                      Repay
                    </th>
                  </tr>
                </thead>
                <tbody>
                  {Array.from({ length: 10 }).map((_, index) => (
                    <tr key={index}>
                      <td className="border-t px-4 py-2 bg-gray-300 dark:bg-gray-700 h-6"></td>
                      <td className="border-t px-4 py-2 bg-gray-300 dark:bg-gray-700 h-6"></td>
                      <td className="border-t px-4 py-2 bg-gray-300 dark:bg-gray-700 h-6"></td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          ) : loansError ? (
            <div className="w-[100%] flex justify-center items-center flex-col p-4">
              <TbFileSad
          className={`text-gray-300 dark:text-[#993d4b] w-[400px] h-[70px] pb-2 block mx-auto`}
          strokeWidth={1}
        />
              <div className="text-red-500 ">{loansError}</div>
            </div>
          ) : (
            <div className="w-[100%] flex justify-center">
              <table className="w-[70%]  h-[85px] bg-white rounded-lg shadow-md text-[12px] dark:bg-dark text-gray-900 dark:text-white">
                <thead>
                  <tr>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400">
                      Loan Money
                    </th>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400">
                      Left to Repay
                    </th>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400">
                      Repay
                    </th>
                  </tr>
                </thead>
                <tbody>
                  {loans.map((loan: any, index) => (
                    <tr key={index}>
                      <td className="border-t px-4 py-2">{loan.loanAmount}</td>
                      <td className="border-t px-4 py-2">
                        {loan.amountLeftToRepay}
                      </td>
                      <td className="border-t px-4 py-2">
                      <button
                          onClick={() => onSubmit(loan.amountLeftToRepay , index)}
                          className= {` text-gray-900 border border-purple-900 rounded-full px-4 py-1 ${selectedIndex === index && isLoading ? 'bg-gray-200 cursor-not-allowed ' : 'bg-white hover:bg-gray-200' }`}
                        >
                           { selectedIndex === index && isLoading ? (
                            <div className="flex justify-center items-center ">
                              <ArrowPathIcon className="h-5 w-5 animate-spin  text-gray-500  " />
                            </div>
                          ) : (
                            "Repay"
                          )}
                        </button>
                      </td>
                    </tr>
                  ))}
                  <tr className="font-bold text-red-500">
                    <td className="border-t px-4 py-2">
                      ${totalLoanMoney.toLocaleString()}
                    </td>
                    <td className="border-t px-4 py-2">
                      ${totalLeftToRepay.toLocaleString()}
                    </td>
                    <td className="border-t px-4 py-2"></td>
                  </tr>
                </tbody>
              </table>
            </div>
          )}
           <Pagination
        currentPage={currentPage}
        totalPages={totalPages}
        onPageChange={setCurrentPage}/>
        </div>
      </div>

      {/* Desktop and Tablet View */}
      <div className="hidden lg:block lg:w-[100%] lg:bg-gary-100 lg:dark:bg-dark dark:text-blue-500 ">
        {loanCardsLoading ? (
          <div className="flex justify-evenly py-10   space-x-6">
            {Array.from({ length: 4 }).map((_, index) => (
              <div
                key={index}
                className="flex items-center space-x-4 animate-pulse p-4 bg-gray-200 dark:bg-gray-800 rounded-md w-[240px] h-[85px] shadow-md"
              >
                <div className="pl-1 w-12 h-12 bg-gray-300 dark:bg-gray-300 rounded-full" />
                <div>
                  <div className="w-20 h-4 bg-gray-300 dark:bg-gray-300 rounded mb-2" />
                  <div className="w-32 h-6 bg-gray-300 dark:bg-gray-300 rounded" />
                </div>
              </div>
            ))}
          </div>
        ) : loanCardsError ? (
          <div className="hidden  lg:flex lg:flex-col lg:items-center lg:justify-center lg:text-center lg:space-y-4 lg:h-200px lg:bg-gray-100 lg:dark:bg-gray-800 lg:py-8">
            <TbFileSad
          className={`text-gray-300 dark:text-[#993d4b] w-[400px] h-[70px] pb-2 block mx-auto`}
          strokeWidth={1}
        />
            <div className="text-red-500 ">
              {loanCardsError}
            </div>
          </div>
        ) : (
          <div className="w-[100%] flex gap-6 pr-6 py-8">
            {loanCards.map((loanItem, index) => (
              <div
                key={index}
                className="w-[100%] h-[120px] shadow-lg rounded-lg flex items-center"
              >
                <div className="w-[100%]">
                  <LoanCard
                    icon={loanItem.icon}
                    title={loanItem.title}
                    description={loanItem.loanAmount}
                  />
                </div>
              </div>
            ))}
          </div>
        )}

        <div className="mt-8">
          <h2 className="text-lg font-bold mb-4 dark:text-blue-500">
            Active Loans Overview
          </h2>
          {loansLoading ? (
            <div className="overflow-x-auto px-10">
              <table className="w-[100%] bg-white rounded-2xl shadow-md table-fixed dark:bg-dark text-gray-900 dark:text-white animate-pulse">
                <thead>
                  <tr>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400 bg-gray-300 dark:bg-gray-700 h-6">
                      SL No
                    </th>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400 bg-gray-300 dark:bg-gray-700 h-6">
                      Loan Money
                    </th>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400 bg-gray-300 dark:bg-gray-700 h-6">
                      Left to Repay
                    </th>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400 bg-gray-300 dark:bg-gray-700 h-6">
                      Duration
                    </th>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400 bg-gray-300 dark:bg-gray-700 h-6">
                      Interest Rate
                    </th>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400 bg-gray-300 dark:bg-gray-700 h-6">
                      Monthly Installment
                    </th>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400 bg-gray-300 dark:bg-gray-700 h-6">
                      Repay
                    </th>
                  </tr>
                </thead>
                <tbody>
                  {Array.from({ length: 10 }).map((_, index) => (
                    <tr key={index}>
                      <td className="border-t px-4 py-2 bg-gray-300 dark:bg-gray-700 h-6"></td>
                      <td className="border-t px-4 py-2 bg-gray-300 dark:bg-gray-700 h-6"></td>
                      <td className="border-t px-4 py-2 bg-gray-300 dark:bg-gray-700 h-6"></td>
                      <td className="border-t px-4 py-2 bg-gray-300 dark:bg-gray-700 h-6"></td>
                      <td className="border-t px-4 py-2 bg-gray-300 dark:bg-gray-700 h-6"></td>
                      <td className="border-t px-4 py-2 bg-gray-300 dark:bg-gray-700 h-6"></td>
                      <td className="border-t px-4 py-2 bg-gray-300 dark:bg-gray-700 h-6"></td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          ) : loansError ? (
            <div className="hidden lg:flex lg:bg-gray-100 lg:dark:bg-gray-800 lg:h-[400px] lg:flex-col lg:items-center lg:justify-center h-20">
              <TbFileSad
          className={`text-gray-300 dark:text-[#993d4b] w-[400px] h-[70px] pb-2 block mx-auto`}
          strokeWidth={1}
        />
              <div className="text-red-500 text-center">{loansError}</div>
            </div>
          ) : (
            <div className="overflow-x-auto">
              <table className="w-[100%] bg-white rounded-2xl shadow-md table-fixed dark:bg-dark text-gray-900 dark:text-white">
                <thead>
                  <tr>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400">
                      SL No
                    </th>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400">
                      Loan Money
                    </th>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400">
                      Left to Repay
                    </th>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400">
                      Duration
                    </th>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400">
                      Interest Rate
                    </th>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400">
                      Monthly Installment
                    </th>
                    <th className="px-4 py-2 text-left font-semibold text-gray-400">
                      Repay
                    </th>
                  </tr>
                </thead>
                <tbody>
                  {loans.map((loan: any, index) => (
                    <tr key={index}>
                      <td className="border-t px-4 py-2 text-sm">
                        {index + 1}
                      </td>
                      <td className="border-t px-4 py-2 text-sm">
                        {loan.loanAmount}
                      </td>
                      <td className="border-t px-4 py-2 text-sm">
                        {loan.amountLeftToRepay}
                      </td>
                      <td className="border-t px-4 py-2 text-sm">
                        {loan.loanDuration} months
                      </td>
                      <td className="border-t px-4 py-2 text-sm">
                        {loan.interestRate}%
                      </td>
                      <td className="border-t px-4 py-2 text-sm">
                        {loan.installment}
                      </td>
                      <td className="border-t px-4 py-2 text-sm">
                        <button
                          onClick={() => onSubmit(loan.amountLeftToRepay , index)}
                          className="text-purple-900 border border-purple-900 rounded-full px-4 py-1 hover:bg-gray-300"
                        >
                           { selectedIndex === index && isLoading ? (
                            <div className="flex justify-center items-center ">
                              <ArrowPathIcon className="h-5 w-5 animate-spin  text-gray-500  " />
                            </div>
                          ) : (
                            "Repay"
                          )}
                        </button>
                      </td>
                    </tr>
                  ))}
                  <tr className="font-bold text-red-500">
                    <td className="border-t px-4 py-2"></td>
                    <td className="border-t px-4 py-2">
                      ${totalLoanMoney.toLocaleString()}
                    </td>
                    <td className="border-t px-4 py-2">
                      ${totalLeftToRepay.toLocaleString()}
                    </td>
                    <td className="border-t px-4 py-2"></td>
                    <td className="border-t px-4 py-2"></td>
                    <td className="border-t px-4 py-2">
                      ${totalInstallment.toLocaleString()}
                    </td>
                    <td className="border-t px-4 py-2"></td>
                  </tr>
                </tbody>
              </table>
            </div>
          )}
           <Pagination
        currentPage={currentPage}
        totalPages={totalPages}
        onPageChange={setCurrentPage}
      />
          
        </div>
      </div>
    </div>
  );
};

export default LoansPage;
