"use client";
import { useGetAllCardsQuery } from "@/lib/redux/slices/cardSlice";
import React from "react";
import MyCard from "./MyCard";
import { CardContentType } from "@/types/card.types";
import CardListCardSkeleton from "../AllSkeletons/CardListSkeleton/CardListSkeleton";

export default function MyCardLists() {
  const { data, isLoading } = useGetAllCardsQuery({ page: 0, size: 5 });
  // console.log(data?.content);
  if (isLoading) {
    return <CardListCardSkeleton />;
  }
  return (
    <>
      {data?.content.map((card: CardContentType, index: number) => (
        <MyCard key={card.id} content={card} index={index} />
      ))}
    </>
  );
}
