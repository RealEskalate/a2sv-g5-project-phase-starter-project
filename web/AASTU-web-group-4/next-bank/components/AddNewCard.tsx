"use client";
import React from "react";
import { useForm, SubmitHandler } from "react-hook-form";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import { FaCalendarAlt } from "react-icons/fa";

type NewCardProps = {
  cardType: string;
  nameOnCard: string;
  balance: string;
  expirationDate: Date;
};

const AddNewCard: React.FC = () => {
  const { register, handleSubmit, setValue, formState: { errors } } = useForm<NewCardProps>();

  const [selectedDate, setSelectedDate] = React.useState<Date | null>(null);

  const onSubmit: SubmitHandler<NewCardProps> = (data) => {
    console.log(data);
  };

  return (
    // w-[330px] md:w-[600px] h-[530px] md:h-[380px]
    <div className='mr-4'>
      <h1 className='text-[20px] mb-3 font-bold text-[#333B69]'>Add New Card</h1>
      <div className="bg-white lg:w-[800px] w-[330px] md:w-[630px] sm:h-[720px] md:h-[430px] p-7 border-[1px] rounded-xl">
        <p className='text-[17px] text-[#718EBF]'>
          Credit Card generally means a plastic card issued by Scheduled Commercial Banks assigned to a Cardholder, with a credit limit, that can be used to purchase goods and services on credit or obtain cash advances.
        </p>

        <form onSubmit={handleSubmit(onSubmit)} className="grid grid-cols-1 md:grid-cols-2 gap-4 mt-4">
          <div>
            <label htmlFor="cardTypeId" className='text-[16px] block pb-2'>Card Type</label>
            <input
              id='cardTypeId'
              {...register('cardType', { required: true })}
              placeholder='Classic'
              className='border-[1px] border-[#DFEAF2] rounded-md text-[14px] p-3 w-full outline-none text-[#718EBF] placeholder-[#718EBF]'
            />
            {errors.cardType && <span className="text-red-500 text-sm">This field is required</span>}
          </div>

          <div>
            <label htmlFor="nameOnCardId" className='text-[16px] block pb-2'>Name On Card</label>
            <input
              id='nameOnCardId'
              {...register('nameOnCard', { required: true })}
              placeholder='My Cards'
              className='border-[1px] border-[#DFEAF2] rounded-md text-[14px] p-3 w-full outline-none text-[#718EBF] placeholder-[#718EBF]'
            />
            {errors.nameOnCard && <span className="text-red-500 text-sm">This field is required</span>}
          </div>

          <div>
            <label htmlFor="cardNumberId" className='text-[16px] block pb-2'>Balance</label>
            <input
                id='balanceId'
                {...register('balance', {
                required: 'Initial is required',
                pattern: {
                    value: /^[-+]?(\d{1,3}(,\d{3})*|\d+)(\.\d{2})?\s*([A-Z]{3}|[$€£¥])$/,
                    message: 'The balance should be in the following format 278,000,000$',
                },
                })}
                placeholder='27,000$'
                className='border-[1px] border-[#DFEAF2] rounded-md text-[14px] p-3 w-full outline-none text-[#718EBF] placeholder-[#718EBF]'
            />
            {errors.balance && <span className="text-red-500 text-sm">{errors.balance.message as string}</span>}
          </div>

          <div className="relative">
            <label htmlFor="expirationDateId" className='text-[16px] block pb-2'>Expiration Date</label>
            <div className="flex items-center">
              <DatePicker
                id="expirationDateId"
                selected={selectedDate}
                onChange={(date) => {
                  setSelectedDate(date);
                  setValue("expirationDate", date as Date);
                }}
                placeholderText='25 January 2025'
                className='border-[1px] border-[#DFEAF2] rounded-md text-[14px] p-3 w-full pr-10 outline-none text-[#718EBF] placeholder-[#718EBF]'
                dateFormat="dd MMMM yyyy"
              />
              <FaCalendarAlt className="absolute right-3 md:right-0pointer-events-none text-[#718EBF]" />
            </div>
            {errors.expirationDate && <span className="text-red-500 text-sm">{errors.expirationDate.message as string}</span>}
          </div>

          <div className="md:col-span-2">
            <button type="submit" className='rounded-xl text-[16px] px-7 text-center bg-[#1814F3] text-white w-[95%] md:w-[auto] mt-4 p-2'>
              Add Card
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default AddNewCard;