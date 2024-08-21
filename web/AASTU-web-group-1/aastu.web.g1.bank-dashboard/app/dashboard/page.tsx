"use client";
import { useEffect, useRef, useState } from "react";
import {
  IoChevronBackCircleOutline,
  IoChevronForwardCircleOutline,
} from "react-icons/io5";

import {
  BalanceData,
  CardDetails,
  QuickTransferData,
  TransactionContent,
  TransactionData,
} from "@/types";
import { PiTelegramLogoLight } from "react-icons/pi";
import CreditCard from "./_components/Credit_Card";
import { Profile } from "./_components/Profile";
import { Transaction } from "./_components/Transaction";
import { Pie_chart } from "@/app/dashboard/_components/Pie_chart";
import { BalanceAreachart } from "./transactions/component/balanceChart";
import { Barchart } from "./transactions/component/weeklyActivityChart";
import getRandomBalance, { addTransactions, getallTransactions, getCreditCards, getExpenses, getIncomes, getQuickTransfer } from "@/lib/api";
import { Loading } from "./_components/Loading";
import {useUser} from "@/contexts/UserContext"
import Link from "next/link";

const MainDashboard = () => {
  const {isDarkMode} = useUser();
  const QuickTransferSection = useRef<HTMLDivElement | null>(null);

  const scrollCards = (scrollOffset: number) => {
    if (QuickTransferSection.current) {
      QuickTransferSection.current.scrollLeft += scrollOffset;
    }
  };

  const [loading, setLoading] = useState(true);
  const [creditCards, setCreditCards] = useState<CardDetails[]>([]);
  const [recentTransactions, setRecentTransactions] = useState<
    TransactionContent[]
  >([]);
  const [transactions, setTransactions] = useState<TransactionContent[]>([]);
  const [balanceHistory, setBalanceHistory] = useState<BalanceData[]>([]);
  const [weeklyIncome, setWeeklyIncome] = useState<TransactionContent[]>([]);
  const [weeklyWithdraw, setWeeklyWithdraw] = useState<TransactionContent[]>([]);
  const [quickTransfer, setQuickTransfer] = useState<QuickTransferData[]>([]);
  const [selectedProfile, setSelectedProfile] =
    useState<QuickTransferData | null>(null);
  const [amount, setAmount] = useState<string>("");
  let totalCreditcardpage;
  const handleProfileSelect = (account: QuickTransferData) => {
    setSelectedProfile(account);
  };

  const handleAmountChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setAmount(e.target.value);
  };

  const handleSend = () => {
    if (selectedProfile) {
      console.log("Sending to:", selectedProfile.username, "Amount:", amount);
      addTransactions({type:"transfer",amount:parseInt(amount),receiverUserName:selectedProfile.username,description:"Quick Transfer",})
    }
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        const res = await getCreditCards(0,2);
        const recent = await getallTransactions(0, 3);
        const statistics = await getallTransactions(0, 100);
        const BalanceHistory = await getRandomBalance();
        const incomes = await getIncomes(0, 7);
        const withdraw = await getExpenses(0, 7);
        const accounts = await getQuickTransfer(7);
        setCreditCards(res?.content || []);
        totalCreditcardpage = res?.totalPages;
        setRecentTransactions(recent?.content || []);
        setTransactions(statistics?.content || []);
        setBalanceHistory(BalanceHistory || []);
        setWeeklyIncome(incomes?.content || []);
        setWeeklyWithdraw(withdraw?.content || []);
        setQuickTransfer(accounts || []);
      } finally {
        setLoading(false);
      }
    };
    fetchData();
  }, []);

  if (loading) {
    return (
      <Loading/>
    );
  }

  return (
    <div
      className={`p-5 space-y-5 ${
        isDarkMode ? "bg-gray-700 text-white" : "bg-[#F5F7FA] text-black"
      }`}
    >
      {/* First Row: My Cards and Recent Transactions */}
      <div className="md:flex sm:grid-cols-2 md:gap-5 space-y-5 md:space-y-0">
        {/* My Cards Section */}
        <div className="md:w-2/3 space-y-5">
          <div className="flex justify-between font-inter text-[16px] font-semibold">
            <h4>My Cards</h4>

            <h4>
              <Link href="/dashboard/credit-cards/">
              
                  See All
               
              </Link>
            </h4>
          </div>
          <div className="flex space-x-5 overflow-x-auto [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]">
            {creditCards.map((card) => (
              <CreditCard
                key={card.id}
                id={card.id}
                balance={card.balance}
                semiCardNumber={card.semiCardNumber}
                cardHolder={card.cardHolder}
                expiryDate={card.expiryDate}
                cardType={card.cardType}
              />
            ))}
          </div>
        </div>

        {/* Recent Transactions Section */}
        <div className="space-y-5 md:space-y-5 w-full md:w-1/3">
          <div className="font-inter text-[16px] font-semibold">
            <h4>Recent Transactions</h4>
          </div>
          <div
            className={`space-y-3 p-3 md:p-5  ${
              isDarkMode
                ? "bg-gray-800 text-white border-gray-600"
                : "bg-white text-black"
            }
        rounded-xl
        md:shadow-lg
          `}
          >
            {recentTransactions.map((transaction) => (
              <Transaction
                key={transaction.transactionId}
                date={transaction.date}
                amount={transaction.amount}
                description={transaction.description}
                type={transaction.type}
                transactionId={transaction.transactionId}
                senderUserName={transaction.senderUserName}
                receiverUserName={transaction.receiverUserName}
              />
            ))}
          </div>
        </div>
      </div>

      {/* Second Row: Weekly Activity and Expense Statistics */}
      <div className="md:flex sm:grid-cols-2 md:gap-5 space-y-5 md:space-y-0">
        {/* Weekly Activity Section */}
        <div className="md:w-2/3 space-y-5">
          <div className="font-inter text-[16px] font-semibold">
            <h4>Weekly Activity</h4>
          </div>
          <div
            className={`${
              isDarkMode
                ? "bg-gray-800 text-white border-gray-600"
                : "bg-white text-black "
            }  md:shadow-lg  rounded-xl `}
          >
            <Barchart
              weeklyDeposit={weeklyIncome}
              weeklyWithdraw={weeklyWithdraw}
            />
          </div>
        </div>

        {/* Expense Statistics Section */}
        <div className="md:w-1/3 space-y-5">
          <div className="font-inter text-[16px] font-semibold">
            <h4>Expense Statistics</h4>
          </div>
          <div className="bg-white rounded-xl md:shadow-lg ">
            <Pie_chart transactions={transactions} />
          </div>
        </div>
      </div>

      {/* Third Row: Quick Transfer and Balance History */}
      <div className="md:grid md:grid-cols-[1fr,2fr] md:gap-10 space-y-5 md:space-y-0">
        {/* Quick Transfer Section */}
        <div className="space-y-5">
          <div className="font-inter text-[16px] font-semibold">
            <h4>Quick Transfer</h4>
          </div>
          <div
            className={`
        ${isDarkMode ? "bg-gray-800 text-white" : "bg-white text-black"}
        rounded-xl
        md:shadow-lg
  
        p-5
        space-y-5
      `}
          >
            <div>
              <button
                className={`
            float-right
            hover:bg-blue-500
            rounded-xl
            ${isDarkMode ? "hover:bg-blue-600" : "hover:bg-blue-500"}
          `}
                onClick={() => scrollCards(200)}
              >
                <IoChevronForwardCircleOutline size={20} />
              </button>
              <button
                className={`
            float-left
            hover:bg-blue-500
            rounded-xl
            ${isDarkMode ? "hover:bg-blue-600" : "hover:bg-blue-500"}
          `}
                onClick={() => scrollCards(-200)}
              >
                <IoChevronBackCircleOutline size={20} />
              </button>
            </div>

            <div
              ref={QuickTransferSection}
              className={`
          flex
          max-w-[300px]
          space-x-5
          overflow-x-auto
          [&::-webkit-scrollbar]:hidden
          [-ms-overflow-style:none]
          [scrollbar-width:none]
        
        `}
            >
              {quickTransfer.map((account) => (
                <Profile
                  key={account.id}
                  image="/images/avatar2.svg"
                  name={account.name}
                  job="Director"
                  isSelected={selectedProfile?.id === account.id}
                  onClick={() => handleProfileSelect(account)}
                />
              ))}
            </div>
            <div className="flex space-x-10 h-[40px] items-center">
              <h4
                className={`
            font-inter
            text-[12px]
            ${isDarkMode ? "text-[#9AA1B4]" : "text-[#718EBF]"}
          `}
              >
                Write Amount
              </h4>
              <div
                className={`
            rounded-3xl
            flex
            items-center
            ${isDarkMode ? "bg-gray-700" : "bg-gray-200"}
          `}
              >
                <input
                  type="number"
                  className={`
              w-[90px]
              h-[40px]
              rounded-full
              ${isDarkMode ? "bg-gray-600" : "bg-gray-200"}
              px-3
              ${isDarkMode ? "text-white" : "text-black"}
            `}
                  placeholder="525.50"
                  value={amount}
                  step="0.01" // Allows decimal numbers with two decimal places
                  min="0"
                  onChange={handleAmountChange}
                />
                <button
                  className={`
              ${isDarkMode ? "bg-[#3B6EE2]" : "bg-[#1814F3]"}
              text-white
              rounded-full
              px-4
              h-[40px]
              ml-2
              flex
              items-center
              space-x-2
            `}
                  onClick={handleSend}
                  disabled={!selectedProfile || !amount}
                >
                  <p>Send</p>
                  <PiTelegramLogoLight />
                </button>
              </div>
            </div>
          </div>
        </div>

        {/* Balance History Section */}
        <div className="space-y-5">
          <div className="font-inter text-[16px] font-semibold">
            <h4>Balance History</h4>
          </div>
          <div
            className={`
        ${isDarkMode ? "bg-gray-800  shadow-md" : "bg-white  shadow-lg"}
        rounded-xl
        md:shadow
        transition-all
        duration-300
        
      `}
          >
            <BalanceAreachart balanceHistory={balanceHistory} />
          </div>
        </div>
      </div>
    </div>
  );
};

export default MainDashboard;