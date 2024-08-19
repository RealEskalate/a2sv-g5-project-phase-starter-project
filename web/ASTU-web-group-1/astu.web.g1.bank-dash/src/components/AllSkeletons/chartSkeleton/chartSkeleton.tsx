import { Skeleton } from "../../ui/skeleton";

const YearlyTotalInvestment = () => {
  return (
    <div className="w-full md:w-1/2">
      <Skeleton className="h-5 w-52 bg-slate-200 mt-3 mb-5 text-20px py-2 font-semibold" />

      <Skeleton className="bg-slate-200 h-64 p-6 rounded-3xl">
        <Skeleton className="bg-white w-full h-full"></Skeleton>
      </Skeleton>
    </div>
  );
};

export default YearlyTotalInvestment;
