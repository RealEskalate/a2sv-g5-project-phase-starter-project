import React from 'react';

interface CardProps {
  balance: string;
  cardHolder: string;
  validThru: string;
  cardNumber: string;
  filterClass?: string;
  bgColor?: string; 
  textColor?: string; 
  iconBgColor?: string;
  showIcon?: boolean; 
}

const Card: React.FC<CardProps> = ({
  balance,
  cardHolder,
  validThru,
  cardNumber,
  filterClass = "filter-white",
  bgColor = 'from-[#4C49ED] to-[#0A06F4]',
  textColor = 'text-white',
  iconBgColor = 'bg-opacity-10',
  showIcon = true
}) => {
  
  const isBlueGradient = bgColor.includes('#4C49ED') || bgColor.includes('#0A06F4');

  return (
    
    <div className={`w-full max-w-2xl bg-gradient-to-b ${bgColor} ${textColor} rounded-3xl shadow-md font-lato pt-6 px-6  md:pt-8 md:px-8`}>
      <div className="flex justify-between items-start">
        <div>
          <p className="text-xs md:text-sm font-light">Balance</p>
          <p className="text-xl md:text-2xl font-medium">{balance}</p>
        </div>
        {showIcon && (
          <div className={`w-8 h-8 md:w-10 md:h-10 ${iconBgColor} rounded-full flex items-center justify-center`}>
            <img src="/sim.svg" alt="sim image" className={filterClass} />
          </div>
        )}
      </div>

      <div className='flex justify-between gap-12 mt-4 md:mt-6'>
        <div>
          <p className="text-xs md:text-sm font-medium text-[#7290bf]">CARD HOLDER</p>
          <p className="font-medium text-base md:text-lg">{cardHolder}</p>
        </div>
        <div className='pr-8 md:pr-12'>
          <p className="text-xs md:text-sm font-medium text-[#7290bf]">VALID THRU</p>
          <p className="font-medium text-base md:text-lg">{validThru}</p>
        </div>
      </div>

      <div className="relative mt-4 md:mt-6 flex flex-row justify-between py-4 md:py-6">
      <div className="absolute top-0 left-0 w-full h-1 bg-gradient-to-b from-transparent via-gray to-transparent shadow-md"></div>
      <div className="text-base md:text-lg font-medium  z-10">
          {cardNumber}
        </div>
    
      <div
        className={`absolute w-6 h-6 md:w-8 md:h-8 rounded-full -right-2 md:-right-3 top-0 ${
          isBlueGradient ? 'bg-white' : 'bg-[#9199AF80]'
        }`}
      ></div>
      <div
        className={`absolute w-6 h-6 md:w-8 md:h-8 rounded-full right-0 md:right-1 top-0 ${
          isBlueGradient ? 'bg-white' : 'bg-[#9199AF80]'
        }`}
      ></div>
      </div>
    </div>
  
  );
};

export default Card;


