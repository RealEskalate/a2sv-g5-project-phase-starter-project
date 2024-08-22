import React from 'react'
import { Table, TableHeader, TableBody, TableHead, TableRow, TableCell } from "@/components/ui/table"

// Define the DataItem interface
interface DataItem {
    Description: string;
    TransactionID: string;
    Type: string;
    Card: string;
    Date: string;
    amount: Number; // Assuming amount is a string, change to number if necessary
    Receipt: string;
}

// Define the Props interface
interface Props {
    Data: DataItem[];
}

const TableUI: React.FC<Props> = ({ Data }) =>{
  return (
    <div className='p-4 rounded-2xl bg-white'>
        <Table>
            <TableHeader>
                <TableRow className='text-slate-400'>
                    <TableHead>Description</TableHead>
                    <TableHead>Transaction ID</TableHead>
                    <TableHead>Type</TableHead>
                    <TableHead>Card</TableHead>
                    <TableHead>Date</TableHead>
                    <TableHead>Amount</TableHead>
                    <TableHead>Receipt</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                {Data.map((item, index) => (
                    <TableRow key={index}>
                        <TableCell className='flex items-center gap-2'>
                            {item.amount.valueOf() > 0 ? <img src='/down.png' alt='down' /> : <img src='/up.png' alt='up' />}
                            {item.Description}
                        </TableCell>
                        <TableCell>{item.TransactionID}</TableCell>
                        <TableCell>{item.Type}</TableCell>
                        <TableCell>{item.Card}</TableCell>
                        <TableCell>{item.Date}</TableCell>
                        <TableCell className={`${item.amount.valueOf() > 0 ? 'text-green-600' : 'text-red-600'}`}>
                            {item.amount.valueOf() > 0 ? `$${item.amount}` : `-$${-1*item.amount.valueOf()}.00` }
                        </TableCell>
                        <TableCell>{item.Receipt}</TableCell>
                    </TableRow>
                ))}
            </TableBody>
        </Table>
    </div>
  )
}

export default TableUI;
