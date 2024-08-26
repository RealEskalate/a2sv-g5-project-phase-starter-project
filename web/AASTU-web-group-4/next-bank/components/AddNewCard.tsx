"use client";
import React, { useState } from "react";
import { useForm, SubmitHandler } from "react-hook-form";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import { FaCalendarAlt } from "react-icons/fa";
import { createCard } from "@/services/cardfetch";
import { convertDateToISOString } from "@/lib/utils";
import Image from "next/image";
import { message } from "antd";

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
  const [loading, setLoading] = useState(false);
  const [messageApi, contextHolder] = message.useMessage();

  const onSubmit: SubmitHandler<NewCardProps> = async (data) => {
    setLoading(true);
    try {
      const apiData = {
        balance: Number(data.balance),
        cardHolder: data.nameOnCard,
        expiryDate: convertDateToISOString(data.expiryDate),
        passcode: data.passcode,
        cardType: data.cardType,
      };
      const fetch = await createCard(apiData, token);
      if (fetch) {
        messageApi.open({
          type: "success",
          content: "Successfully created your new card",
          duration: 4,
        });
        setSelectedDate(null); // Reset selected date
        reset(); // Reset the form after successful submission
      }
    } catch (error) {
      messageApi.open({
        type: "error",
        content: "Card creation was not successful. Please try again.",
        duration: 4,
      });
    } finally {
      setLoading(false);
    }
  };

  return (
    <>
      {contextHolder}
      <div
        className={`mr-4 ${
          loading ? "animate-pulse opacity-50 pointer-events-none" : ""
        }`}
      >
        <div className="bg-white 2xl:w-[800px] xl:w-[600px] lg:w[600px] w-[330px] md:w-[400px] sm:h-[880px] md:h-[520px] lg:h-[470px] p-7 dark:border-[1px] dark:border-gray-700 rounded-xl dark:bg-dark dark:text-white">
          <p className="text-[17px] md:text-[15px] text-[#718EBF]">
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
              <label
                htmlFor="cardTypeId"
                className="text-[16px] md:text-[15px] block pb-2"
              >
                Card Type
              </label>
              <input
                id="cardTypeId"
                {...register("cardType", {
                  required: "Card Type is required",
                })}
                placeholder="Classic"
                className="border-[1px] border-[#DFEAF2] rounded-md text-[15px] md:text-[14px] p-3 w-full outline-none text-[#718EBF] dark:text-white placeholder-[#718EBF]"
              />
              {errors.cardType && (
                <span className="text-red-500 text-sm">
                  {errors.cardType.message}
                </span>
              )}
            </div>

            <div>
              <label
                htmlFor="nameOneCardId"
                className="text-[16px] md:text-[15px] block pb-2"
              >
                Name On Card
              </label>
              <input
                id="nameOnCardId"
                {...register("nameOnCard", {
                  required: "Name on Card is required",
                })}
                placeholder="My Cards"
                className="border-[1px] border-[#DFEAF2] rounded-md text-[15px] md:text-[14px] p-3 w-full outline-none text-[#718EBF] dark:text-white placeholder-[#718EBF]"
              />
              {errors.nameOnCard && (
                <span className="text-red-500 text-sm">
                  {errors.nameOnCard.message}
                </span>
              )}
            </div>

            <div>
              <label
                htmlFor="balanceId"
                className="text-[16px] md:text-[15px] block pb-2"
              >
                Balance
              </label>
              <input
                id="balanceId"
                {...register("balance", {
                  required: "Balance is required",
                  pattern: {
                    value: /^\d+$/,
                    message: "Only numbers are allowed in this field",
                  },
                })}
                placeholder="27,000$"
                className="border-[1px] border-[#DFEAF2] rounded-md text-[15px] md:text-[14px] p-3 w-full outline-none text-[#718EBF] dark:text-white placeholder-[#718EBF]"
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
                className="text-[16px] md:text-[15px] block pb-2"
              >
                Expiration Date
              </label>
              <div className="relative">
                <DatePicker
                  id="expirationDateId"
                  selected={selectedDate}
                  onChange={(date) => {
                    if (date) {
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
              <label
                htmlFor="passcodeId"
                className="text-[16px] md:text-[15px] block pb-2"
              >
                Passcode
              </label>
              <input
                id="passcodeId"
                {...register("passcode", {
                  required: "Passcode is required",
                })}
                placeholder="******"
                className="border-[1px] border-[#DFEAF2] rounded-md text-[15px] md:text-[14px] p-3 w-full outline-none text-[#718EBF] dark:text-white placeholder-[#718EBF]"
              />
              {errors.passcode && (
                <span className="text-red-500 text-sm">
                  {errors.passcode.message}
                </span>
              )}
            </div>

            <div className="md:mt-2 lg:mt-4 md:ml-10 2xl:ml-56">
              <button
                type="submit"
                className="rounded-xl text-[16px] px-7 md:px-4 xl:px-7 text-center bg-[#1814F3] dark:bg-blue-700 text-white w-[95%] md:w-[auto] mt-4 py-2"
              >
                Add Card
              </button>
            </div>
          </form>
        </div>
      </div>
    </>
  );
};

export default AddNewCard;
