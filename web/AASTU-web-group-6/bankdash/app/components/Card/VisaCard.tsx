import React from "react";
import Image from "next/image";
import simUrl from "../../../public/assets/sim-icon.png";
import simBlackUrl from "../../../public/assets/sim-black-icon.png";
interface ColorType {
  isBlack: boolean;
  isFade: boolean;
}
const VisaCard: React.FC<ColorType> = ({ isBlack, isFade }) => {
  return (
    <div
      className={`w-full font-Lato flex flex-col gap-2 grow rounded-3xl ${
        isBlack
          ? "text-colorBody-1 bg-white border-solid border-[1px] border-gray-200"
          : isFade
          ? "text-white bg-card-gradient-2"
          : "text-white bg-card-gradient-1"
      } `}
    >
      <div className="flex flex-col gap-6 p-6">
        <div className="flex text-sm">
          <div className="balance-box flex flex-col grow">
            <div className="label font-normal">Balance</div>
            <div className="balance text-xl font-medium">$5,756</div>
          </div>
          <Image
            width={48}
            height={24}
            src={isBlack ? simBlackUrl : simUrl}
            alt="SIM Icon"
            className="simIcon"
          />
        </div>
        <div className="flex">
          <div className="holder-box text-sm font-normal flex flex-col grow">
            <div className="label opacity-70">CARD HOLDER</div>
            <div className="name text-base font-medium">Eddy Cusuma</div>
          </div>
          <div className="valid-box">
            <div className="label opacity-70">VALID THRU</div>
            <div className="name text-base font-medium">12/22</div>
          </div>
        </div>
      </div>
      <div
        className={`flex items-center gap-2 card-box rounded-b-3xl p-4 bg-card-box-light ${
          isBlack ? "border-solid border-t-2 border-gray-200" : ""
        }`}
      >
        <div className="number flex grow font-medium text-xl">
          3778*** **** 1234
        </div>
        <div className="number flex">
          <div
            className={`circle w-8 h-8 ${
              isBlack ? "bg-colorBody-2" : "bg-white"
            } opacity-50 rounded-full`}
          ></div>
          <div
            className={`circle w-8 h-8 ${
              isBlack ? "bg-colorBody-2" : "bg-white"
            } opacity-50 -ml-4 rounded-full`}
          ></div>
        </div>
      </div>
    </div>
  );
};

export default VisaCard;
