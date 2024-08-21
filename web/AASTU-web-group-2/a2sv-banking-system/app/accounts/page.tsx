"use client";
import React, { useEffect, useState } from "react";
import {
  MdHome,
  MdSettings,
  MdAttachMoney,
  MdAccountBalance,
} from "react-icons/md";
import ListCard from "./components/ListCard";
import { IconType } from "react-icons";
import BarChartForAccounts from "./components/BarChartForAccounts";
import Card from "../components/Page2/Card";
import { getSession } from "next-auth/react";
import { useRouter } from "next/navigation";
import { getCards } from "@/lib/api/cardController";
import { Card as CardType } from "@/types/cardController.Interface";
import { getCurrentUser } from "@/lib/api/userControl";
import { UserInfo } from "@/types/userInterface";
import Refresh from "../api/auth/[...nextauth]/token/RefreshToken";
import {
  getTransactionIncomes,
  getTransactionsExpenses,
} from "@/lib/api/transactionController";
import Loading from "./components/Loading";
type DataItem = {
  heading: string;
  text: string;
  headingStyle: string;
  dataStyle: string;
};

type Column = {
  icon: IconType;
  iconStyle: string;
  data: DataItem[];
};

type Data = {
  access_token: string;
  data: string;
  refresh_token: string;
};

type SessionDataType = {
  user: Data;
};

const Page = () => {
  const [session, setSession] = useState<Data | null>(null);
  const [access_token, setAccess_token] = useState("");
  const router = useRouter();
  const [loading, setLoading] = useState(true);
  const [getCard, setGetCards] = useState<CardType[]>();
  const [currentUser, setCurrentUser] = useState<UserInfo>();
  const [income, setIncome] = useState(0);
  const [expense, setExpense] = useState(0);

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
        setLoading(false);
      }
    };

    fetchSession();
  }, [router]);

  // Combined fetching data to reduce multiple useEffect hooks
  useEffect(() => {
    const fetchData = async () => {
      if (!access_token) return;

      try {
        // Fetch Cards
        const cardData = await getCards(access_token);
        setGetCards(cardData.content);

        // Fetch Balance
        const currentUser = await getCurrentUser(access_token);
        setCurrentUser(currentUser);

        // Fetch Income
        const incomeData = await getTransactionIncomes(0, 1, access_token);
        const totalIncome = incomeData.data.content.reduce(
          (sum: number, item: any) => sum + item.amount,
          0
        );
        setIncome(totalIncome);

        // Fetch Expense
        const expenseData = await getTransactionsExpenses(0, 1, access_token);
        const totalExpense = expenseData.data.content.reduce(
          (sum: number, item: any) => sum + item.amount,
          0
        );
        setExpense(totalExpense);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    fetchData();
  }, [access_token]);

  // Example data for the first ListCard
  const ReusableCard: Column = {
    icon: MdHome,
    iconStyle: "text-[#FFBB38] bg-[#FFF5D9]",
    data: [
      {
        heading: "My Balance",
        text: String(currentUser?.accountBalance) ?? "",
        headingStyle: "text-sm font-bold text-nowrap text-[#718EBF]",
        dataStyle: "text-xs text-nowrap",
      },
    ],
  };

  // Example data for the second ListCard
  const card1: Column = {
    icon: MdAttachMoney, // Updating the icon
    iconStyle: "text-[#396AFF] bg-[#E7EDFF]", // Updating the iconStyle
    data: ReusableCard.data.map((item) => ({
      ...item,
      text: String(income),
      heading: "Income", // Updating the heading
    })),
  };

  // Example data for the third ListCard
  const card2: Column = {
    icon: MdSettings, // Updating the icon
    iconStyle: "text-[#FF82AC] bg-[#FFE0EB]", // Updating the iconStyle
    data: ReusableCard.data.map((item) => ({
      ...item,
      text: String(expense),
      heading: "Expense", // Updating the heading
    })),
  };

  // Example data for the fourth ListCard
  const card3: Column = {
    icon: MdAccountBalance, // Updating the icon
    iconStyle: "text-[#16DBCC] bg-[#DCFAF8]", // Updating the iconStyle
    data: ReusableCard.data.map((item) => ({
      ...item,
      heading: "Total Savings", // Updating the heading
    })),
  };

  // First column with multiple data items
  const ReusableLastTransaction: Column = {
    icon: MdHome,
    iconStyle: "text-[#FFBB38] bg-[#FFF5D9]",
    data: [
      {
        heading: "Spotify Subscription",
        text: "25 Jan 2021",
        headingStyle: "text-sm font-bold text-nowrap",
        dataStyle: "text-xs text-nowrap text-[#718EBF]",
      },
      {
        heading: "-$150",
        text: "",
        headingStyle: "text-xs font-bold text-[#FE5C73]",
        dataStyle: "text-xs text-nowrap",
      },
    ],
  };

  // First transaction example
  const transaction1: Column = {
    icon: MdAccountBalance, // Different icon
    iconStyle: "text-[#16DBCC] bg-[#DCFAF8]", // Different iconStyle
    data: ReusableLastTransaction.data.map((item, index) => ({
      ...item,
      heading: index === 0 ? "Mobile Services" : item.heading, // Custom heading for the first item
    })),
  };

  const transaction2: Column = {
    icon: MdAttachMoney, // Updating the icon
    iconStyle: "text-[#16DBCC] bg-[#DCFAF8]", // Updating the iconStyle
    data: ReusableLastTransaction.data.map((item, index) => ({
      ...item,
      heading: index === 0 ? "Emilly Wilson " : "+$780",
      headingStyle:
        index === 0 ? item.headingStyle : "text-xs font-bold text-[#16DBAA]",
    })),
  };

  if (loading) {
    return <Loading></Loading>;
  }

  // Don't render anything while loading

  if (!session) {
    router.push(
      `./api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`
    );
    return null;
  }

  return (
    <>
      <div className="flex flex-col h-full bg-[#F5F7FA] px-3 py-3 gap-5">
        <div>
          <div className="flex flex-wrap gap-2">
            <ListCard column={ReusableCard} width={"w-[48%] md:w-[23%]"} />
            <ListCard column={card1} width={"w-[48%] md:w-[23%]"} />
            <ListCard column={card2} width={"w-[48%] md:w-[23%]"} />
            <ListCard column={card3} width={"w-[48%] md:w-[23%]"} />
          </div>
        </div>

        <div className="flex flex-col md:flex-row gap-5">
          <div className="flex flex-col gap-5 md:w-1/2">
            <span className="text-xl text-[#333B69] font-semibold">
              Last Transaction
            </span>
            <div className="bg-white flex flex-col justify-between rounded-2xl">
              <ListCard column={ReusableLastTransaction} width={"w-full"} />
              <ListCard column={transaction1} width={"w-full"} />
              <ListCard column={transaction2} width={"w-full"} />
            </div>
          </div>

          <div className="md:w-1/2 gap-1 flex flex-col">
            <div className="flex justify-between mr-2">
              <span className="text-xl text-[#333B69] font-semibold">
                My Card
              </span>
              <span className="text-sm text-[#333B69] font-semibold">
                See All
              </span>
            </div>
            {getCard ? (
              getCard.map((items) => (
                <Card
                  key={items.id}
                  balance={String(items.balance)}
                  cardHolder={items.cardHolder}
                  validThru={formatDate(items.expiryDate)}
                  cardNumber="3778 **** **** 1234"
                  filterClass=""
                  bgColor="from-[#4C49ED] to-[#0A06F4]"
                  textColor="text-white"
                  iconBgColor="bg-opacity-10"
                  showIcon={true}
                ></Card>
              ))
            ) : (
              <div className="border rounded-3xl my-4 mx-2 animate-pulse">
                <div className="relative w-full bg-gradient-to-b from-gray-300 to-gray-500 text-transparent rounded-3xl shadow-md h-[230px] min-w-[350px]">
                  <div className="flex justify-between items-start px-6 pt-6">
                    <div>
                      <p className="text-xs font-semibold bg-gray-400 rounded w-16 h-4 mb-2"></p>
                      <p className="text-xl font-medium bg-gray-400 rounded w-24 h-6"></p>
                    </div>
                    <div className="w-8 h-8 bg-gray-400 rounded-full"></div>
                  </div>

                  <div className="flex justify-between gap-12 mt-4 px-6">
                    <div>
                      <p className="text-xs font-medium bg-gray-400 rounded w-16 h-4 mb-2"></p>
                      <p className="font-medium text-base bg-gray-400 rounded w-24 h-6"></p>
                    </div>
                    <div className="pr-8">
                      <p className="text-xs font-medium bg-gray-400 rounded w-16 h-4 mb-2"></p>
                      <p className="font-medium text-base md:text-lg bg-gray-400 rounded w-24 h-6"></p>
                    </div>
                  </div>

                  <div className="relative mt-8 flex justify-between py-4 items-center">
                    <div className="absolute inset-0 w-full h-full bg-gradient-to-b from-white/30 to-transparent z-0"></div>
                    <div className="relative z-10 text-base font-medium px-6 bg-gray-400 rounded w-40 h-6"></div>
                    <div className="flex justify-end relative z-10 px-6">
                      <div className="w-12 h-12 bg-gray-400 rounded-full"></div>
                    </div>
                  </div>
                </div>
              </div>
            )}
          </div>
        </div>

        <div className="flex flex-col md:flex-row gap-5">
          <div className="flex flex-col gap-5 md:w-1/2">
            <span className="text-xl text-[#333B69] font-semibold">
              Debit & Credit Overview
            </span>
            <BarChartForAccounts></BarChartForAccounts>
          </div>
          <div className="flex flex-col gap-5 md:w-1/2">
            <span className="text-xl text-[#333B69] font-semibold">
              Invoice Sent
            </span>
            <div className="bg-white flex flex-col justify-between rounded-2xl">
              <ListCard column={ReusableLastTransaction} width={"w-full"} />
              <ListCard column={transaction1} width={"w-full"} />
              <ListCard column={transaction2} width={"w-full"} />
              <ListCard column={transaction2} width={"w-full"} />
            </div>
          </div>
        </div>
      </div>
    </>
  );
};
const formatDate = (dateString: string): string => {
  const date = new Date(dateString);

  const options: Intl.DateTimeFormatOptions = {
    year: "numeric",
    month: "short",
    day: "numeric",
  };

  return date.toLocaleDateString("en-US", options);
};

export default Page;
