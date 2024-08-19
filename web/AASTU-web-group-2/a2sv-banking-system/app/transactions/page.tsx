"use client";
import React, { useState } from 'react';
import TransactionsList from '../components/Page2/TransactionsList';
import WhiteCard from '../components/Page2/WhiteCard';
import Card from '../components/Page2/Card';
import Tabs from '../components/Tabs';
import BarChart from '../components/Page2/BarChart';

const Page = () => {
    const [activeTab, setActiveTab] = useState('All Transactions');

    const handleTabChange = (tab: string) => {
        setActiveTab(tab);
    };

    return (
        <div className="bg-[#f5f7fa] py-4 px-8 max-w-full">
            <div className="mb-4">
                <div className='flex flex-col md:flex-row md:space-x-8 '>
                    <div className="w-full md:w-1/3 lg:w-3/5">
                        <div className="pt-4 flex items-center justify-between">
                            <h2 className="text-xl font-bold text-[#343C6A]">My Cards</h2>
                            <button className="px-4 py-2 text-sm font-bold text-[#343C6A] border border-none">
                                + Add Card
                            </button>
                        </div>

                        <div className='flex overflow-x-auto space-x-6 scrollbar-hide gap-16 [&::-webkit-scrollbar]:hidden'>
                            <div className="flex-shrink-0 w-72">
                                <Card
                                    balance="$5,756"
                                    cardHolder="Eddy Cusuma"
                                    validThru="12/22"
                                    cardNumber="3778 **** **** 1234"
                                    filterClass=""
                                    bgColor="from-[#4C49ED] to-[#0A06F4]"
                                    textColor="text-white"
                                    iconBgColor="bg-opacity-10"
                                    showIcon={true}
                                />
                            </div>
                            <div className="flex-shrink-0 w-72">
                                <WhiteCard />
                            </div>
                             <div className="flex-shrink-0 w-72">
                                <Card
                                    balance="$5,756"
                                    cardHolder="Eddy Cusuma"
                                    validThru="12/22"
                                    cardNumber="3778 **** **** 1234"
                                    filterClass=""
                                    bgColor="from-[#4C49ED] to-[#0A06F4]"
                                    textColor="text-white"
                                    iconBgColor="bg-opacity-10"
                                    showIcon={true}
                                />
                            </div>
                        </div>
                    </div>
                    <div className="w-full md:w-1/3 lg:w-1/5 mt-8 md:mt-0 pt-4 pb-8">
                        <h2 className="text-xl font-bold text-[#343C6A]">My Expense</h2>
                        <div className="w-full max-h-[200px] flex-grow pt-6">
                            <BarChart />
                        </div>
                    </div>
                </div>
            </div>
            <div className="mb-4 md:w-4/5 lg:w-10/12">
                <h2 className="text-xl font-bold mb-4 pt-6 text-[#343C6A]">Recent Transactions</h2>
                <Tabs
                    tabs={['All Transactions', 'Income', 'Expense']}
                    activeTab={activeTab}
                    onTabChange={handleTabChange}
                />
                <TransactionsList />
            </div>
        </div>
    );
};

export default Page;
