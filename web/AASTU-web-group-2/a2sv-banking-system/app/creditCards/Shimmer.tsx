export const ShimmerCreditCard = () => {
  return (
    <div className="flex justify-around items-center bg-gray-200 p-3 lg:p-4 rounded-2xl shadow-sm animate-pulse">
      <div className="w-10 h-10 bg-gray-300 rounded-full"></div>
      <div className="flex gap-9">
        <div>
          <div className="bg-gray-300 h-4 w-24 rounded"></div>
          <div className="bg-gray-200 h-3 w-20 mt-2 rounded"></div>
        </div>
        <div className="hidden lg:block">
          <div className="bg-gray-300 h-4 w-24 rounded"></div>
          <div className="bg-gray-200 h-3 w-20 mt-2 rounded"></div>
        </div>
      </div>
      <div className="bg-gray-300 h-5 w-16 rounded"></div>
    </div>
  );
};

export const ShimmerMainCreditCard = () => {
  return (
    <div className="border rounded-3xl my-4 mx-2">
      <div className="relative w-full bg-gradient-to-b from-gray-200 to-gray-300 rounded-3xl shadow-md h-[230px] min-w-[350px] animate-pulse">
        <div className="flex justify-between items-start px-6 py-4">
          <div className="space-y-2">
            <div className="bg-gray-300 h-4 w-24 rounded"></div>
            <div className="bg-gray-400 h-6 w-32 rounded"></div>
          </div>
          <div className="bg-gray-300 h-8 w-8 rounded-full"></div>
        </div>

        <div className="flex justify-between gap-12 mt-4 px-6">
          <div className="space-y-2">
            <div className="bg-gray-300 h-3 w-16 rounded"></div>
            <div className="bg-gray-400 h-4 w-20 rounded"></div>
          </div>
          <div className="pr-8 space-y-2">
            <div className="bg-gray-300 h-3 w-16 rounded"></div>
            <div className="bg-gray-400 h-4 w-20 rounded"></div>
          </div>
        </div>

        <div className="relative mt-8 py-4 px-6">
          <div className="bg-gray-400 h-8 w-full rounded"></div>
        </div>
      </div>
    </div>
  );
};

export const ShimmerPieChartPage = () => {
  return (
    <div className="flex flex-col gap-5 bg-white rounded-2xl py-3 justify-center px-4 items-center shadow-sm h-full animate-pulse">
      <div className="flex flex-col w-52 shadow-none bg-transparent border-none">
        <div className="mx-auto aspect-square w-full max-w-[300px] bg-gray-200 rounded-full"></div>
        <div className="mt-6 space-y-3">
          <div className="flex justify-center gap-5">
            <div className="w-3 h-3 bg-gray-300 rounded-full"></div>
            <div className="w-20 h-4 bg-gray-300 rounded"></div>
          </div>
          <div className="flex justify-center gap-5">
            <div className="w-3 h-3 bg-gray-300 rounded-full"></div>
            <div className="w-20 h-4 bg-gray-300 rounded"></div>
          </div>
        </div>
      </div>
    </div>
  );
};
