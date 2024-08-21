'use client'
import React, { useState, useEffect } from "react";
import { colors } from "@/constants/index";
import Image from "next/image";
import { createTransaction, getLatestTransfers } from '@/services/transactionfetch';
import { useForm } from "react-hook-form";
import Cookie from 'js-cookie'

const dummyData = [
  { name: "Natnael Worku", position: "CEO", imageSrc: "/Images/pp.jpg" },
  { name: "John Doe", position: "CTO", imageSrc: "/Images/pp.jpg" },
  { name: "Jane Smith", position: "CFO", imageSrc: "/Images/pp.jpg" },
  { name: "Michael Johnson", position: "COO", imageSrc: "/Images/pp.jpg" },
  { name: "Emily Davis", position: "CMO", imageSrc: "/Images/pp.jpg" },
];

interface UserType {
  id: string,
  name: string,
  username: string,
  city: string,
  country: string,
  profilePicture: string
}


const QuickTransfer = () => {
  const {register, reset,handleSubmit,formState:{errors}} = useForm();
  const [currentIndex, setCurrentIndex] = useState(0);
  const [users, setUsers] = useState<UserType[]>([]);
  const [selectedUser, setSelectedUser] = useState<string>('');
  const accessToken = Cookie.get('accessToken') ?? ''

  useEffect(() => {
    const getUsersData = async () => {
      const data = await getLatestTransfers(accessToken, 6);
      setUsers(data.data);
    };
    getUsersData();
  });

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
    console.log(userId);
  }



  const onSubmit  = async (data: { amount: string; }) =>{
    const objectData = {
      type: "transfer",
      description: ` Transfer to ${selectedUser}`,
      amount: data.amount,
      receiverUserName: `${selectedUser}`
    }
    const res = await createTransaction( objectData ,accessToken)
    if (res.status === 200){
      console.log('Transaction created successfully')
      
      reset()
    }
    console.log(res)

  }

  return (
    <div className="flex bg-white rounded-2xl flex-col py-4 items-start gap-3 w-[100%]">
      <div className="flex justify-between items-center h-35 w-[100%] px-2">
        <div
          onClick={handlePrev}
          className={`w-[40px] h-[40px] rounded-full ${colors.white} flex justify-center items-center shadow-lg cursor-pointer ${users.length <= 3 ? "opacity-50 cursor-not-allowed" : ""}`}
        >
          <span className="text-gray-500">&lt;</span>
        </div>

        <div className="flex py-6 overflow-hidden gap-2">
          {users.length > 3 ? (
            users.slice(currentIndex, currentIndex + 3).map((item, index) => (
              <div key={item.id}  data-id={item.username}
              className={`flex flex-col gap-2 flex-1 cursor-pointer  `}
              onClick={handleUserSelect}
              >
                <Image
                  src={'/Images/pp.jpg'}
                  width={70}
                  height={70}
                  className={`rounded-full ${item.username === selectedUser ? `border border-linear-gradient(to bottom,blue-500,blue-300)`: ''}`}
                  alt={item.name}
                />
                <div className="flex flex-col items-center">
                  <div
                    className={`font-normal text-[12px] ${colors.textblack} text-center whitespace-normal`}
                  >
                    {item.name}
                  </div>
                </div>
              </div>
            ))
          ) : (
            users.map((item, index) => (
              <div key={item.id} data-id= {item.username} 
              className={`flex flex-col gap-2 flex-1 cursor-pointer  `}
              onClick={handleUserSelect}
              >
                <Image
                  src={'/Images/pp.jpg'}
                  width={70}
                  height={70}
                  className={`rounded-full ${item.username === selectedUser ? `border-4 border-blue-500  shadow-md`: ''}`}
                  alt={item.name}
                />
                <div className="flex flex-col">
                  <div
                    className={`font-normal text-[12px] ${colors.textblack} text-center whitespace-normal`}
                  >
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
        <p
          className={`text-center px-2 text-nowrap text-[12px] font-normal ${colors.textgray}`}
        >
          Write amount
        </p>
        <div
          className={`col-span-2 flex w-[100%] ${colors.lightblue} rounded-3xl`}
        >
          <input
            type="number"
            {...register("amount", { required: "Username is required" })}
            placeholder="535.35"
            className={`w-[50%] border-1 rounded-3xl py-2 border-black ${colors.lightblue} focus:${colors.lightblue} focus:outline-none px-3 focus:border-none`}
          />
          <button
            className={`${colors.blue} text-white w-[60%] py-2 rounded-3xl`}
            onClick={handleSubmit(onSubmit as any)}
          >
            Send ✈
          </button>
        </div>
      </div>
    </div>
  );
};

export default QuickTransfer;
