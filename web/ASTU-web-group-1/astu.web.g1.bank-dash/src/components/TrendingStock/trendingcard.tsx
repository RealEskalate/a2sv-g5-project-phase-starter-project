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
    <tr className="">
      <td className="font-[400] font-Inter text-sm py-1.5 px-1">{No}</td>
      <td className="font-[400] font-Inter text-sm py-1.5 px-1 truncate">
        {Name}
      </td>
      <td className="font-[400] font-Inter text-sm py-1.5 px-1">{price}</td>
      <td
        className={`font-[500] py-1.5 px-1 text-16px ${
          Color ? "text-mintGreen" : "text-candyPink"
        }`}
      >
        {Return}
      </td>
    </tr>
  );
};

export default Trendingcard;
