import ActiveLoansOverviewTable from "@/components/ActiveLoansOverviewTable/ActiveLoansOverviewTable";
import Loansitem from "@/components/LoansItems/Loansitem";
import React from "react";
import StoreProvider from "@/providers/StoreProvider";
import { Store } from "lucide-react";
export default function page() {
  return (
    <div className="flex flex-col gap-5">
      <Loansitem />
      <StoreProvider>
        <ActiveLoansOverviewTable />
      </StoreProvider>
    </div>
  );
}
