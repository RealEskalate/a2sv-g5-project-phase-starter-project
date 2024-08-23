"use client";
import { useUser } from "@/contexts/UserContext";
import Link from "next/link";
import React, { useEffect, useState } from "react";
import CreditCard from "../../_components/Credit_Card";
import { CardDetails, TransactionContent } from "@/types";
import { getCreditCards, getExpenses } from "@/lib/api";
import { ExpenseChart } from "./ExpenseChart";

// Shimmer component
const Shimmer = () => (
  <div className="animate-pulse flex space-x-4">
    <div className="rounded-lg bg-gray-300 h-24 w-full"></div>
  </div>
);

export const TransactionCards = ({
  onLoadingComplete,
}: {
  onLoadingComplete: any;
}) => {
  const { isDarkMode } = useUser();
  const [creditCards, setCreditCards] = useState<CardDetails[]>([]);
  const [expenses, setExpenses] = useState<TransactionContent[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const [cards, initialExpenses] = await Promise.all([
          getCreditCards(0, 2),
          getExpenses(0, 6),
        ]);
        setCreditCards(cards?.content || []);
        setExpenses(initialExpenses?.content || []);
        setLoading(false);
        onLoadingComplete(false);
      } catch (error) {
        console.error("Error fetching data:", error);
      } finally {
        
      }
    };
    fetchData();
  }, [onLoadingComplete]);

  return (
    <div className="md:flex sm:grid-cols-2 md:gap-5 space-y-5 md:space-y-0">
      <div className="md:w-2/3 space-y-5">
        <div className="flex justify-between font-inter text-[16px] font-semibold">
          <h4>My Cards</h4>
          <h4>
            <Link href="/dashboard/credit-cards/#add-card">+Add Card</Link>
          </h4>
        </div>
        <div className="flex space-x-5 overflow-x-auto [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]">
          {loading
            ? Array(2).map((_, i) => <Shimmer key={i} />)
            : creditCards.map((card) => (
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
      <div className="md:w-1/3 md:space-y-5 w-full">
        <div className="font-inter text-[16px] font-semibold">
          <h4>My Expense</h4>
        </div>
        <div
          className={`rounded-xl pt-1 ${
            isDarkMode ? "bg-gray-800" : "bg-white s"
          }hadow-lg`}
        >
          {loading ? (
            <div className="p-2">
              <Shimmer />
            </div>
          ) : (
            <ExpenseChart expenses={expenses} />
          )}
        </div>
      </div>
    </div>
  );
};
