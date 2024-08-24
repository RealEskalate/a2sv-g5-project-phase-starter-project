import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import RecentTransctionSkeleton from "../../../components/AllSkeletons/RecentTransactionSkeleton/recentTransactionSkeleton";

const Other = () => {
  return (
    <div>
      <h1 className="text-16px md:text-18px xl:text-22px text-[#333B69] font-semibold">
        Recent Transaction
      </h1>
      <div className="bg-white p-8 rounded-3xl">
        <Tabs defaultValue="all_transaction" className="w-full">
          <TabsList className="w-full mb-6 md:justify-start bg-white">
            <TabsTrigger
              value="all_transaction"
              className="data-[state=active]:border-b-transparent border-b-4 border-b-transparent data-[state=active]:text-[#1814f3] data-[state=active]:border-b-[#1814f3]"
            >
              All Transaction
            </TabsTrigger>
            <TabsTrigger
              value="income"
              className=" data-[state=active]:border-b-transparent border-b-4 border-b-transparent data-[state=active]:text-[#1814f3] data-[state=active]:border-b-[#1814f3]"
            >
              Income
            </TabsTrigger>
            <TabsTrigger
              value="expence"
              className=" data-[state=active]:border-b-transparent border-b-4 border-b-transparent data-[state=active]:text-[#1814f3] data-[state=active]:border-b-[#1814f3]"
            >
              Expense
            </TabsTrigger>
          </TabsList>

          <TabsContent value="all_transaction">
            <RecentTransctionSkeleton />
          </TabsContent>
          <TabsContent value="income">
            <RecentTransctionSkeleton />
          </TabsContent>
          <TabsContent value="expence">
            <RecentTransctionSkeleton />
          </TabsContent>
        </Tabs>
      </div>
    </div>
  );
};

export default Other;
