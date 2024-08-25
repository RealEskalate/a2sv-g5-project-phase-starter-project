import { Mastercard } from "../icons/CreditCard";
import Image from "next/image";

export default function CardB() {
  return (
    <div className="flex gap-12 justify-between pl-6 ">
      <div className=" w-full rounded-xl bg-gradient-to-r from-violet-700 to-blue-600 text-white shadow-2xl transition-transform sm:h-56  sm:hover:scale-110">
        <div className="flex flex-col  w-full px-8 py-8 sm:top-8">
          <div className="flex justify-between">
            <div className="">
              <p className="font-light">Name</p>
              <p className="font-medium tracking-widest">Carter Mullen</p>
            </div>
            <Image
              className="h-14 w-14 object-contain"
              src={Mastercard}
              alt="Mastercard"
              height={56}
              width={56}
            />
          </div>
          <div className="pt-1">
            <p className="font-light">Card Number</p>
            <p className="tracking-more-wider font-medium">
              4312 567 7890 7864
            </p>
          </div>
          <div className="pt-4 pr-6 sm:pt-6">
            <div className="flex justify-between">
              <div className="">
                <p className="text-xs font-light">Valid From</p>
                <p className="text-base font-medium tracking-widest">11/15</p>
              </div>
              <div className="">
                <p className="text-xs font-light">Expiry</p>
                <p className="text-base font-medium tracking-widest">03/25</p>
              </div>

              <div className="">
                <p className="text-xs font-light">CVV</p>
                <p className="tracking-more-wider text-sm font-bold">521</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div className=" w-full rounded-xl bg-gradient-to-r from-slate-500 to-slate-400 text-white shadow-2xl transition-transform sm:h-56 sm:hover:scale-110">
        <div className=" top-4 w-full px-8 py-8 sm:top-8">
          <div className="flex justify-between">
            <div className="">
              <p className="font-light">Name</p>
              <p className="font-medium tracking-widest">Carter Mullen P</p>
            </div>
            <Image
              className="h-14 w-14 object-contain"
              src={Mastercard}
              alt="Mastercard"
              height={56}
              width={56}
            />
          </div>
          <div className="pt-1">
            <p className="font-light">Card Number</p>
            <p className="tracking-more-wider font-medium">
              0006 2345 7453 2345
            </p>
          </div>
          <div className="pt-4 pr-6 sm:pt-6">
            <div className="flex justify-between">
              <div className="">
                <p className="text-xs font-light">Valid From</p>
                <p className="text-base font-medium tracking-widest">11/15</p>
              </div>
              <div className="">
                <p className="text-xs font-light">Expiry</p>
                <p className="text-base font-medium tracking-widest">03/25</p>
              </div>

              <div className="">
                <p className="text-xs font-light">CVV</p>
                <p className="tracking-more-wider text-sm font-bold">521</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
