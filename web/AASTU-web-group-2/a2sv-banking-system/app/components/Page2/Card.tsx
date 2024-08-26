import React from 'react';

interface CardProps {
  balance?: string;
  cardHolder?: string;
  validThru?: string;
  cardNumber?: string;
  filterClass?: string;
  bgColor?: string;
  textColor?: string;
  iconBgColor?: string;
  showIcon?: boolean;
  loading?: boolean; 
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
  showIcon = true,
  loading = false, 
}) => {

  const isBlueGradient = bgColor.includes('#4C49ED') || bgColor.includes('#0A06F4');
  const ellipseImageSrc = isBlueGradient ? '/group17.svg' : '/group18.svg';
  const iconSrc = isBlueGradient ? '/sim.svg' : '/blackSim.svg';

  const cardHolderTextColor = isBlueGradient ? 'text-[rgba(255,255,255,0.7)]' : 'text-[#718EBF]';

  if (loading) {
    return (
      <div className="border dark:border-[#333B69] rounded-3xl my-4 mx-2 animate-pulse">
                <div className="relative w-full bg-gradient-to-r from-gray-200 dark:from-[#333B69] to-gray-300 dark:to-[#555B85] text-transparent rounded-3xl shadow-md min-w-[350px]">
                  <div className="flex justify-between items-start px-6 pt-6">
                    <div>
                      <p className="text-xs font-semibold bg-gray-300 dark:bg-[#555B85] rounded w-16 h-4 mb-2"></p>
                      <p className="text-xl font-medium bg-gray-300 dark:bg-[#555B85] rounded w-24 h-6"></p>
                    </div>
                    <div className="w-8 h-8 bg-gray-300 dark:bg-[#555B85] rounded-full"></div>
                  </div>

                  <div className="flex justify-between gap-12 mt-4 px-6">
                    <div>
                      <p className="text-xs font-medium bg-gray-300 dark:bg-[#555B85] rounded w-16 h-4 mb-2"></p>
                      <p className="font-medium text-base bg-gray-300 dark:bg-[#555B85] rounded w-24 h-6"></p>
                    </div>
                    <div className="pr-8">
                      <p className="text-xs font-medium bg-gray-300 dark:bg-[#555B85] rounded w-16 h-4 mb-2"></p>
                      <p className="font-medium text-base md:text-lg bg-gray-300 dark:bg-[#555B85] rounded w-24 h-6"></p>
                    </div>
                  </div>

                  <div className="relative mt-8 flex justify-between py-4 items-center">
                    <div className="absolute inset-0 w-full h-full bg-gradient-to-b from-white/20 dark:from-gray-700/20 to-transparent z-0"></div>
                    <div className="ml-4 relative z-10 text-base font-medium px-6 bg-gray-300 dark:bg-[#555B85] rounded w-40 h-6"></div>
                    <div className="flex justify-end relative z-10 px-6">
                      <div className="w-10 h-10 bg-gray-300 dark:bg-[#555B85] rounded-full "></div>
                    </div>
                  </div>
                </div>
              </div>
    );
  }

  return (
    <div className='border rounded-3xl my-4 mx-2 '>
      {/* Set position relative here to contain the absolute positioned element */}
      {/* <div className={`relative w-full bg-gradient-to-r ${bgColor} ${textColor} rounded-3xl shadow-md font-lato pt-6  min-w-[350px]`}>*/}
      <div className={`relative w-full bg-gradient-to-r ${bgColor} ${textColor} rounded-3xl shadow-md font-lato pt-6 min-w-[300px] md:min-w-[200px] lg:min-w-[350px]`}>
        <div className="flex justify-between items-start px-6">
          <div>
            <p className="text-xs font-semibold text-[#7b99cd]">Balance</p>
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
          {/* The absolute element stays within the bounds of the relative parent */}
          <div className="absolute inset-0 w-full h-full bg-gradient-to-r from-white/30 to-transparent dark:from-gray-900/30 z-0 "></div>
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


