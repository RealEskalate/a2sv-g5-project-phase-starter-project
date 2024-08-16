'use client';
import { useGetAllCardsQuery } from '@/lib/redux/slices/cardSlice';
import React from 'react';

export default function TestNaol() {
  const { data, isLoading, isError } = useGetAllCardsQuery();
  console.log(data);
  return <div>TestNaol</div>;
}
