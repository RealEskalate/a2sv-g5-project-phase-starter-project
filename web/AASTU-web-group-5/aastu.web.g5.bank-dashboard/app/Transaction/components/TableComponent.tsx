// components/TableComponent.tsx
"use client";
import React from "react";
import {
  ColumnDef,
  useReactTable,
  getCoreRowModel,
  getPaginationRowModel,
  flexRender,
} from "@tanstack/react-table";
import { ChevronLeft, ChevronRight } from "lucide-react";

interface TableProps {
  columns: ColumnDef<any, any>[];
  data: any[];
}

export function TableComponent({ columns, data }: TableProps) {
  console.log(data,22222)
  const tableInstance = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    initialState: {
      pagination: { pageSize: 5 }, // Limit to 5 rows per page
    },
    debugTable: true,
  });

  return (
    <div className="overflow-x-auto ">
      <table className="min-w-full bg-white border rounded-[25px]">
        <thead className="text-[#718EBF] font-Inter ">
          {tableInstance.getHeaderGroups().map((headerGroup) => (
            <tr key={headerGroup.id}>
              {headerGroup.headers.map((header) => (
                <th
                  key={header.id}
                  className="text-left p-2 border-b border-gray-300"
                >
                  {header.isPlaceholder
                    ? null
                    : flexRender(
                        header.column.columnDef.header,
                        header.getContext()
                      )}
                </th>
              ))}
            </tr>
          ))}
        </thead>
        <tbody>
          {tableInstance.getRowModel().rows.map((row) => (
            <tr key={row.id}>
              {row.getVisibleCells().map((cell) => (
                <td
                  key={cell.id}
                  className="p-2 border-b border-gray-300 text-sm text-gray-700"
                >
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>

      <Pagination table={tableInstance} />
    </div>
  );
}

interface PaginationProps {
  table: ReturnType<typeof useReactTable>;
}

function Pagination({ table }: PaginationProps) {
  const pageCount = table.getPageCount();
  const { pageIndex } = table.getState().pagination;

  // Calculate the range of page numbers to display
  const totalPageNumbersToShow = 4;
  let startPage = Math.max(0, pageIndex - Math.floor(totalPageNumbersToShow / 2));
  let endPage = startPage + totalPageNumbersToShow - 1;

  // Ensure we don't exceed the page count
  if (endPage >= pageCount) {
    endPage = pageCount - 1;
    startPage = Math.max(0, endPage - totalPageNumbersToShow + 1);
  }

  // Generate the array of page numbers to render
  const pageNumbers = [];
  for (let i = startPage; i <= endPage; i++) {
    pageNumbers.push(i);
  }

  return (
    <div className="flex justify-end items-center mt-4 space-x-2">
      <button
        className="flex items-center text-gray-600 p-2 hover:text-blue-500 disabled:opacity-50"
        onClick={() => table.previousPage()}
        disabled={!table.getCanPreviousPage()}
      >
        <ChevronLeft size={20} />
        <span className="ml-1">Prev</span>
      </button>

      <div className="flex items-center space-x-1">
        {pageIndex > 0 && startPage > 0 && (
          <button
            onClick={() => table.setPageIndex(0)}
            className={`p-2 text-sm ${pageIndex === 0 ? 'bg-blue-500 text-white' : 'bg-gray-200 text-gray-700'} hover:bg-blue-600 hover:text-white transition-colors`}
          >
            1
          </button>
        )}
        {pageNumbers.map((page) => (
          <button
            key={page}
            onClick={() => table.setPageIndex(page)}
            className={`p-2 text-sm ${page === pageIndex ? 'bg-blue-500 text-white' : 'bg-gray-200 text-gray-700'} hover:bg-blue-600 hover:text-white transition-colors`}
          >
            {page + 1}
          </button>
        ))}
        {pageCount > endPage + 1 && (
          <button
            onClick={() => table.setPageIndex(pageCount - 1)}
            className={`p-2 text-sm ${pageIndex === pageCount - 1 ? 'bg-blue-500 text-white' : 'bg-gray-200 text-gray-700'} hover:bg-blue-600 hover:text-white transition-colors`}
          >
            {pageCount}
          </button>
        )}
      </div>

      <button
        className="flex items-center text-gray-600 p-2 hover:text-blue-500 disabled:opacity-50"
        onClick={() => table.nextPage()}
        disabled={!table.getCanNextPage()}
      >
        <span className="mr-1">Next</span>
        <ChevronRight size={20} />
      </button>
    </div>
  );
}
export default TableComponent;