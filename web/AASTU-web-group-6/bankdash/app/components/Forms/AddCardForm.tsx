"use client";
import { useAppDispatch } from "@/app/Redux/store/store";
import React from "react";
import { useForm } from "react-hook-form";
import {
  addCard,
  setStatus,
  setError,
  Card,
} from "../../Redux/slices/cardSlice";
import CardService from "@/app/Services/api/CardService";
import { useSession } from "next-auth/react";

const AddCardForm = () => {
  const { data: session } = useSession();
  const form = useForm<Card>();
  const { register, handleSubmit, formState } = form;
  const { errors } = formState;
  const dispatch = useAppDispatch();
  const accessToken = session?.accessToken as string;
  const onSubmit = async (data: Card) => {
    const cardData = {
      balance: data.balance,
      cardHolder: data.cardHolder,
      expiryDate: data.expiryDate,
      passcode: "123456",
      cardType: data.cardType,
    };

    try {
      dispatch(setStatus("loading"));
      const res: any = await CardService.addCard(cardData, accessToken);
      if (res) {
        dispatch(addCard(res));
        dispatch(setStatus("succeeded"));
      }
    } catch (error) {
      dispatch(setError("Failed to fetch cards"));
      dispatch(setStatus("failed"));
    }
  };

  return (
    <div className="p-4 sm:p-8">
      <form
        className="w-full max-w-4xl mx-auto"
        onSubmit={handleSubmit(onSubmit)}
        noValidate
      >
        <div className="flex flex-col gap-6 sm:flex-row sm:gap-10">
          <div className="flex flex-col w-full sm:w-1/2">
            <label className="text-[#232323] text-base font-semibold pb-2 dark:text-gray-300">
              Card Type
            </label>
            <input
              className="dark:bg-gray-700 border-solid border-2 border-blue-50 dark:border-gray-600 placeholder:text-base placeholder:text-[#8BA3CB] dark:placeholder:text-gray-400 focus:outline-none focus:border-blue-200 dark:focus:border-gray-500 p-2 px-3 rounded-lg w-full"
              type="text"
              placeholder="Classic"
              {...register("cardType", { required: "Card type is required" })}
            />
            {errors.cardType && (
              <span className="text-red-500 text-sm">
                {errors.cardType.message}
              </span>
            )}
          </div>
          <div className="flex flex-col w-full sm:w-1/2">
            <label className="text-[#232323] text-base font-semibold pb-2 dark:text-gray-300">
              Balance
            </label>
            <input
              className="dark:bg-gray-700 border-solid border-2 border-blue-50 dark:border-gray-600 placeholder:text-base placeholder:text-[#8BA3CB] dark:placeholder:text-gray-400 focus:outline-none focus:border-blue-200 dark:focus:border-gray-500 p-2 px-3 rounded-lg w-full"
              type="number"
              placeholder="0"
              {...register("balance", {
                required: "Balance is required",
                min: { value: 0, message: "Balance must be at least 0" },
              })}
            />
            {errors.balance && (
              <span className="text-red-500 text-sm">
                {errors.balance.message}
              </span>
            )}
          </div>
        </div>
        <div className="flex flex-col gap-6 mt-6 sm:flex-row sm:gap-10">
          <div className="flex flex-col w-full sm:w-1/2">
            <label className="text-[#232323] text-base font-semibold pb-2 dark:text-gray-300">
              Name on Card
            </label>
            <input
              className="dark:bg-gray-700 border-solid border-2 border-blue-50 dark:border-gray-600 placeholder:text-base placeholder:text-[#8BA3CB] dark:placeholder:text-gray-400 focus:outline-none focus:border-blue-200 dark:focus:border-gray-500 p-2 px-3 rounded-lg w-full"
              type="text"
              placeholder="John Doe"
              {...register("cardHolder", {
                required: "Cardholder's name is required",
              })}
            />
            {errors.cardHolder && (
              <span className="text-red-500 text-sm">
                {errors.cardHolder.message}
              </span>
            )}
          </div>
          <div className="flex flex-col w-full sm:w-1/2">
            <label className="text-[#232323] text-base font-semibold pb-2 dark:text-gray-300">
              Expiration Date
            </label>
            <input
              className="dark:bg-gray-700 border-solid border-2 border-blue-50 dark:border-gray-600 placeholder:text-base placeholder:text-[#8BA3CB] dark:placeholder:text-gray-400 focus:outline-none focus:border-blue-200 dark:focus:border-gray-500 p-2 px-3 rounded-lg w-full"
              type="date"
              {...register("expiryDate", {
                required: "Expiration date is required",
              })}
            />
            {errors.expiryDate && (
              <span className="text-red-500 text-sm">
                {errors.expiryDate.message}
              </span>
            )}
          </div>
        </div>
        <button
          type="submit"
          className="mt-6 px-8 py-3 bg-[#1814F3] text-white text-lg font-semibold rounded-[9px] w-full"
        >
          Add Card
        </button>
      </form>
    </div>
  );
};

export default AddCardForm;
