import { Skeleton } from "@/components/ui/skeleton";
import AddNewCardFormSkeleton from "./AddNewCardFormSkeleton";

const AddNewCardSkeleton = () => {
  return (
    <div className=" w-full md:w-8/12">
      <Skeleton className="text-[#333B69] pb-2 bg-slate-200 w-48 h-6 mb-2" />
      <div className="bg-white p-5 rounded-3xl space-y-6">
        <div>
          <Skeleton className="text-15px py-2 bg-slate-200 w-full h-5 mb-1" />
          <Skeleton className="text-15px py-2 bg-slate-200 w-full h-5" />
        </div>
        <AddNewCardFormSkeleton />
      </div>
    </div>
  );
};

export default AddNewCardSkeleton;
