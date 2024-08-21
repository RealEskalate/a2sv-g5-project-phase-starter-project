'use client'
import Image from "next/image";
import { IconType } from "react-icons";
import { useEffect, useState } from "react";
import ImageComponent from "./components/ImageComponent";
import Reviving from "./components/QuickTransfer";
import { BalanceHistory } from "./components/BalanceHistory";
import { WeeklyActivity } from "./components/WeeklyActivity";
import { ExpenseStatistics } from "./components/ExpenseStatistics";
import RecentTransaction from "./components/RecentTransaction";
import CreditCard from "./components/CreditCard";
import { getSession } from "next-auth/react";
import { useRouter } from "next/navigation";
import { getCards } from "@/lib/api/cardController";
import { GetCardsResponse, Card as CardType } from "@/types/cardController.Interface";


// import {RecentTransaction} from "@/components/RecentTransaction"
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

export default function Home() {
  const [session, setSession] = useState<Data | null>(null);
  const router = useRouter();
  const [loading, setLoading] = useState(true);
  const [getCard, setGetCards] = useState<CardType[]>();
  const route = useRouter();

  // getting the session ends here
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
        const cardData = await getCards(session?.access_token, 0, 3);
        console.log("Fetching Complete", cardData.content)
        setGetCards(cardData.content);
      }
    };
    addingData();
  });

  if (loading) return null; // Don't render anything while loading


  if (!session) {
    router.push(
      `./api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`
    );
    return null;
  }
  return (
    <div className="h-screen w-screen ">

    <div className="flex flex-col">
      {/* Mobile Version */}
      <div className="flex flex-col md:hidden">
      <div className="flex items-center justify-between">
              <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-2xl">My Cards</h1>
              <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-lg">See All</h1>
            </div>
        <div className="flex overflow-x-auto [&::-webkit-scrollbar]:hidden">

          <div className="flex-col">

            <div className="flex">
              <div className="flex min-w-max min-h-max [&::-webkit-scrollbar]:hidden">
              {getCard &&
              getCard.map((items, index) => (
                <CreditCard
                  key={items.id}
                  balance={String(items.balance)}
                  cardHolder={items.cardHolder}
                  validThru={formatDate(items.expiryDate)}
                  cardNumber="3778 **** **** 1234"
                  filterClass=""
                  bgColor={index % 2 === 0 ? "from-[#4C49ED] to-[#0A06F4]" : "bg-white"}
                  textColor={index%2 == 0 ? "text-white": "text-black"}
                  iconBgColor="bg-opacity-10"
                  showIcon={true}
                ></CreditCard>
              ))}  
              </div>
            </div>
          </div>
        </div>
        <RecentTransaction />
        <WeeklyActivity />
        <ExpenseStatistics />
        <Reviving />
        <BalanceHistory />
      </div>

      {/* Web Version */}
      <div className="hidden md:flex flex-col  px-6 py-4 bg-[#f5f7fa] h-[130vh]">
           {/* <div className="flex items-center justify-between">
              <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-2xl">My Cards</h1>
              <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-lg">See All</h1>
            </div> */}
        <div className="flex">
          <div className="flex flex-col w-1/2">
            {/* My Cards Section */}
            <div className="flex items-center justify-between">
              <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-2xl">My Cards</h1>
              <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-lg">See All</h1>
            </div>
            <div className="flex space-x-6 overflow-x-auto [&::-webkit-scrollbar]:hidden">
              {getCard &&
              getCard.map((items, index) => (
                <CreditCard
                  key={items.id}
                  balance={String(items.balance)}
                  cardHolder={items.cardHolder}
                  validThru={formatDate(items.expiryDate)}
                  cardNumber="3778 **** **** 1234"
                  filterClass=""
                  bgColor={index % 2 === 0 ? "from-[#4C49ED] to-[#0A06F4]" : "bg-white"}
                  textColor={index%2 == 0 ? "text-white": "text-black"}
                  iconBgColor="bg-opacity-10"
                  showIcon={true}
                ></CreditCard>
              ))}   
            </div>
          </div>
  
  <div className="flex flex-col justify-between w-1/2 flex-grow-0">
    <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-2xl">Recent Transaction</h1>
    <RecentTransaction />
  </div>
</div>

        <div className="flex space-x-6">
            <div className=" w-1/2">
              <h1 className="flex mx-4 my-4 font-bold  text-[#343C6A] text-2xl"> Weekly Activity</h1>
              <WeeklyActivity />
            </div>
            <div className=" w-1/3">
              <h1 className="flex mx-4 my-4 font-bold  text-[#343C6A] text-2xl"> Expense Statistics</h1>
              <ExpenseStatistics  />
            </div>
        </div>
        <div className="flex justify-between space-x-6 w-full h-24">
          <div className=" w-1/3 ">
          <h1 className="flex mx-4 my-4 font-bold  text-[#343C6A] text-2xl">Quick Transfers</h1>
            <Reviving />
          </div>
          <div className="w-2/3 h-5" >
          <h1 className="flex mx-4 my-4 font-bold  text-[#343C6A] text-2xl">Balance History</h1>
          <BalanceHistory />
          </div>
        </div>
      </div>
    </div>
    </div>

  
  );

}
const formatDate = (dateString: string): string => {
  const date = new Date(dateString);

  const options: Intl.DateTimeFormatOptions = {
    year: "numeric",
    month: "short",
    day: "numeric",
  };

  return date.toLocaleDateString("en-US", options);
};