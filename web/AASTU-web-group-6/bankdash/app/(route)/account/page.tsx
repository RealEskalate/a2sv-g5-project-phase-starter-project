"use client";
import Card from "../../components/Accounts/account";
import LastTransList from "@/app/components/Accounts/LastTransList";
import { DebitCreditOver } from "../../components/Accounts/DebitCreditOver";
import InvoiceCard from "../../components/Accounts/InvoiceCard";
import VisaCard from "@/app/components/Card/VisaCard";
import { Card as CardType } from "@/app/Redux/slices/cardSlice";
import { useAppSelector } from "@/app/Redux/store/store";
import LastTrans from "@/app/components/Accounts/Last_trans";
import { TransactionType } from "@/app/Redux/slices/TransactionSlice";
import UserService from "@/app/Services/api/userService";
import { useSession } from "next-auth/react";
import { useEffect, useState } from "react";
import { ShimmerVisaCard } from "@/app/components/Shimmer/ShimmerVisa";
export default function Home() {
  const { data: session } = useSession();
  const accessToken = session?.accessToken as string;
  const CardData: CardType[] = useAppSelector((state) => state.cards.cards);
  const [totalBalance, setTotalBalance] = useState<number>(0);
  const formattedAmount = `$${totalBalance.toLocaleString()}`;
  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await UserService.current(accessToken);
        setTotalBalance(data.data.accountBalance);
        console.log(data.data.accountBalance, "current balance");
      } catch (error) {
        console.log("error", error);
        console.error("Error fetching current  data:", error);
      }
    };

    fetchData();
  }, [accessToken]);
  return (
    <div className="w-[96%] flex flex-col grow gap-6 p-5 lg:p-8 pt-11 md:pt-6 ">
      <div className="flex flex-col lg:flex-row gap-6 xl:gap-7 pt-0">
        <div className="flex w-full gap-4 xl:gap-7 flex-col md:flex-row">
          <div className="flex gap-7 w-full">
            <Card
              title="My Balance"
              amount={formattedAmount}
              color="#FFF5D9"
              icon="/assets/money-tag 1.svg"
              width="w-full"
            />
            <Card
              title="Income"
              amount="$5,600"
              color="#E7EDFF"
              icon="/assets/expense.svg"
              width="w-full"
            />
          </div>
          <div className="flex gap-7 w-full">
            <Card
              title="Expense"
              amount="$3,460"
              color="#FFE0EB"
              icon="/assets/income.svg"
              width="w-full"
            />
            <Card
              title="Total Saving"
              amount="$7,920"
              color="#DCFAF8"
              icon="/assets/saving.svg"
              width="w-full"
            />
          </div>
        </div>
      </div>

      <div className="flex  flex-col lg:flex-row my-5 justify-between">
        <div className="lg:w-[65%]">
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5 dark:text-gray-300">
            Last Transaction
          </p>
          <div className=" bg-white dark:bg-[#232328]  rounded-3xl p-3 shadow-lg border-gray-300 sm:text-[12px]">
            <LastTransList />
          </div>
        </div>
        <div className="lg:w-[30%] lg:h-[250px]">
          <div className="flex justify-between">
            <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5 dark:text-gray-300">
              My Card
            </p>
            <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5 dark:text-gray-300">
              See All
            </p>
          </div>
          <>
            {CardData.length > 0 ? (
              CardData?.slice(0, 1).map((item, index) => (
                <VisaCard
                  key={index}
                  data={item}
                  isBlack={false}
                  isFade={false}
                  isSimGray={false}
                />
              ))
            ) : (
              <div className="w-full flex gap-6 ">
                <ShimmerVisaCard />
              </div>
            )}
          </>
        </div>
      </div>
      <div className="flex flex-col lg:flex-row justify-between my-5">
        <div className="lg:w-[65%] lb:h-[364px]">
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5 dark:text-gray-300">
            Debit & Credit Overview
          </p>
          <DebitCreditOver />
        </div>
        <div className="lg:w-[30%]">
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5 dark:text-gray-300">
            Invoices Sent
          </p>
          <div className="rounded-3xl p-9 lg:p-3 xl:p-9  bg-white dark:bg-[#232328]">
            <InvoiceCard
              title="Apple Store"
              date="5h ago"
              amount="$450"
              icon="/assets/apple-icon.svg"
              color="#DCFAF8"
            />
            <InvoiceCard
              title="Michael"
              date="2 days ago"
              amount="$160"
              icon="/assets/userr.svg"
              color="#FFF5D9"
            />
            <InvoiceCard
              title="Playstation"
              date="5 days ago"
              amount="$1085"
              icon="/assets/Group-m.svg"
              color="#E7EDFF"
            />
            <InvoiceCard
              title="William"
              date="10 days ago"
              amount="$90"
              icon="/assets/userr.svg"
              color="#FFE0EB"
            />
          </div>
        </div>
      </div>
    </div>
  );
}
