import React from "react";
interface TableButtonProps {
  classname?: string;
  text: string;
  type?: "button" | "submit" | "reset";
}
const TableButton = ({ classname, text, type }: TableButtonProps) => {
  return (
    <button
      className={`${classname} py-1 md:py-2 outline-none border rounded-3xl font-medium text-12px lg:text-15px text-black border-gray-dark hover:text-blue-bright hover:border-blue-bright focus:text-blue-bright focus:border-blue-bright`}
      type={type}
    >
      {text}
    </button>
  );
};

export default TableButton;
