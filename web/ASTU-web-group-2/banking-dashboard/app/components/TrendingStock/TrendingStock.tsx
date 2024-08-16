import React from "react";

const TrendingStock = () => {
  return (
    <div className="w-[445px] h-[300px] bg-white rounded-[25px] max-sm:w-[325px] max-sm:h-[213px] max-sm:rounded-[15px]">
      <table className="">
        <thead className="">
          <tr className="">
            <th className="font-medium text-[16px] leading-[19.36px]  text-[#718EBF]  max-sm:text-[12px] max-sm:font-normal pl-[30px] pt-[27px]  max-sm:pl-[12px]">
              SL No
            </th>
            <th className="font-medium text-[16px] leading-[19.36px]  pl-[36px] text-[#718EBF] max-sm:pl-[30px] pt-[27px] max-sm:text-[12px] max-sm:font-normal">
              Name
            </th>
            <th className="font-medium text-[16px] pt-[27px] leading-[19.36px]  pl-[86px] text-[#718EBF]  max-sm:text-[12px] max-sm:font-normal max-sm:pl-[59px]">
              Price
            </th>
            <th className="font-medium text-[16px] pt-[27px] leading-[19.36px]  pl-[65px] text-[#718EBF] max-sm:pl-[50px] max-sm:text-[12px] max-sm:font-normal ">
              Return
            </th>
          </tr>
          <tr className="h-[1px] bg-[#F4F5F7] ">
            <td colspan="4" className="p-0 pl-[4px]"></td>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td className="pl-[30px] pt-[22px] font-normal text-[16px] leading-[19.36px] max-sm:text-[12px] max-sm:font-normal max-sm:pt-[12px] max-sm:pl-[14px]">
              01.
            </td>
            <td className="pt-[22px] pl-[52px] font-normal text-[16px] leading-[19.36px] max-sm:text-[12px] max-sm:font-normal max-sm:pt-[12px] max-sm:pl-[37px]">
              Trivago
            </td>
            <td className="pt-[22px] pl-[86px] font-normal text-[16px] leading-[19.36px] max-sm:text-[12px] max-sm:font-normal max-sm:pl-[55px] max-sm:pt-[12px]">
              $520
            </td>
            <td className="pt-[22px] pl-[64px] font-normal text-[16px] leading-[19.36px] max-sm:text-[12px] max-sm:font-normal max-sm:pl-[53px] max-sm:pt-[12px]">
              +5%
            </td>
          </tr>
          <tr>
            <td className="pl-[30px] pt-[28px] font-normal text-[16px] leading-[19.36px]  max-sm:text-[12px] max-sm:font-normal max-sm:pt-[12px] max-sm:pl-[14px]">
              02.
            </td>
            <td className="pl-[52px] pt-[28px] font-normal text-[16px] leading-[19.36px] max-sm:text-[12px] max-sm:font-normal max-sm:pt-[12px] max-sm:pl-[37px]">
              Expedia
            </td>
            <td className="pl-[86px] pt-[28px] font-normal text-[16px] leading-[19.36px] max-sm:text-[12px] max-sm:font-normal max-sm:pl-[55px] max-sm:pt-[12px]">
              $450
            </td>
            <td className="pl-[64px] pt-[28px] font-normal text-[16px] leading-[19.36px] max-sm:text-[12px] max-sm:font-normal max-sm:pl-[53px] max-sm:pt-[12px]">
              +3%
            </td>
          </tr>
          <tr>
            <td className="pl-[30px] pt-[28px] font-normal text-[16px] leading-[19.36px]  max-sm:text-[12px] max-sm:font-normal max-sm:pt-[12px] max-sm:pl-[14px]">
              03.
            </td>
            <td className="pl-[52px] pt-[28px] font-normal text-[16px] leading-[19.36px] max-sm:text-[12px] max-sm:font-normal max-sm:pt-[12px] max-sm:pl-[37px]">
              Airbnb
            </td>
            <td className="pl-[86px] pt-[28px] font-normal text-[16px] leading-[19.36px] max-sm:text-[12px] max-sm:font-normal max-sm:pl-[55px] max-sm:pt-[12px]">
              $670
            </td>
            <td className="pl-[64px] pt-[28px] font-normal text-[16px] leading-[19.36px] max-sm:text-[12px] max-sm:font-normal max-sm:pl-[53px] max-sm:pt-[12px]">
              +7%
            </td>
          </tr>
          <tr>
            <td className="pl-[30px] pt-[28px] font-normal text-[16px] leading-[19.36px]  max-sm:text-[12px] max-sm:font-normal max-sm:pt-[12px] max-sm:pl-[14px]">
              04.
            </td>
            <td className="pl-[52px] pt-[28px] font-normal text-[16px] leading-[19.36px] max-sm:text-[12px] max-sm:font-normal max-sm:pt-[12px] max-sm:pl-[37px]">
              Booking
            </td>
            <td className="pl-[86px] pt-[28px] font-normal text-[16px] leading-[19.36px] max-sm:text-[12px] max-sm:font-normal max-sm:pl-[55px] max-sm:pt-[12px]">
              $490
            </td>
            <td className="pl-[64px] pt-[28px] font-normal text-[16px] leading-[19.36px] max-sm:text-[12px] max-sm:font-normal max-sm:pl-[53px] max-sm:pt-[12px]">
              +4%
            </td>
          </tr>
          <tr className="">
            <td className="pl-[30px] pt-[28px] font-normal text-[16px] leading-[19.36px] max-sm:pl-[14px] max-sm:text-[12px] max-sm:font-normal max-sm:pt-[12px]">
              05.
            </td>
            <td className="pl-[52px] pt-[28px] font-normal text-[16px] leading-[19.36px] max-sm:text-[12px] max-sm:font-normal max-sm:pt-[12px] max-sm:pl-[37px]">
              Kayak
            </td>
            <td className="pl-[86px] pt-[28px] font-normal text-[16px] leading-[19.36px] max-sm:text-[12px] max-sm:font-normal max-sm:pl-[55px] max-sm:pt-[12px]">
              $380
            </td>
            <td className="pl-[64px] pt-[28px] font-normal text-[16px] leading-[19.36px] max-sm:text-[12px] max-sm:font-normal max-sm:pl-[53px] max-sm:pt-[12px]">
              +2%
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  );
};

export default TrendingStock;
