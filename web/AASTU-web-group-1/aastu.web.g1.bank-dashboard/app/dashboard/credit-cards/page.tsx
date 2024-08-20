import { Donut } from "@/components/ui/Piechart";
import Image from "next/image";
import CreditCard from "../_components/Credit_Card";
import Cardinfo from "./components/Cardinfo";
import CardSetting from "./components/CardSetting";
import InputForm from "./components/InputForm";

const CreditCards = () => {
  return (
    <div>
      <div className="p-3">
        <div className="p-3">
          <h1 className="text-2xl">My Cards</h1>
          <div className="flex flex-row max-y-[200px] overflow-y-auto gap-6 sm:max-x-[500px] md:max-x-[600px]">
            <CreditCard
              isBlue={true}
              balance={5894}
              creditNumber="3778*** ****1234"
              name="Eddy Cusuma"
              textColor="text-white"
            />
            <CreditCard
              isBlue={false}
              balance={5894}
              creditNumber="3778*** ****1234"
              name="Eddy Cusuma"
              textColor="text-black"
            />
            {/* Add more <CreditCard /> components as needed */}
          </div>
        </div>
        <div className="p-3">
          <div className="md:grid md:grid-cols-2">
            <div className="max-w-screen-sm">
              <h1 className="text-2xl"> Card Expense Statistics </h1>
              <div
                style={{ borderRadius: "5px", overflow: "hidden" }}
                className="rounded-full"
              >
                <Donut />
              </div>
            </div>
            <div className="p-3">
              <h1 className="text-2xl">Card List</h1>
              <Cardinfo />
              <Cardinfo />
              <Cardinfo />
            </div>
          </div>
        </div>
        <h1 className="m-3 text-2xl">Add New Card</h1>
        <div className="md:grid md:grid-cols-[2fr,1fr] gap-4">
          <div className="bg-white rounded-xl py-2 max-w-fit ">
            <p className="m-3 text-[#718EBF] max-h-[440px]">
              Credit Card generally means a plastic card issued by Scheduled
              Commercial Banks assigned to a Cardholder, with a credit limit,
              that can be used to purchase goods and services on credit or
              obtain cash advances.
            </p>
            <div className="m-3">
              <InputForm />
            </div>
          </div>
          <div>
            <h1 className="mb-2 mt-3 text-2xl">Card Setting</h1>
            <div className="max-h-[440px]">
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
    </div>
  );
};

export default CreditCards;


// Adding new Card Fo
// the Card settings

