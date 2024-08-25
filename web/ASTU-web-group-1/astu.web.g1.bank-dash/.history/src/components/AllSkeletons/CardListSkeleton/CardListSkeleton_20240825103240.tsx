import React from "react";
import CardlistcardSkeleton from "./CardlistcardSkeleton";

const CardListSkeleton = () => {
  return (
    <div className="flex flex-col  w-full  ">
      <p className="text-[#333B69] pb-2 font-semibold">Card List</p>
      <div className="h-80 lg:h-[16.5rem] xl:h-80 overflow-y-scroll whitespace-nowrap scroll-smooth">
        {data?.content?.map((card, index) => (
          <CardlistcardSkeleton key={card.id} />
        ))}
      </div>
    </div>
  );
};

export default CardListSkeleton;
