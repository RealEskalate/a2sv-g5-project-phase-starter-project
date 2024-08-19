import ActiveLoansOverviewTable from "@/components/ActiveLoansOverviewTable/ActiveLoansOverviewTable";
import Loansitem from "@/components/LoansItems/Loansitem";

export default function page() {
  return (
    <div className="flex flex-col gap-5">
      <Loansitem />
      <ActiveLoansOverviewTable />
    </div>
  );
}
