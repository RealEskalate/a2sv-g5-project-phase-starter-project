"use client";
import React, { useEffect } from "react";
import ToggleInput from "../Form/ToggleInput";
import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useAppDispatch, useAppSelector } from "@/hooks/hoooks";
import {
  useGetProfileQuery,
  useUpdatePreferenceMutation,
} from "@/lib/redux/api/profileAPI";
import { setProfile } from "@/lib/redux/slices/profileSlice";
import { UserPreferenceType } from "@/types/user.types";
import { Loader } from "lucide-react";

const preferencesSchema = z.object({
  currency: z.enum(["USD", "ETB", "Yen"]).default("USD"), // Assuming these are the available options
  timeZone: z.string().min(1, "Time Zone is required"),
  sentOrReceiveDigitalCurrency: z.boolean(),
  receiveMerchantOrder: z.boolean(),
  accountRecommendations: z.boolean(),
  twoFactorAuthentication: z.boolean(),
});

const Preferences = () => {
  const dispatch = useAppDispatch();
  const getData = useAppSelector((state) => state.profile);

  const { refetch, data, error, isSuccess } = useGetProfileQuery();

  console.log(data);

  useEffect(() => {
    if (isSuccess && data) {
      dispatch(setProfile(data?.data));
      console.log(data.data);
    }
  }, [data, dispatch]);

  const {
    register,
    handleSubmit,
    formState: { errors },
    watch,
  } = useForm({
    resolver: zodResolver(preferencesSchema),
    defaultValues: {
      currency: getData.preference.currency,
      timeZone: getData.preference.timeZone,
      sentOrReceiveDigitalCurrency:
        getData.preference.sentOrReceiveDigitalCurrency,
      receiveMerchantOrder: getData.preference.receiveMerchantOrder,
      accountRecommendations: getData.preference.accountRecommendations,
      twoFactorAuthentication: getData.preference.twoFactorAuthentication,
    },
  });

  if (error) {
    return <h1>An Error Occured..</h1>;
  }

  const [updatePreference, { isLoading }] = useUpdatePreferenceMutation();

  const onSubmit = (data: UserPreferenceType) => {
    updatePreference(data).then((res: any) => {
      dispatch(setProfile(res?.data?.data));
      alert("Preferences updated successfully");
      refetch();
    });
  };

  return (
    <div>
      <form action="" onSubmit={handleSubmit(onSubmit)}>
        <div className="flex flex-col md:flex-row md:space-x-5">
          <div className=" w-full lg:w-6/12 space-y-3 my-3">
            <label htmlFor="select" className="gray-dark text-16px">
              Country
            </label>
            <select
              id="currency"
              {...register("currency")}
              className="w-full border-2 border-[#DFEAF2] p-5 py-3 rounded-xl placeholder:text-blue-steel focus:border-blue-steel outline-none"
            >
              <option value="USD">USD</option>
              <option value="ETB">ETB</option>
              <option value="GPY">Yen</option>
            </select>
          </div>
          <div className=" w-full lg:w-6/12 space-y-3 my-3">
            <label htmlFor="select" className="gray-dark text-16px">
              Time Zone
            </label>
            <select
              id="timeZone"
              {...register("timeZone")}
              className="w-full border-2 border-[#DFEAF2] p-5 py-3 rounded-xl placeholder:text-blue-steel focus:border-blue-steel outline-none"
            >
              <option value="GMT3p">GMT 3+</option>
              <option value="GMT4p">GMT 4+</option>
              <option value="GMT5p">GMT 5+</option>
              <option value="GMT6p">GMT 6+</option>
            </select>
          </div>
        </div>

        <label className="gray-dark text-16px">Notification</label>

        <ToggleInput
          label="I send or receive digital currency"
          inputType="checkbox"
          id="sentOrReceiveDigitalCurrency"
          registerName="sentOrReceiveDigitalCurrency"
          register={register}
          placeholder="Sent Or Receive Digital Currency"
          currentState={watch("sentOrReceiveDigitalCurrency") as boolean}
        />
        <ToggleInput
          label="I recieve merchant order"
          inputType="checkbox"
          id="receiveMerchantOrder"
          registerName="receiveMerchantOrder"
          register={register}
          placeholder="Receive Merchant Order"
          currentState={watch("receiveMerchantOrder") as boolean}
        />
        <ToggleInput
          label="There are recommendation for my account"
          inputType="checkbox"
          id="accountRecommendations"
          registerName="accountRecommendations"
          register={register}
          placeholder="Account Recommendations"
          currentState={watch("accountRecommendations") as boolean}
        />

        <div className="flex justify-end">
          <button
            type="submit"
            className="bg-[#1814f3] text-white px-10 py-2 rounded-lg w-full md:w-auto mt-4"
          >
            {isLoading ? <Loader className="animate-spin" /> : "Save"}
          </button>
        </div>
      </form>
    </div>
  );
};

export default Preferences;
