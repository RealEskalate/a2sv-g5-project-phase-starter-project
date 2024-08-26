"use client";
import React from "react";
import Image from "next/image";
import x from "../../../public/assets/next-icon.svg";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faPaperPlane } from "@fortawesome/free-regular-svg-icons";
import { AreaComp } from "../Charts/AreaComp";
import { useAppSelector } from "@/app/Redux/store/store";
import { useState, useEffect, useRef } from "react";
import {
  faArrowRight,
  faGreaterThan,
  faLessThan,
} from "@fortawesome/free-solid-svg-icons";
import { BalanceType } from "@/app/Redux/slices/TransactionSlice";
import ModalTrans from "../Forms/SendMoneyForm";
import TransactionService from "@/app/Services/api/transactionApi";
import { useSession } from "next-auth/react";
import ShimmerProfileItem from "../Shimmer/ShimmerProfileItem";
interface quickType {
  id: string;
  name: string;
  username: string;
  city: string;
  country: string;
  profilePicture: string;
}

const Bottom = () => {
  const containerRef = useRef<HTMLDivElement>(null);
  const { data: session } = useSession();
  const handleScroll = () => {
    if (containerRef.current) {
      containerRef.current.scrollLeft += 200;
      console.log("scrolled");
    } else {
      console.log("not scrolled");
    }
  };
  const people: any = [1, 2, 3];
  const BalanceData: BalanceType[] = useAppSelector(
    (state) => state.transactions.balanceHist
  );
  const [isModalOpen, setIsModalOpen] = useState(false);

  const handleModalToggle = () => {
    setIsModalOpen(!isModalOpen);
  };
  const [quickData, setQuickData] = useState<quickType[]>([]);
  const [username, setUsername] = useState<string>("");
  const [amount, setAmount] = useState<number>(0);
  const accessToken = session?.accessToken as string;

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await TransactionService.getQuickTransfer(accessToken);

        setQuickData(data);
        console.log(quickData, "quick");
      } catch (error) {
        console.log("error", error);
        console.error("Error fetching quick data:", error);
      }
    };

    fetchData();
  }, [accessToken]);

  return (
    <section className="Botom flex gap-6 xs:flex-col lg:flex-row ">
      <div className="cards-container sm:w-full lg:w-[45%]  center-content flex flex-col gap-6">
        <h1 className="text-xl font-semibold text-colorBody-1 dark:text-gray-300">
          Quick Transfer
        </h1>
        <div className="flex gap-6 bg-white dark:bg-[#232328] rounded-3xl  p-6">
          <div className="profle-box w-full flex flex-col gap-4">
            <div className="w-full flex gap-2 items-center   ">
              <div className=" flex overflow-hidden" ref={containerRef}>
                {quickData.length > 0 ? (
                  quickData.map((account) => (
                    <button
                      key={account.id}
                      className={`profile-item flex flex-col gap-1 p-6 items-center justify-center   dark:text-gray-300 ${
                        username == account.username
                          ? `border-2 border-solid rounded-[30px] border-blue-600 `
                          : ""
                      }`}
                      onClick={() => {
                        setUsername(account.username);
                      }}
                    >
                      <Image
                        className="rounded-full"
                        src={"/assets/profile-1.png"}
                        width={70}
                        height={70}
                        alt={account.name}
                      />
                      <div className="name text-base font-semibold">
                        {account.name}
                      </div>
                      <div className="role text-base text-[#718EBF]">CEO</div>
                    </button>
                  ))
                ) : (
                  <div className="flex gap-4">
                    <ShimmerProfileItem />
                    <ShimmerProfileItem />
                    <ShimmerProfileItem />
                  </div>
                )}
              </div>
              <button className="relative flex p-6 py-7 items-center justify-center bg-white dark:bg-gray-700 dark:shadow-gray-500 shadow-sm shadow-blue-300 rounded-full">
                <FontAwesomeIcon
                  icon={faGreaterThan}
                  className="w-5 font-normal dark:text-gray-300"
                  onClick={handleScroll}
                />
              </button>
            </div>

            <div className="bottom flex gap-4 items-center">
              <div className="text-gray- text-lg text-[#718EBF] dark:text-gray-400">
                Write Amount
              </div>

              <div
                className={`flex items-center text-base text-[#718EBF] bg-[#EDF1F7] dark:bg-gray-700 dark:text-gray-400 rounded-[50px] py-1 pl-6 grow justify-end  
                  `}
              >
                <input
                  className="flex w-full grow bg-[#EDF1F7] dark:bg-transparent  focus:outline-none"
                  onChange={(e) => setAmount(Number(e.target.value))}
                ></input>
                <button
                  onClick={handleModalToggle}
                  className="flex gap-2 w-full grow p-4 text-white bg-[#1814f6] hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-[50px] text-sm px-6  text-center dark:hover:bg-blue-700 dark:focus:ring-blue-800 text-6 items-center justify-center"
                  type="button"
                >
                  Send
                  <FontAwesomeIcon
                    icon={faPaperPlane}
                    className="w-5 font-normal"
                  />
                </button>
              </div>
              {isModalOpen && (
                <div
                  className="fixed inset-0 z-50 flex justify-center items-center bg-black bg-opacity-50 backdrop-blur-sm"
                  onClick={handleModalToggle}
                >
                  <div
                    className="relative bg-white p-6 rounded-lg shadow-lg max-w-md w-full"
                    onClick={(e) => e.stopPropagation()}
                  >
                    <ModalTrans
                      isOpen={isModalOpen}
                      onClose={handleModalToggle}
                      userName={username}
                      amount={amount}
                      accessToken={accessToken}
                    />
                  </div>
                </div>
              )}
            </div>
          </div>
        </div>
      </div>
      <div className="cards-container center-content flex flex-col gap-6 xs:w-full lg:w-[55%]">
        <h1 className="page text-xl font-semibold text-colorBody-1 dark:text-gray-300">
          Balance History
        </h1>

        <div className="flex w-full gap-6 p-8 bg-white dark:bg-[#232328] dark:border-gray-500 rounded-3xl border-solid border-gray-200 border-[0.5px] shadow-sm xs:w-[85%] sm:w-full">
          <div className="leftCanva pb-10 flex flex-col items-end justify-between text-sm text-[#718EBF]">
            <span>400</span>
            <span>300</span>
            <span>200</span>
            <span>100</span>
            <span>0</span>
          </div>
          <AreaComp data={BalanceData} />
        </div>
      </div>
    </section>
  );
};

export default Bottom;
