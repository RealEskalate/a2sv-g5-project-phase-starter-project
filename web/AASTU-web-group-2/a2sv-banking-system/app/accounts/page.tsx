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
import User from "@/types/userInterface";
import Preference from "@/types/userInterface";
import {
  getTransactionIncomes,
  getTransactionsExpenses,
} from "@/lib/api/transactionController";
// import { PaginatedTransactionsResponse } from "@/types/transactionController.interface";

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
  const router = useRouter();
  const [loading, setLoading] = useState(true);
  const [getCard, setGetCards] = useState<CardType[]>();
  const [currentUser, setCurrentUser] = useState();
  const [income, setIncome] = useState(0);
  const [expense, setExpense] = useState(0);
  // Getting the session from the server
  useEffect(() => {
    const fetchSession = async () => {
      const sessionData = (await getSession()) as SessionDataType | null;
      if (sessionData && sessionData.user) {
        setSession(sessionData.user);
      } else {
        router.push(
          `./api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`
        );
      }
      setLoading(false);
    };

    fetchSession();
  }, [router]);

  // Fetching cards
  useEffect(() => {
    const addingData = async () => {
      if (session?.access_token) {
        const cardData = await getCards(session?.access_token);
        console.log("Fetching Complete", cardData.content);
        setGetCards(cardData.content);
      }
    };
    addingData();
  });

  // Fetching Balance
  useEffect(() => {
    const addingData = async () => {
      if (session?.access_token) {
        const current = await getCurrentUser(session?.access_token);
        setCurrentUser(current);
      }
    };
    addingData();
  });

  // Fetching Income
  useEffect(() => {
    const addingData = async () => {
      if (session?.access_token) {
        const current = await getTransactionIncomes(
          0,
          1,
          session?.access_token
        );
        console.log("INCOME", current.data);
        current.data.content.map((items: any) => {
          setIncome(income + items.amount);
        });
      }
    };
    addingData();
  });

  // Fetching Expense
  useEffect(() => {
    const addingData = async () => {
      if (session?.access_token) {
        const current = await getTransactionsExpenses(
          0,
          1,
          session?.access_token
        );
        console.log("Expense", current.data);
        current.data.content.map((items: any) => {
          setExpense(expense + items.amount);
        });
      }
    };
    addingData();
  });
  // console.log("USER, ", currentUser)
  // Example data for the first ListCard
  const ReusableCard: Column = {
    icon: MdHome,
    iconStyle: "text-[#FFBB38] bg-[#FFF5D9]",
    data: [
      {
        heading: "My Balance",
        text: currentUser?.accountBalance ?? "",
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

  if (loading) return null; // Don't render anything while loading

  if (!session) {
    router.push(
      `./api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`
    );
    return null;
  }

  console.log("get Card", getCard);
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
            {getCard &&
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
              ))}
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
