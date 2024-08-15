import CreditCard from "../_components/CreditCard";
import React from "react";
import Image from "next/image";
import { Donut } from "@/components/ui/Piechart";

const CreditCards = () => {
  return (
    <div>
      <div className="p-3">
        <div className="p-3">
          <h1 className="text-2xl">My Cards</h1>
          <div className="flex flex-row max-y-[200px] overflow-y-auto gap-6 sm:max-x-[500px] md:max-x-[600px]">

            <CreditCard />
            <CreditCard />
            {/* Add more <CreditCard /> components as needed */}
          </div>
        </div>

        <div className="p-3">
          <h1> Card Expense Statistics </h1>
          <div className="max-w-screen-sm">
            <Donut />
          </div>
        </div>
        <div className="p-3">
          <h1>Cardlist</h1>
          <Cardinfo />
          <Cardinfo />
          <Cardinfo />
        </div>
        <div className="bg-white rounded-xl py-2">
          <h1 className="m-3">Add New Card</h1>
          <p className="m-3 text-[#718EBF]">
            Credit Card generally means a plastic card issued by Scheduled
            Commercial Banks assigned to a Cardholder, with a credit limit, that
            can be used to purchase goods and services on credit or obtain cash
            advances.
          </p>
          <div className="m-3">
            <InputForm />
          </div>
        </div>
        <div>
          <h1 className="mb-2 mt-3">Card Setting</h1>
          <div>
            <CardSetting
              image="/images/BlockCard.png"
              title="Block Card"
              description="Instantly block your card"
              color="bg-yellow-200"
            />
            <CardSetting
              image="/images/Lock.png"
              title="Change Pic Code"
              description="Withdraw without any card"
              color="bg-violet-200"
            />
            <CardSetting
              image="/images/Google.png"
              title="Add to Google Pay"
              description="Withdraw without any card"
              color="bg-pink-200"
            />
            <CardSetting
              image="/images/Apple.png"
              title="Add to Apple Pay"
              description="Withdraw without any card"
              color="bg-green-200"
            />
            <CardSetting
              image="/images/Apple.png"
              title="Added to Apple store"
              description="Withdraw without any card"
              color="bg-green-200"
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default CreditCards;

// list of Cards with type and related bank
const Cardinfo = () => {
  return (
    <div>
      <div className="flex p-5 bg-white rounded-xl mb-2 mt-4 max-w-screen-sm justify-between">
        <div className="flex-initial w-[2/12] m-3">
          <Image
            src={`/images/CardType.png`}
            alt={"Cards"}
            width={27}
            height={18}
          />
        </div>
        <div className="flex-initial w-[4/12] m-3">
          <div>
            <h2>Card Type</h2>
            <p className="text-gray-500">Secondary</p>
          </div>
        </div>
        <div className="flex-initial w-[3/12] m-3">
          <div>
            <h2>Bank</h2>
            <p className="text-gray-500">DBL Bank</p>
          </div>
        </div>
        <div className="flex-initial w-[3/12] m-3">
          <p className="text-[#1814F3]"> view details</p>
        </div>
      </div>
    </div>
  );
};

// Adding new Card Form
const InputForm = () => {
  return (
    <div>
      <form>
        <div className="flex flex-col my-4">
          <label className="text-[#515B6F] font-semibold">Card Type</label>
          <input
            type="text"
            placeholder="Classic"
            className="inputField mb-2 rounded-xl py-2 px-2 border border-gray-300"
          />
          {/* <p className="text-red-500 text-center mt-2">
              {errors.name?.message}
            </p> */}
        </div>
        <div className="flex flex-col my-4">
          <label className="text-[#515B6F] font-semibold">Name On Card</label>
          <input
            type="text"
            placeholder="My Cards"
            className="inputField mb-2 rounded-xl py-2 px-2 border border-gray-300"
            // {...register("email")}
          />
          {/* <p className="text-red-500 text-center mt-2">
              {errors.email?.message}
            </p> */}
        </div>
        <div className="flex flex-col my-4">
          <label className="text-[#515B6F] font-semibold">Card Number</label>
          <input
            type="number"
            placeholder="**** **** **** ****"
            className="inputField mb-2 rounded-xl py-2 px-2 border border-gray-300"
            // {...register("password")}
          />
          {/* <p className="text-red-500 text-center mt-2">
              {errors.password?.message}
            </p> */}
        </div>
        <div className="flex flex-col my-4">
          <label className="text-[#515B6F] font-semibold">
            Expiration Date
          </label>
          <input
            type="date"
            placeholder="00:00:00 UTC on 1st January 1970"
            className="inputField mb-2 rounded-xl py-2 px-2 border border-gray-300"
            // {...register("confirmPassword")}
          />
          {/* <p className="text-red-500 text-center mt-2">
              {errors.confirmPassword?.message}
            </p> */}
        </div>
        {/* {error && (
            <p className="text-red-500 text-center mt-2 mb-5">{error}</p>
          )} */}
        <button className="bg-[#1814F3] w-[100%] text-white p-2 rounded-full mb-2">
          Add to Cart
        </button>
      </form>
    </div>
  );
};

// the Card settings
interface props {
  image: string;
  color: string;
  title: string;
  description: string;
}

const CardSetting = ({ image, color, title, description }: props) => {
  return (
    <div className="flex bg-white mb-3 ">
      <div className="flex-initial w-[5/12] m-3">
        <div
          className={`${color} bg-opacity-25 font-semibold py-2 px-4 rounded-lg  text-sm w-[45px] `}
        >
          <Image src={image} alt={title} width={20} height={20} />
        </div>
      </div>
      <div className="flex-initial w-[7/12] m-3">
        <div>
          <h1>{title}</h1>
          <p>{description}</p>
        </div>
      </div>
    </div>
  );
};
