"use client"
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
        <div className="min-h-screen bg-[#f5f7fa] p-4">
            <div className="bg-[#f5f7fa] p-4 rounded-lg shadow-md">

                <div className="mb-4">
                    <div className="flex justify-between items-center">
                        <h2 className="text-xs font-semibold text-[#343C6A]">My Cards</h2>
                        <button className="px-4 py-2 text-sm font-semibold text-[#343C6A] fixed top-0 right-0 mr-6 mt-6">
                            + Add Card
                        </button>
                    </div>

                    <div className="grid grid-cols-1 md:grid-cols-2 gap-6 pb-4">
                        <div className="flex space-x-6 pb-4">
                            <div className="flex-shrink-0">
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
                            <div className="flex-shrink-0">
                                <WhiteCard />
                            </div>
                        </div>

                        <div className="hidden md:block">
                            <h2 className="text-xl font-bold mb-4 text-[#343C6A] font-inter">My Expense</h2>
                            <BarChart />
                        </div>
                    </div>

                    </div>

                    <div className="block md:hidden mb-4">
                        <h2 className="text-xl font-bold mb-4 text-[#343C6A] font-inter">My Expense</h2>
                        <BarChart />
                </div>
                

                <div className="mb-4">
                    <div>
                        <h2 className="text-xl font-bold mb-4 text-[#343C6A] font-inter">Recent Transactions</h2>
                    </div>
                    <div>
                        <Tabs
                            tabs={['All Transactions', 'Income', 'Expense']}
                            activeTab={activeTab}
                            onTabChange={handleTabChange}
                        />
                    </div>
                    <div>
                        <TransactionsList />
                    </div>
                </div>

            </div>
        </div>
    );
};

export default Page;
