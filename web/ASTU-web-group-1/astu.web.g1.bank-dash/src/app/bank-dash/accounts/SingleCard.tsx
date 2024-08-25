"use client";
import CardSkeleton from "@/components/AllSkeletons/CardSkeleton/CardSkeleton";
import MyCard from "@/components/MyCard/MyCard";
import { useGetAllCardsQuery } from "@/lib/redux/slices/cardSlice";
import { CardContentType } from "@/types/card.types";
import { Plus } from "lucide-react";
import Link from "next/link";
import React from "react";

export default function SingleCard() {
  const { data, isLoading } = useGetAllCardsQuery({ page: 0, size: 5 });
  const card = data?.content[0];
  return (
    <>
      {isLoading ? (
        <>
          <CardSkeleton />
        </>
      ) : card ? (
        <MyCard key={card.id} content={card} index={0} />
      ) : (
        <div className="w-[295px] h-[175px] bg-gray-200 rounded-3xl justify-center items-center flex flex-shrink-0">
          <Link href="/bank-dash/credit-card">
            <Plus size={32} />
          </Link>
        </div>
      )}
    </>
  );
}
