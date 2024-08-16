import React from 'react';

interface Column<T> {
  Header: string;
  accessor: keyof T;
  Cell?: (props: { [key: string]: any }) => React.ReactNode;
}

interface TableProps<T> {
  columns: Column<T>[];
  data: T[];
}

const Table = <T extends {}>({ columns, data }: TableProps<T>) => {
  return (
    <div className=" rounded-3xl shadow-md p-4 bg-white ">
      <div className="overflow-x-auto">
        <table className="min-w-full divide-y divide-[#E6EFF5] rounded-lg bg-white">
          <thead className="bg-white font-inter font-medium hidden sm:table-header-group">
            <tr>
              {columns.map((column, index) => (
                <th
                  key={index}
                  className="px-6 py-3 text-left text-xs font-medium text-[#718EBF]"
                >
                  {column.Header}
                </th>
              ))}
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-[#E6EFF5]">
            {data.map((row, rowIndex) => (
              <tr key={rowIndex}>
                {columns.map((column, colIndex) => (
                  <td
                    key={colIndex}
                    className={`px-6 py-4 whitespace-nowrap font-normal ${
                      column.accessor === 'description' || column.accessor === 'amount'
                        ? ''
                        : 'hidden sm:table-cell'
                    }`}
                  >
                    {column.Cell
                      ? column.Cell(row)
                      : (row[column.accessor] as unknown as React.ReactNode)}
                  </td>
                ))}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default Table;
