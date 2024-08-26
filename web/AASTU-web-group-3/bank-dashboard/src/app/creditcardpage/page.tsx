"use client";
import Image from "next/image";
import DoughnutChart from "../components/Chart/Doughnut";
import React from "react";
import { useState } from "react";
import CardListPage from "../components/cardList/CardList";
import CreditCard from "../components/CreditCard";
import {
  useDeleteCardMutation,
  useGetCardsQuery,
  useCreateCardMutation,
} from "../../lib/redux/api/cardsApi";
import { Card } from "../../lib/redux/types/cards";
import Loading from "../loading";
import Link from "next/link";

const cardStyles = {
  Debit: {
    backgroundImg:
      "bg-[linear-gradient(107.38deg,#2D60FF_2.61%,#539BFF_101.2%)]",
    textColor: "text-white",
  },
  Primary: {
    backgroundImg:
      "bg-[linear-gradient(107.38deg,#4C49ED_2.61%,#0A06F4_101.2%)]",
    textColor: "text-white",
  },
  Visa: {
    backgroundImg: "bg-black",
    textColor: "text-white",
  },
  Secondary: {
    backgroundImg: "bg-gray-200",
    textColor: "text-black",
  },
};

const CreditCardPage = () => {
  const { data, error, isLoading } = useGetCardsQuery({ page: 0, size: 10 });
  const [createCard, { isLoading: isCreating }] = useCreateCardMutation();

  const [cardType, setCardType] = useState("");
  const [nameOnCard, setNameOnCard] = useState("");
  const [password, setPassword] = useState("");

  const [expirationDate, setExpirationDate] = useState("");
  const [balance, setBalance] = useState(0);

  const handleAddCard = async (event: React.FormEvent) => {
    event.preventDefault();
    try {
      console.log("Creating card...");
      await createCard({
        cardType: cardType,
        cardHolder: nameOnCard,
        expiryDate: expirationDate,
        balance: balance,
        passcode: password,
      }).unwrap();

      // Reset form fields
      setCardType("");
      setPassword("");
      setNameOnCard("");
      setExpirationDate("");
      setBalance(0);
      window.location.reload();
    } catch (error) {
      console.error("Failed to add new card", error);
    }
  };

  const [deleteCard] = useDeleteCardMutation();

  const handleBlockCard = async () => {
    if (data?.content.length === 0) return;

    const cardIndex = prompt(
      "Enter the card number you want to delete (e.g., 1st, 2nd, 3rd):"
    );
    const index = parseInt(cardIndex ?? "") - 1; // Convert to zero-based index

    if (index >= 0 && data?.content && index < data.content.length) {
      const cardId = data?.content[index].id;
      try {
        await deleteCard(data?.content[index].id).unwrap();

        // Reload the page after successful card deletion
        window.location.reload();
      } catch (error) {
        console.error("Failed to delete card", error);
      }
    } else {
      alert("Invalid card number.");
    }
  };

  if (isLoading) return <Loading />;
  if (error) return <div>Error fetching cards</div>;
  return (
    <div className="body bg-[#F5F7FA] dark:bg-darkPage w-full h-full overflow-y-auto pb-5 m-0">
      <div className="cards m-2 bg-[#F5F7FA] dark:bg-darkPage">
        <div className="credit-card-info flex md:pl-0 px-3 h-20 items-center">
          <h1 className="font-semibold text-[#343C6A] dark:text-white text-[16px] md:text-[22px] ml-4">
            My cards
          </h1>
        </div>
  
        <div className="creditcards flex gap-5 lg:flex-row overflow-x-auto overflow-y-hidden no-scrollbar h-56 lg:justify-start lg:px-4">
          {data?.content.map((card: Card, index: number) => {
            const style =
              cardStyles[card.cardType as keyof typeof cardStyles] ||
              cardStyles.Primary;
  
            return (
              <div
                key={index}
                className="credit-card min-h-80 w-[360px] max-w-72 md:max-w-96 flex-shrink-0"
              >
                <CreditCard
                  name={card.cardHolder}
                  balance={String(card.balance)}
                  cardNumber={card.semiCardNumber}
                  validDate={card.expiryDate}
                  backgroundImg={style.backgroundImg}
                  textColor={style.textColor}
                />
              </div>
            );
          })}
        </div>
      </div>
  
      <div className="statandlist md:flex">
        <div className="md:w-[30%] md:mx-4 dark:dark:bg-darkComponent">
          <p className="p-4 md:pl-0 font-semibold text-[16px] leading-[19.36px] text-[#343C6A]  dark:text-white mx-2">
            Card Expense Statistics
          </p>
          <div className="piechart flex md:h-[230px] md:mx-2 items-center h-auto bg-white  dark:bg-darkComponent rounded-[15px] mx-5 px-8">
            <DoughnutChart />
          </div>
        </div>
  
        <div className="md:w-[70%]">
          <p className="p-4 md:pl-0 md:pb-2 font-semibold text-[16px] leading-[19.36px] text-[#343C6A] dark:text-white mx-2">
            Card List
          </p>
          <div className="cardList w-auto h-auto mx-4 md:mx-0 dark:bg-darkComponent">
            <CardListPage />
          </div>
        </div>
      </div>
  
      <div className="addnewandcardsetting md:flex md:flex-row md:w-full">
        <div className="addnewcard flex flex-col md:w-[65%]">
          <div className="p-5 font-semibold text-[16px] leading-[19.36px] text-[#343C6A] dark:text-white mx-2">
            Add New Card
          </div>
          <form className="newcard flex flex-col md:flex md:flex-row md:flex-wrap md:w-auto md:h-[321px] justify-between w-auto h-[527px] p-4 border-2 bg-white dark:bg-darkComponent rounded-[15px] mx-6">
            <p className="description text-[#718EBF] dark:text-gray-400 text-[12px] leading-[22px]">
              Credit Card generally means a plastic card issued by Scheduled
              Commercial Banks assigned to a Cardholder, with a credit limit,
              that can be used to purchase goods and services on credit or
              obtain cash advances.
            </p>
            <div className="flex flex-col md:w-[50%]">
              <label htmlFor="CardType" className="md:text-xs md:font-normal dark:text-gray-400">
                Card Type
              </label>
              <select
                name="CardType"
                id="CardType"
                value={cardType}
                onChange={(e) => setCardType(e.target.value)}
                className="border-[1px] dark:border-gray-700 md:w-[90%] w-auto h-[40px] rounded-[10px] p-2 md:text-xs bg-white dark:bg-gray-800 text-gray-800 dark:text-gray-200"
              >
                <option value="" disabled>
                  Select Card Type
                </option>
                <option value="Primary">Primary</option>
                <option value="Secondary">Secondary</option>
                <option value="Visa">Visa</option>
                <option value="Debit">Debit</option>
              </select>
            </div>
  
            <div className="flex flex-col md:w-[50%]">
              <label htmlFor="nameoncard" className="md:text-xs md:font-normal dark:text-gray-400">
                Name On Card
              </label>
              <input
                type="text"
                name="nameoncard"
                id="nameoncard"
                placeholder="My Cards"
                value={nameOnCard}
                onChange={(e) => setNameOnCard(e.target.value)}
                className="border-[1px] dark:border-gray-700 md:w-[90%] w-auto h-[40px] rounded-[10px] p-3 md:text-xs bg-white dark:bg-gray-800 text-gray-800 dark:text-gray-200"
              />
            </div>
            <div className="flex flex-col md:w-[50%]">
              <label htmlFor="Password" className="md:text-xs md:font-normal dark:text-gray-400">
                Password
              </label>
              <input
                type="password"
                name="Password"
                id="Password"
                placeholder="******"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="border-[1px] dark:border-gray-700 md:w-[90%] w-auto h-[40px] rounded-[10px] p-3 md:text-xs bg-white dark:bg-gray-800 text-gray-800 dark:text-gray-200"
              />
            </div>
            <div className="flex flex-col md:w-[50%]">
              <label htmlFor="expirationdate" className="md:text-xs md:font-normal dark:text-gray-400">
                Expiration Date
              </label>
              <input
                type="date"
                name="expirationdate"
                id="expirationdate"
                value={expirationDate}
                onChange={(e) => setExpirationDate(e.target.value)}
                className="border-[1px] dark:border-gray-700 md:w-[90%] w-auto h-[40px] rounded-[10px] p-3 md:text-xs bg-white dark:bg-gray-800 text-gray-800 dark:text-gray-200"
              />
            </div>
            <div className="md:flex md:w-full">
              <div className="flex flex-col md:w-[50%]">
                <label htmlFor="balance" className="md:text-xs md:font-normal dark:text-gray-400">
                  Balance
                </label>
                <input
                  type="text"
                  name="balance"
                  id="balance"
                  placeholder="1000.00"
                  onChange={(e) => setBalance(parseFloat(e.target.value))}
                  className="border-[1px] dark:border-gray-700 md:w-[90%] w-auto h-[40px] rounded-[10px] p-3 md:text-xs bg-white dark:bg-gray-800 text-gray-800 dark:text-gray-200"
                />
              </div>
              <button
                type="submit"
                onClick={handleAddCard}
                disabled={isCreating}
                className="mt-6 md:mx-0 border-[1px] md:font-normal md:w-[45%] w-full h-[40px] md:my-4 bg-[#1814F3] text-white rounded-[10px] md:text-xs dark:bg-[#3B82F6]"
              >
                {isCreating ? "Adding..." : "Add Card"}
              </button>
            </div>
          </form>
        </div>
        <div className="card setting md:w-[35%] md:mr-8 m-auto">
          <div className="p-5 md:pl-0 font-semibold text-[16px] leading-[19.36px] text-[#343C6A] dark:text-white mx-2">
            Card Setting
          </div>
          <div className="flex flex-col justify-between md:w-full md:mx-2 w-auto p-4 h-[325px] border-[1px] dark:border-gray-700 rounded-[15px] bg-white dark:bg-darkComponent mx-6">
            <div
              onClick={handleBlockCard}
              className="blockcard cursor-pointer md:pl-0 flex w-auto hover:w-[75%] hover:rounded-md hover:bg-slate-100 dark:hover:bg-gray-700 h-[45px] pl-5"
            >
              <div className="left">
                {/* SVG remains unchanged */}
              </div>
              <div className="right flex-row w-auto h-auto p-2 pl-3">
                <div className="w-auto font-normal text-[#343C6A] dark:text-white">
                  Block My Card
                </div>
                <div className="w-auto h-auto text-[#718EBF] dark:text-gray-400">
                  Stop all payment
                </div>
              </div>
            </div>
            {/* Similar updates for other sections */}
          </div>
        </div>
      </div>
    </div>
  );
  
};

export default CreditCardPage;
