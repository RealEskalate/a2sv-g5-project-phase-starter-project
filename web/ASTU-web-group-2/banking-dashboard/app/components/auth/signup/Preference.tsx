import React, { useState } from "react";
import { useForm, Controller } from "react-hook-form";
import * as Yup from "yup";
import { yupResolver } from "@hookform/resolvers/yup";
import ToggleSwitch from "./ToggleSwitch";
import { MainData } from "./Signup";
import { useSignUpMutation } from "@/lib/service/TransactionService";
import { useRouter } from "next/navigation";
import { signIn } from "next-auth/react";
import notify from "@/utils/notify";
const schema = Yup.object().shape({
  currency: Yup.string().required("Currency is required"),
  timeZone: Yup.string().required("TimeZone is required"),
  sentOrReceiveDigitalCurrency: Yup.boolean().required(),
  receiveMerchantOrder: Yup.boolean().required(),
  accountRecommendations: Yup.boolean().required(),
});

interface FormValues {
  currency: string;
  timeZone: string;
  sentOrReceiveDigitalCurrency: boolean;
  receiveMerchantOrder: boolean;
  accountRecommendations: boolean;
}

interface YourFormComponentProps {
  mainData: MainData;
  setMainData: React.Dispatch<React.SetStateAction<MainData>>;
}

interface SignUpResponse {
  data?: any; // Replace `any` with the actual type if known
  error?: {
    data?: {
      message: string;
    };
  };
}

const YourFormComponent: React.FC<YourFormComponentProps> = ({
  setMainData,
  mainData,
}) => {
  const router = useRouter();
  const {
    control,
    handleSubmit,
    register,
    formState: { errors },
  } = useForm<FormValues>({
    resolver: yupResolver(schema),
    defaultValues: {
      currency: mainData.preference?.currency || "",
      timeZone: mainData.preference?.timeZone || "",
      sentOrReceiveDigitalCurrency:
        mainData.preference?.sentOrReceiveDigitalCurrency || false,
      receiveMerchantOrder: mainData.preference?.receiveMerchantOrder || false,
      accountRecommendations:
        mainData.preference?.accountRecommendations || false,
    },
  });
  const [signUp, { isLoading }] = useSignUpMutation();
  const [posting, setIsPosting] = useState(false);
  const onSubmit = async (data: FormValues) => {
    setIsPosting(true);
    const formattedData = {
      ...mainData,
      preference: {
        currency: data.currency,
        sentOrReceiveDigitalCurrency: data.sentOrReceiveDigitalCurrency,
        receiveMerchantOrder: data.receiveMerchantOrder,
        accountRecommendations: data.accountRecommendations,
        timeZone: data.timeZone,
        twoFactorAuthentication: true,
      },
    };

    setMainData(formattedData);

    const res = await signUp(formattedData);

    if ("data" in res) {
      const loginRes = await signIn("credentials", {
        userName: mainData.username,
        password: mainData.password,
        redirect: false,
      });
      if (!loginRes?.ok) {
        notify.error("Invalid Credentials");
      } else {
        router.push("/dashboard");
        notify.success("Signup successful");
      }
    } else if ("error" in res) {
      let errorMessage = "Signup failed";
      if (res.error && "data" in res.error) {
        const errorData = (res.error as { data?: { message: string } }).data;
        errorMessage = errorData?.message || "Signup failed";
      }
      notify.error(errorMessage);
      console.error("Signup failed", res.error);
    }
    setIsPosting(false);
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div className="w-full">
          <label className="block mb-1 font-400 text-[16px] text-[#232323] capitalize">
            Currency
          </label>
          <input
            placeholder={`USD`}
            className="w-full p-2 border border-[#DFEAF2] rounded-[15px] focus:outline-none focus:ring-2 focus:ring-blue-200"
            {...register("currency")}
          />
          {errors.currency && (
            <p className="text-red-500">{errors.currency.message}</p>
          )}
        </div>
        <div className="w-full">
          <label className="block mb-1 font-400 text-[16px] text-[#232323] capitalize">
            TimeZone
          </label>
          <input
            placeholder={`{GMT-12:00} International Data Line West`}
            className="w-full p-2 border border-[#DFEAF2] rounded-[15px] focus:outline-none focus:ring-2 focus:ring-blue-200"
            {...register("timeZone")}
          />
          {errors.timeZone && (
            <p className="text-red-500">{errors.timeZone.message}</p>
          )}
        </div>
      </div>

      <div className="space-y-4">
        <label className="block mb-1 font-400 text-[16px] text-[#232323] capitalize">
          Notification
        </label>
        <div className="md:grid gap-3">
          <Controller
            name="sentOrReceiveDigitalCurrency"
            control={control}
            render={({ field }) => (
              <ToggleSwitch
                id="sentOrReceiveDigitalCurrency"
                label="I send or receive digital currency"
                checked={field.value}
                onChange={field.onChange}
              />
            )}
          />
          <Controller
            name="receiveMerchantOrder"
            control={control}
            render={({ field }) => (
              <ToggleSwitch
                id="receiveMerchantOrder"
                label="I receive merchant order"
                checked={field.value}
                onChange={field.onChange}
              />
            )}
          />
          <Controller
            name="accountRecommendations"
            control={control}
            render={({ field }) => (
              <ToggleSwitch
                id="accountRecommendations"
                label="There are recommendations for my account"
                checked={field.value}
                onChange={field.onChange}
              />
            )}
          />
        </div>
      </div>
      <div className="flex justify-end">
        <button
          type="submit"
          className="w-full md:w-auto px-4 py-2 bg-[#083E9E] text-white rounded-lg"
          disabled={posting}
        >
          {posting ? (
            <div className="w-8 h-8 border-4 border-dashed rounded-full animate-spin [animation-duration:3s]  border-white mx-auto"></div>
          ) : (
            "Register"
          )}
        </button>
      </div>
    </form>
  );
};

export default YourFormComponent;
