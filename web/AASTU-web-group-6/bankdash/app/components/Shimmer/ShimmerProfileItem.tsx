const ShimmerProfileItem = () => {
  return (
    <div className="profile-item flex flex-col gap-1 p-6 items-center justify-center dark:text-gray-300 animate-pulse">
      {/* Shimmer for Image */}
      <div className="shimmer-image rounded-full bg-gray-300 dark:bg-gray-700 w-[70px] h-[70px]"></div>

      {/* Shimmer for Name */}
      <div className="shimmer-text name h-4 w-24 bg-gray-300 dark:bg-gray-700 mt-4 rounded"></div>

      {/* Shimmer for Role */}
      <div className="shimmer-text role h-3 w-20 bg-gray-300 dark:bg-gray-700 mt-2 rounded"></div>
    </div>
  );
};

export default ShimmerProfileItem;
