"use client";
import { useGetAllCardsQuery } from "@/lib/redux/slices/cardSlice";
import React from "react";
import MyCard from "./MyCard";
import { CardContentType } from "@/types/card.types";
import CardSkeleton from "../AllSkeletons/CardSkeleton/CardSkeleton";

export default function MyCardLists() {
  const { data, isLoading } = useGetAllCardsQuery({ page: 0, size: 5 });
  // console.log(data?.content);
  if (isLoading) {
    return <CardSkeleton />;
  }
  return (
    <>
      {data?.content.map((card: CardContentType, index: number) => (
        <MyCard key={card.id} content={card} index={index} />
      ))}
    </>
  );
}
