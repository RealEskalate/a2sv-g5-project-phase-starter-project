import React from "react";
import CardSetting from "./CardSetting";

const CardList = () => {
  return (
    <div className="bg-white rounded-2xl p-3 shadow h-full dark:bg-[#050914] dark:border dark:border-[#333B69]">
      <CardSetting
        data={[["Block Card", "Instantly block your card"]]}
        icon={<img src="cardsetting1.svg" />}
      />
      <CardSetting
        data={[["Change Pic Code", "Withdraw without any card"]]}
        icon={<img src="cardsetting2.svg" />}
      />
      <CardSetting
        data={[["Add to Google Pay", "Withdraw without any card"]]}
        icon={<img src="cardsetting3.svg" />}
      />
      <CardSetting
        data={[["Add to Apple Pay", "Withdraw without any card"]]}
        icon={<img src="cardsetting4.svg" />}
      />
      <CardSetting
        data={[["Add to Apple Store", "Withdraw without any card"]]}
        icon={<img src="cardsetting4.svg" />}
      />
    </div>
  );
};

export default CardList;
