import { CardSkeleton } from "./CardSkeleton";

const MyCardsLoad = () => {
  return (
    <div className="flex overflow-x-auto space-x-4 md:pr-3 pr-1 scrollbar-none">
      <CardSkeleton />
      <CardSkeleton />
      <CardSkeleton />
    </div>
  );
};

export default MyCardsLoad;
