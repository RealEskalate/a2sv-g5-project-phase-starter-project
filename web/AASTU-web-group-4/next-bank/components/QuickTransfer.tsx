'use client'
import React, { useState, useEffect, useContext } from "react";
import { colors } from "@/constants/index";
import Image from "next/image";
import { createTransaction, getLatestTransfers } from '@/services/transactionfetch';
import { useForm } from "react-hook-form";
import Cookie from 'js-cookie';
import { TbFileSad } from "react-icons/tb";
import { message } from 'antd';


import AccountContext from "./account_balance_context";
import { UserData } from "@/types";
import { currentuser } from "@/services/userupdate";
import { ArrowPathIcon } from "@heroicons/react/24/outline";
interface UserType {
  id: string;
  name: string;
  username: string;
  city: string;
  country: string;
  profilePicture: string;
}

const QuickTransfer: React.FC = () => {
  const { register, reset, handleSubmit, formState: { errors } } = useForm();
  const [currentIndex, setCurrentIndex] = useState(0);
  const [users, setUsers] = useState<UserType[]>([]);
  const [selectedUser, setSelectedUser] = useState<string>('');
  const [status, setStatus] = useState<'loading' | 'success' | 'error'>('loading');
  const accessToken = Cookie.get('accessToken') ?? '';
  const [messageApi, contextHolder] = message.useMessage();
  const success = (amount: string, username: string) => {
    messageApi.open({
      type: 'success',
      content: `Successfully transferred ${amount} to ${username}`,
      duration: 4
    });
  };

  const errormessage = () => {
    messageApi.open({
      type: 'error',
      content: 'Transaction was not successful',
      duration:4
    } as any);
  };

  const nouser = () => {
    messageApi.open({
      type: 'error',
      content: 'Please select user',
      duration:4} as any);};


  const lowbalance = () => {
    messageApi.open({
      type: 'error',
      content: 'Insufficient funds',
      duration:4} as any);};

  useEffect(() => {
    const fetchUsers = async () => {
      setStatus('loading');
      try {

        const data = await getLatestTransfers(6);
        if (data.success) {
          setUsers(data.data);
          setStatus('success');
        } else {
          setStatus('error');
        }
        console.log(data,status)
      } catch (error) {
        console.error("Error fetching the users: ", error);
        setStatus('error');
      }
    };

    fetchUsers();
  }, []);

  const [accountBalance, setAccountBalance ] = useState(0);
  const [info, setinfo] = useState<UserData>();
  const [visible , setvisible] = useState(false)
  useEffect(() => {
    const fetch = async () => {
      try {
        const data = await currentuser();
        setinfo(data.data || []);
        setAccountBalance(data.data.accountBalance);
      } catch (error) {
        console.error("Error:", error);
      }
    };
    fetch();
  } , []);

  const handleNext = () => {
    if (currentIndex < users.length - 3) {
      setCurrentIndex(currentIndex + 1);
    }
  };

  const handlePrev = () => {
    if (currentIndex > 0) {
      setCurrentIndex(currentIndex - 1);
    }
  };

  const handleUserSelect = (event: React.MouseEvent<HTMLDivElement>) => {
    const userId = event.currentTarget.dataset.id;
    setSelectedUser(userId || '');
  };
  const[loading , setloading] = useState(false)
  const onSubmit = async (data: { amount: string }) => {
    setloading(true)
    
    const transactionData = {
      type: "transfer",
      description: `Transfer to ${selectedUser}`,
      amount: data.amount,
      receiverUserName: selectedUser
    };

    try {
      const res = await createTransaction(transactionData, accessToken);
      if (res.success && parseInt(transactionData.amount) < accountBalance) {
        success(transactionData.amount, transactionData.receiverUserName);
        reset();
      } 
      else if (parseInt(transactionData.amount) > accountBalance){
        lowbalance();
        console.error('Insufficient funds , typeof(accountBalance):', (accountBalance) , (transactionData.amount));
      }
      else if (selectedUser === '') {
        nouser();
      }
      else {
        errormessage();
        console.error('Failed to create transaction', res);
      }
    } catch (error) {
      errormessage();
      console.error('Error creating transaction:', error);
    }
    setloading(false)
  };

  

  if (status === 'loading') {
    return (
      <div className="flex  rounded-2xl flex-col py-4 items-start gap-3 w-[100%] bg-gray-200 dark:bg-dark text-gray-900 dark:text-white animate-pulse">
        {/* Loading skeleton */}
        <div className="flex justify-between items-center h-35 w-[100%] px-2">
          <div className="w-[40px] h-[40px] rounded-full bg-gray-300 dark:bg-gray-600 flex justify-center items-center shadow-lg"></div>
          <div className="flex py-6 gap-2">
            {Array.from({ length: 3 }).map((_, index) => (
              <div key={index} className="flex flex-col gap-2 flex-1 cursor-pointer">
                <div className="w-[70px] h-[70px] bg-gray-300 dark:bg-gray-600 rounded-full"></div>
                <div className="flex flex-col items-center">
                  <div className="w-24 h-4 bg-gray-300 dark:bg-gray-600"></div>
                </div>
              </div>
            ))}
          </div>
          <div className="w-[40px] h-[40px] rounded-full bg-gray-300 dark:bg-gray-600 flex justify-center items-center shadow-lg"></div>
        </div>
        <div className="grid grid-cols-5 w-[100%] justify-center items-center px-2">
          <div className="w-24 h-4 bg-gray-300 dark:bg-gray-600 col-span-2"></div>
          <div className="w-44 h-10 bg-gray-300 rounded-xl dark:bg-gray-600 col-span-3"></div>
        </div>
      </div>
    );
  }

  if (status === 'error') {
    return (
      <div className="flex flex-col items-center justify-center h-full text-red-500">
        <TbFileSad
          className={`text-gray-300 dark:text-[#993d4b] w-[400px] h-[70px] pb-2 block mx-auto`}
          strokeWidth={1}
        />
        <div>Error fetching the users</div>
      </div>
    );
  }

  return (
    <div className="flex bg-white rounded-2xl flex-col py-4 items-start gap-3 w-[100%] dark:bg-dark text-gray-900 dark:text-white">
      {contextHolder}
      <div className="flex justify-between items-center h-35 w-[100%] px-2">
        <div
          onClick={handlePrev}
          className={`w-[40px] h-[40px] rounded-full ${colors.white} flex justify-center items-center shadow-lg cursor-pointer ${users.length <= 3 ? "opacity-50 cursor-not-allowed" : ""}`}
        >
          <span className="text-gray-500">&lt;</span>
        </div>

        <div className="flex py-6 overflow-hidden gap-2">
          {users.length === 0 ? (
            <div className="flex flex-col items-center justify-center h-full text-white dark:text-blue-500">
              <div>No data to display</div>
            </div>
          ) : (
            users.slice(currentIndex, currentIndex + 3).map((item) => (
              <div key={item.id} data-id={item.username}
                className={`flex flex-col gap-2 flex-1 cursor-pointer`}
                onClick={handleUserSelect}
              >
                <Image
                  src={'/Images/pp.jpg'}
                  width={70}
                  height={70}
                  className={`rounded-full ${item.username === selectedUser ? `border-4 border-blue-500 shadow-md` : ''}`}
                  alt={item.name}
                />
                <div className="flex flex-col items-center">
                  <div className={`font-normal text-[12px] ${colors.textblack} text-center whitespace-normal`}>
                    {item.name}
                  </div>
                </div>
              </div>
            ))
          )}
        </div>

        <div
          onClick={handleNext}
          className={`w-[40px] h-[40px] rounded-full ${colors.white} flex justify-center items-center shadow-lg cursor-pointer ${users.length <= 3 ? "opacity-50 cursor-not-allowed" : ""}`}
        >
          <span className="text-gray-500">&gt;</span>
        </div>
      </div>

      <div className="grid grid-cols-3 w-[100%] justify-center items-center px-2">
        <p className={`text-center px-2 text-nowrap text-[12px] font-normal ${colors.textgray}`}>
          Write amount
        </p>
        <div className={`col-span-2 flex w-[100%] ${colors.lightblue} rounded-3xl`}>
          <input
            type="number"
            {...register("amount", { required: "Amount is required" })}
            placeholder="535.35"
            className={`w-[50%] border-1 rounded-3xl py-2 border-black ${colors.lightblue} focus:${colors.lightblue} focus:outline-none px-3 focus:border-none`}
          />
          <button
            className={`${colors.blue} text-white w-[60%] py-2 rounded-3xl ${loading ? 'bg-indigo-500' :'bg-[#1814F9]'}`}
            onClick={handleSubmit(onSubmit as any )}
          >
          {loading ?(
            <div className="flex justify-center items-center ">
            <ArrowPathIcon className="h-5 w-5 animate-spin  text-white  " />
          </div>
          ) :("Send ✈")}
            
          </button>
        </div>
      </div>
    </div>
  );
};

export default QuickTransfer;
