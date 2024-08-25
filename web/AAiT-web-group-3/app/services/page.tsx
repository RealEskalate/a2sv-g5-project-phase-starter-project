import Card from "./(components)/card";
import { FaShieldHeart } from "react-icons/fa6";
import { GiShoppingBag } from "react-icons/gi";
import { AiFillSafetyCertificate } from "react-icons/ai";
import { FaHandHoldingDollar } from "react-icons/fa6";
import ServiceList from "./(components)/ServiceList";
import { FaBriefcase } from "react-icons/fa";

const page = () => {
  return (
    <>
      <div className="bg-gray-100 px-[25px] md:px-[50px]">
        <div className="mt-10 overflow-x-auto no-scrollbar ">
          <div className="flex space-x-4 md:space-x-16 ">
            <Card
              icon={FaShieldHeart}
              title="Life Insurance"
              sub_title="Unlimited Protection"
              icon_bg="#E7EDFF"
              icon_color="#396AFF"
            />
            <Card
              icon={GiShoppingBag}
              title="Shopping"
              sub_title="Buy.Think.Grow."
              icon_bg="#FFF5D9"
              icon_color="#FFBB38"
            />
            <Card
              icon={AiFillSafetyCertificate}
              title="Safety"
              sub_title="We are your allies"
              icon_bg="#DCFAF8"
              icon_color="#16DBCC"
            />
          </div>
        </div>
        <p className="font-inter font-semibold text-[#343C6A] text-[22px] leading-[26.63px] my-5">
          Bank Services List
        </p>
        <ServiceList
          
        />
     
      </div>
    </>
  );
};

export default page;
