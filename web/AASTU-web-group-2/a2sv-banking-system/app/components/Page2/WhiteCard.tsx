import React from 'react';
import Card from './Card'; 

const WhiteCard = () => {
  return (
    <Card
      balance="$5,756"
      cardHolder="Eddy Cusuma"
      validThru="12/22"
      cardNumber="3778 **** **** 1234"
      filterClass = "filter-black"
      bgColor="from-white to-gray-200"
      textColor="text-black"
      iconBgColor="bg-opacity-10"
      showIcon={true}
    />
  );
};

export default WhiteCard;
