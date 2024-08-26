"use client";
import ResponsiveCreditCard from "@/components/CreditCard";
import { colors } from "@/constants";
import Component from "@/components/DoughnutChart";
import AddNewCard from "@/components/AddNewCard";
import CardSetting from "@/components/CardSetting";
import CardList from "@/components/CardList";
import { useEffect, useState } from "react";
import { getAllCards } from "@/services/cardfetch";
import Cookies from "js-cookie";
import Image from "next/image";
import CardListLoad from "@/components/loadingComponents/CardListLoad";
import MyCardsLoad from "@/components/loadingComponents/MyCardsLoad";
import { TbFileSad } from "react-icons/tb";

const CreditCard = () => {
  const [cards, setCards] = useState<any[]>([]);
  const [token, setToken] = useState<any>(null); // Initialize with null to indicate it's being fetched
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(true); // Add a loading state

  useEffect(() => {
    const fetchCards = async () => {
      try {
        const storedToken = Cookies.get("accessToken");
        setToken(storedToken);

        if (!storedToken) {
          setLoading(false); // Stop loading if no token is found
          throw new Error("Token not found. Please log in again.");
        }

        const data = await getAllCards(storedToken);
        setCards(data);
        setError(null);
      } catch (err) {
        setError("Failed to fetch cards data!");
      } finally {
        setLoading(false); // Stop loading once the fetch is done
      }
    };

    fetchCards();
  }, [cards]);

  return (
    <div className="lg:ml-72 ml-5 overflow-x-hidden mx-auto">
      <div className="myCards max-w-[97%] mt-4">
        <h1 className="text-[19px] mb-3 font-bold text-[#333B69] dark:text-blue-500">My Cards</h1>
        <div className="flex overflow-x-auto space-x-4 md:pr-3 pr-1 scrollbar-thin scrollbar-track-[#F5F7FA] dark:scrollbar-track-dark scrollbar-thumb-[#92a7c5] scrollbar-thumb-rounded-full">
          {loading ? (
            <MyCardsLoad count={3}/>
          ) : Array.isArray(cards) && cards.length > 0 ? (
            cards.map((card: any, index: number) => (
              <span key={index} className="p-3">
                <ResponsiveCreditCard
                  backgroundColor={index % 2 === 0 ? colors.blue : colors.white}
                  balance={card.balance}
                  cardHolder={card.cardHolder}
                  expiryDate={card.expiryDate.slice(0, 10)}
                  cardNumber={card.semiCardNumber}
                />
              </span>
            ))
          ) : token ? (
            <div className="max-h-[400px] lg:w-[730px] md:w-[487px] bg-white py-16 rounded-xl flex flex-col justify-center dark:bg-dark dark:border-[1px] dark:border-gray-700">
              <TbFileSad
                className={`text-gray-300 dark:text-[#993d4b] w-[400px] h-[70px] pb-2 block mx-auto`}
                strokeWidth={1}
              />
              <span className="mx-auto my-auto md:text-xl text-sm text-red-500 mb-5">
                {error ? error : "There are no cards for now!"}
              </span>
            </div>
          ) : (
            <MyCardsLoad count={3}/>
          )}
        </div>
      </div>

      <div className="flex flex-col md:flex-row gap-16">
        <div className="doughnutChart lg:w-[360px] md:w-[231] w-[325px] my-6">
          <h1 className="text-[19px] mb-6 font-bold text-[#333B69] dark:text-blue-500">
            Card Expense Statistics
          </h1>
          <Component />
        </div>
        <div className="cardlist lg:w-[730px] md:w-[487px] sm-w-[325] my-6">
          <h1 className="text-[19px] mb-3 font-bold text-[#333B69] dark:text-blue-500">Card List</h1>
          {loading ? (
            <CardListLoad />
          ) : token ? (
            error ? (
              <div className="pr-6 py-32 bg-white max-h-[400px] lg:w-[730px] md:w-[487px] w-[325] flex flex-col justify-center align-middle rounded-xl scrollbar-none dark:bg-dark dark:border-[1px] dark:border-gray-700 ">
                <TbFileSad
                    className={`text-gray-300 dark:text-[#993d4b] w-[400px] h-[70px] pb-2 block mx-auto`}
                    strokeWidth={1}
                  />
                <span className="mx-auto my-auto md:text-xl text-sm text-red-500">
                  {error}
                </span>
              </div>
            ) : (
              <CardList card_list={cards} />
            )
          ) : (
            <CardListLoad />
          )}
        </div>
      </div>

      <div className="flex flex-col md:flex-row w-[80%] mb-16">
        <div className="md:mb-2 mb-0 md:mr-5 lg:mr-10">
          <h1 className="text-[20px] mb-3 font-bold text-[#333B69] dark:text-blue-500">Add New Card</h1>
          <AddNewCard token={token} />
        </div>

        <div>
          <h1 className="text-[19px] mb-3 font-bold text-[#333B69] md:mt-0 mt-6 dark:text-blue-500">Card Setting</h1>
          <CardSetting />
        </div>
      </div>
    </div>
  );
};

export default CreditCard;
