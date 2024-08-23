"use client";
import { Donut } from "@/components/ui/Piechart";
import { CardDetails } from "@/types";
import CreditCard from "../_components/Credit_Card";
import Cardinfo from "./components/Cardinfo";
import CardSetting from "./components/CardSetting";
import InputForm from "./components/InputForm";
import { useUser } from "@/contexts/UserContext";
import { useEffect, useRef, useState } from "react";
import { useRouter } from "next/navigation";
import { BalanceData, TransactionContent } from "@/types";
import getRandomBalance, {
  getallTransactions,
  getCreditCards,
} from "@/lib/api";
import { Loading } from "../_components/Loading";

const CreditCards = () => {
  const { isDarkMode } = useUser();
  const formSectionRef = useRef<HTMLDivElement>(null);

  // Check for hash in URL and scroll to the target section
  if (window.location.hash === "#add-card") {
    formSectionRef.current?.scrollIntoView({ behavior: "smooth" });
  }

  const [recentTransactions, setRecentTransactions] = useState<
    TransactionContent[]
  >([]);
  const [creditCards, setCreditCards] = useState<CardDetails[]>([]);
  const [balance, setBalanceHistory] = useState<BalanceData[]>([]);
  const [loading, setLoading] = useState(true);
  let totalCreditcardpage;

  useEffect(() => {
    const fetchData = async () => {
      try {
        const res = await getCreditCards(0, 50);
        const recent = await getallTransactions(0, 3);
        const balance = await getRandomBalance();

        setCreditCards(res?.content || []);
        totalCreditcardpage = res?.totalPages;
        setBalanceHistory(balance || []);
        setRecentTransactions(recent?.content || []);
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
      className={`p-3 ${
        isDarkMode ? "bg-gray-700 text-gray-200" : " text-black"
      }`}
    >
      <div className="p-3">
        <h1
          className={`text-2xl ${isDarkMode ? "text-white" : "text-[#333B69]"}`}
        >
          My Cards
        </h1>
        <div className="flex flex-row max-y-[200px] overflow-y-auto gap-6 sm:max-x-[500px] md:max-x-[600px] ">
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
      <div className="p-3">
        <div className="md:grid md:grid-cols-2">
          <div className="max-w-screen-sm">
            <h1
              className={`text-2xl ${
                isDarkMode ? "text-white" : "text-[#333B69]"
              }`}
            >
              Card Expense Statistics
            </h1>
            <div
              style={{ borderRadius: "5px", overflow: "hidden" }}
              className="rounded-full bg-white"
            >
              <Donut />
            </div>
          </div>
          <div className="p-3">
            <h1
              className={`text-2xl ${
                isDarkMode ? "text-white" : "text-[#333B69]"
              }`}
            >
              Card List
            </h1>
            <Cardinfo />
            <Cardinfo />
            <Cardinfo />
          </div>
        </div>
      </div>
      <h1
        className={`m-3 text-2xl ${
          isDarkMode ? "text-white" : "text-[#333B69]"
        }`}
      >
        Add New Card
      </h1>
      <div className="md:grid md:grid-cols-[2fr,1fr] gap-4">
        <div
          className={`rounded-xl py-2 max-w-fit ${
            isDarkMode ? "bg-gray-800" : "bg-white"
          }`}
        >
          <p
            className={`m-3 max-h-[440px] ${
              isDarkMode ? "text-gray-400" : "text-[#718EBF]"
            }`}
          >
            Credit Card generally means a plastic card issued by Scheduled
            Commercial Banks assigned to a Cardholder, with a credit limit, that
            can be used to purchase goods and services on credit or obtain cash
            advances.
          </p>
          <div className="m-3">
            <InputForm />
          </div>
        </div>
        <div>
          <h1
            className={`mb-2 mt-3 text-2xl ${
              isDarkMode ? "text-white" : "text-[#333B69]"
            }`}
          >
            Card Setting
          </h1>
          <div className="max-h-[440px]">
            <CardSetting
              image="/images/BlockCard.png"
              title="Block Card"
              description="Instantly block your card"
              color="bg-yellow-200"
            />
            <CardSetting
              image="/images/Lock.png"
              title="Change Pic Code"
              description="Withdraw without any card"
              color="bg-violet-200"
            />
            <CardSetting
              image="/images/Google.png"
              title="Add to Google Pay"
              description="Withdraw without any card"
              color="bg-pink-200"
            />
            <CardSetting
              image="/images/Apple.png"
              title="Add to Apple Pay"
              description="Withdraw without any card"
              color="bg-green-200"
            />
            <CardSetting
              image="/images/Apple.png"
              title="Added to Apple store"
              description="Withdraw without any card"
              color="bg-green-200"
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default CreditCards;
