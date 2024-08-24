const TableSkeleton = () => {
    return (
      <div className="bg-white rounded-3xl animate-pulse">
        <table className="w-full border-collapse">
          <thead>
            <tr>
              <th className="font-medium text-lg text-gray-300 p-4">SL No</th>
              <th className="font-medium text-lg text-gray-300 p-4">Name</th>
              <th className="font-medium text-lg text-gray-300 p-4">Price</th>
              <th className="font-medium text-lg text-gray-300 p-4">Return</th>
            </tr>
            <tr className="bg-gray-100 h-px">
              <td colSpan={4}></td>
            </tr>
          </thead>
          <tbody>
            {Array(5).fill("#cccccc").map((_, index) => (
              <tr key={index}>
                <td className="p-4">
                  <div className="h-4 bg-gray-200 rounded w-10"></div>
                </td>
                <td className="p-4">
                  <div className="h-4 bg-gray-200 rounded w-24"></div>
                </td>
                <td className="p-4">
                  <div className="h-4 bg-gray-200 rounded w-20"></div>
                </td>
                <td className="p-4">
                  <div className="h-4 bg-gray-200 rounded w-12"></div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    );
  };
  
  export default TableSkeleton;
  