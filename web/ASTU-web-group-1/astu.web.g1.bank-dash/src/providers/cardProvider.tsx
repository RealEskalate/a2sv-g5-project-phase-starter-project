'use client';
import { store } from '@/store/store';
import React from 'react';
import { Provider } from 'react-redux';

export default function CardProvider({ children }: any) {
  return <Provider store={store}>{children}</Provider>;
}
