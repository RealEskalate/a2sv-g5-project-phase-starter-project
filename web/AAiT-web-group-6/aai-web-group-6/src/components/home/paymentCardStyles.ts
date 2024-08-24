export const cardBackground = (isWhite: boolean) =>
  isWhite ? "bg-white border" : "bg-gradient-to-r from-[#4c49ed] to-[#0a06f4]";

export const cardTextColor = (isWhite: boolean) =>
  isWhite ? "text-black" : "text-white";

export const cardLightTextColor = (isWhite: boolean) =>
  isWhite ? "text-slate-500" : "text-slate-400";

export const cardBottomBackground = (isWhite: boolean) =>
  isWhite ? "bg-white" : "bg-gradient-to-b from-[#4c49ed] to-[#0a06f4]";

export const chipImage = (isWhite: boolean) =>
  isWhite ? "/images/chip_card_white.png" : "/images/chip_card.png";

export const logoImage = (isWhite: boolean) =>
  isWhite ? "/images/rounded_white.png" : "/images/rounded.png";
