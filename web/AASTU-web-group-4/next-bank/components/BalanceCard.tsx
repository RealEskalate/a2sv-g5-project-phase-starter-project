// import React from 'react';
// import { useCurrency } from '@/context/CurrencyContext';

// interface BalanceCardProps {
//   iconSrc: string;
//   altText: string;
//   title: string;
//   amount: number;
//   index: number;
// }

// export const BalanceCard: React.FC<BalanceCardProps> = ({ iconSrc, altText, title, amount, index }) => {
//   const { currency } = useCurrency();

//   // Utility to format amount based on currency
//   const formatAmount = (amount: number) => {
//     switch (currency) {
//       case 'EUR':
//         return `€${(amount * 0.85).toFixed(2)}`; // Example conversion rate
//       case 'GBP':
//         return `£${(amount * 0.75).toFixed(2)}`; // Example conversion rate
//       default:
//         return `$${amount.toFixed(2)}`;
//     }
//   };

//   return (
//     <div>
//       <img src={iconSrc} alt={altText} />
//       <h3>{title}</h3>
//       <p>{formatAmount(amount)}</p>
//     </div>
//   );
// };


import React, { useContext } from 'react';
import { CurrencyContext } from '../context/CurrencyContext';

interface BalanceCardProps {
  iconSrc: string;
  altText: string;
  title: string;
  amount: string;
  index: number;
}

export const BalanceCard: React.FC<{ balance: BalanceCardProps }> = ({ balance }) => {
  const { currency, exchangeRate } = useContext(CurrencyContext);
  
  const originalAmount = parseFloat(balance.amount.replace(/[$,]/g, ''));
  const convertedAmount = (originalAmount * exchangeRate).toFixed(2);

  return (
    <div className="balance-card">
      <img src={balance.iconSrc} alt={balance.altText} />
      <h3>{balance.title}</h3>
      <p>{currency} {convertedAmount}</p>
    </div>
  );
};
