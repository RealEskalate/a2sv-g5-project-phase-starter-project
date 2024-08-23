import React from "react";
import InfoboxCard from "./InfoboxCard";
import { infoboxListItemsService } from "./infoboxListItemsService";

const InfoboxForServicePage = () => {
  return (
    <div className="grid grid-cols-1 sm:grid-cols-3 md:grid-cols gap-4 p-4 w-auto">
      {infoboxListItemsService.map((item, index) => (
        <InfoboxCard
          key={index}
          name={item.name}
          icon={item.icon}
          value={item.value}
          classNameText1="text-[#232323] font-semibold text-lg max-lg:text-[12px] max-md:text-lg"
          classNameText2="text-[#718EBF] font-normal text-sm max-lg:text-[12px] max-md:text-sm"
        />
      ))}
    </div>
  );
};

export default InfoboxForServicePage;
