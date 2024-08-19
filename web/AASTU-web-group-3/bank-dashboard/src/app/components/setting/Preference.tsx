"use client"
import React, { useRef,useState } from "react";
import { useForm } from "react-hook-form";
import SwitchButton from "./switch";

interface FormInput {
  currency: string;
  sentOrReceiveDigitalCurrency: boolean;
  receiveMerchantOrder: boolean;
  accountRecommendations: boolean;
  timeZone: string;
}

const Preference = () => {
    
    const { register, handleSubmit, setValue } = useForm<FormInput>({
        defaultValues: {
            currency: "",
            timeZone: "",
            sentOrReceiveDigitalCurrency: false,
            receiveMerchantOrder: false,
            accountRecommendations: false,
        },
    });
    
    const  [sentResDigitalCurrency, setSentResDigitalCurrency] = useState(true)
    const  [recMerchantOrder, setRecMerchantOrder] = useState(true)
    const  [accRecommendations, setAccRecommendations] = useState(false)
        

  const fileInputRef = useRef<HTMLInputElement>(null);

  const onSubmit = async (data: FormInput) => {
    console.log("Form submitted:", data);
    const formData = new FormData();
    // Submit the formData object to your server or API
  };

  return (
    <div className="p-4 ">

        <div className="flex flex-wrap items-end justify-between w-full font-semibold">
          <div className="mb-3 w-full md:w-[45%]">
            <label className="block text-black text-base mb-2">Currency</label>
            <input
              className="w-full p-3 md:p-2 text-[#718EBF] border-2 text-sm border-[#DFEAF2] rounded-lg focus:outline-none"
              type="text"
              id="currency"
              placeholder="USD"
              {...register("currency", {
                required: {
                  value: true,
                  message: `currency is required`,
                },
              })}
            />
          </div>
          <div className="mb-3 w-full md:w-[45%]">
            <label className="block text-black text-sm mb-2">Time Zone</label>
            <input
              className="w-full p-3 md:p-2 text-[#718EBF] border-2 text-sm border-[#DFEAF2] rounded-lg focus:outline-none"
              type="text"
              id="Time Zone"
              placeholder="(GMT-12:00) International Date Line West"
              {...register("timeZone", {
                required: {
                  value: true,
                  message: `currency is required`,
                },
              })}
            />
          </div>
        </div>
        <div className=" flex flex-col gap-4 font-semibold">
            <div className="font-semibold text-xl p-2 ">Notification</div>
            <div className="flex w-full md:w-1/2 gap-2 ">
                <div className="w-1/5">
                <SwitchButton isOn={sentResDigitalCurrency} onToggle = {setSentResDigitalCurrency}/> 

                </div>
                I send or receive digita currency
            </div>
            <div className="flex w-full md:w-1/2 gap-2">
                <div className="w-1/5">
                <SwitchButton isOn={recMerchantOrder} onToggle = {setRecMerchantOrder}/> 

                </div>
                I receive merchant order
            </div>
            <div className="flex w-full md:w-[52%] gap-2">
                <div className="w-1/5">

                <SwitchButton isOn={accRecommendations} onToggle = {setAccRecommendations}/>  
                </div>
                There are recommendation for my account
            </div>
            
        </div>

        <div className="flex justify-end w-full my-6">
          <button
            className=" w-full md:w-1/5 bg-[#1814F3] text-white font-semibold py-2 px-4 rounded-lg focus:outline-none"
            type="submit" onClick ={handleSubmit(onSubmit)}
          >
            Save
          </button>
        </div>
    </div>
  );
};

export default Preference;
