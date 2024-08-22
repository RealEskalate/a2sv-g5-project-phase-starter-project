import React from "react";
import Image, { StaticImageData } from "next/image";

interface CreditCardColorProps {
  cardBgColor: string;
  bottomBgColor: string;
  imageCreditCard: StaticImageData;
  grayCircleColor: boolean;
}

interface CardProps {
  cardData: any;
  cardColor: CreditCardColorProps;
}

const Card: React.FC<CardProps> = ({ cardData, cardColor }) => {
  function formatExpiryDate(expiryDate: string) {
    const date = new Date(expiryDate);
    const day = String(date.getDate()).padStart(2, "0");
    const month = String(date.getMonth() + 1).padStart(2, "0");
    return `${day}/${month}`;
  }

  // Format card number
  const formatCardNumber = (semiCardNumber: string) => {
    const firstNum = semiCardNumber.slice(0, 4);
    const lastNum = semiCardNumber.slice(-4);

    return `${firstNum} **** **** ${lastNum}`;
  };

  const dateFormat = formatExpiryDate(cardData.expiryDate);
  const formattedCardNumber = formatCardNumber(cardData.semiCardNumber);

  return (
    <div className="w-full">
      <div className={cardColor.cardBgColor}>
        <div className="flex justify-between p-5">
          <div>
            <div className="text-sm opacity-70">Balance</div>
            <div className="text-lg">${cardData.balance}</div>
          </div>
          <div>
            <Image src={cardColor.imageCreditCard} alt="chip card" />
          </div>
        </div>
        <div className="flex gap-16 p-4">
          <div className="pl-2">
            <div className="text-sm opacity-70">CARD HOLDER</div>
            <div>{cardData.cardHolder}</div>
          </div>
          <div>
            <div className="text-sm opacity-70">VALID THRU</div>
            <div>{dateFormat}</div>
          </div>
        </div>
        <div className={cardColor.bottomBgColor}>
          <div className="text-xl">{formattedCardNumber}</div>
          <div className="flex">
            <div
              className={`w-8 h-8 rounded-full ${
                cardColor.grayCircleColor ? "bg-gray-400" : "bg-gray-100"
              } opacity-50`}
            ></div>
            <div
              className={`w-8 h-8 rounded-full ${
                cardColor.grayCircleColor ? "bg-gray-400" : "bg-gray-100"
              } -ml-4 opacity-50`}
            ></div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Card;
