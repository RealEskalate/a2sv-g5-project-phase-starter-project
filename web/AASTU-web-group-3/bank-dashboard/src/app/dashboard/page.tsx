"use client";
import React from "react";
import Image from "next/image";
import {
  useCreateTransactionMutation,
  useGetAllTransactionsQuery,
} from "@/lib/redux/api/transactionsApi";
import CreditCard from "../components/CreditCard";
import { BarChartComponent } from "../components/Chart/Barchart";
import { PieChartComponent } from "../components/Chart/PieChart";
import { AreaChartComponent } from "../components/Chart/AreaChartComponent";
import { Transaction } from "@/lib/redux/types/transactions";
import { useDispatch, useSelector } from "react-redux";
import dollar from "../../../public/images/iconfinder_6_4753731 1.png";
import Link from "next/link";
import { useGetCardsQuery } from "@/lib/redux/api/cardsApi";
import { RootState } from "@/lib/redux/store";
import { setCards, setLoading, setError } from "@/lib/redux/slices/cardsSlice";
import { useEffect } from "react";
import { cardStyles } from "@/lib/CardColor";
import { Card } from "../../lib/redux/types/cards";
import Loading from "../loading";
import { CardData } from "@/types/cardData";
import { DollarSignIcon } from "lucide-react";

import { useGetQuickTransfersQuery } from "@/lib/redux/api/transactionsApi";

const getTransactionAmount = (transaction: Transaction) => {
  switch (transaction.type) {
    case "shopping":
    case "transfer":
    case "service":
      return `-$${Math.abs(transaction.amount)}`;
    case "deposit":
      return `+$${transaction.amount}`;
    default:
      return `$${transaction.amount}`;
  }
};

const getAmountStyle = (transaction: Transaction) => {
  switch (transaction.type) {
    case "shopping":
    case "transfer":
    case "service":
      return "text-red-500";
    case "deposit":
      return "text-green-500";
    default:
      return "";
  }
};

const imageData = [
  {
    src: "https://via.placeholder.com/48",
    alt: "Placeholder 1",
    name: "Name 1",
    position: "Position 1",
  },
  {
    src: "https://via.placeholder.com/48",
    alt: "Placeholder 2",
    name: "Name 2",
    position: "Position 2",
  },
  {
    src: "https://via.placeholder.com/48",
    alt: "Placeholder 3",
    name: "Name 3",
    position: "Position 3",
  },
];

const HomePage = () => {
  const { data: transactionsData, isLoading } = useGetAllTransactionsQuery({
    size: 3,
    page: 0,
  });
  const { data: quickTransfersData, isLoading: isTransfersLoading } =
    useGetQuickTransfersQuery({
      num: 3,
    });

  const [selectedUser, setSelectedUser] = React.useState(null);
  const [amount, setAmount] = React.useState("");
  const [createTransaction] = useCreateTransactionMutation();

  const handleTransferClick = (user: any) => {
    setSelectedUser(user);
    setAmount("");

    const handleSend = async () => {
      if (!selectedUser || !amount) return;

      const transactionData = {
        type: "transfer",
        description: "Quick transfer",
        amount: parseFloat(amount),
        receiverUserName: selectedUser.username,
      };

      try {
        await createTransaction(transactionData);
        alert("Transfer successful!");
      } catch (error) {
        console.error("Transfer failed:", error);
      }
    };

    const dispatch = useDispatch();
    const { cards, loading, error } = useSelector(
      (state: RootState) => state.cards
    );

    const {
      data: cardsData,
      isLoading: cardsLoading,
      isError: errorCard,
    } = useGetCardsQuery({ size: 5, page: 0 });

    useEffect(() => {
      dispatch(setLoading(cardsLoading));
      if (cardsData) {
        dispatch(setCards(cardsData.content));
      }
      if (errorCard) {
        dispatch(setError("Error on fetching data"));
      }
    }, [cardsData, errorCard, cardsLoading, dispatch]);

    if (isLoading) return <div>Loading...</div>;

    if (isLoading) return <Loading />;
    if (error) return <div>Error fetching cards</div>;
    return (
      <div className="bg-[#F5F7FA] min-h-screen p-5">
        <div className="lg:flex lg:justify-between lg:border-2">
          <div className="lg:w-[60%] rounded-xl bg-[#F5F7FA] ">
            <div className="credit-card-info flex justify-between  h-16 items-center ">
              <h3 className="font-bold text-[#343C6A] text-[16px] md:text-[22px] ">
                My cards
              </h3>
              <Link href="/creditcardpage#add-new-card">
                <h3 className="font-bold text-[#343C6A] text-[16px] md:text-[22px] ">
                  See All
                </h3>
              </Link>
            </div>
            <div className="creditcards flex  gap-5 lg:flex-row overflow-x-auto overflow-y-hidden no-scrollbar  h-56  lg:justify-start lg:px-4 ">
              {cardsData?.content.map((card: Card, index: number) => {
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

          <div className="lg:w-[35%] lg:pt-3 lg:pb-2">
            <div>
              <p className="font-bold text-[#343C6A] text-[16px] md:text-[22px] mb-3">
                Recent Transaction
              </p>
            </div>
            <div className="transactionData mt-5  h-[210px] lg:mt-3 rounded-lg lg:h-[215px] bg-white lg:p-5 lg:rounded-lg lg:shadow-md ">
              {transactionsData!.data.content.map((transaction, index) => (
                <div key={index} className="flex justify-around  items-center ">
                  <div className="flex left  p-3 items-center space-x-2">
                    <div className="relative left w-12 h-10 rounded-full flex items-center justify-center bg-gray-200">
                      <Image
                        src={dollar}
                        alt="transaction"
                        height={100}
                        width={100}
                      />
                    </div>
                    <div className="flex  flex-col">
                      <p className="text-sm font-medium">
                        {transaction.description}
                      </p>
                      <small className="text-xs text-gray-500">
                        {transaction.date}
                      </small>
                    </div>
                  </div>
                  <p
                    className={`font-semibold m-auto mr-5 ${getAmountStyle(
                      transaction
                    )}`}
                  >
                    {getTransactionAmount(transaction)}
                  </p>
                </div>
              ))}
            </div>
          </div>
        </div>

        <div className="my-5 lg:space-y-0 lg:flex lg:justify-between  ">
          <div className="weeklyActivities  h-auto rounded-lg">
            <div>
              <p className="font-bold text-[#343C6A] text-[16px] md:text-[22px] mb-3">
                Weekly Activities
              </p>
            </div>
            <BarChartComponent />
          </div>

          <div className="expense my-5 rounded-lg h-auto lg:p-0 lg:m-0">
            <div>
              <p className="font-bold text-[#343C6A] text-[16px] md:text-[22px] mb-3">
                Expenses Stasitics
              </p>
            </div>
            <PieChartComponent />
          </div>
        </div>

        <div className="quickandBalance flex flex-col lg:flex-row items-start lg:items-center lg:justify-between gap-6 lg:gap-8  lg:w-full">
          <div className="quicktransfer  w-full lg:w-1/2  p-4 lg:p-6 lg:h-[450px]">
            <div className="my-6">
              <p className="font-bold text-[#343C6A] text-[16px] md:text-[22px]">
                Quick Transfer
              </p>
            </div>
            <div className="w-full bg-white p-5 rounded-xl shadow-md">
              <div className="flex justify-center">
                <div className="flex mt-3 space-x-16">
                  {quickTransfersData?.data.map((user, index) => (
                    <div
                      key={index}
                      className="flex flex-col items-center cursor-pointer"
                      onClick={() => handleTransferClick(user)}
                    >
                      <div className="w-12 h-12 bg-gray-300 rounded-full overflow-hidden">
                        <img
                          src={user.profilePicture} // Assuming profilePicture is a valid image URL
                          alt={user.name}
                          className="w-full h-full object-cover"
                        />
                      </div>
                      <p className="text-sm font-medium">{user.name}</p>
                      <small className="text-xs text-gray-500">
                        {user.username}
                      </small>
                    </div>
                  ))}
                </div>
                <div className="svg flex items-center ml-4">
                  <svg
                    width="72"
                    height="72"
                    viewBox="0 0 50 72"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <g filter="url(#filter0_d_150_11)">
                      <circle cx="32" cy="32" r="20" fill="white" />
                    </g>
                    <path
                      d="M29 26L35.5 32.5L29 39"
                      stroke="#718EBF"
                      stroke-width="1.5"
                    />
                    <defs>
                      <filter
                        id="filter0_d_150_11"
                        x="0"
                        y="0"
                        width="72"
                        height="72"
                        filterUnits="userSpaceOnUse"
                        color-interpolation-filters="sRGB"
                      >
                        <feFlood
                          flood-opacity="0"
                          result="BackgroundImageFix"
                        />
                        <feColorMatrix
                          in="SourceAlpha"
                          type="matrix"
                          values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
                          result="hardAlpha"
                        />
                        <feMorphology
                          radius="2"
                          operator="erode"
                          in="SourceAlpha"
                          result="effect1_dropShadow_150_11"
                        />
                        <feOffset dx="4" dy="4" />
                        <feGaussianBlur stdDeviation="9" />
                        <feColorMatrix
                          type="matrix"
                          values="0 0 0 0 0.904427 0 0 0 0 0.893194 0 0 0 0 0.908333 0 0 0 0.8 0"
                        />
                        <feBlend
                          mode="normal"
                          in2="BackgroundImageFix"
                          result="effect1_dropShadow_150_11"
                        />
                        <feBlend
                          mode="normal"
                          in="SourceGraphic"
                          in2="effect1_dropShadow_150_11"
                          result="shape"
                        />
                      </filter>
                    </defs>
                  </svg>
                </div>
              </div>
              <div className="bottomSend mt-5 lg:mt-8 flex justify-between items-center">
                <h3 className="text-[#718EBF] pl-3">Write Amount</h3>
                <div className="relative flex items-center mt-3">
                  <input
                    type="text"
                    placeholder="525.20"
                    className="w-full p-3 h-12 rounded-3xl text-black border bg-gray-100"
                  />
                  <button className="flex items-center justify-center absolute top-0 right-0 h-full px-3 bg-[#1814F3] text-white rounded-3xl">
                    <span className="mr-3">Send</span>
                    <svg
                      width="17"
                      height="14"
                      viewBox="0 0 17 14"
                      fill="none"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <path
                        d="M16.0963 0.572034C16.1748 0.206511 15.8164 -0.101694 15.4669 0.0316926L0.303581 5.82017C0.121087 5.88985 0.000378397 6.06476 8.88017e-07 6.26009C-0.000376621 6.45546 0.11964 6.63084 0.301883 6.70121L4.56154 8.34637V13.5281C4.56154 13.7467 4.71163 13.9366 4.9243 13.9872C5.13554 14.0374 5.35582 13.9369 5.45485 13.7404L7.2166 10.2445L11.5159 13.4351C11.7774 13.6292 12.1533 13.5058 12.2485 13.1939C16.2627 0.0327312 16.0891 0.605098 16.0963 0.572034ZM12.3532 2.2305L4.96655 7.49106L1.78829 6.26359L12.3532 2.2305ZM5.50531 8.26599L11.944 3.6806C6.4036 9.52539 6.69296 9.21775 6.6688 9.25028C6.6329 9.29857 6.73125 9.11035 5.50531 11.5431V8.26599ZM11.541 12.2785L7.75659 9.46993L14.5993 2.25126L11.541 12.2785Z"
                        fill="white"
                      />
                    </svg>
                  </button>
                </div>
              </div>
            </div>
          </div>

          <div className="BalanceHistoryw-full lg:w-1/2 p-4 lg:p-6">
            <div className="my-6">
              <p className="font-bold text-[#343C6A] text-[16px] md:text-[22px]">
                Balance History
              </p>
            </div>
            <div className="w-auto h-full bg-white rounded-xl">
              <AreaChartComponent />
            </div>
          </div>
        </div>
      </div>
    );
  };
};

export default HomePage;
