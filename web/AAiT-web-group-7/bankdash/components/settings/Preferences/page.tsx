"use client";
import React, { useState } from "react";
import { Switch } from "@/components/ui/switch";
import { useForm, Controller } from "react-hook-form";
import ErrorMessage from "@/components/Message/ErrorMessage";
import { Currencies } from "@/components/constants/currency";
import { timezones } from "@/components/constants/timezones";
interface FormData {
  timezone: string;
  currency: string;
  transaction: boolean;
  merchant: boolean;
  recommendation: boolean;
  twoFactorAuth: boolean;
}

const PrefPage = () => {
  const form = useForm<FormData>({
    defaultValues: {
      currency: "",
      timezone: "",
      transaction: false,
      merchant: false,
      recommendation: false,
      twoFactorAuth: false,
    },
  });
  const { control, handleSubmit, formState } = form;
  const { errors } = formState;

  const onsubmit = (data: FormData) => {
    console.log(data);
  };

  return (
    <form
      onSubmit={handleSubmit(onsubmit)}
      className="flex flex-col mt-10 text-sm space-y-10"
    >
      <div className="flex justify-between">
        <div className="flex flex-col gap-3">
          <div className="text-[#232323]">Currency</div>
          <Controller
            name="currency"
            control={control}
            rules={{ required: "Select a currency" }}
            render={({ field }) => (
              <select
                {...field}
                className="text-[#718EBF] rounded-xl w-[510px] border border-[#DFEAF2] py-3 px-5"
              >
                <option value="">Select a currency</option>
                {Currencies.map((currency) => (
                  <option key={currency.value} value={currency.value}>
                    {currency.label}
                  </option>
                ))}
              </select>
            )}
          />

          <ErrorMessage message={errors.currency?.message} />
        </div>
        <div className="flex flex-col gap-3">
          <div className="text-[#232323]">Time Zone</div>

          <Controller
            name="timezone"
            control={control}
            rules={{ required: "Select a timezone" }}
            render={({ field }) => (
              <select
                {...field}
                className="text-[#718EBF] rounded-xl w-[510px] border border-[#DFEAF2] py-3 px-5"
              >
                <option value="">Select a timezone</option>
                {timezones.map((time) => (
                  <option key={time.name} value={time.offset}>
                    {`(${
                      time.offset === 0
                        ? "GMT"
                        : time.offset > 0
                        ? "GMT+"
                        : "GMT-"
                    }${
                      Math.abs(time.offset) > 0 ? Math.abs(time.offset) : ""
                    }) ${time.name}`}
                  </option>
                ))}
              </select>
            )}
          />
          <ErrorMessage message={errors.timezone?.message} />
        </div>
      </div>
      <div className="flex flex-col gap-5">
        <div className="font-semibold">Notification</div>
        <div className="flex items-center gap-4">
          <Controller
            name="transaction"
            control={control}
            render={({ field }) => (
              <Switch checked={field.value} onCheckedChange={field.onChange} />
            )}
          />
          <div>I send or receive digita currency</div>
        </div>
        <div className="flex items-center gap-4">
          <Controller
            name="merchant"
            control={control}
            render={({ field }) => (
              <Switch checked={field.value} onCheckedChange={field.onChange} />
            )}
          />
          <div>I receive merchant order</div>
        </div>
        <div className="flex items-center gap-4">
          <Controller
            name="recommendation"
            control={control}
            render={({ field }) => (
              <Switch checked={field.value} onCheckedChange={field.onChange} />
            )}
          />
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
