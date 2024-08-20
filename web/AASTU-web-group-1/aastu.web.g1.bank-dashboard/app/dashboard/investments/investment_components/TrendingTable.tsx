import React from 'react'
import {Table,TableBody, TableCell, TableFooter, TableHead,TableHeader,TableRow,} from "@/components/ui/table"
import { trendingArray } from '@/constants';
export default function TrendingTable(){
  return (
    <div className='text-sm md:text-lg' >
      <Table className='bg-white rounded-2xl'>
        <TableHeader>
            <TableRow>
                <TableHead className='text-sm md:text-lg' >SL No</TableHead>
                <TableHead className='text-sm md:text-lg' >Name</TableHead>
                <TableHead className='text-sm md:text-lg' >Price</TableHead>
                <TableHead className='text-sm md:text-lg' >Return</TableHead>
            </TableRow>
        </TableHeader>
        <TableBody>
            {trendingArray.map((item) => (
                <TableRow key={item.id}>
                    <TableCell>{item.id}</TableCell>
                    <TableCell>{item.name}</TableCell>
                    <TableCell>{item.price}</TableCell>
                    <TableCell>{item.return}</TableCell>
                </TableRow>    
            ))}
        </TableBody>
      </Table>
    </div>
  )
}
