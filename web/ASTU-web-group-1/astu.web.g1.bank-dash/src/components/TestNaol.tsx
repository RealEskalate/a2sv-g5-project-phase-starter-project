'use client';
import { useGetAllCardsQuery } from '@/lib/redux/slices/cardSlice';
import { useGetAllTransactionsQuery } from '@/lib/redux/slices/transactionSlice';
import React from 'react';

export default function TestNaol() {
  const { data, isLoading, isError } = useGetAllTransactionsQuery('1');
  console.log(data);
  return <div>TestNaol</div>;
}
