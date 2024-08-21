'use client'
import {Table,TableBody, TableCell, TableFooter, TableHead,TableHeader,TableRow,} from "@/components/ui/table"
import {getLoansAll} from '@/lib/loanApies';
import { useState, useEffect, Suspense } from "react";

function Loading(){
  return (
    <div className="flex justify-center items-center mt-10">
      <div className="animate-spin rounded-full h-32 w-32 border-t-4 border-dotted border-blue-600"></div>
    </div>
  );
}
export function TableDemo() {
    const [table, setTable] = useState<any>([])
    const [loading, setLoading] = useState(true)
    useEffect(() => {
      const fetchData = async () => {
        try {
          const tableData = await getLoansAll();
          if (tableData==null){
            throw new Error("failed to get data");
          }
          setTable(tableData);
          setLoading(false);
          
        }catch (error) {
          console.error("An error occurred on card:", error);
        }
      };
      fetchData();
    }, []);

    return (
      loading ? (<Loading/>): 

        (<Table className="bg-white w-[90%] mx-auto rounded-2xl my-4" >
          <TableHeader className="p-10">
            <TableRow className="text-[#243a61] font-[600] p-10">
              <TableHead className="hidden md:table-cell text-center" >SL No</TableHead>
              <TableHead className="text-center" >Loan Money</TableHead>
              <TableHead className="text-center">Left to repay</TableHead>
              <TableHead className="hidden md:table-cell text-center">Duration</TableHead>
              <TableHead className="hidden md:table-cell text-center">Interest rate</TableHead>
              <TableHead className="hidden md:table-cell text-center">Installment</TableHead>
              <TableHead className="text-center">Repay</TableHead>
            </TableRow>
          </TableHeader>

          <TableBody >
            {table.map((Invoice:any,idx:number) => (
              <TableRow key={idx} className={ idx%2==0 ? ("bg-[#dfdfdf59]"):("bg-white")} >
                <TableCell className="hidden md:table-cell text-center ">{idx+1}</TableCell>
                <TableCell className="text-center">{Invoice.loanAmount}</TableCell>
                <TableCell className="text-center">{Invoice.amountLeftToRepay}</TableCell>
                <TableCell className="hidden md:table-cell text-center" >{Invoice.duration}</TableCell>
                <TableCell className="hidden md:table-cell text-center">{Invoice.interestRate}</TableCell>
                <TableCell className="hidden md:table-cell text-center">{Invoice.installment}</TableCell>
                <TableCell className="text-center">
                  <button className="border border-1 border-gray-800 rounded-full m-auto hover:text-blue-700 hover:border-blue-700 text-[10px] md:text-[15px] p-2 w-[65px] md:w-[75px]">
                    Repay
                  </button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>

          <TableFooter className="text-red-500">
            <TableRow className="table-cell md:hidden">
            <TableCell className="hidden md:table-cell">Total</TableCell>
            </TableRow>
            <TableRow >  
              <TableCell className="hidden md:table-cell text-center">Total</TableCell> 
              <TableCell className="text-center">{table.map((a:any) => a.loanAmount).reduce(function(a:number, b:number){return a + b;})}</TableCell>
              <TableCell colSpan={3} className="hidden md:table-cell text-center">{table.map((a:any) => a.amountLeftToRepay).reduce(function(a:number, b:number){return a + b;})}</TableCell>
              <TableCell className="text-center" >{table.map((a:any) => a.installment).reduce(function(a:number, b:number){return a + b;})}</TableCell>
            </TableRow>
          </TableFooter>

        </Table>
    ))
    
}
