'use client'
import {Table,TableBody, TableCell, TableFooter, TableHead,TableHeader,TableRow,} from "@/components/ui/table"
import {getLoansAll} from '@/lib/loanApies';
import { useState, useEffect } from "react";
import { useUser } from "@/contexts/UserContext";
import { Loading } from "@/app/dashboard/_components/Loading";
import {
  Pagination,
  PaginationContent,
  PaginationEllipsis,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination";

export function TableDemo() {
    const { isDarkMode } = useUser();
    const rowsPerPage = 6;
    const [table, setTable] = useState<any>([])
    const [loading, setLoading] = useState(true)
    const [totalPages, setTotalPages] = useState<number>(5);
    const [currentPage, setCurrentPage] = useState(1);


    const handlePageChange = (page: number) => {
      setCurrentPage(page);
    };
  
    const handleNextPage = () => {
      if (currentPage < totalPages) {
        setCurrentPage(currentPage + 1);
      }
    };
  
    const handlePreviousPage = () => {
      if (currentPage > 1) {
        setCurrentPage(currentPage - 1);
      }
    };
  
    const renderPageButtons = () => {
      const pagesToShow = Math.min(totalPages, 3);
      const startPage =
        currentPage <= 2 || totalPages <= 4
          ? 1
          : currentPage > totalPages - 2
          ? totalPages - 3
          : currentPage - 1;
  
      return Array.from({ length: pagesToShow }, (_, index) => {
        const page = startPage + index;
        return (
          <PaginationItem key={page} onClick={() => handlePageChange(page)} className={`${ isDarkMode ? "bg-gray-600 text-gray-50" : " bg-gray-200"} rounded-lg mx-1`} >
            <PaginationLink {...(page === currentPage ? { isActive: true } : {})} className={`${page === currentPage ? (
              isDarkMode ? "text-gray-900" : "bg-slate-700 text-white"
            ):("") }`}>
            {page}
            </PaginationLink>
          </PaginationItem>
        );
      });
    };

    useEffect(() => {
      const fetchData = async () => {
        try {
          const data = await getLoansAll(currentPage - 1, rowsPerPage);
          if (data==null){
            throw new Error("failed to get data");
          }
          setTable(data.content);
          setTotalPages(data?.totalPages || 7);
          setLoading(false);
          
        }catch (error) {
          console.error("An error occurred on card:", error);
        }
      };
      fetchData();
    }, [currentPage]);


    return (
      loading ? (<Loading/>): 
        (
        <div className="flex flex-col justify-center rounded-2xl bg-transparent">
          <Table className="bg-white w-[90%] mx-auto rounded-2xl my-4" >
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
          <Pagination className="mx-auto">
            <PaginationContent className="hover:cursor-pointer">
              <PaginationItem className={`${ isDarkMode ? "bg-gray-600 text-gray-50" : "bg-gray-200"} rounded-xl`} >
                <PaginationPrevious onClick={handlePreviousPage} aria-disabled={currentPage === 1}/>
              </PaginationItem>
              {renderPageButtons()}
              <PaginationItem className={`${ isDarkMode ? "bg-gray-600 text-gray-50" : "bg-gray-200"} rounded-lg`}>
                <PaginationNext onClick={handleNextPage} aria-disabled={currentPage === totalPages} />
              </PaginationItem>
            </PaginationContent>
          </Pagination>
        </div>
    ))
    
}
