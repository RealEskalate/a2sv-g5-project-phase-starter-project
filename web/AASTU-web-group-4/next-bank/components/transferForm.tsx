'use client';
import React, { useState } from 'react';
import { colors } from '@/constants';
import { FaEyeSlash, FaEye } from "react-icons/fa";
import { useForm } from 'react-hook-form';
import { createTransaction } from '@/services/transactionfetch';
import Cookies from 'js-cookie';

const TransferPage: React.FC = () => {
  const [visible, setVisible] = useState(false); 
  const { register, reset, handleSubmit, formState: { errors  } } = useForm();
  let account = 5000;

  const onSubmit = async (data: any) => {
    const Token = Cookies.get('accessToken') ?? "";
    console.log(data)
    const res = await createTransaction(data, Token);
    if (res.status === 200) {
      
      reset();
    }
  };

  const handleVisibility = () => {
    setVisible(!visible);
  };

  return (
    <div className="flex justify-center items-center p-6 lg:p-14">
      <div className="bg-white p-6 rounded-lg shadow-lg w-full max-w-lg">
        <div className="text-center text-2xl font-bold mb-4 flex justify-center items-center gap-3">
          <p>
            {visible ? `$${account}` : '$****'}
          </p>
          <button
            onClick={handleVisibility}
            aria-label={visible ? 'Hide account balance' : 'Show account balance'}
          >
            {visible ? (
              <FaEyeSlash className="text-green-200 hover:text-green-500 text-2xl" />
            ) : (
              <FaEye className="text-green-400 hover:text-green-700 text-3xl" />
            )}
          </button>
        </div>

        <div className="border-t border-gray-200 pt-4">
          <h2 className="text-xl font-bold mb-4">Transfer Details</h2>
          <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-4">
            <div>
              <label htmlFor="type" className="block text-black text-sm font-bold mb-1">Type:</label>
              <select
                id="type"
                {...register('type', { required: "This field is required" })}
                className="shadow appearance-none border placeholder-current rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                defaultValue=""
              >
                <option value="" disabled>Select a transaction type</option>
                <option value="transfer">Transfer</option>
                <option value="shopping">Shopping</option>
                <option value="deposit">Deposit</option>
                <option value="service">Service</option>
              </select>
              {errors.type && <p className="text-red-500 text-xs italic">{typeof errors.type.message === "string" && errors.type.message}</p>}
            </div>
            <div>
              <label htmlFor="description" className="block text-black text-sm font-bold mb-1">Description (optional):</label>
              <textarea
                {...register('description')}
                id="description"
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                placeholder="Enter a brief description"
              />
            </div>
            <div>
              <label htmlFor="amount" className="block text-black text-sm font-bold mb-1">Amount:</label>
              <input
                {...register('amount', { required: "This field is required" })}
                type="number"
                id="amount"
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                placeholder="Enter amount"
                required
              />
              {errors.amount && <p className="text-red-500 text-xs italic">{errors.amount.message as string}</p>}
            </div>
            <div>
              <label htmlFor="receiverUserName" className="block text-black text-sm font-bold mb-1">Receiver Username:</label>
              <input
                {...register('receiverUserName', { required: "This field is required" })}
                type="text"
                id="receiverUserName"
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                placeholder="Enter receiver's username"
                required
              />
              {errors.receiverUserName && <p className="text-red-500 text-xs italic">{errors.receiverUserName.message as string}</p>}
            </div>
            <div className="flex items-center justify-center mt-4">
              <button
                type="submit"
                className={`${colors.blue} hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline`}
              >
                Transfer
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default TransferPage;