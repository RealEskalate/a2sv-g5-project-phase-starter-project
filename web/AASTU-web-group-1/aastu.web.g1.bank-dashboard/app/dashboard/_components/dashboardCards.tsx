import { CardDetails } from "@/types";
import Link from "next/link";
import React, { useEffect, useState } from "react";
import CreditCard from "./Credit_Card";
import { getCreditCards } from "@/lib/api";

// Shimmer component for skeleton loading effect
const Shimmer = () => {
  return (
    <div className="animate-pulse">
      <div className="bg-gray-200 h-40 w-64 rounded-lg"></div>
    </div>
  );
};

export const Cards = ({onLoadingComplete}: {onLoadingComplete: any}) => {
  const [creditCards, setCreditCards] = useState<CardDetails[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try{
        const res = await getCreditCards(0, 2);
        setCreditCards(res?.content || []);
           onLoadingComplete(false);
           setLoading(false);
      }finally{
   }
      
    };
    fetchData();
  }, [onLoadingComplete]);

  return (
    <div className="md:w-2/3 space-y-5">
      <div className="flex justify-between font-inter text-[16px] font-semibold">
        <h4>My Cards</h4>
        <h4>
          <Link href="/dashboard/credit-cards/">See All</Link>
        </h4>
      </div>
      <div className="flex space-x-5 overflow-x-auto [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]">
        {loading
          ? [1, 2].map((index) => <Shimmer key={index} />) // Render shimmer loaders while loading
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
  );
};
