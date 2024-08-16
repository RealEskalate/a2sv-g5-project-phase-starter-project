// components/columns.ts

import React from 'react';  // Import React
import { ColumnDef } from "@tanstack/react-table";

export const columns: ColumnDef<any>[] = [
  { header: "Description", accessorKey: "column1" },
  { header: "Transaction ID", accessorKey: "column2" },
  { header: "Type", accessorKey: "column3" },
  { header: "Card", accessorKey: "column4" },
  { header: "Date", accessorKey: "column5" },
  { header: "Amount", accessorKey: "column6" },
  {
    header: "Receipt",
    accessorKey: "column7",
  }
];

export default columns;
