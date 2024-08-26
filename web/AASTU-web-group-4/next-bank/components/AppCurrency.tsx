import React from 'react';
import { CurrencyProvider } from '@/context/CurrencyContext';
import BalanceCard from '@/components/AccountSmallCard';

const App: React.FC = () => {
  return (
    <CurrencyProvider>
      <BalanceCard />
    </CurrencyProvider>
  );
};

export default App;
