"use client";
import React, { useState } from "react";
import { Switch } from "@/components/ui/switch";
import { useForm } from "react-hook-form";

interface FormData {
  timezone: string;
  currency: string;
}

const PrefPage = () => {
  const form = useForm<FormData>();
  const { register, handleSubmit, formState } = form;
  const { errors } = formState;

  const onsubmit = (data: FormData) => {
    console.log(data);
  };

  const Currencies = [
    { value: "USD", label: "United States Dollar - $" },
    { value: "EUR", label: "Euro - €" },
    { value: "GBP", label: "British Pound - £" },
    { value: "JPY", label: "Japanese Yen - ¥" },
    { value: "CNY", label: "Chinese Renminbi - ¥" },
    { value: "AUD", label: "Australian Dollar - AU$" },
    { value: "CAD", label: "Canadian Dollar - CA$" },
    { value: "CHF", label: "Swiss Franc - CHF" },
    { value: "HKD", label: "Hong Kong Dollar - HK$" },
    { value: "INR", label: "Indian Rupee - ₹" },
    { value: "MXN", label: "Mexican Peso - $" },
    { value: "RUB", label: "Russian Ruble - ₽" },
    { value: "SGD", label: "Singapore Dollar - SGD" },
    { value: "ZAR", label: "South African Rand - R" },
    { value: "ETB", label: "Ethiopian Birr - Br" },
  ];
  const timezones = [
    { offset: -12, name: "International Date Line West" },
    { offset: -11, name: "Samoa" },
    { offset: -10, name: "Hawaii" },

    { offset: -9, name: "Alaska" },
    { offset: -8, name: "Pacific Time" },
    { offset: -7, name: "Mountain Time" },
    { offset: -6, name: "Central Time" },
    { offset: -5, name: "Eastern Time" },

    { offset: -4, name: "Atlantic Time" },

    { offset: -3, name: "West Africa Time" },
    { offset: -2, name: "Mid-Atlantic" },
    { offset: -1, name: "Azores" },
    { offset: 0, name: "Greenwich Mean Time" },
    { offset: 1, name: "Central European Time" },
    { offset: 2, name: "Eastern European Time" },
    { offset: 3, name: "Moscow Time" },

    { offset: 4, name: "Gulf Standard Time" },

    { offset: 5, name: "Pakistan Time" },

    { offset: 6, name: "Bangladesh Time" },

    { offset: 7, name: "Krasnoyarsk Time" },
    { offset: 8, name: "China Standard Time" },

    { offset: 9, name: "Japan Standard Time" },

    { offset: 10, name: "Australian Eastern Standard Time" },

    { offset: 11, name: "Solomon Islands" },

    { offset: 12, name: "Kiribati Time" },

    { offset: 13, name: "Phoenix Islands" },
    { offset: 14, name: "Line Islands" },
  ];
  return (
    <form
      onSubmit={handleSubmit(onsubmit)}
      className="flex flex-col mt-10 text-sm space-y-10"
    >
      <div className="flex justify-between">
        <div className="flex flex-col gap-3">
          <div className="text-[#232323]">Currency</div>
          <select
            className="text-[#718EBF] rounded-xl w-[510px]  border border-[#DFEAF2] py-3 px-5"
            {...register("currency", {
              required: {
                value: true,
                message: "Select a currency",
              },
            })}
          >
            <option value="">Select a currency</option>
            {Currencies.map((currency) => (
              <option value={currency.value}>{currency.label}</option>
            ))}
          </select>
          <p
            className="text-red-600 flex font-semibold gap-1
            "
          >
            {errors.currency && (
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                className="size-4"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z"
                />
              </svg>
            )}
            {errors.currency?.message}{" "}
          </p>
        </div>
        <div className="flex flex-col gap-3">
          <div className="text-[#232323]">Time Zone</div>

          <select
            className="text-[#718EBF] rounded-xl w-[510px]  border border-[#DFEAF2] py-3 px-5"
            {...register("timezone", {
              required: {
                value: true,
                message: "Select a timezone",
              },
            })}
          >
            <option value="">Select a currency</option>
            {timezones.map((time) => (
              <option value={time.offset}>{`(${
                time.offset > 0 ? "GMT+" : "GMT-"
              }${Math.abs(time.offset)}) ${time.name}`}</option>
            ))}
          </select>
          <p
            className="text-red-600 flex font-semibold gap-1
            "
          >
            {errors.timezone && (
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                className="size-4"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z"
                />
              </svg>
            )}
            {errors.timezone?.message}{" "}
          </p>
        </div>
      </div>
      <div className="flex flex-col gap-5">
        <div className="font-semibold">Notification</div>
        <div className="flex items-center gap-4">
          <Switch />
          <div>I send or receive digita currency</div>
        </div>
        <div className="flex items-center gap-4">
          <Switch />
          <div>I receive merchant order</div>
        </div>
        <div className="flex items-center gap-4">
          <Switch className="bg-[#16DBCC]" />
          <div>There are recommendation for my account</div>
        </div>
      </div>
      <div className="flex w-full justify-end mt-10  px-[30px] ">
        <button
          type="submit"
          className="px-10 py-3 text-white rounded-xl bg-[#1814F3]"
        >
          Save
        </button>
      </div>
    </form>
  );
};

export default PrefPage;
