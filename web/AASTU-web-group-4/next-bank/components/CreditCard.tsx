import React from 'react'
import Image from 'next/image'
import { colors } from '@/constants';
import { textColors } from '@/constants';

interface ResponsiveCreditCardProps {
  backgroundColor: string;
}

const ResponsiveCreditCard: React.FC<ResponsiveCreditCardProps> = ({ backgroundColor }) => {
  return (
    <div className={`${backgroundColor} md:w-[231px] md:h-[170px] w-[265px] h-[170px]  rounded-xl relative ${backgroundColor === colors.white ? 'border-[1px] border-gray-300' : ''}`}>
      <div className="flex justify-between w-[95%]">
        <div className='mt-1 ml-3 p-2'>
            <span className={`text-[11px]  ${backgroundColor == colors.blue ? textColors.textWhite : textColors.textDimmed}`}>Balance</span>
            <span className={`block text-[16px]  font-bold ${backgroundColor == colors.blue ? 'text-white' : 'text-black'}`}>$5,756</span>
        </div>
        <Image src={backgroundColor === colors.blue ? '/icons/chip.png' : '/icons/blackChip.png'} width={30} height={29} alt="chip card" className='h-[29px] mt-4 mr-2' />
      </div>

      <div className="flex justify-between w-[90%]">
        <div className='ml-3 pl-1.5 '>
            <span className={`text-[10px]  ${backgroundColor == colors.blue ? textColors.textWhite : textColors.textDimmed}`}>CARD HOLDER</span>
            <span className={`block text-[13px]  font-bold ${backgroundColor == colors.blue ? 'text-white' : 'text-black'}`}>Tekola Chane</span>
        </div>

        <div className='mr-3 '>
            <span className={`text-[10px]  ${backgroundColor == colors.blue ? textColors.textWhite : textColors.textDimmed}`}>VALID THRU</span>
            <span className={`block text-[13px]  font-bold ${backgroundColor == colors.blue ? 'text-white' : 'text-black'}`}>12/22</span>
        </div>
      </div>
      
      <div className={`flex justify-between child-div absolute bottom-0 left-0 right-0 ${backgroundColor == colors.blue ? 'bg-gradient-to-b from-blue-600' : 'border-t-[1px]'}`}>
        <span className={`p-3 ml-2 text-[15px] md:text-[15px]  ${backgroundColor == colors.blue ? 'text-white' : 'text-black'}`}>3778 **** **** 1234</span>
        <Image src={'/icons/masterCard.png'} width={35} height={33} alt="card icon" className='mt-0.5 mr-3' />
      </div>

    </div>
  )
}

export default ResponsiveCreditCard
