import React from 'react';
import { CurrencyProvider } from './context/CurrencyContext';
import BalanceCard from '@/components/AccountBarChart';
// import Sidebar from '@/components/Sidebar';
import LastTransactionCard from './components/LastTransactionCard';

const App: React.FC = () => {
  return (
    <CurrencyProvider>
      <div className="app-container">
        <LastTransactionCard />
        <BalanceCard />
      </div>
    </CurrencyProvider>
  );
};

export default App;
