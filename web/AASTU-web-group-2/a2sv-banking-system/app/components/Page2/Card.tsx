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

const CreditCard: React.FC<CardProps> = ({
  balance,
  cardHolder,
  validThru,
  cardNumber,
  filterClass = "filter-white",
  bgColor = '', 
  textColor = 'text-white',
  iconBgColor = 'bg-opacity-10',
  showIcon = true
}) => {

  const isBlueGradient = bgColor.includes('#4C49ED') || bgColor.includes('#0A06F4');
  const ellipseImageSrc = isBlueGradient ? '/group17.svg' : '/group18.svg';
  const iconSrc = isBlueGradient ? '/sim.svg' : '/blackSim.svg'; 

  const cardHolderTextColor = isBlueGradient ? 'text-[rgba(255,255,255,0.7)]' : 'text-[#718EBF]';

  return (
    <div className='border rounded-3xl my-4 mx-2'>
      <div className={`relative w-full bg-gradient-to-b ${bgColor} ${textColor} rounded-3xl shadow-md font-lato pt-6  h-[230px]  min-w-[300px]`}>
        <div className="flex justify-between items-start px-6">
          <div>
            <p className="text-xs font-semibold text-[#718EBF]">Balance</p>
            <p className="text-xl font-medium">{balance}</p>
          </div>
          {showIcon && (
            <div className={`w-8 h-8 ${iconBgColor} rounded-full flex items-center justify-center`}>
              <img src={iconSrc} alt="sim image" />
            </div>
          )}
        </div>

        <div className='flex justify-between gap-12 mt-4 px-6'>
          <div>
            <p className={`text-xs font-medium ${cardHolderTextColor}`}>CARD HOLDER</p>
            <p className="font-medium text-base">{cardHolder}</p>
          </div>
          <div className='pr-8'>
            <p className={`text-xs font-medium ${cardHolderTextColor}`}>VALID THRU</p>
            <p className="font-medium text-base md:text-lg">{validThru}</p>       
          </div>
        </div>

        <div className="relative mt-8 flex justify-between py-4 items-center">
        <div className="absolute inset-0 w-full h-full bg-gradient-to-b from-white/30 to-transparent z-0"></div>
          <div className="relative z-10 text-base font-medium px-6">
            {cardNumber}
          </div>
          <div className='flex justify-end relative z-10 px-6'>
            <img src={ellipseImageSrc} alt="ellipse background" />
          </div>
        </div>
      </div>
    </div>
  );
};

export default CreditCard;
