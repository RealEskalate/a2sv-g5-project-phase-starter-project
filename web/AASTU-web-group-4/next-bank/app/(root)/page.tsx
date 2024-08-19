"use client"
{/*
import React from "react";
import { Color } from "chart.js";
import { colors, logo } from "@/constants";
import DesktopCreditCart from "@/components/DesktopCreditCard";
import RecentTransaction from "@/components/Recent Transaction";
import ExpensesChart from "@/components/ExpensesCart";
import { icons, Import } from "lucide-react";
import { text } from "stream/consumers";
import BarChart from "@/components/BarChart";
import PieChart from "@/components/PieChart";
import QuickTransfer from "@/components/QuickTransfer";
import LineChart from "@/components/LineChart";



const page = () => {
  return (
    <div className={`${colors.graybg} p-6 md:ml-64 max-w-full md:p-12`}>
      <div className="flex flex-col justify-between max--full  md:flex-row max-w-full gap-10 ">
        <div className="max-w-full py-4 md:w-3/5">
          <div className={`${colors.navbartext} flex justify-between `}>
            <h1 className="font-bold text-2xl">My Cards</h1>
            <p className="my-2"> See All</p>
          </div>

          <div className="max-w-sm md:max-w-full">
            <div className="flex gap-3 overflow-x-auto max-w-full md:w-auto">
              <div className=" py-3 ">
                <DesktopCreditCart
                  bgColor={colors.blue}
                  textColor={colors.textwhite}
                />
              </div>
              <div className=" py-3 ">
                <DesktopCreditCart
                  bgColor={colors.white}
                  textColor={colors.textblack}
                />
              </div>
            </div>
          </div>
        </div>
        <div className="w-full md:w-2/5">
          <div className={`${colors.navbartext} flex justify-between  py-4`}>
            <h1 className="font-bold text-2xl">Recent Transaction</h1>
          </div>
          <div className="">
            <RecentTransaction
              icon={logo.RT1}
              color={colors.lightorange}
              colortext={colors.textred}
              name="deposit from my"
              date="23 january 2023"
              amount="-850"
            />
            <RecentTransaction
              icon={logo.RT2}
              color={colors.lightblue}
              colortext={colors.textgreen}
              name="deposit from my"
              date="23 january 2023"
              amount="+2500"
            />
            <RecentTransaction
              icon={logo.RT3}
              color={colors.lightgreen}
              colortext={colors.textgreen}
              name="deposit from my"
              date="23 january 2023"
              amount="+5400"
            />
          </div>
        </div>
      </div>
      <div className="flex flex-col justify-between gap-10 md:flex-row max-w-full">
        <div className="w-full md:w-3/5">
          <div className={`${colors.navbartext} flex justify-between py-4`}>
            <h1 className="font-bold text-2xl">Weekly Activity</h1>
          </div>
          <div>
            <BarChart />
          </div>
        </div>
        <div className="w-full md:w-2/5 py-6">
          <div className={`${colors.navbartext}`}>
            <h1 className="font-bold text-2xl">Expense Statstics</h1>
          </div>
          <div className="py-4">
            <PieChart />
          </div>
        </div>
      </div>

      <div className="flex flex-col justify-between w-full  md:flex-row max-w-full gap-10 ">
        <div className=" md:w-2/5 py-4 ">
          <div className={`${colors.navbartext} flex justify-between `}>
            <h1 className="font-bold text-2xl">Quick Transfer</h1>
          </div>
          <div className="flex justify- gap-3 overflow-x-auto md: max-w-full">
            <div className="flex py-3 ">
              {" "}
              <QuickTransfer />
            </div>
          </div>
        </div>
        <div className="md:w-3/5">
          <div className={`${colors.navbartext} flex justify-between  py-4`}>
            <h1 className="font-bold text-2xl">Balance History</h1>
          </div>
          <div className="">
            <LineChart />
          </div>
        </div>
      </div>
    </div>
  );
};

export default page;

*/}

import React from 'react';
import {
  updateUserDetails,
  updatePreference,
  fetchUserDetails,
  randominvestmentdata,
  currentuser,
} from '@/services/userupdate';

const ApiTestComponent = () => {

  const handleUpdateUserDetails = async () => {
    const userData = {
      "name": "John wick",
      "email": "johnwick@example.com",
      "dateOfBirth": "1990-01-01T00:00:00.000Z",
      "permanentAddress": "123 Main St",
      "postalCode": "45678",
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
    try {
      const response = await updateUserDetails(userData);
      console.log('Updated User Details Response:', response);
    } catch (error) {
      console.error('Error:', error);
    }
  };

  const handleUpdatePreference = async () => {
    const userData = { /* Your preference data here */ };
    try {
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

export default ApiTestComponent;
