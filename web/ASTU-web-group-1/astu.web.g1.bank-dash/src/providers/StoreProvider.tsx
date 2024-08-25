'use client';
import { store } from '@/store/store';
import React from 'react';
import { Provider } from 'react-redux';

export default function StoreProvider({ children }: any) {
  return <Provider store={store}>{children}</Provider>;
}
