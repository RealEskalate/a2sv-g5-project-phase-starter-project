import React from 'react';
import { TransactionTableProps } from "@/constants/transactions/types";
import { Button } from '@mui/material';
import Image from 'next/image';
import { LightTextColor, NormalTextColor } from '@/constants/transactions/colors';

type Props = {
    TableData: TransactionTableProps[];
}

export default function ExpenseTable({ TableData }: Props) {
    return (
        <div className="overflow-x-auto bg-white rounded-3xl">
            <table className="min-w-full  border border-gray-200">
                <thead className=" text-gray-700">
                    <tr>
                        <th className={"px-6 py-3 border-b text-left text-x-[16px] font-medium  text-[#718EBF]"}>Description</th>
                        <th className={"px-6 py-3 border-b text-left text-x-[16px] font-medium  text-[#718EBF]"}>Transaction ID</th>
                        <th className={"px-6 py-3 border-b text-left text-x-[16px] font-medium  text-[#718EBF]"}>Type</th>
                        <th className={"px-6 py-3 border-b text-left text-x-[16px] font-medium  text-[#718EBF]"}>Card</th>
                        <th className={"px-6 py-3 border-b text-left text-x-[16px] font-medium  text-[#718EBF]"}>Date</th>
                        <th className={"px-6 py-3 border-b text-left text-x-[16px] font-medium  text-[#718EBF]"}>Amount</th>
                        <th className={"px-6 py-3 border-b text-left text-x-[16px] font-medium  text-[#718EBF]"}>Recite</th>
                    </tr>
                </thead>
                <tbody>
                    {TableData.map((value, ind) => (
                        <tr key={ind} className="hover:bg-gray-50">
                            <td className="px-6 py-4 border-b">
                                <div className="flex items-center gap-x-4">
                                    <Image
                                        src={value.amount < 0 ? "/icons/expenseTransction.svg" : "/icons/incomeTransction.svg"}
                                        alt="icon"
                                        width={30}
                                        height={30}
                                    />
                                    <span className="text-sm  text-[#232323]">{value.description}</span>
                                </div>
                            </td>
                            <td className="px-6 py-4 border-b text-sm text-[#232323]">{value.transactionId}</td>
                            <td className="px-6 py-4 border-b text-sm text-[#232323]">{value.type}</td>
                            <td className="px-6 py-4 border-b text-sm text-[#232323]">{value.card}</td>
                            <td className="px-6 py-4 border-b text-sm text-[#232323]">{value.date}</td>
                            <td className={`px-6 py-4 border-b text-sm text-left ${value.amount < 0 ? "text-[#FE5C73]" : "text-[#16DBAA]"}`}>
                                {value.amount < 0 ? `- ${Math.abs(value.amount).toLocaleString('en-US', { style: 'currency', currency: 'USD' })}` : `+ ${value.amount.toLocaleString('en-US', { style: 'currency', currency: 'USD' })}`}
                            </td>
                            <td className="px-6 py-4 border-b text-sm">
                                <Button variant="outlined" size="small" className='rounded-[50px] text-[#123288]'>Download</Button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    )
}
