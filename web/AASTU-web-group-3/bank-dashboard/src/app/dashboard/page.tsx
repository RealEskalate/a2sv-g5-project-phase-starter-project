"use client";
import React, { useEffect } from "react";
import Image from "next/image";
import {
  useCreateTransactionMutation,
  useGetQuickTransfersQuery,
  useGetAllTransactionsQuery,
} from "@/lib/redux/api/transactionsApi";
import { Transaction } from "@/lib/redux/types/transactions";
import { useDispatch, useSelector } from "react-redux";
import dollar from "../../../public/images/iconfinder_6_4753731 1.png";
import Link from "next/link";
import { useGetCardsQuery } from "@/lib/redux/api/cardsApi";
import { RootState } from "@/lib/redux/store";
import { setCards, setLoading, setError } from "@/lib/redux/slices/cardsSlice";
import CreditCard from "../components/CreditCard";
import Loading from "../loading";
import { cardStyles } from "@/lib/CardColor";
import { BarChartComponent } from "../components/Chart/Barchart";
import { PieChartComponent } from "../components/Chart/PieChart";
import { AreaChartComponent } from "../components/Chart/AreaChartComponent";
import { Card } from "../../lib/redux/types/cards";
import { toast, ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

// import { Transaction } from "@/lib/redux/types/transactions";

// import { useGetQuickTransfersQuery, useCreateTransactionMutation } from "@/lib/redux/api/transactionsApi";
interface User {
  username: string;
}
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

const HomePage = () => {
  const { data: transactionsData, isLoading: isTransactionsLoading } =
    useGetAllTransactionsQuery({
      size: 3,
      page: 0,
    });
  const { data: quickTransfersData, isLoading: isTransfersLoading } =
    useGetQuickTransfersQuery({
      num: 3,
    });

  const [selectedUser, setSelectedUser] = React.useState<User | null>(null);
  const [amount, setAmount] = React.useState("");
  const [createTransaction, { isLoading: isCreatingTransaction }] =
    useCreateTransactionMutation();

  const handleTransferClick = (user: any) => {
    setSelectedUser(user);
    setAmount(""); // Reset the amount input when a new user is selected
  };

  const handleSend = async () => {
    if (!selectedUser || !amount) return;

    const transactionData = {
      type: "transfer",
      description: "Quick transfer",
      amount: parseFloat(amount),
      receiverUserName: selectedUser.username,
    };
    try {
      const response = await createTransaction(transactionData).unwrap();
      if (response.success) {
        toast.success(`Transaction successful: ${response.message}`);
      } else {
        toast.error(`Transaction failed: ${response.message}`);
      }
    } catch (error) {
      toast.error(`Transaction failed: ${error}`);
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

  if (isTransactionsLoading || isTransfersLoading || cardsLoading)
    return <Loading />;
  if (error) return <div>Error fetching cards</div>;

  return (
    <div className="bg-[#F5F7FA] min-h-screen p-5 dark:bg-darkPage">
      <ToastContainer />
      <div className="lg:flex lg:justify-between">
        <div className="lg:w-[60%] rounded-xl bg-[#F5F7FA]  dark:bg-darkPage">
          <div className="credit-card-info flex justify-between  h-16 items-center ">
            <h3 className="font-bold text-[#343C6A] text-[16px] md:text-[22px] dark:text-white">
              My cards
            </h3>
            <Link href="/creditcardpage#add-new-card">
              <h3 className="font-bold text-[#343C6A] text-[16px] md:text-[22px]  dark:text-blue-600">
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

        <div className="lg:w-[35%] lg:pt-3 lg:pb-2 ">
          <div>
            <p className="font-bold text-[#343C6A] text-[16px] md:text-[22px] mb-3 dark:text-white">
              Recent Transaction
            </p>
          </div>
          <div className="transactionData mt-5  h-[210px] lg:mt-3 rounded-lg lg:h-[215px] bg-white lg:p-5 lg:rounded-lg lg:shadow-md dark:bg-darkComponent">
            {transactionsData?.data?.content?.map((transaction, index) => (
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
            <p className="font-bold text-[#343C6A] text-[16px] md:text-[22px] mb-3 dark:text-white">
              Weekly Activities
            </p>
          </div>
          <BarChartComponent />
        </div>

        <div className="expense my-5 rounded-lg h-auto lg:p-0 lg:m-0">
          <div>
            <p className="font-bold text-[#343C6A] text-[16px] md:text-[22px] mb-3 dark:text-white">
              Expenses Statistics
            </p>
          </div>
          <PieChartComponent />
        </div>
      </div>

      <div className="quickandBalance flex flex-col lg:flex-row items-start lg:items-center lg:justify-between gap-6 lg:gap-8  lg:w-full">
        <div className="quicktransfer w-full lg:w-1/2 p-4 lg:p-6 lg:h-[450px]">
          <div className="my-6">
            <p className="font-bold text-[#343C6A] text-[16px] md:text-[22px] dark:text-white">
              Quick Transfer
            </p>
          </div>
          <div className="w-full bg-white dark:bg-darkBackground p-5 rounded-xl shadow-md">
            <div className="flex justify-center">
              <div className="flex mt-3 space-x-16">
                {quickTransfersData?.data.map((user, index) => (
                  <div
                    key={index}
                    className="flex flex-col items-center space-y-2"
                    onClick={() => handleTransferClick(user)}
                  >
                    <div
                      className={`relative w-12 h-12 md:w-16 md:h-16 rounded-full cursor-pointer ${
                        selectedUser?.username === user.username
                          ? "border-2 border-blue-500"
                          : ""
                      }`}
                    >
                      <Image
                        src="https://via.placeholder.com/48"
                        alt={user.username}
                        className="rounded-full object-cover"
                        fill
                      />
                    </div>
                    <p className="text-xs text-center">{user.username}</p>
                  </div>
                ))}
              </div>
            </div>
            <div className="mt-6">
              <input
                type="number"
                placeholder="Enter amount"
                value={amount}
                onChange={(e) => setAmount(e.target.value)}
                className="w-full px-3 py-2 text-center border border-gray-300 rounded-md"
              />
              <button
                className="w-full py-2 mt-4 text-white bg-blue-500 rounded-md"
                onClick={handleSend}
                disabled={isCreatingTransaction || !selectedUser || !amount}
              >
                Send
              </button>
            </div>
          </div>
        </div>

        <div className="balance w-full lg:w-1/2 p-4 lg:p-6 lg:h-[450px]">
          <div className="my-6">
            <p className="font-bold text-[#343C6A] text-[16px] md:text-[22px] dark:text-white">
              Balance Statistics
            </p>
          </div>
          <AreaChartComponent />
        </div>
      </div>
    </div>
  );
};

export default HomePage;
