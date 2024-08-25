import React from 'react';
import Image from 'next/image';
import { CardProps } from '@/types/index.';

const CreditCard: React.FC<CardProps> = ({
  balance,
  cardHolder,
  validThru,
  cardNumber,
  gradientFrom,
  gradientTo,
  chipImage,
  borderStyle = '',
  bottomBackground,
  textColor = 'text-white'
}) => {
  const inlineGradientStyle = {
    backgroundImage: `linear-gradient(to right, ${gradientFrom}, ${gradientTo})`,
  };

  return (
    <div
      className={`flex flex-col w-[265px] h-[170px] md:w-[350px] md:h-[235px] rounded-3xl ${borderStyle}`}
    >
      <div
        style={inlineGradientStyle}
        className="flex flex-col justify-between gap-4 md:gap-6 rounded-t-3xl"
      >
        <div className={`flex justify-between mt-4 md:mt-6 mx-4 md:mx-6 ${textColor}`}>
          <div>
            <p className="text-xs md:text-sm">Balance</p>
            <p className="text-xl md:text-2xl font-bold">{balance}</p>
          </div>
          <Image src={chipImage} alt="Chip Card" width={30} height={30} />
        </div>
        <div className={`flex items-start space-x-4 md:space-x-8 mx-4 md:mx-6 mb-2 md:mb-4 ${textColor}`}>
          <div>
            <p className="uppercase tracking-wider text-xs md:text-sm">Card Holder</p>
            <p className="uppercase tracking-wider font-semibold text-xs md:text-sm">{cardHolder}</p>
          </div>
          <div>
            <p className="uppercase tracking-wider text-xs md:text-sm">Valid Thru</p>
            <p className="uppercase tracking-wider font-semibold text-xs md:text-sm">{validThru}</p>
          </div>
        </div>
      </div>
      <div
        className={`flex justify-between items-center align-bottom p-4 md:p-6 rounded-b-3xl ${bottomBackground}`}
      >
        <div className={`text-lg md:text-xl tracking-widest ${textColor}`}>{cardNumber}</div>
        <Image src="/Group17.png" alt="MasterCard" width={25} height={25} />
      </div>
    </div>
  );
};

export default CreditCard;