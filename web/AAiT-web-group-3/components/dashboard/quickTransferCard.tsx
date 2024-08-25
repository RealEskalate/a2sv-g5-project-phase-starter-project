import * as React from "react";

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "@/components/ui/carousel";

const QuickTransferCard = () => {
  return (
    <Carousel
      opts={{
        align: "center",
      }}
      className="w-8/12"
    >
      <CarouselContent className="flex items-center justify-between">
        {Array.from({ length: 5 }).map((_, index) => (
          <CarouselItem
            key={index}
            className="flex flex-col items-center justify-between basis-24"
          >
            <Avatar className="w-14 h-14">
              <AvatarImage src="/images/profile.jpeg" />
              <AvatarFallback>CN</AvatarFallback>
            </Avatar>

            <span className="text-xs text-primary-color-800">Livia Bator</span>
            <span className="text-xs text-primary-color-200">CEO</span>
          </CarouselItem>
        ))}
      </CarouselContent>
      <CarouselPrevious />
      <CarouselNext />
    </Carousel>
  );
};

export default QuickTransferCard;
