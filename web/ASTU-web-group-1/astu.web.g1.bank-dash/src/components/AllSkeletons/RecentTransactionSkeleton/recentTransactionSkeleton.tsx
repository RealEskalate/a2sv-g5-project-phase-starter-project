import { Skeleton } from "@/components/ui/skeleton";

const RecentTransctionSkeleton = () => {
  const currentData = [1, 2, 3, 4, 5];
  return (
    <div>
      <div className="flex flex-col gap-4">
        {currentData?.length == 0 ? (
          <div>No transactions found.</div>
        ) : (
          <div className="relative overflow-x-auto bg-white px-4 md:px-6 pt-5 md:pt-6 rounded-2xl md:rounded-2xl">
            <table className="bg-white px-5 lg:px-11 w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
              <thead className=" text-12px md:text-16px font-Lato font-medium text-blue-steel bg-white border-b">
                <tr className="">
                  <th scope="col" className="hidden md:table-cell pb-2">
                    <Skeleton className="h-5 w-full bg-slate-300" />
                  </th>
                  <th scope="col" className=" hidden md:table-cell pb-2">
                    <Skeleton className="h-5 w-full bg-slate-300" />
                  </th>
                  <th scope="col" className="hidden lg:table-cell pb-2">
                    <Skeleton className="h-5 w-full bg-slate-300" />
                  </th>
                  <th scope="col" className="hidden lg:table-cell pb-2">
                    <Skeleton className="h-5 w-full bg-slate-300" />
                  </th>
                  <th scope="col" className="hidden lg:table-cell pb-2">
                    <Skeleton className="h-5 w-full bg-slate-300" />
                  </th>
                  <th scope="col" className="hidden md:table-cell pb-2">
                    <Skeleton className="h-5 w-full bg-slate-300" />
                  </th>
                  <th scope="col" className="hidden lg:table-cell pb-2 w-fit">
                    <Skeleton className="h-5 w-full bg-slate-300" />
                  </th>
                </tr>
              </thead>
              <tbody className="text-12px xl:text-16px text-gray-dark cursor-pointer  hover:bg-gray-100 dark:hover:bg-gray-700">
                {currentData?.map((datax, index) => (
                  <tr
                    key={index}
                    className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-100 dark:hover:bg-gray-700"
                  >
                    <td className="py-3">
                      <Skeleton className="h-4 w-full bg-slate-200" />
                    </td>
                    <td className="hidden md:table-cell py-3">
                      <Skeleton className="h-4 w-full bg-slate-200" />
                    </td>
                    <td className="hidden lg:table-cell py-3">
                      <Skeleton className="h-4 w-full bg-slate-200" />
                    </td>
                    <td className="hidden lg:table-cell py-3">
                      <Skeleton className="h-4 w-full bg-slate-200" />
                    </td>
                    <td className="hidden lg:table-cell py-3">
                      <Skeleton className="h-4 w-full bg-slate-200" />
                    </td>
                    <td className="">
                      <Skeleton className="h-4 w-full bg-slate-200" />
                    </td>
                    <td className="hidden lg:table-cell py-3 w-24 md:w-32">
                      <Skeleton className="h-4 w-full rounded-3xl bg-slate-200" />
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  );
};

export default RecentTransctionSkeleton;
