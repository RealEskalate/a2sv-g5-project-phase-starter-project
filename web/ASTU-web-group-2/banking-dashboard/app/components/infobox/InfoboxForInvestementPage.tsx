import React from "react";
import { infoboxListItemsInvestement } from "./infoboxListItemsInvestement"; 
import InfoboxCard from "./InfoboxCard";

const InfoboxForInvestementPage = () => {
  return (
    <div className="grid grid-cols-1 sm:grid-cols-3 md:grid-cols gap-4 p-4 w-auto">
      {infoboxListItemsInvestement.map((item, index) => (
        <InfoboxCard
          key={index}
          name={item.name}
          icon={item.icon}
          value={item.value}
        />
      ))}
    </div>
  );
};

export default InfoboxForInvestementPage;
