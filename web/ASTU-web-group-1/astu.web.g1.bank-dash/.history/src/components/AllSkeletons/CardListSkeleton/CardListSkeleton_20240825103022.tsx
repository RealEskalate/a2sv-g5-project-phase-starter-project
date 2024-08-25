import React from "react";

const CardListSkeleton = () => {
  return (
    <div className="flex flex-col  w-full  ">
      <p className="text-[#333B69] pb-2 font-semibold">Card List</p>
      <div className="h-80 lg:h-[16.5rem] xl:h-80 overflow-y-scroll whitespace-nowrap scroll-smooth">
        {data?.content?.map((card, index) => (
          <CardListCard
            key={card.id}
            cardType={card.cardType}
            bank={bankList[index % bankList.length]}
            cardNumber={formatCardNumber(card.semiCardNumber)}
            imageUrl={imageList[index % imageList.length]}
            namainCard={card.cardHolder}
          />
        ))}
      </div>
    </div>
  );
};

export default CardListSkeleton;
