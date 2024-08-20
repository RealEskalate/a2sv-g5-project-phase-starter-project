"use client";
import React, { useState } from 'react';
import { colors , logo } from '@/constants';
import Image from 'next/image';
import { FaEyeSlash } from "react-icons/fa";
import { FaEye } from "react-icons/fa";
import { useForm } from 'react-hook-form';
import { createTransaction } from '@/services/transactionfetch';
import Cookies from 'js-cookie';
const TransferPage: React.FC = () => {
  const[visible , setvisible] = useState(false)

  const {register, reset,handleSubmit,formState:{errors}} = useForm();
  let account = 5000
  const onSubmit = async (data: any) => {
    const Token = Cookies.get('accessToken') ?? "";
    
    
    console.log(data);
    const res = await createTransaction( data ,Token)
    console.log(res)
    if (res.status == 200) {
      reset()
    }



  };

  
  
  const handleVisibility = () => {
    setvisible(!visible);

  }
  
  return (
    <div className=" container  flex flex-col items-center max-h-screen h-screen  md:flex-row  md:h-full bg-gray-100 p-0  md:ml-45 md:p-8 md:max-h-full md:gap-10 md:mb-24 ">
      
        <Image
        src={logo.transfer}
        alt="next logo"
        width={500}
        height={100}
        className="md:ml-72  hidden lg:block"
      />
      
      <div className=" p-8 rounded-lg max-w-full md:w-2/6 ">
        {visible && (
         <div className='text-center text-2xl font-bold p-4 flex justify-center gap-6  '>
            <p >
              ${account}
            </p>
            <button onClick={handleVisibility}> <FaEyeSlash className='text-orange-400 hover:text-orange-800 text-3xl' /> </button>
          </div>
        )
         }
         {
          !visible && (
            <div className='text-center text-2xl font-bold p-4 flex justify-center gap-6  '>
            <p >
              $****
            </p>
            <button onClick={handleVisibility}> <FaEye className='text-orange-400 hover:text-orange-800 text-3xl' /> </button>
          </div>
          )
         }
        <h1 className="text-center text-gray-800 mb-4">User name</h1>

        <form onSubmit={handleSubmit(onSubmit)} className=''>
          <div className="mb-4">
            <label htmlFor="type" className="block text-gray-700 text-sm font-bold mb-2">
              Type:
            </label>
            <select
              id="type"
              // value={type}
              {...register('type',{
                required: "This field is required"
              })}
              // onChange={(e) => setType(e.target.value)}
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              // required
            >
              <option value="transfer">Transfer</option>
              <option value="transfer">Shopping</option>
              <option value="transfer">Deposit</option>
              <option value="transfer">Service</option>
            </select>
          </div>

          <div className="mb-4">
            <label htmlFor="description" className="block text-gray-700 text-sm font-bold mb-2">
              Description (optional):
            </label>
            <textarea
              {...register('description',{
                required: "This field is required"
              } )}
              id="description"
              // value={description}/
              // onChange={(e) => setDescription(e.target.value)}
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              placeholder="Enter a brief description"
            />
          </div>

          <div className="mb-4">
            <label htmlFor="amount" className="block text-gray-700 text-sm font-bold mb-2">
              Amount:
            </label>
            <input
              {...register('amount',{
                required: "This field is required"
              })}
              type="number"
              id="amount"
              // value={amount}
              // onChange={(e) => setAmount(e.target.value)}
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
              {...register('receiverUserName',{
                required: "This field is required"
              })}
              type="text"
              id="receiverUserName"
              // value={receiverUserName}
              // onChange={(e) => setReceiverUserName(e.target.value)}
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              placeholder="Enter receiver's username"
              required
            />
          </div>

          <div className="flex items-center justify-between ">
            <button
              type="submit"
              className={`${colors.blue} hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline `} 
            >
              Transfer
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default TransferPage;