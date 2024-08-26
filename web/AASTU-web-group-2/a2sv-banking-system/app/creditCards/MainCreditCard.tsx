import React from "react";

interface Props {
  color1: string;
  color2: string;
  balance: number;
  validThru: string;
  cardNumber: string;
  cardHolder: string;
}

const MainCreditCard = ({
  color1,
  color2,
  balance,
  validThru,
  cardNumber,
  cardHolder,
}: Props) => {
  return (
    <div style={{ backgroundColor: color1 }} className="p-4 rounded-lg">
      <div>
        <div>
          <p>Balance</p>
          <p>
            {"$"}
            {balance}
          </p>
        </div>
        <div>
          <div>
            <p>CARD HOLDER</p>
            <p>{cardHolder}</p>
          </div>
          <div>
            <p>VALID THRU</p>
            <p>{validThru}</p>
          </div>
        </div>
        <div style={{ backgroundColor: color2 }} className="p-2 rounded-lg">
          <p>{cardNumber}</p>
          <img src="two_circle.svg" alt="" />
        </div>
        <img src="Chip_Card.svg" alt="" />
      </div>
    </div>
  );
};

export default MainCreditCard;
