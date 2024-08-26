// import React, { createContext, useState, useContext } from 'react';

// interface CurrencyContextProps {
//   currency: string;
//   setCurrency: (currency: string) => void;
// }

// export const CurrencyContext = createContext<CurrencyContextProps | undefined>(undefined);

// export const CurrencyProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
//   const [currency, setCurrency] = useState<string>('USD');

//   return (
//     <CurrencyContext.Provider value={{ currency, setCurrency }}>
//       {children}
//     </CurrencyContext.Provider>
//   );
// };

// export const useCurrency = () => {
//   const context = useContext(CurrencyContext);
//   if (!context) {
//     throw new Error('useCurrency must be used within a CurrencyProvider');
//   }
//   return context;
// };


import React, { createContext, useState, useEffect, ReactNode, FC } from 'react';

interface CurrencyContextType {
  currency: string;
  setCurrency: (currency: string) => void;
  exchangeRate: number;
}

export const CurrencyContext = createContext<CurrencyContextType>({
  currency: 'USD',
  setCurrency: () => {},
  exchangeRate: 1,
});

interface CurrencyProviderProps {
  children: ReactNode;
}

export const CurrencyProvider: FC<CurrencyProviderProps> = ({ children }) => {
  const [currency, setCurrency] = useState<string>('USD');
  const [exchangeRate, setExchangeRate] = useState<number>(1);

  useEffect(() => {
    const fetchExchangeRate = async () => {
      try {
        const response = await fetch(`https://api.exchangerate-api.com/v4/latest/${currency}`);
        const data = await response.json();
        setExchangeRate(data.rates[currency] || 1);
      } catch (error) {
        console.error('Failed to fetch exchange rate', error);
      }
    };

    fetchExchangeRate();
  }, [currency]);

  return (
    <CurrencyContext.Provider value={{ currency, setCurrency, exchangeRate }}>
      {children}
    </CurrencyContext.Provider>
  );
};
