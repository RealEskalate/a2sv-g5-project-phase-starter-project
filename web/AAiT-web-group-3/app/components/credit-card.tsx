import Image from "next/image";
import React from "react";

const CreditCard = ({ color }: { color: string }) => {
  return (
    <div className="w-full">
      {color === "blue" && (
        <div className="flex flex-col justify-between overflow-hidden items-star rounded-2xl min-w-[265px] sm:min-w-fit bg-gradient-to-r from-[#4C49ED] to-[#0A06F4] text-white sm:h-[230px] md:h-fit xl:min-h-[230px]">
          <section className="flex justify-between items-center p-2 px-3 md:p-4 pb-2">
            <div>
              <h3 className="text-sm">Balance</h3>
              <p className="text-lg font-medium">$5,756</p>
            </div>
            <Image
              src={"/images/whiteSimCardChip.png"}
              className="rounded-[10px] bg-blue-500 p-0"
              alt="blackSimCardChip"
              width={40}
              height={30}
            />
          </section>

          <section className="flex justify-between items-center p-2 px-3 md:p-4">
            <div>
              <h3 className="font-light text-sm">CARD HOLDER</h3>
              <p className="font-medium">Bemnet Adu</p>
            </div>

            <div>
              <h3 className="font-light text-sm">VALID THRU</h3>
              <p className="font-medium">12/23</p>
            </div>
          </section>

          <section className="p-2 px-3 md:p-4 flex justify-between items-center pt-4 bg-gradient-to-br from-[#6360df] to-[#130fee] ">
            <h3 className="font-medium ">3778 **** **** 1234</h3>
            <span className="flex gap-[-4px]">
              <span className="bg-white/60 w-6 h-6 rounded-full"></span>
              <span className="bg-white/60 w-6 h-6 rounded-full ml-[-25%]"></span>
            </span>
          </section>
        </div>
      )}

      {color === "white" && (
        <div className="flex flex-col justify-between overflow-hidden items-star rounded-2xl min-w-[265px] sm:min-w-fit bg-white border border-[#DFEAF2] sm:h-[230px] md:h-fit xl:min-h-[230px]">
          <section className="flex justify-between items-center p-2 px-3 md:p-4 pb-2">
            <div>
              <h3 className="text-sm text-primary-color-200">Balance</h3>
              <p className="text-lg font-medium text-primary-color-800">
                $5,756
              </p>
            </div>
            <Image
              src={"/images/blackSimCardChip.png"}
              className="rounded-lg p-0"
              alt="blackSimCardChip"
              width={40}
              height={30}
            />
          </section>

          <section className="flex justify-between items-center p-2 px-3 md:p-4">
            <div>
              <h3 className=" text-sm text-primary-color-200">CARD HOLDER</h3>
              <p className="font-medium text-primary-color-800">Bemnet Adu</p>
            </div>

            <div>
              <h3 className=" text-sm text-primary-color-200">VALID THRU</h3>
              <p className="font-medium text-primary-color-800">12/23</p>
            </div>
          </section>

          <section className="p-2 px-3 md:p-4 flex justify-between items-center pt-4 border-t-2 ">
            <h3 className="font-medium text-primary-color-800">
              3778 **** **** 1234
            </h3>
            <span className="flex gap-[-4px]">
              <span className="bg-gray-500/60 w-6 h-6 rounded-full"></span>
              <span className="bg-gray-500/60 w-6 h-6 rounded-full ml-[-25%]"></span>
            </span>
          </section>
        </div>
      )}
    </div>
  );
};

export default CreditCard;
