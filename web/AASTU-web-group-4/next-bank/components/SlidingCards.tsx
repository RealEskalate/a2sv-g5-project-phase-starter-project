import React from 'react';
import DesktopCreditCard from '@/components/DesktopCreditCard';
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from '@/components/ui/carousel'; // Adjust the import path according to your setup
import Link from 'next/link';

const SlidingCards: React.FC = () => {
  const cards = [
    { id: 1, bgColor: 'bg-blue-700', textColor: 'text-white' },
    { id: 2, bgColor: 'bg-indigo-700', textColor: 'text-white' },
    { id: 3, bgColor: 'bg-green-700', textColor: 'text-white' },
  ];

  return (
    <div>
      {/* Desktop and Tablet View */}
      <div className="hidden lg:flex lg:justify-center lg:space-x-8 w-full max-w-[750px] mx-auto">
        {cards.slice(0, 2).map((card) => (
          <DesktopCreditCard
            key={card.id}
            bgColor={card.bgColor}
            textColor={card.textColor}
          />
        ))}
      </div>

      {/* Mobile View - Carousel */}
      <div className="lg:hidden pt-8 w-full max-w-[350px] h-[235px] mx-auto relative">
        <Carousel>
          <CarouselContent>
            {cards.map((card) => (
              <CarouselItem key={card.id} className="relative w-full h-full flex justify-center items-center">
                <DesktopCreditCard
                  bgColor={card.bgColor}
                  textColor={card.textColor}
                />
              </CarouselItem>
            ))}
          </CarouselContent>
          <div className="absolute top-1/2 left-0 transform -translate-y-1/2 flex items-center">
            <CarouselPrevious className="p-2 text-gray-500 hover:text-gray-700">
              <span className="text-2xl">&lt;</span>
            </CarouselPrevious>
          </div>
          <div className="absolute top-1/2 right-0 transform -translate-y-1/2 flex items-center">
            <CarouselNext className="p-2 text-gray-500 hover:text-gray-700">
              <span className="text-2xl">&gt;</span>
            </CarouselNext>
          </div>
        </Carousel>

        {/* Add Card Link */}
        <Link href="#" className="absolute -top-3 -right-4 bg-transparent hover:text-blue-700 font-bold">
          + Add Card
        </Link>
      </div>
    </div>
  );
};

export default SlidingCards;
