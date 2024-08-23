"use client"


import React from "react";
import { Color } from "chart.js";
import { colors, logo } from "@/constants";
import DesktopCreditCart from "@/components/DesktopCreditCard";
import ResponsiveCreditCard from "@/components/CreditCard";
import RecentTransaction from "@/components/Recent Transaction";
import ExpensesChart from "@/components/ExpensesCart";
import { icons, Import } from "lucide-react";
import { text } from "stream/consumers";
import BarChart from "@/components/BarChart";
import PieChart from "@/components/PieChart";
import QuickTransfer from "@/components/QuickTransfer";
import LineChart from "@/components/LineChart";
import Link from "next/link";



const page = () => {
  return (
    <div className={`${colors.graybg} p-6 md:ml-64 md:max-w-full md:p-12  dark:bg-dark text-gray-900 dark:text-white`}>
      <div className="flex flex-col justify-between md:flex-row  gap-10 ">
        <div className=" py-4 md:w-3/5 md:max-w-full">
          <div className={`${colors.navbartext} flex justify-between `}>
            <h1 className="font-bold text-2xl dark:text-blue-500">My Cards</h1>
            <Link href="/credit-card" className="py-2 dark:text-blue-500"> See All</Link>
          </div>

          <div className="max-w-[345px] md:max-w-full">
            <div className="flex gap-3 overflow-x-auto max-w-full md:w-auto">
              <div className=" py-3 ">
                <ResponsiveCreditCard
                  backgroundColor={colors.blue}
                />
              </div>
              <div className=" py-3 ">
              <ResponsiveCreditCard
                  backgroundColor={colors.white}
                />
              </div>
            </div>
          </div>
        </div>
        <div className="  md:w-2/5  flex flex-col ">
          <div className={`${colors.navbartext} flex justify-between  py-4`}>
            <h1 className="font-bold text-2xl dark:text-blue-500 ">Recent Transaction</h1>
          </div>
          <div className="flex flex-col rounded-2xl pr-2 w-[100%]">
            <RecentTransaction/>
          </div>
        </div>
      </div>
      <div className=" w-[100%] flex flex-col justify-between  md:grid md:grid-cols-5 md:gap-10 ">
        <div className=" md:col-span-3 ">
          <div className={`${colors.navbartext} flex justify-between py-4`}>
            <h1 className="font-bold text-2xl dark:text-blue-500">Weekly Activity</h1>
          </div>
          <div className="w-[100%]">
            <BarChart />
          </div>
        </div>
        <div className=" w-[100%] py-5 flex flex-col gap-5 md:col-span-2 ">
          <div className={`${colors.navbartext}`}>
            <h1 className="font-bold text-2xl dark:text-blue-500">Expense Statstics</h1>
          </div>
          <div className="w-[100%] pr-6">
            <PieChart />
          </div>
        </div>
      </div>

      <div className="flex flex-col justify-between w-full  md:grid md:grid-cols-5 md:gap-10 ">
        <div className=" md:col-span-2 py-4  ">
          <div className={`${colors.navbartext} flex justify-between `}>
            <h1 className="font-bold text-2xl dark:text-blue-500">Quick Transfer</h1>
          </div>
          <div className="flex  gap-3 ">
            <div className="flex py-3 ">
              {" "}
              <QuickTransfer />
            </div>
          </div>
        </div>
        <div className=" md:col-span-3 ">
          <div className={`${colors.navbartext} flex justify-between  py-4`}>
            <h1 className="font-bold text-2xl dark:text-blue-500">Balance History</h1>
          </div>
          <div className="pr-6">
            <LineChart />
          </div>
        </div>
      </div>
    </div>
  );
};

export default page;



{/*import React from 'react';
import {
  updateUserDetails,
  updatePreference,
  fetchUserDetails,
  randominvestmentdata,
  currentuser,
} from '@/services/userupdate';

const userData = {
  "name": "John Doe",
  "email": "johndoe@example.com",
  "dateOfBirth": "1990-01-01T00:00:00.000Z",
  "permanentAddress": "123 Main St",
  "postalCode": "12345",
  "username": "johndoe",
  "password": "securepassword123",
  "presentAddress": "456 Elm St",
  "city": "Sample City",
  "country": "Sample Country",
  "profilePicture": "profile-pic-url",
  "preference": {
    "currency": "USD",
    "sentOrReceiveDigitalCurrency": true,
    "receiveMerchantOrder": true,
    "accountRecommendations": true,
    "timeZone": "America/New_York",
    "twoFactorAuthentication": true,
  },
}
const ApiTestComponent = () => {

  const handleUpdateUserDetails = async () => {
    try {
      const response = await updateUserDetails(userData);
      console.log('Update User Details Response:', response);
    } catch (error) {
      console.error('Error:', error);
    }
  };

  const handleUpdatePreference = async () => {
    try {
      console.log(userData)
      const response = await updatePreference(userData);
      console.log('Update Preference Response:', response);
    } catch (error) {
      console.error('Error:', error);
    }
  };

  const handleFetchUserDetails = async () => {
    const userId = 'your-user-id';
    try {
      const response = await fetchUserDetails(userId);
      console.log('Fetch User Details Response:', response);
    } catch (error) {
      console.error('Error:', error);
    }
  };

  const handleRandomInvestmentData = async () => {
    const userId = 'your-user-id';
    try {
      const response = await randominvestmentdata(userId);
      console.log('Random Investment Data Response:', response);
    } catch (error) {
      console.error('Error:', error);
    }
  };

  const handleCurrentUser = async () => {
    const userId = 'your-user-id';
    try {
      const response = await currentuser(userId);
      console.log('Current User Response:', response);
    } catch (error) {
      console.error('Error:', error);
    }
  };

  return (
    <div className='flex flex-col'>
      <button onClick={handleUpdateUserDetails}>Update User Details</button>
      <button onClick={handleUpdatePreference}>Update Preference</button>
      <button onClick={handleFetchUserDetails}>Fetch User Details</button>
      <button onClick={handleRandomInvestmentData}>Random Investment Data</button>
      <button onClick={handleCurrentUser}>Current User</button>
    </div>
  );
};

export default ApiTestComponent;*/}
