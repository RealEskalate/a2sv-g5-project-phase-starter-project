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
import getRandomBalance, {
  addTransactions,
  getallTransactions,
  getCreditCards,
  getExpenses,
  getIncomes,
  getQuickTransfer,
} from "@/lib/api";
import { Loading } from "./_components/Loading";
import { useUser } from "@/contexts/UserContext";
import Link from "next/link";
import { toast } from "sonner";
import { ModalTrans } from "./_components/ModalTrans";
const MainDashboard = () => {
  const { isDarkMode } = useUser();
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
  const [weeklyWithdraw, setWeeklyWithdraw] = useState<TransactionContent[]>(
    []
  );
  const [quickTransfer, setQuickTransfer] = useState<QuickTransferData[]>([]);
  const [selectedProfile, setSelectedProfile] =
    useState<QuickTransferData | null>(null);
  const [amount, setAmount] = useState<string>("");
  const [sendLoading, setSendLoading] = useState(false);
  const [isModalOpen, setIsModalOpen] = useState(false);
  let totalCreditcardpage;
  const handleProfileSelect = (account: QuickTransferData) => {
    setSelectedProfile(account);
  };

  const handleAmountChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setAmount(e.target.value);
  };

  const handleModalToggle = () => {
    setIsModalOpen(!isModalOpen);
  };

  const handleSend = async () => {
    if (selectedProfile) {
      console.log("Sending to:", selectedProfile.username, "Amount:", amount);
      setSendLoading(true);
      const result: boolean | undefined = await addTransactions({
        type: "transfer",
        amount: parseInt(amount),
        receiverUserName: selectedProfile.username,
        description: "Quick Transfer",
      });
      setLoading(false);
      if (result) {
        toast("sucess sending");
      } else {
        toast("failed sending");
      }
      setLoading(true);
    }
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        const res = await getCreditCards(0, 2);
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
    return <Loading />;
  }

  return (
    <div
      className={` relative p-5 space-y-5 ${
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
              <Link href="/dashboard/credit-cards/">See All</Link>
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
            // min-w-[300px] sm:h-56 md:min-w-[320px]  ${bgColor} rounded-2xl pt-3 space-y-6
            className={`space-y-5 p-3 md:p-5 md:h-[200px] lg:w-[350px] lg:h-[220px] ${
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
    ${
      isDarkMode
        ? "bg-[#3B6EE2] hover:bg-[#2A56B8]"
        : "bg-[#1814F3] hover:bg-[#0F0DC7]"
    }
    text-white
    rounded-full
    px-4
    h-[40px]
    ml-2
    flex
    items-center
    space-x-2
    transition-all duration-300 ease-in-out
  `}
                  onClick={handleModalToggle}
                  disabled={!selectedProfile || !amount}
                >
                  <p>Send</p>
                  <PiTelegramLogoLight />
                </button>
              </div>
              {isModalOpen && (
                <div
                  className="fixed inset-0 z-50 flex justify-center items-center bg-black bg-opacity-50 backdrop-blur-sm"
                  onClick={handleModalToggle}
                >
                  <div
                    className="relative bg-white p-6 rounded-lg shadow-lg max-w-md w-full"
                    onClick={(e) => e.stopPropagation()} // Prevent modal from closing when clicking inside it
                  >
                    <ModalTrans
                      isOpen={isModalOpen}
                      onClose={handleModalToggle}
                    />
                  </div>
                </div>
              )}
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
      {/* {sendLoading && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 backdrop-blur-md z-50">
          <div role="status">
            <svg
              aria-hidden="true"
              className="w-12 h-12 text-gray-200 animate-spin dark:text-gray-600 fill-blue-600"
              viewBox="0 0 100 101"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
                fill="currentColor"
              />
              <path
                d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
                fill="currentFill"
              />
            </svg>
            <span className="sr-only">Loading...</span>
          </div>
        </div>
      )} */}
    </div>
  );
};

export default MainDashboard;
