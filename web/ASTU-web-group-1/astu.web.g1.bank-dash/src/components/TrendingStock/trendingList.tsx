import React from "react";
import Trendingcard from "./trendingcard";

const TrendingList = () => {
  const data = [
    { name: "Product A", price: "$25.00", pers: "+5%", color: true },
    { name: "Product B", price: "$45.00", pers: "+10%", color: true },
    { name: "Product C", price: "$30.00", pers: "-8%", color: false },
    { name: "Product D", price: "$60.00", pers: "+12%", color: true },
    { name: "Product E", price: "$15.00", pers: "-4%", color: false },
    { name: "Product F", price: "$75.00", pers: "+15%", color: true },
    { name: "Product G", price: "$20.00", pers: "-6%", color: false },
    { name: "Product H", price: "$50.00", pers: "-11%", color: false },
    { name: "Product I", price: "$40.00", pers: "+9%", color: true },
    { name: "Product J", price: "$55.00", pers: "+13%", color: true },
  ];
  return (
    <div className="w-full md:w-2/5 px-2">
      <h1 className="py-2 mb-4 w-fit  font-Inter font-[600] text-xl text-deepNavy">
        Trending Stock
      </h1>
      <div
        className="overflow-x-auto bg-white md:px-4 max-h-[241px] min[890px]:max-h-[290px] lg:max-h-[243px] rounded-2xl md:rounded-2xl"

        style={{
          scrollbarWidth: "none",
          msOverflowStyle: "none",
        }}
      >
        <table className="w-full text-left divide-y">
          <thead className="sticky top-0 bg-white z-10">
            <tr>
              <th className="font-[500] font-Inter text-14px text-blue-steel py-3 pl-4">
                <span className=" md:hidden lg:inline">SL</span>
                No
              </th>
              <th className="font-[500] font-Inter text-14px text-blue-steel py-3 pl-4">
                Name
              </th>
              <th className="font-[500] font-Inter text-14px text-blue-steel py-3 pl-4">
                Price
              </th>
              <th className="font-[500] font-Inter text-14px text-blue-steel py-3 pl-4">

                Return
              </th>
            </tr>
          </thead>
          <tbody className="w-full">
            {data.map((datas, index) => (
              <>
                <Trendingcard
                  key={index}
                  No={`${index + 1}`}
                  Name={datas.name}
                  price={datas.price}
                  Return={datas.pers}
                  Color={datas.color}
                />
              </>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default TrendingList;
