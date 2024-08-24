"use client";
import { useUser } from "@/contexts/UserContext";
import Link from "next/link";
import React, { useEffect, useState } from "react";
import CreditCard from "../../_components/Credit_Card";
import { CardDetails, TransactionContent } from "@/types";
import { getCreditCards, getExpenses } from "@/lib/api";
import { ExpenseChart } from "./ExpenseChart";
import { CreditCardShimmer, RecentTransactionShimmer } from "../../_components/Shimmer";
import { AddCardModal } from "./AddCardModal";




export const TransactionCards = ({
  onLoadingComplete,
}: {
  onLoadingComplete: any;
}) => {
  const { isDarkMode } = useUser();
  const [creditCards, setCreditCards] = useState<CardDetails[]>([]);
  const [expenses, setExpenses] = useState<TransactionContent[]>([]);
  const [loading, setLoading] = useState(true);
  
  const [isModalOpen, setIsModalOpen] = useState(false);
  useEffect(() => {
    const fetchData = async () => {
      try {
        const [cards, initialExpenses] = await Promise.all([
          getCreditCards(0, 2),
          getExpenses(0, 6),
        ]);
        setCreditCards(cards?.content || []);
        setExpenses(initialExpenses?.content || []);
     
        onLoadingComplete(false);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
      finally{   setLoading(false);}
    };
    fetchData();
  }, [onLoadingComplete]);

  const handleModalToggle = () => {
    setIsModalOpen(!isModalOpen);
  };
  return (
    <div className="md:flex sm:grid-cols-2 md:gap-5 space-y-5 md:space-y-0">
      <div className="md:w-2/3 space-y-5">
        <div className="flex justify-between font-inter text-[16px] font-semibold">
          <h4>My Cards</h4>
          <h4>
            <button onClick={handleModalToggle}>+Add Card</button>
          </h4>
          {isModalOpen && (
            <div onClick={handleModalToggle}>
              <div
                className="relative bg-white p-6 rounded-lg shadow-lg max-w-md w-full"
                onClick={(e) => e.stopPropagation()}
              >
                <AddCardModal
                  isOpen={isModalOpen}
                  onClose={handleModalToggle}
                 
                />
              </div>
            </div>
          )}
        </div>
        <div className="flex space-x-5 overflow-x-auto [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]">
          {loading || creditCards.length === 0
            ? [1, 2].map((index) => <CreditCardShimmer key={index} />)
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
          className={`rounded-xl space-y-5 p-3 md:h-[200px] lg:w-[365px] lg:h-[220px] ${
            isDarkMode ? "bg-gray-800" : "bg-white "
          } shadow-lg`}
        >
          {loading || expenses.length === 0 ? (
            [1, 2, 3].map((index) => <RecentTransactionShimmer key={index} />)
          ) : (
            <ExpenseChart expenses={expenses} />
          )}
        </div>
      </div>
    </div>
  );
};
