'use client';
import MyCard from '@/components/MyCard/MyCard';
import { useGetAllCardsQuery } from '@/lib/redux/slices/cardSlice';
import { CardContentType } from '@/types/card.types';
import React from 'react';

export default function SingleCard() {
  const { data, isLoading } = useGetAllCardsQuery({ page: 0, size: 5 });
  console.log(data?.content);
  const card = data?.content[0];
  return (
    <>
      {isLoading ? (
        <>loading card</>
      ) : card ? (
        <MyCard key={card.id} content={card} index={0} />
      ) : (
        <>no card</>
      )}
    </>
  );
}