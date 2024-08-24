'use client'
import React, { useRef } from 'react';
import ResponsiveCreditCard from '@/components/CreditCard';
import { colors } from '@/constants';
import Component from '@/components/DoughnutChart';
import AddNewCard from '@/components/AddNewCard';
import CardSetting from '@/components/CardSetting';
import CardList from '@/components/CardList';
import { FaArrowLeft, FaArrowRight } from 'react-icons/fa';

const CreditCard = () => {
  const cards = [colors.blue, colors.white, colors.blue, colors.white];
  const scrollContainerRef = useRef(null);

  const handleScroll = (direction: string) => {
    if (scrollContainerRef.current) {
      scrollContainerRef.current.scrollBy({
        left: direction === 'left' ? -300 : 300,
        behavior: 'smooth',
      });
    }
  };

  return (
    <div className='lg:ml-72 ml-5 overflow-x-hidden mx-auto'>
      <div className="relative myCards max-w-[97%] mt-4">
        <h1 className='text-[19px] mb-3 font-bold text-[#333B69] dark:text-blue-500'>My Cards</h1>
        <div className="flex items-center">
          <button
            onClick={() => handleScroll('left')}
            className="absolute left-0 z-10 flex items-center justify-center w-10 h-10 bg-green-300 rounded-full shadow-md hover:bg-green-500 focus:outline-none border-none"
            aria-label="Scroll left"
          >
            <FaArrowLeft />
          </button>
          <div
            ref={scrollContainerRef}
            className="flex overflow-x-auto space-x-4 md:pr-3 pr-1 scrollbar-none"
          >
            {cards.map((bg, index) => (
              <span key={index}>
                <ResponsiveCreditCard backgroundColor={bg} />
              </span>
            ))}
          </div>
          <button
            onClick={() => handleScroll('right')}
            className="absolute right-0 z-10 flex items-center justify-center w-10 h-10 bg-green-300 rounded-full shadow-md hover:bg-green-500 focus:outline-none"
            aria-label="Scroll right"
          >
            <FaArrowRight />
          </button>
        </div>
      </div>

      <div className="flex flex-col md:flex-row gap-16 mt-8">
        <div className="doughnutChart lg:w-[350px] md:w-[231px] w-[325px] my-6">
          <h1 className='text-[19px] mb-3 font-bold text-[#333B69] dark:text-blue-500'>Card Expense Statistics</h1>
          <Component />
        </div>
        <div className="cardlist lg:w-[730px] md:w-[487px] sm:w-[325px] my-6">
          <h1 className='text-[19px] mb-3 font-bold text-[#333B69] dark:text-blue-500'>Card List</h1>
          <CardList />
        </div>
      </div>

      <div className="flex flex-col md:flex-row w-[80%] mb-16">
        <div className='md:mb-2 mb-0 md:mr-10'>
          <AddNewCard />
        </div>

        <div>
          <h1 className='text-[19px] mb-3 font-bold text-[#333B69] dark:text-blue-500'>Card Setting</h1>
          <CardSetting />
        </div>
      </div>
    </div>
  );
}

export default CreditCard;


// import React from 'react'
// import ResponsiveCreditCard from '@/components/CreditCard'
// import { colors } from '@/constants'
// import Component from '@/components/DoughnutChart'
// import AddNewCard from '@/components/AddNewCard'
// import CardSetting from '@/components/CardSetting'
// import CardList from '@/components/CardList'

// const CreditCard = () => {
//   const cards:string[] = [colors.blue, colors.white, colors.blue, colors.white]
//   return (
//     <div className='lg:ml-72 ml-5 overflow-x-hidden mx-auto '>
//       <div className="myCards max-w-[97%] mt-4">
//         <h1 className='text-[19px] mb-3 font-bold text-[#333B69] dark:text-blue-500'>My Cards</h1>
//         <div className="flex overflow-x-auto space-x-4 md:pr-3 pr-1 scrollbar-none">
//           {cards.map((bg, index) => (
//             <span key={index} className='p-3'>
//               <ResponsiveCreditCard backgroundColor={bg} />
//             </span>
//           ))}
//         </div>
//       </div>
      
//       <div className="flex flex-col md:flex-row gap-16">
//         <div className="doughnutChart lg:w-[350px] md:w-[231] w-[325px] my-6">
//           <h1 className='text-[19px] mb-3 font-bold text-[#333B69] dark:text-blue-500'>Card Expense Statistics</h1>
//           <Component />
//         </div>
//         <div className="cardlist lg:w-[730px] md:w-[487px] sm-w-[325] my-6">
//           <h1 className='text-[19px] mb-3 font-bold text-[#333B69] dark:text-blue-500'>Card List</h1>
//           <CardList />
//         </div>
//       </div>

//       <div className="flex flex-col md:flex-row w-[80%] mb-16">
//         <div className='md:mb-2 mb-0 md:mr-10'>
//           <AddNewCard/>
//         </div>

//         <div>
//         <h1 className='text-[19px] mb-3 font-bold text-[#333B69] dark:text-blue-500'>Card Setting</h1>
//           <CardSetting />
//         </div>
//       </div>
//     </div>
//   )
// }

// export default CreditCard