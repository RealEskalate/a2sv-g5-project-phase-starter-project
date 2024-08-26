import Link from "next/link";
import { CardData } from "@/types/cardData";
import NotFound from "@/app/not-found";

export default function CardBox({
  cardType,
  bank,
  cardNumber,
  NamainCard,
  detailsLink,
  svgBgColor,
  svgColor,
}: CardData) {
 return (
  <div className="body flex md:w-auto w-auto h-auto p-2 border-[1px] rounded-[10px] m-2 bg-white dark:bg-darkComponent border-gray-300 dark:border-gray-700">
    <div className="left">
      <svg
        width="45"
        height="45"
        viewBox="0 0 45 45"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <rect width="45" height="45" rx="12" fill={svgBgColor} />
        <g clipPath="url(#clip0_196_454)">
          <path
            d="M30.6621 18.5101V18.2365C30.6621 16.9788 29.6389 15.9556 28.3812 15.9556H15.2809C14.0232 15.9556 13 16.9788 13 18.2365V18.5101H30.6621Z"
            fill={svgColor}
          />
          <path
            d="M23.4347 25.8473C23.4347 24.6642 23.8152 23.5396 24.5178 22.6128H13V25.7945C13 27.0522 14.0232 28.0754 15.2809 28.0754H23.9162C23.601 27.385 23.4347 26.629 23.4347 25.8473ZM21.8311 25.156H19.8625V23.9841H21.8311V25.156ZM15.6273 23.9841H18.6906V25.156H15.6273V23.9841Z"
            fill={svgColor}
          />
          <path
            d="M25.7344 21.441C26.6281 20.8157 27.6898 20.4788 28.8033 20.4788C29.4478 20.4788 30.0748 20.592 30.6621 20.8083V19.6821H13V21.441H25.7344Z"
            fill={svgColor}
          />
          <path
            d="M33 25.8471C33 23.5293 31.1211 21.6504 28.8033 21.6504C26.4855 21.6504 24.6066 23.5293 24.6066 25.8471C24.6066 28.1648 26.4855 30.0438 28.8033 30.0438C31.1211 30.0438 33 28.1648 33 25.8471ZM29.369 28.0891V28.566H28.783V28.566V28.566H28.1971V28.0926C27.8429 27.9726 27.553 27.7616 27.27 27.5546L27.9618 26.6087C28.342 26.8868 28.5366 27.0205 28.8033 27.0205C28.9541 27.0205 29.0761 26.9487 29.1219 26.8332C29.1773 26.6934 29.099 26.5645 28.9124 26.4886C28.9124 26.4886 28.9124 26.4886 28.0748 26.2092C27.3627 25.4832 27.2594 25.0164 27.3628 24.5724C27.467 24.1254 27.7688 23.7758 28.1971 23.6012V23.1281H29.369V23.5816C29.6666 23.6643 29.9185 23.7864 30.0634 23.865L29.5043 24.8949C29.1336 24.6937 28.7925 24.6333 28.6598 24.6787C28.531 24.7227 28.5114 24.8067 28.5041 24.8383C28.4936 24.8831 28.488 24.9519 28.5597 25.0317C28.6286 25.1086 29.3542 25.4032 29.3542 25.4032C30.1371 25.7219 30.5055 26.5223 30.2113 27.2649C30.0585 27.6509 29.7508 27.9432 29.369 28.0891Z"
            fill={svgColor}
          />
        </g>
        <defs>
          <clipPath id="clip0_196_454">
            <rect
              width="20"
              height="20"
              fill="white"
              transform="translate(13 13)"
            />
          </clipPath>
        </defs>
      </svg>
    </div>
    <div className="right w-full flex justify-between items-center p-2">
      <div>
        <div className="font-medium text-sm md:text-[12px] text-gray-900 dark:text-darkText">Card Type</div>
        <div className="font-normal text-xs text-[#718EBF] dark:text-gray-400">{cardType}</div>
      </div>
      <div>
        <div className="font-medium text-sm md:text-[12px] text-gray-900 dark:text-darkText">Bank</div>
        <div className="font-normal text-xs text-[#718EBF] dark:text-gray-400">{bank}</div>
      </div>
      <div className="hidden md:block">
        <div className="font-medium text-sm md:text-[12px] text-gray-900 dark:text-darkText">Card Number</div>
        <div className="font-normal text-xs text-[#718EBF] dark:text-gray-400">{cardNumber}</div>
      </div>
      <div className="hidden md:block">
        <div className="font-medium text-sm md:text-[12px] text-gray-900 dark:text-darkText">Namain Card</div>
        <div className="font-normal text-xs text-[#718EBF] dark:text-gray-400">{NamainCard}</div>
      </div>
      <div>
        <Link
          href={detailsLink}
          className="font-normal text-[11px] text-[#1814F3] dark:text-darkAccent"
        >
          View Details
        </Link>
      </div>
    </div>
  </div>
);

}
