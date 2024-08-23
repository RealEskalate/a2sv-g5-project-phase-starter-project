import React from "react";
import { useForm } from "react-hook-form";
import { PostCardRequest } from "../../types/cardController.Interface";
import { postCard } from "../../lib/api/cardController";

export const InputLabel = ({ label }: { label: string }) => {
  return (
    <label htmlFor="" className="text-sm lg:text-base">
      {label}
    </label>
  );
};

const AddCardForm = ({
  access_token,
  handleAddition,
}: {
  access_token: string;
  handleAddition: Function;
}) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm<PostCardRequest>();

  const onSubmit_ = handleSubmit(async (data) => {
    console.log(data.passcode);
    data.passcode = "192930";
    const response = await postCard(data, access_token);
    handleAddition(response);
    reset();
  });

  return (
    <div className="bg-white p-5 flex-col gap-7 rounded-xl flex shadow h-full dark:bg-[#050914] dark:border dark:border-[#333B69]">
      <p className="text-xs text-[#718EBF] leading-5 lg:text-base">
        Credit Card generally means a plastic card issued by Scheduled
        Commercial Banks assigned to a Cardholder, with a credit limit, that can
        be used to purchase goods and services on credit or obtain cash
        advances.
      </p>

      <form
        action=""
        className="flex flex-col items-start gap-4"
        onSubmit={onSubmit_}
      >
        <div className="flex flex-col md:flex-row gap-4 w-full">
          <div className="flex flex-col gap-2">
            <InputLabel label="Card Type" />
            <input
              type="text"
              placeholder="Card Type"
              className={`border px-3 py-3 rounded-xl lg:w-80 text-xs lg:text-base dark:bg-[#313245]
 dark:border-[#333B69] 
${errors.cardType ? "border-red-500" : "border-indigo-50"}`}
              {...register("cardType", {
                required: "Card Type is required",
                minLength: {
                  value: 3,
                  message: "Card Type must be at least 3 characters",
                },
              })}
            />
            {errors.cardType && (
              <span className={`text-red-600 text-xs`}>
                {errors.cardType.message}
              </span>
            )}
          </div>
          <div className="flex flex-col gap-2">
            <InputLabel label="Name On Card" />
            <input
              type="text"
              placeholder="My Cards"
              className={`border px-3 py-3 rounded-xl lg:w-80 text-xs lg:text-base dark:bg-[#313245]
 dark:border-[#333B69]  ${
   errors.cardHolder ? "border-red-500" : "border-indigo-50"
 }`}
              {...register("cardHolder", {
                required: "Card Holder name is required",
                minLength: {
                  value: 3,
                  message: "Name must be at least 3 characters",
                },
              })}
            />
            {errors.cardHolder && (
              <span className="text-red-600 text-xs">
                {errors.cardHolder.message}
              </span>
            )}
          </div>
        </div>
        <div className="flex flex-col md:flex-row gap-4 w-full">
          <div className="flex flex-col gap-2">
            <InputLabel label="Balance" />
            <input
              type="number"
              placeholder="27,000$"
              className={`border px-3 py-3 rounded-xl lg:w-80 text-xs lg:text-base dark:bg-[#313245]
 dark:border-[#333B69] ${
   errors.balance ? "border-red-500" : "border-indigo-50"
 }`}
              {...register("balance", {
                required: "Balance is required",
                min: {
                  value: 0,
                  message: "Balance cannot be negative",
                },
                validate: (value) =>
                  value <= 1000000 || "Balance cannot exceed $1,000,000",
              })}
            />
            {errors.balance && (
              <span className="text-red-600 text-xs">
                {errors.balance.message}
              </span>
            )}
          </div>
          <div className="flex flex-col gap-2 w-full md:w-auto">
            <InputLabel label="Expiration Date" />
            <input
              type="date"
              className={`border px-3 py-3 rounded-xl lg:w-80 text-xs lg:text-base dark:bg-[#313245]
 dark:border-[#333B69] text-slate-500 ${
   errors.expiryDate ? "border-red-500" : "border-indigo-50"
 }`}
              {...register("expiryDate", {
                required: "Expiration Date is required",
                validate: (value) => {
                  const today = new Date();
                  const selectedDate = new Date(value);
                  return (
                    selectedDate > today ||
                    "Expiration date must be in the future"
                  );
                },
              })}
            />
            {errors.expiryDate && (
              <span className="text-red-600 text-xs">
                {errors.expiryDate.message}
              </span>
            )}
          </div>
        </div>

        <button
          type="submit"
          className="bg-[#1814F3] hover:bg-[#423fef] text-white px-5 py-3 rounded-xl w-full lg:w-auto text-sm"
        >
          Add Card
        </button>
      </form>
    </div>
  );
};

export default AddCardForm;
