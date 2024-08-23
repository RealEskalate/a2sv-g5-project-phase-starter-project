"use client";
import React, { useState, useEffect } from 'react';
import { colors, logo } from '@/constants';
import Image from 'next/image';
import { FaEyeSlash, FaEye } from "react-icons/fa";
import { useForm } from 'react-hook-form';
import { createTransaction } from '@/services/transactionfetch';
import Cookies from 'js-cookie';
import { message } from 'antd';

const TransferPage: React.FC = () => {
  const [visible, setVisible] = useState(false);
  const [messageApi, contextHolder] = message.useMessage();
  const [status, setStatus] = useState<'success' | 'error' | null>(null);
  const { register, reset, handleSubmit, formState: { errors } } = useForm();
  let account = 5000;

  useEffect(() => {
    if (status === 'success') {
      messageApi.open({
        type: 'success',
        content: `Successfully transferred amount to the receiver`,
        duration:4
      });
    } else if (status === 'error') {
      messageApi.open({
        type: 'error',
        content: 'Transaction was not successful',
        duration:4
      });
    }
    setStatus(null); // Reset status after showing the message
  }, [status]);

  const onSubmit = async (data: any) => {
    const Token = Cookies.get('accessToken') ?? "";
    
    const transactionData = {
      ...data,
      type: data.type,
      description: data.description,
      amount: data.amount,
      receiverUserName: data.receiverUserName,
    };

    try {
      const res = await createTransaction(transactionData, Token);
      console.log(res);
      // console.log(res)
      if (res.success) {
        setStatus('success');
        reset();
      } else {
        setStatus('error');
      }
    } catch (error) {
      console.error("Error creating transaction:", error);
      setStatus('error');
    }
  };

  const handleVisibility = () => {
    setVisible(!visible);
  };
  
  return (
    <>
      {contextHolder}
      <div className="container flex flex-col items-center max-h-screen h-screen md:flex-row md:h-full bg-gray-100 p-0 md:ml-45 md:p-8 md:max-h-full md:gap-10 md:mb-24 dark:bg-dark text-gray-900 dark:text-white">
        <Image
          src={logo.transfer}
          alt="next logo"
          width={500}
          height={100}
          className="md:ml-72 hidden lg:block"
        />
        
        <div className="p-8 rounded-lg max-w-full md:w-2/6">
          <div className='text-center text-2xl font-bold p-4 flex justify-center gap-6'>
            <p>{visible ? `$${account}` : "$****"}</p>
            <button onClick={handleVisibility}>
              {visible ? <FaEyeSlash className='text-orange-400 hover:text-orange-800 text-3xl' /> : <FaEye className='text-orange-400 hover:text-orange-800 text-3xl' />}
            </button>
          </div>
          
          <h1 className="text-center text-gray-800 mb-4">User name</h1>

          <form onSubmit={handleSubmit(onSubmit)}>
            <div className="mb-4">
              <label htmlFor="type" className="block text-gray-700 text-sm font-bold mb-2">
                Type:
              </label>
              <select
                id="type"
                {...register('type', { required: "This field is required" })}
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              >
                <option value="transfer">Transfer</option>
                <option value="shopping">Shopping</option>
                <option value="deposit">Deposit</option>
                <option value="service">Service</option>
              </select>
            </div>

            <div className="mb-4">
              <label htmlFor="description" className="block text-gray-700 text-sm font-bold mb-2">
                Description (optional):
              </label>
              <textarea
                {...register('description', { required: "This field is required" })}
                id="description"
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                placeholder="Enter a brief description"
              />
            </div>

            <div className="mb-4">
              <label htmlFor="amount" className="block text-gray-700 text-sm font-bold mb-2">
                Amount:
              </label>
              <input
                {...register('amount', { required: "This field is required" })}
                type="number"
                id="amount"
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                placeholder="Enter amount"
                required
              />
            </div>

            <div className="mb-6">
              <label htmlFor="receiverUserName" className="block text-gray-700 text-sm font-bold mb-2">
                Receiver Username:
              </label>
              <input
                {...register('receiverUserName', { required: "This field is required" })}
                type="text"
                id="receiverUserName"
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                placeholder="Enter receiver's username"
                required
              />
            </div>

            <div className="flex items-center justify-between">
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
    </>
  );
};

export default TransferPage;
