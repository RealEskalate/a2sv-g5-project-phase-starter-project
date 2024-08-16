// components/ServicesCards/ServicesCardApp.tsx
import React from 'react';
import ServicesCard from './servicesCard';
import Insuranceicon from '../../../public/assets/icons/insurance-icon.svg';
import Shoppingicon from '../../../public/assets/icons/shopping-icon.svg';
import Safetyicon from '../../../public/assets/icons/safety-icon.svg';

const ServicesCardApp = () => {
  return (
    <div className="overflow-x-auto whitespace-nowrap p-4" style={{ 
        scrollbarWidth: 'none', 
        msOverflowStyle: 'none', 
      }}>
      <div className="inline-flex gap-6 min-w-full justify-between">
        <ServicesCard
          children = {<Insuranceicon/>}
          title='Life Insurance'
          description='Unlimited Protection'
          color='white'
        />
        <ServicesCard
          children = {<Shoppingicon/>}
          title='Shopping'
          description='Buy. Think. Grow'
          color='white'
        />
        <ServicesCard
          children = {<Safetyicon/>}
          title='Safety'
          description='We are your allies'
          color='white'
        />
      </div>
    </div>
  );
}

export default ServicesCardApp;
