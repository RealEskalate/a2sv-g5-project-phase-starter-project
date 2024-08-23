"use client";
import React, { useEffect } from "react";
import Example from "../components/AccountBarChart";
import CreditCard from "../components/CreditCard";
import Image from "next/image";
import {
  spotify,
  user,
  mobileService,
  returnValue,
  nameinvestment,
  totalinvestment,
  apple,
} from "@/../../public/Icons";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "@/lib/redux/store";
import { useGetAllTransactionsQuery } from "@/lib/redux/api/transactionsApi";
import {
  setAllTransactions,
  setLoading,
  setError,
} from "@/lib/redux/slices/transactionsSlice";
import { setCards } from "@/lib/redux/slices/cardsSlice";
import { useGetCardsQuery } from "@/lib/redux/api/cardsApi";

interface CardProps {
  title: string;
  salary: number;
  index: number;
  icon: string;
}

interface Transaction {
  title: string;
  jobtitle: string;
  creditcard: string;
  status: string;
  value: number;
  date: string;
  icon: string;
}

const dataCorner: CardProps[] = [
  { icon: nameinvestment, title: "My Balance", salary: 12000, index: 1 },
  { icon: spotify, title: "Income", salary: 5600, index: 2 },
  { icon: totalinvestment, title: "Expense", salary: 3460, index: 3 },
  { icon: returnValue, title: "Total Saving", salary: 7920, index: 4 },
];

const transactions: Transaction[] = [
  {
    icon: spotify,
    title: "Spotify Subscription",
    jobtitle: "Shopping",
    creditcard: "1234****",
    status: "Pending",
    value: 150,
    date: "25 Jan 2021",
  },
  {
    icon: mobileService,
    title: "Mobile Service",
    jobtitle: "Service",
    creditcard: "1234****",
    status: "Pending",
    value: 1200,
    date: "15 Feb 2021",
  },
  {
    icon: user,
    title: "Grocery Shopping",
    jobtitle: "Supermarket",
    creditcard: "1234****",
    status: "Completed",
    value: 350,
    date: "10 Mar 2021",
  },
];

function formatDate(dateString: string): string {
  const date = new Date(dateString);
  const month = (date.getMonth() + 1).toString().padStart(2, "0"); // Add 1 because months are zero-indexed
  const year = date.getFullYear().toString()?.slice(-2); // Get the last two digits of the year
  return `${month}/${year}`;
}

function formatCardNumber(cardNumber: string): string {
  const start = cardNumber?.slice(0, 4);
  const end = cardNumber?.slice(-4);
  return `${start} **** **** ${end}`;
}

function formatBalance(balance: number): string {
  return new Intl.NumberFormat("en-US", {
    style: "currency",
    currency: "USD",
    minimumFractionDigits: 2,
  }).format(balance);
}

const Page = () => {
  const dispatch = useDispatch();
  const { allTransactions, loading, error } = useSelector(
    (state: RootState) => state.transactions
  );
  const { cards } = useSelector((state: RootState) => state.cards);

  const {
    data: allData,
    isLoading: isLoadingAll,
    isError: isErrorAll,
  } = useGetAllTransactionsQuery({ size: 3, page: 0 });

  const {
    data,
    error: cardError,
    isLoading,
  } = useGetCardsQuery({ size: 1, page: 0 });

  useEffect(() => {
    if (!isLoading && !cardError && data) {
      dispatch(setCards(data.content));
    }
  }, [data, isLoading, cardError, dispatch]);

  useEffect(() => {
    dispatch(setLoading(isLoadingAll));

    if (allData) {
      dispatch(setAllTransactions(allData.data.content));
    }

    if (isErrorAll) {
      console.error("Error fetching data", {
        isErrorAll,
      });
      dispatch(setError("Error loading transactions"));
    }
  }, [allData, isLoadingAll, isErrorAll, dispatch]);

  console.log("cards", cards);
  console.log("all transaction", allTransactions);
  return (
    <div className="max-w-screen max-h-screen mx-auto px-4 py-4 sm:px-6 sm:py-3 lg:px-4 lg:py-8">
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-4">
        {dataCorner.map((card) => (
          <div
            key={card.index}
            className="flex justify-center items-center h-[85px] rounded-2xl shadow-xl bg-white p-4"
          >
            <div className="flex items-center">
              <Image
                width={28}
                height={28}
                className=" text-gray-500 mr-4"
                src={card.icon}
                alt="Image Icon"
              />
              <div>
                <h3 className="text-gray-500 text-sm">{card.title}</h3>
                <p className="text-black text-2xl font-bold">
                  ${card.salary.toLocaleString()}
                </p>
              </div>
            </div>
          </div>
        ))}
      </div>

      <div className="flex flex-col md:flex-row justify-between">

        <div className=" w-full md:w-[53%] lg:w-[63%]">
        <h1 className="text-[22px] font-bold leading-[26.63px] text-[rgba(51,59,105,1)] text-left px-4 py-2">
          Last transaction
        </h1>
          <div className="rounded-3xl bg-white shadow-xl w-full">
            <div className="flex flex-col space gap-y-2">
              {allTransactions.map((transaction, index) => (
                <div
                  key={index}
                  className="flex items-center p-3 bg-white rounded-lg shadow-sm"
                >
                  <div className=" p-3 rounded-full">
                    <Image
                      width={20}
                      height={20}
                      src={transactions[index].icon}
                      alt={`${transaction.type} Icon`}
                    />
                  </div>
                  <div className=" flex w-3/4 gap-2 items-center  justify-between ">
                    <div className="w-1/2 md:w-1/3 lg:w-1/4">
                      <p className="text-gray-800 font-medium">
                        {transaction.senderUserName}
                      </p>
                      <p className="text-gray-400 hidden text-xs sm:block">
                        {transaction.date}
                      </p>
                    </div>
                    <div className="w-1/3 md:w-1/5 lg:w-1/4 flex">
                      <p className="text-gray-400 text-xs w-10/12 flex-shrink-0 text-left truncate">
                        {transaction.description}
                      </p>
                    </div>
                    <div className="w-1/5 lg:w-1/4 flex">
                      <p className="text-gray-400 w-10/12 flex-shrink-0 truncate hidden text-xs lg:block">
                        {transaction.transactionId}
                      </p>
                    </div>
                    <div className="w-1/5 lg:w-1/4 flex">
                      <p
                        className={` style={{ color: '#718EBF' }} hidden lg:block font-medium w-10/12 flex-shrink-0 truncate text-xs ${
                          transaction.type
                            ? "style={{ color: '#718EBF' }}"
                            : "text-green-500"
                        }`}
                      >
                        {transaction.type}
                      </p>
                      {/* <p className="text-gray-400 w-10/12 flex-shrink-0 truncate hidden text-xs sm:block">{transaction.transactionId}</p> */}
                    </div>
                    <div className="w-1/5 lg:w-1/4 flex">
                      <p
                        className={`text-400 font-medium w-10/12 text-right ${
                          transaction.amount < 0
                            ? "text-red-500"
                            : "text-green-500"
                        }`}
                      >
                        ${Math.abs(transaction.amount).toLocaleString()}
                      </p>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
            </div>

          <div className="w-full md:w-[43%] lg:w-[35%] rounded-3xl overflow-hidden col-span-1 ">
            <h1 className="text-[22px] font-bold leading-[26.63px] text-[rgba(51,59,105,1)] text-left px-4 py-2">
              my card
            </h1>
            {/* <div className="flex items-center justify-center"> */}

            <div className="min-w-60 w-73 lg:w-72 lg:h-48 xl:w-96  xl:h-56  ">
              <CreditCard
                name={cards[1]?.cardHolder}
                balance={formatBalance(cards[1]?.balance)}
                cardNumber={formatCardNumber(cards[1]?.semiCardNumber)}
                validDate={formatDate(cards[1]?.expiryDate)}
                backgroundImg="bg-[linear-gradient(107.38deg,#2D60FF_2.61%,#539BFF_101.2%)]"
                textColor="text-white"
                />
                
            </div>
                {/* </div> */}
          </div>
      </div>

      <h1 className="text-[22px] font-bold leading-[26.63px] text-[rgba(51,59,105,1)] px-4 py-7">
        Debit & Credit Overview
      </h1>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-4 ">
        <div className="rounded-3xl bg-white shadow-xl p-6 lg:col-span-2 space gap-y-3 h-80">
          <span className="text-left font-inter text-sm font-normal leading-4 px-6 py-2">
            $7,560 Debited & $5,420 Credited in this Week
          </span>
          <Example />
        </div>

        <div className="rounded-3xl bg-white shadow-md p-6">
          <h3 className="text-lg font-semibold">Invoices Sent</h3>

          <div className="flex flex-col space-y-3 mt-3">
            {/* Invoice 1 */}
            <div className="flex justify-between items-center">
              <div className="flex items-center space-x-2">
                <div className="bg-[#DCFAF8] w-12 h-12 rounded-2xl flex items-center justify-center">
                  <Image
                    width={20}
                    height={20}
                    src={apple}
                    alt="Deposit Icon"
                  />
                </div>
                <div className="flex flex-col">
                  <p className="text-sm font-light text-gray-400">
                    Apple Store
                  </p>
                  <small className="text-xs text-blue-500">5h ago</small>
                </div>
              </div>
              <p className="text-blue-500 font-light ml-auto">$450</p>
            </div>

            {/* Invoice 2 */}
            <div className="flex justify-between items-center">
              <div className="flex items-center space-x-2">
                <div className="bg-[#FFF5D9] w-12 h-12 rounded-2xl flex items-center justify-center">
                  <Image
                    width={20}
                    height={20}
                    src={spotify}
                    alt="Spotify"
                    className="w-6 h-6 object-contain"
                  />
                </div>
                <div className="flex flex-col">
                  <p className="text-sm font-light text-gray-400">Michael</p>
                  <small className="text-xs text-blue-500">2 days ago</small>
                </div>
              </div>
              <p className="text-blue-500 font-light ml-auto">$160</p>
            </div>

            {/* Invoice 3 */}
            <div className="flex justify-between items-center">
              <div className="flex items-center space-x-2">
                <div className="bg-[#E7EDFF] w-12 h-12 rounded-2xl flex items-center justify-center">
                  <Image
                    width={20}
                    height={20}
                    src={spotify}
                    alt="Playstation"
                    className="w-6 h-6 object-contain"
                  />
                </div>
                <div className="flex flex-col">
                  <p className="text-sm font-light text-gray-400">
                    Playstation
                  </p>
                  <small className="text-xs text-blue-500">5 days ago</small>
                </div>
              </div>
              <p className="text-blue-500 font-light ml-auto">$1085</p>
            </div>

            {/* Invoice 4 */}
            <div className="flex justify-between items-center">
              <div className="flex items-center space-x-2">
                <div className="bg-[#FFE0EB] w-12 h-12 rounded-2xl flex items-center justify-center">
                  <Image
                    width={20}
                    height={20}
                    src={spotify}
                    alt="William"
                    className="w-6 h-6 object-contain"
                  />
                </div>
                <div className="flex flex-col">
                  <p className="text-sm font-light text-gray-400">William</p>
                  <small className="text-xs text-blue-500">10 days ago</small>
                </div>
              </div>
              <p className="text-blue-500 font-light ml-auto">$90</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Page;
