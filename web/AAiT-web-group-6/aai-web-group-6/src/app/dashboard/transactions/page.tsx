"use client";

import PaymentCard from '@/components/transactions/credidCard'
import { ExpenseChart } from '@/components/transactions/transactionCharts'
import { NormalTextColor } from '@/constants/transactions/colors'
import { CredidtCardsDemo, transactionTableData } from '@/constants/transactions/constants'
import { Button } from '@mui/material'
import React, { useState } from 'react'
import { TransactionTableProps } from "@/constants/transactions/types";
import ExpenseTable from '@/components/transactions/expenseTable'


type Props = {}

export default function Transasction({ }: Props) {
    const [activeTab, setActiveTab] = useState('all')
    return (
        <div className="transactions px-[40px]">
            <div className="topside flex gap-x-[30px]  py-6 ">
                <div className="mycards">
                    <div className='flex  justify-between'>
                        <h1 className={'font-semibold ' + NormalTextColor}>My Cards</h1>
                        <Button variant='text' className='text-[#343C6A] font-semibold'>+ Add Card</Button>
                    </div>
                    <div className='flex gap-x-5 md:gap-x-6 lg:gap-x-7 min-w-[376px] max-w-[730px] '>
                        {CredidtCardsDemo.map((value, ind) => {
                            return (
                                <PaymentCard key={ind} {...value} />)
                        })}
                    </div>
                </div>
                <div className="expenses">
                    <h1 className={'font-semibold px-2 py-[6px] ' + NormalTextColor}>My Expense</h1>
                    <div className=' bg-white max-w-[351px] h-[225px] rounded-[20px] flex justify-center items-center'>
                        <ExpenseChart />
                    </div>
                </div>
            </div>
            <div className="bottomside">
                <h1 className={"text-2xl font-semibold " + NormalTextColor}>Recent Transactions</h1>
                <div className="tableData">
                    <ul className='flex gap-x-6 py-6 border-b'>
                        <li><Button variant='text' className={activeTab === "all" ? "  " : ""} >All Transactions</Button></li>
                        <li><Button variant='text'>Income</Button></li>
                        <li><Button variant='text'>Expense</Button></li>
                    </ul>
                    <ExpenseTable TableData={transactionTableData} />
                </div>

            </div>

        </div>
    )
}