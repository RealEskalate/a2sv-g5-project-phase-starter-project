// components/ServicesCards/ServicesCardApp.tsx
import React from 'react';
import ServicesCard from './servicesCard';

const ServicesCardApp = () => {
  return (
    <div className="overflow-x-auto whitespace-nowrap p-4" style={{ 
        scrollbarWidth: 'none', 
        msOverflowStyle: 'none', 
      }}>
      <div className="inline-flex gap-6 min-w-max justify-center">
        <ServicesCard
          image='/assets/icons/life-insurance.svg'
          title='Life Insurance'
          description='Unlimited Protection'
          color='white'
        />
        <ServicesCard
          image='/assets/icons/shopping.svg'
          title='Shopping'
          description='Buy. Think. Grow'
          color='white'
        />
        <ServicesCard
          image='/assets/icons/safety.svg'
          title='Safety'
          description='We are your allies'
          color='white'
        />
      </div>
    </div>
  );
}

export default ServicesCardApp;
