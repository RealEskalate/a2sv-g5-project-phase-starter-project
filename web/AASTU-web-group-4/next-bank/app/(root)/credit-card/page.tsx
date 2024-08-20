import ResponsiveCreditCard from "@/components/CreditCard";
import { colors } from "@/constants";
import Component from "@/components/DoughnutChart";
import AddNewCard from "@/components/AddNewCard";
import CardSetting from "@/components/CardSetting";
import CardList from "@/components/CardList";
import { useFetchCards } from "@/components/AllCards";
import { cardType } from "@/types";

const token = "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJuYXRpIiwiaWF0IjoxNzI0MTU2NzgyLCJleHAiOjE3MjQyNDMxODJ9.QbgA09xcnOGWFNOXKiYfkUQFZ0WxqNFc7u5OojUQEZJfd4DUdJVJ_pCj8b2Ebmzr";

const CreditCard = async () => {
  const cards = await useFetchCards(token);
  console.log("alemitu", cards)
  return (
    <div className="lg:ml-72 ml-5 overflow-x-hidden mx-auto">
      <div className="myCards max-w-[97%] mt-4">
        <h1 className="text-[19px] mb-3 font-bold text-[#333B69]">My Cards</h1>
        <div className="flex overflow-x-auto space-x-4 md:pr-3 pr-1 scrollbar-none">
          {cards.map((card: any, index: number) => (
            <span key={index} className="p-3">
              <ResponsiveCreditCard
                backgroundColor={index % 2 == 0 ? colors.blue : colors.white}
                balance={card.balance}
                cardHolder={card.cardHolder}
                expiryDate={card.expiryDate.slice(0, 10)}
                cardNumber={card.semiCardNumber}
              />
            </span>
          ))}
        </div>
      </div>

      <div className="flex flex-col md:flex-row gap-16">
        <div className="doughnutChart lg:w-[350px] md:w-[231] w-[325px] my-6">
          <h1 className="text-[19px] mb-3 font-bold text-[#333B69]">
            Card Expense Statistics
          </h1>
          <Component />
        </div>
        <div className="cardlist lg:w-[730px] md:w-[487px] sm-w-[325] my-6">
          <h1 className="text-[19px] mb-3 font-bold text-[#333B69]">
            Card List
          </h1>
          <CardList card_list ={cards} />
        </div>
      </div>

      <div className="flex flex-col md:flex-row w-[80%] mb-16">
        <div className="md:mb-2 mb-0 md:mr-10">
          <AddNewCard />
        </div>

        <div>
          <h1 className="text-[19px] mb-3 font-bold text-[#333B69]">
            Card Setting
          </h1>
          <CardSetting />
        </div>
      </div>
    </div>
  );
};

export default CreditCard;
