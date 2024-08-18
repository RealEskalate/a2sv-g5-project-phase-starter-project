import React, { useState } from "react";

const Trendingcard = ({
  No,
  Name,
  price,
  Return,
  Color,
}: {
  No: string;
  Name: string;
  price: string;
  Return: string;
  Color: boolean;
}) => {
  return (
    <tr>
      <td className="font-[400] font-Inter text-sm py-1.5 pl-4">{No}</td>
      <td className="font-[400] font-Inter text-sm py-1.5 pl-4 truncate">
        {Name}
      </td>
      <td className="font-[400] font-Inter text-sm py-1.5 pl-4 truncate">{price}</td>
      <td
        className={`font-[500] py-1.5 pl-4 truncate text-16px ${

          Color ? "text-mintGreen" : "text-candyPink"
        }`}
      >
        {Return}
      </td>
    </tr>
  );
};

export default Trendingcard;
