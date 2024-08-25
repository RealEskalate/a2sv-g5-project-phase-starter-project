"use client";
import { useEffect, useState } from "react";
import { Loading } from "./_components/Loading";
import { useUser } from "@/contexts/UserContext";
import { Cards } from "./_components/dashboardCards";
import { RecentTransaction } from "./_components/RecentTransaction";
import { WeeklyActivity } from "./_components/WeeklyActivity";
import { ExpenseStatistics } from "./_components/ExpenseStatistics";
import { BalanceHistory } from "./_components/BalanceHistory";
import { QuickTransfer } from "./_components/QuickTransfer";

const MainDashboard = () => {
  const { isDarkMode } = useUser();
  const [cardLoaded, setCardLoaded] = useState(false);
  const [recentTransactionLoaded, setRecentTransactionLoaded] = useState(false);
  const [weeklyActivityLoaded, setWeeklyActivityLoaded] = useState(false);
  const [expenseStatisticsLoaded, setExpenseStatisticsLoaded] = useState(false);
  const [quickTransferLoaded, setQuickTransferLoaded] = useState(false);
  const [balanceHistoryLoaded, setBalanceHistoryLoaded] = useState(false);
  const [dataFetched, setDataFetched] = useState(false);

  // Check if all data has been fetched
  useEffect(() => {
    if (
      cardLoaded &&
      recentTransactionLoaded &&
      weeklyActivityLoaded &&
      expenseStatisticsLoaded &&
      quickTransferLoaded &&
      balanceHistoryLoaded
    ) {
      setDataFetched(true);
    }
  }, [
    cardLoaded,
    recentTransactionLoaded,
    weeklyActivityLoaded,
    expenseStatisticsLoaded,
    quickTransferLoaded,
    balanceHistoryLoaded,
  ]);

  return (
    <div
      className={`min-h-screen 
        md:p-10 p-5 space-y-5 ${
          isDarkMode ? "bg-gray-700 text-white" : "bg-[#F5F7FA] text-black"
        }`}
    >
      {/* First Row: My Cards and Recent Transactions */}
      <div className="md:flex sm:grid-cols-2 md:gap-5 md:space-y-0 space-y-5 ">
        <Cards onLoadingComplete={() => setCardLoaded(true)} />
        <RecentTransaction
          onLoadingComplete={() => setRecentTransactionLoaded(true)}
        />
      </div>

      {/* First Row: My Cards and Recent Transactions */}
      <div className="md:flex sm:grid-cols-2 md:gap-5 space-y-5 md:space-y-0">
        <Cards onLoadingComplete={() => setCardLoaded(true)} />
        <RecentTransaction
          onLoadingComplete={() => setRecentTransactionLoaded(true)}
        />
      </div>

      {/* Second Row: Weekly Activity and Expense Statistics */}
      <div className="md:flex sm:grid-cols-2 md:gap-5 space-y-5 md:space-y-0">
        <WeeklyActivity
          onLoadingComplete={() => setWeeklyActivityLoaded(true)}
        />
        <ExpenseStatistics
          onLoadingComplete={() => setExpenseStatisticsLoaded(true)}
        />
      </div>

      {/* Third Row: Quick Transfer and Balance History */}
      <div className="md:grid md:grid-cols-[1fr,2fr] md:gap-10 space-y-5 md:space-y-0">
        <QuickTransfer onLoadingComplete={() => setQuickTransferLoaded(true)} />
        <BalanceHistory
          onLoadingComplete={() => setBalanceHistoryLoaded(true)}
        />
      </div>
    </div>
  );
};

export default MainDashboard;
