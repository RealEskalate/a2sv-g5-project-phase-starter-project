import React from "react";
import GoogleIcon from "../../../public/assets/icons/google-icon.svg";
import ChangePinIcon from "../../../public/assets/icons/change-pin-icon.svg";
import BlockIcon from "../../../public/assets/icons/block-card-icon.svg";
import AppleIcon from "../../../public/assets/icons/apple-icon.svg";
import CardSettingItem from "./CardSettingItem";

const data = [
  {
    Icon: BlockIcon,
    title: "Block Card",
    description: "Instantly block your card",
  },
  {
    Icon: ChangePinIcon,
    title: "Change Pic Code",
    description: "Withdraw without any card",
  },
  {
    Icon: GoogleIcon,
    title: "Add to Google Pay",
    description: "Withdraw without any card",
  },
  {
    Icon: AppleIcon,
    title: "Add to Apple Pay",
    description: "Withdraw without any card",
  },
  {
    Icon: AppleIcon,
    title: "Add to Apple Store",
    description: "Withdraw without any card",
  },
];

export default function CardSettings() {
  return (
    <>
      {/* <div className="w-full lg:w-1/3 ">
      <CardSettings />
    </div> */}
      <p className="font-semibold text-lg lg:text-20px text-deepNavy py-3">
        Card Setting
      </p>
      <div className="bg-white w-full px-5 py-6 rounded-3xl space-y-5">
        {data.map((ele) => (
          <CardSettingItem
            Icon={ele.Icon}
            title={ele.title}
            description={ele.description}
          />
        ))}
      </div>
    </>
  );
}
