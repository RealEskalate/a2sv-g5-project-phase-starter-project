"use client";
import React, { useState } from "react";
import { useForm, SubmitHandler } from "react-hook-form";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import { FaCalendarAlt } from "react-icons/fa";
import { createCard } from "@/services/cardfetch";
import { convertDateToISOString } from "@/lib/utils";
import Image from "next/image";
import { TbFileSad } from "react-icons/tb";

type NewCardProps = {
  cardType: string;
  nameOnCard: string;
  balance: string;
  expiryDate: Date;
  passcode: string;
};

type TokenProp = {
  token: string;
};

const AddNewCard: React.FC<TokenProp> = ({ token }) => {
  const {
    register,
    handleSubmit,
    setValue,
    reset,
    formState: { errors },
  } = useForm<NewCardProps>();

  const [selectedDate, setSelectedDate] = useState<Date | null>(null);
  const [Loading, setLoading] = useState(false);
  const [error, setError] = useState(false);
  const [submissionResult, setSubmissionResult] = useState<string | null>(null);

  const onSubmit: SubmitHandler<NewCardProps> = async (data) => {
    setLoading(true);
    setSubmissionResult(null);
    try {
      const apiData = {
        balance: Number(data.balance),
        cardHolder: data.nameOnCard,
        expiryDate: convertDateToISOString(data.expiryDate),
        passcode: data.passcode,
        cardType: data.cardType,
      };
      const fetch = await createCard(apiData, token);
      setSubmissionResult("Card added successfully!");
      console.log(fetch);
      setError(false);
    } catch (error) {
      console.error("Date conversion error: ", error);
      setSubmissionResult("Failed to add the card. Please try again!");
      setError(true);
    } finally {
      setLoading(false);
    }
  };

  const handleReset = () => {
    setSubmissionResult(null);
    reset();
    setSelectedDate(null);
  };

  if (submissionResult) {
    return (
      <div className="mr-4 lg:w-[800px] w-[330px] md:w-[630px] p-6 border-[1px] m:h-[720px] md:h-[470px] rounded-xl md:px-20 md:py-40 mb-5 flex flex-col items-center justify-center dark:border-[1px] dark:border-gray-700">
        {error ? (
          <></>
          // HEY YO NATI BOY UP THE ERROR COMPONENT HERE 
        ) : (
          <>
            <Image
              src="/icons/check.png"
              width={80}
              height={80}
              alt="null"
              className="pb-2 block"
            />
            <p className="md:text-[20px] text-[14px] font-bold text-center text-black-1 dark:text-gray-400">
              {submissionResult}
            </p>
          <button onClick={handleReset} className="mt-5 px-5 py-2 bg-blue-500 text-white rounded-lg">
            Enter New Card
          </button>
          </>
        )}
        
      </div>
    );
  }

  return (
    <div className={`mr-4 ${Loading ? "animate-pulse opacity-50 pointer-events-none" : ""}`}>
      <div className="bg-white lg:w-[800px] w-[330px] md:w-[630px] sm:h-[720px] md:h-[470px] p-7 dark:border-[1px] dark:border-gray-700 rounded-xl dark:bg-dark dark:text-white">
        <p className="text-[17px] text-[#718EBF]">
          Credit Card generally means a plastic card issued by Scheduled
          Commercial Banks assigned to a Cardholder, with a credit limit, that
          can be used to purchase goods and services on credit or obtain cash
          advances.
        </p>

        <form
          onSubmit={handleSubmit(onSubmit)}
          className="grid grid-cols-1 md:grid-cols-2 gap-4 mt-4 dark:bg-dark text-gray-900 dark:text-white"
        >
          <div>
            <label htmlFor="cardTypeId" className="text-[16px] block pb-2">
              Card Type
            </label>
            <input
              id="cardTypeId"
              {...register("cardType", { required: "Card Type is required" })}
              placeholder="Classic"
              className="border-[1px] border-[#DFEAF2] rounded-md text-[15px] p-3 w-full outline-none text-[#718EBF] dark:text-white placeholder-[#718EBF]"
            />
            {errors.cardType && (
              <span className="text-red-500 text-sm">
                {errors.cardType.message}
              </span>
            )}
          </div>

          <div>
            <label htmlFor="nameOneCardId" className="text-[16px] block pb-2">
              Name On Card
            </label>
            <input
              id="nameOnCardId"
              {...register("nameOnCard", {
                required: "Name on Card is required",
              })}
              placeholder="My Cards"
              className="border-[1px] border-[#DFEAF2] rounded-md text-[15px] p-3 w-full outline-none text-[#718EBF] dark:text-white placeholder-[#718EBF]"
            />
            {errors.nameOnCard && (
              <span className="text-red-500 text-sm">
                {errors.nameOnCard.message}
              </span>
            )}
          </div>

          <div>
            <label htmlFor="balanceId" className="text-[16px] block pb-2">
              Balance
            </label>
            <input
              id="balanceId"
              {...register("balance", {
                required: "Balance is required",
                pattern: {
                  value: /^\d+$/,
                  message: "only numbers are inserted in this field",
                },
              })}
              placeholder="27,000$"
              className="border-[1px] border-[#DFEAF2] rounded-md text-[15px] p-3 w-full outline-none text-[#718EBF] dark:text-white placeholder-[#718EBF]"
            />
            {errors.balance && (
              <span className="text-red-500 text-sm">
                {errors.balance.message}
              </span>
            )}
          </div>

          <div className="relative">
            <label
              htmlFor="expirationDateId"
              className="text-[16px] block pb-2"
            >
              Expiration Date
            </label>
            <div className="relative">
              <DatePicker
                id="expirationDateId"
                selected={selectedDate}
                onChange={(date) => {
                  if (date) {
                    // Check if the date is not null
                    setSelectedDate(date);
                    setValue("expiryDate", date, { shouldValidate: true });
                  }
                }}
                placeholderText="dd MMMM yyyy"
                className="border-[1px] border-[#DFEAF2] rounded-md text-[15px] p-3 w-full pr-40 outline-none text-[#718EBF] dark:text-white placeholder-[#718EBF]"
                dateFormat="dd MMMM yyyy"
              />
              <FaCalendarAlt className="absolute right-3 top-1/2 transform -translate-y-1/2 pointer-events-none text-[#718EBF] dark:text-white" />
            </div>
            {errors.expiryDate && (
              <span className="text-red-500 text-sm">
                {errors.expiryDate?.message}
              </span>
            )}
          </div>

          <div>
            <label htmlFor="passcodeId" className="text-[16px] block pb-2">
              Passcode
            </label>
            <input
              id="passcodeId"
              {...register("passcode", { required: "Passcode is required" })}
              placeholder="******"
              className="border-[1px] border-[#DFEAF2] rounded-md text-[15px] p-3 w-full outline-none text-[#718EBF] dark:text-white placeholder-[#718EBF]"
            />
            {errors.passcode && (
              <span className="text-red-500 text-sm">
                {errors.passcode.message}
              </span>
            )}
          </div>

          <div className="md:mt-4 md:ml-36 lg:ml-56">
            <button
              type="submit"
              className="rounded-xl text-[16px] px-7 text-center bg-[#1814F3] dark:bg-blue-700 text-white w-[95%] md:w-[auto] mt-4 p-2"
            >
              Add Card
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default AddNewCard;
