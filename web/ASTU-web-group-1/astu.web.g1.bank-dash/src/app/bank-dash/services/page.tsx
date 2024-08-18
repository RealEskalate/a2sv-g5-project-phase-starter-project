import BankServicesList from '@/components/BankServicesList/BankServicesList';
import ServicesCardApp from '@/components/ServicesCards/servicesCardApp';
import React from 'react';

export default function page() {
  return (
    <div>
      <ServicesCardApp />
      <BankServicesList/>
    </div>
  );
}
