import { CardSkeleton } from "./CardSkeleton";

const MyCardsLoad = ({ count }: { count: number }) => {
  return (
    <div className="flex overflow-x-auto space-x-4 md:pr-3 pr-1 scrollbar-none">
      {Array.from({ length: count }).map((_, index) => (
        <CardSkeleton key={index} />
      ))}
    </div>
  );
};

export default MyCardsLoad;
