'use client';
import { useGetAllCardsQuery } from '@/lib/redux/slices/cardSlice';
import React from 'react';
import MyCard from './MyCard';
import { CardContentType } from '@/types/card.types';

export default function MyCardLists() {
  const { data, isLoading } = useGetAllCardsQuery({ page: 0, size: 5 });
  console.log(data?.content);
  return (
    <>
      {data?.content.map((card: CardContentType, index: number) => (
        <MyCard key={card.id} content={card} index={index} />
      ))}
    </>
  );
}